package goeverynet_test

import (
	"encoding/json"
	"testing"

	"github.com/douglaszuqueto/goeverynet"
)

func TestDownlink(t *testing.T) {
	packet := `
	{
		"params": {
		  "payload": "A0UBAHEDBf8AAQMFAABBBAAFCGjijA==",
		  "radio": {
			"gps_time": 1353754115175,
			"delay": 0.04933524131774902,
			"datarate": 2,
			"modulation": {
			  "bandwidth": 500000,
			  "type": "LORA",
			  "coderate": "4/5",
			  "spreading": 10,
			  "inverted": true
			},
			"datr": "SF10BW500",
			"hardware": {
			  "status": 1,
			  "chain": 0,
			  "power": 26,
			  "tmst": 1892997124,
			  "snr": -0.8,
			  "rssi": -107,
			  "channel": 6,
			  "gps": {
				"lat": -25.46276092529297,
				"lng": -49.2612190246582,
				"alt": 926
			  }
			},
			"time": 1669718902.1950538,
			"freq": 926.9,
			"size": 38
		  },
		  "counter_down": 2,
		  "lora": {
			"header": {
			  "confirmed": false,
			  "adr": true,
			  "ack": false,
			  "version": 0,
			  "type": 3,
			  "pending": true
			},
			"mac_commands": [
			  {
				"LinkADRReq": {
				  "Redundancy": {
					"ChMaskCntl": 7,
					"NbTrans": 1
				  },
				  "DataRate_TXPower": {
					"TXPower": 5,
					"DataRate": 4
				  },
				  "ChMask": 1
				}
			  },
			  {
				"LinkADRReq": {
				  "Redundancy": {
					"ChMaskCntl": 0,
					"NbTrans": 1
				  },
				  "DataRate_TXPower": {
					"TXPower": 5,
					"DataRate": 0
				  },
				  "ChMask": 255
				}
			  },
			  {
				"LinkADRReq": {
				  "Redundancy": {
					"ChMaskCntl": 4,
					"NbTrans": 1
				  },
				  "DataRate_TXPower": {
					"TXPower": 5,
					"DataRate": 0
				  },
				  "ChMask": 0
				}
			  },
			  {
				"DutyCycleReq": {
				  "MaxDCycle": 0
				}
			  },
			  {
				"RXParamSetupReq": {
				  "Frequency": 9233000,
				  "DLsettings": {
					"RX2DataRate": 8,
					"RX1DRoffset": 0
				  }
				}
			  }
			]
		  },
		  "port": 0,
		  "encrypted_payload": "E1DxlxHYzl6mYSP1wCJ0ngWx6kPYlQ=="
		},
		"meta": {
		  "network": "358b679dd88d47948f6e108677b4b79a",
		  "tags": ["jvtech"],
		  "packet_hash": "ea834f866877e0c159d3221d6713adb4",
		  "application": "6caabba5aa141e23",
		  "device_addr": "162a9521",
		  "time": 1669718899.733,
		  "device": "000516800011abf6",
		  "packet_id": "4335a9f364fae746e0b433067bc26ae5",
		  "gateway": "b0fd0b7002640000"
		},
		"type": "downlink"
	  }
	`

	var event goeverynet.DownlinkFrame

	err := json.Unmarshal([]byte(packet), &event)
	if err != nil {
		t.Fatal(err)
	}
}
