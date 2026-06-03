# smsutils

Golang common SMS utilities.

`smsutils` is the **core module** of the smsutils ecosystem — a pluggable, multi-provider
SMS sending library for Go. It defines the common interfaces, models, error types and
phone-number helpers that every cloud provider driver implements. Applications depend on
this core module plus one or more driver modules, and send messages through a single,
unified API regardless of the underlying provider.

## Architecture

The ecosystem follows a **plugin driver pattern**. The core module exposes interfaces and a
global driver registry; each provider ships as a separate module that registers itself via
an `init()` function when imported.

```
                ┌──────────────────────────┐
                │     Application code     │
                └────────────┬─────────────┘
                             │ client.NewClient(config)
                             ▼
        ┌────────────────────────────────────────────┐
        │            smsutils/v3 (core)              │
        │  model/   interfaces: Client, Driver,Config│
        │  client/  NewClient factory                │
        │  driver/  Register()                        │
        │  errors/  SmsutilsError                     │
        │  utils/   phone number parsing              │
        │  internal/state/  global driver registry   │
        └────────────────────────────────────────────┘
                             ▲ driver.Register(Name, Driver{})
        ┌──────────┬─────────┼──────────┬──────────┬──────────┐
        │          │         │          │          │          │
     aliyun     qcloud      bce       ucloud      volc     jdcloud
   (Alibaba)  (Tencent)   (Baidu)   (UCloud)  (Volcengine) (planned)
```

### Module map

| Module | Import path | Provider |
|--------|-------------|----------|
| smsutils (this) | `go.gh.ink/smsutils/v3` | Core library |
| smsutils-aliyun | `go.gh.ink/smsutils/aliyun/v3` | Alibaba Cloud |
| smsutils-qcloud | `go.gh.ink/smsutils/qcloud/v3` | Tencent Cloud |
| smsutils-bce | `go.gh.ink/smsutils/bce/v3` | Baidu Cloud Engine |
| smsutils-ucloud | `go.gh.ink/smsutils/ucloud/v3` | UCloud |
| smsutils-volc | `go.gh.ink/smsutils/volc/v3` | Volcengine |
| smsutils-jdcloud | `go.gh.ink/smsutils/jdcloud/v3` | JD Cloud (planned, not yet implemented) |

## Installation

```bash
go get go.gh.ink/smsutils/v3
# plus one or more drivers, e.g.
go get go.gh.ink/smsutils/aliyun/v3
```

## Usage

1. Blank-import the drivers you want so they self-register.
2. Build a `model.Config` mapping each driver name to its credential map.
3. Call `client.NewClient` to obtain a `model.Client` per driver.
4. Call `SendMessage` on the client.

```go
package main

import (
	"go.gh.ink/smsutils/v3/client"
	"go.gh.ink/smsutils/v3/model"

	// Blank-import drivers to register them.
	_ "go.gh.ink/smsutils/aliyun/v3"
	_ "go.gh.ink/smsutils/qcloud/v3"
)

func main() {
	clients, err := client.NewClient(model.Config{
		Credentials: model.C{
			"aliyun": {
				"accessKeyID":     "<your-access-key-id>",
				"accessKeySecret": "<your-access-key-secret>",
			},
			"qcloud": {
				"secretID":    "<your-secret-id>",
				"secretKey":   "<your-secret-key>",
				"smsSdkAppID": "<your-app-id>",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// Send a templated message through the Aliyun driver.
	err = clients["aliyun"].SendMessage(
		"+8617601205205",       // destination phone number
		"YourSignName",         // sender / signature
		"SMS_123456789",        // template ID / code
		model.Vars{
			{Key: "code", Value: "1234"},
		},
	)
	if err != nil {
		panic(err)
	}
}
```

## Core API

### `client.NewClient`

```go
func NewClient(config model.Config) (clients map[string]model.Client, err error)
```

Iterates over `config.Credentials`, looks each driver up in the global registry, and
constructs a client per driver. Returns a map keyed by driver name. If a referenced driver
was never registered (i.e. you forgot to import it), it returns
`errors.ErrDriverNotRegistered`.

### `model.Config`

```go
type Config struct {
	Credentials C // map[driverName]map[credentialKey]credentialValue
	Marshal     func(v any) ([]byte, error)        // optional, defaults to json.Marshal
	Unmarshal   func(data []byte, v any) error      // optional, defaults to json.Unmarshal
}

type C = map[string]map[string]string
```

