package goeverynet

import "encoding/json"

type UplinkFrame struct {
	Type   string       `json:"type"`
	Params UplinkParams `json:"params"`
	Meta   Meta         `json:"meta"`
}

type UplinkParams struct {
	RxTime           float64    `json:"rx_time"`
	Port             int        `json:"port"`
	Duplicate        bool       `json:"duplicate"`
	Radio            Radio      `json:"radio"`
	CounterUp        int        `json:"counter_up"`
	Lora             UplinkLora `json:"lora"`
	Payload          string     `json:"payload"`
	EncryptedPayload string     `json:"encrypted_payload"`
}

type UplinkHeader struct {
	ClassB    bool `json:"class_b"`
	Confirmed bool `json:"confirmed"`
	Adr       bool `json:"adr"`
	Ack       bool `json:"ack"`
	AdrAckReq bool `json:"adr_ack_req"`
	Version   int  `json:"version"`
	Type      int  `json:"type"`
}

type UplinkLora struct {
	Header      UplinkHeader  `json:"header"`
	MacCommands []interface{} `json:"mac_commands"`
}

// NewUplinkFrame creates a new UplinkFrame
func NewUplinkFrame(payload []byte) (*UplinkFrame, error) {
	frame := &UplinkFrame{}

	err := json.Unmarshal(payload, frame)
	if err != nil {
		return nil, err
	}

	err = frame.Validate()
	if err != nil {
		return nil, err
	}

	return frame, nil
}

// MarshalJSON returns the JSON encoding of a UplinkFrame
func (frame *UplinkFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(*frame)
}

// Validate validates a UplinkFrame
func (frame *UplinkFrame) Validate() error {
	return nil
}
