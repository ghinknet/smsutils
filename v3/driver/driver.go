package driver

import (
	"github.com/ghinknet/smsutils/v3/internal/state"
	"github.com/ghinknet/smsutils/v3/model"
)

func Register(name string, driver model.Driver) {
	state.Drivers[name] = driver
}
