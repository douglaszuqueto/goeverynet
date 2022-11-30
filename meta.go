package goeverynet

type Meta struct {
	Network     string   `json:"network"`
	Tags        []string `json:"tags"`
	PacketHash  string   `json:"packet_hash"`
	Application string   `json:"application"`
	DeviceAddr  string   `json:"device_addr"`
	Time        float64  `json:"time"`
	Device      string   `json:"device"`
	PacketID    string   `json:"packet_id"`
	Gateway     string   `json:"gateway"`
}
