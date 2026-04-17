package client

import (
	"encoding/json"

	"github.com/ghinknet/smsutils/v3/errors"
	"github.com/ghinknet/smsutils/v3/internal/state"
	"github.com/ghinknet/smsutils/v3/model"
)

func NewClient(config model.Config) (map[string]model.Client, error) {
	// Prepare JSON
	if config.Marshal == nil {
		config.Marshal = json.Marshal
	}
	if config.Unmarshal == nil {
		config.Unmarshal = json.Unmarshal
	}

	clients := make(map[string]model.Client)

	for driverName, driverCredential := range config.Credentials {
		// Check driver registered
		newDriverClient, ok := state.Drivers[driverName]
		if !ok {
			return nil, errors.ErrDriverNotRegistered
		}

		// Create driver client
		clients[driverName] = newDriverClient.NewClient(model.DriverClientParam{
			Credential: driverCredential,
			// JSON
			Marshal:   config.Marshal,
			Unmarshal: config.Unmarshal,
		})
	}

	return clients, nil
}
