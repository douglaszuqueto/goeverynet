package goeverynet

type Linkadrreq struct {
	Redundancy      Redundancy      `json:"Redundancy"`
	DatarateTxpower DatarateTxpower `json:"DataRate_TXPower"`
	Chmask          int             `json:"ChMask"`
}

type Redundancy struct {
	Chmaskcntl int `json:"ChMaskCntl"`
	Nbtrans    int `json:"NbTrans"`
}

type DatarateTxpower struct {
	Txpower  int `json:"TXPower"`
	Datarate int `json:"DataRate"`
}

type Dutycyclereq struct {
	Maxdcycle int `json:"MaxDCycle"`
}

type Dlsettings struct {
	Rx2Datarate int `json:"RX2DataRate"`
	Rx1Droffset int `json:"RX1DRoffset"`
}

type Rxparamsetupreq struct {
	Frequency  int        `json:"Frequency"`
	Dlsettings Dlsettings `json:"DLsettings"`
}

type MacCommands struct {
	Linkadrreq      Linkadrreq      `json:"LinkADRReq,omitempty"`
	Dutycyclereq    Dutycyclereq    `json:"DutyCycleReq,omitempty"`
	Rxparamsetupreq Rxparamsetupreq `json:"RXParamSetupReq,omitempty"`
}
