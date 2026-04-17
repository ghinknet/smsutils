package model

type Client interface {
	SendMessage(dest string, sender string, template string, vars Vars) error
}
