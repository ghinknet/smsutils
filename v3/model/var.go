package model

type Var struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Vars []*Var
