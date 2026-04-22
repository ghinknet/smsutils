package client

import (
	"encoding/json"

	"go.gh.ink/smsutils/v3/errors"
	"go.gh.ink/smsutils/v3/internal/state"
	"go.gh.ink/smsutils/v3/model"
)

func NewClient(config model.Config) (clients map[string]model.Client, err error) {
	// Prepare JSON
	if config.Marshal == nil {
		config.Marshal = json.Marshal
	}
	if config.Unmarshal == nil {
		config.Unmarshal = json.Unmarshal
	}

	clients = make(map[string]model.Client)

	for driverName, driverCredential := range config.Credentials {
		// Check driver registered
		newDriverClient, ok := state.Drivers[driverName]
		if !ok {
			return nil, errors.ErrDriverNotRegistered.WithDriverName(driverName)
		}

		// Create driver client
		clients[driverName], err = newDriverClient.NewClient(model.DriverClientParam{
			Credential: driverCredential,
			// JSON
			Marshal:   config.Marshal,
			Unmarshal: config.Unmarshal,
		})
		if err != nil {
			return nil, err
		}
	}

	return clients, nil
}
