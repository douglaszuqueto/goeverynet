package goeverynet

type LocationFrame struct {
	Type   string         `json:"type"`
	Params LocationParams `json:"params"`
	Meta   Meta           `json:"meta"`
}

type LocationParams struct {
	Solutions []Solutions `json:"solutions"`
}

type Solutions struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Quality   int     `json:"quality"`
	Precision int     `json:"precision"`
	Method    string  `json:"method"`
}
