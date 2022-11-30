package goeverynet

type DownlinkFrame struct {
	Type   string         `json:"type"`
	Params DownlinkParams `json:"params"`
	Meta   Meta           `json:"meta"`
}

type DownlinkHeader struct {
	Confirmed bool `json:"confirmed"`
	Adr       bool `json:"adr"`
	Ack       bool `json:"ack"`
	Version   int  `json:"version"`
	Type      int  `json:"type"`
	Pending   bool `json:"pending"`
}

type DownlinkLora struct {
	Header      DownlinkHeader `json:"header"`
	MacCommands []MacCommands  `json:"mac_commands"`
}

type DownlinkParams struct {
	Payload          string       `json:"payload"`
	Radio            Radio        `json:"radio"`
	CounterDown      int          `json:"counter_down"`
	Lora             DownlinkLora `json:"lora"`
	Port             int          `json:"port"`
	EncryptedPayload string       `json:"encrypted_payload"`
}
