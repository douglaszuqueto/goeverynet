package goeverynet

import "encoding/json"

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

// NewDownlinkFrame creates a new DownlinkFrame from a byte array
func NewDownlinkFrame(data []byte) (*DownlinkFrame, error) {
	var frame DownlinkFrame

	err := json.Unmarshal(data, &frame)
	if err != nil {
		return nil, err
	}

	err = frame.Validate()
	if err != nil {
		return nil, err
	}

	return &frame, nil
}

// MarshalJSON returns the JSON encoding of a DownlinkFrame
func (frame *DownlinkFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(*frame)
}

// Validate validates a DownlinkFrame
func (frame *DownlinkFrame) Validate() error {
	return nil
}
