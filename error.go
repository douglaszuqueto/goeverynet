package goeverynet

type ErrorFrame struct {
	Type   string      `json:"type"`
	Params ErrorParams `json:"params"`
	Meta   ErrorMeta   `json:"meta"`
}

type ErrorParams struct {
	Message string `json:"message"`
}

type ErrorMeta struct {
	Time float64 `json:"time"`
}
