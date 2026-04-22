package driver

import (
	"go.gh.ink/smsutils/v3/internal/state"
	"go.gh.ink/smsutils/v3/model"
)

func Register(name string, driver model.Driver) {
	state.Drivers[name] = driver
}
