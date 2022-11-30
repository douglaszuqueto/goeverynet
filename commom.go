package goeverynet

type Gps struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Alt int     `json:"alt"`
}

type Hardware struct {
	Status  int     `json:"status"`
	Chain   int     `json:"chain"`
	Power   int     `json:"power"`
	Tmst    int     `json:"tmst"`
	Snr     float64 `json:"snr"`
	Rssi    int     `json:"rssi"`
	Channel int     `json:"channel"`
	Gps     Gps     `json:"gps"`
}

type Radio struct {
	GpsTime    int64      `json:"gps_time"`
	Delay      float64    `json:"delay"`
	Datarate   int        `json:"datarate"`
	Modulation Modulation `json:"modulation"`
	Datr       string     `json:"datr"`
	Hardware   Hardware   `json:"hardware"`
	Time       float64    `json:"time"`
	Freq       float64    `json:"freq"`
	Size       int        `json:"size"`
}

type Modulation struct {
	Bandwidth int    `json:"bandwidth"`
	Type      string `json:"type"`
	Coderate  string `json:"coderate"`
	Spreading int    `json:"spreading"`
	Inverted  bool   `json:"inverted"`
}