`Marshal`/`Unmarshal` are passed down to drivers (used to encode template parameters and
decode any structured credential fields). Leave them nil to use the standard library JSON.

### `model.Client`

```go
type Client interface {
	SendMessage(dest string, sender string, template string, vars Vars) error
}
```

- `dest` — destination phone number. Accepts E.164 (`+8617601205205`), bare international
  (`8617601205205`), or bare Chinese mobile (`17601205205`); normalization is handled by
  the `utils` package (see below).
- `sender` — provider signature / sign name.
- `template` — provider template ID or code.
- `vars` — ordered template variables.

### `model.Vars`

```go
type Var struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Vars []*Var
```

Note that ordering matters for providers (such as Tencent Cloud and UCloud) that bind
template parameters positionally rather than by key.

### `model.Driver` (for driver authors)

```go
type Driver interface {
	NewClient(params DriverClientParam) (Client, error)
}

type DriverClientParam struct {
	Credential map[string]string
	Unmarshal  func(data []byte, v any) error
	Marshal    func(v any) ([]byte, error)
}
```

Drivers register themselves with `driver.Register`:

```go
func Register(name string, driver model.Driver)
```

Typically called from a package `init()`:

```go
func init() {
	driver.Register(Name, Driver{})
}
```

## Errors

The `errors` package provides `SmsutilsError`, a rich error type carrying provider context.
Three sentinel errors are exported:

| Sentinel | Meaning |
|----------|---------|
| `errors.ErrDriverNotRegistered` | Credentials referenced a driver that was not imported/registered |
| `errors.ErrDriverCredentialInvalid` | Required credential fields were missing or empty |
| `errors.ErrDriverSendFailed` | The provider returned a non-success response |

`SmsutilsError` supports `errors.Is` against these sentinels even after being decorated, and
exposes provider diagnostics via chained `With*` builders / getters:

```go
err := clients["aliyun"].SendMessage(dest, sender, template, vars)
if err != nil {
	var smsErr *errors.SmsutilsError
	if errors.As(err, &smsErr) && errors.Is(err, smserrors.ErrDriverSendFailed) {
		log.Printf("driver=%s code=%s message=%s requestID=%s",
			smsErr.DriverName(),
			smsErr.DriverCode(),
			smsErr.DriverMessage(),
			smsErr.DriverRequestID(),
		)
	}
}
```

Builder / accessor pairs: `WithDriverName`/`DriverName`, `WithDriverCode`/`DriverCode`,
`WithDriverMessage`/`DriverMessage`, `WithDriverRequestID`/`DriverRequestID`,
`WithDriverResponse`/`DriverResponse`. Each `With*` call clones the error so the exported
sentinels stay immutable and `errors.Is` continues to work.

## Phone number utilities (`utils`)

```go
func ParseNumber(number string) (countryCode int64, nationalNumber int64, regionCode string, err error)
func ProcessNumberForChinese(to string) (toProcessed string, countryCode int64, nationalNumber int64, regionCode string, err error)
```

- `ParseNumber` wraps [nyaruka/phonenumbers](https://github.com/nyaruka/phonenumbers),
  prepending `+` when missing.
- `ProcessNumberForChinese` normalizes numbers with China-friendly heuristics. Because a bare
  11-digit number like `17601205205` can be ambiguously parsed as a US number, it detects the
  Chinese mobile pattern (11 digits, leading `1`, second digit `3`–`9`) and forces the `+86`
  country code / `CN` region. All bundled drivers call this before sending so that bare
  Chinese mobile numbers route correctly.

## Driver development

To add a new provider:

1. Create a new module (e.g. `go.gh.ink/smsutils/<provider>/v3`).
2. Define `const Name = "<provider>"` and credential-key constants.
3. Implement `model.Driver` (a `NewClient` factory) and `model.Client` (a `SendMessage`).
4. Validate required credentials, returning `errors.ErrDriverCredentialInvalid` if missing.
5. Call `utils.ProcessNumberForChinese` to normalize the destination.
6. Map provider failures to `errors.ErrDriverSendFailed` with the `With*` decorators.
7. Register the driver from `init()` via `driver.Register(Name, Driver{})`.

See any of the sibling driver modules for reference implementations.

## Requirements

- Go 1.25.0+
- Dependencies: [`github.com/nyaruka/phonenumbers`](https://github.com/nyaruka/phonenumbers),
  `go.gh.ink/toolbox`

## License

[Apache License 2.0](LICENSE)
