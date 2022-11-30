package goeverynet_test

import (
	"encoding/json"
	"testing"

	"github.com/douglaszuqueto/goeverynet"
)

func TestLocationFrame(t *testing.T) {
	packet := `
	{
		"params": {
		  "solutions": [
			{
			  "lat": -25.48264,
			  "lng": -49.24762,
			  "quality": 1,
			  "precision": 4000,
			  "method": "simple:moving"
			}
		  ]
		},
		"meta": {
		  "tags": ["jvtech"],
		  "packet_hash": "2d12d09d48669879a40f9779f12324a4",
		  "application": "6caabba5aa141e23",
		  "device_addr": "162a9521",
		  "time": 1669718932.881,
		  "device": "000516800011abf6"
		},
		"type": "location"
	  }	  
	`

	frameType, err := goeverynet.GetFrameType([]byte(packet))

	if err != nil {
		t.Fatal(err)
	}

	switch frameType {
	case goeverynet.LocationEvent:
		var payload goeverynet.LocationFrame

		err = json.Unmarshal([]byte(packet), &payload)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUplinkFrame(t *testing.T) {
	packet := `
	{
		"params": {
		  "rx_time": 1669718927.7951567,
		  "port": 1,
		  "duplicate": false,
		  "radio": {
			"gps_time": 1353754145752,
			"delay": 0.05233168601989746,
			"datarate": 0,
			"modulation": {
			  "bandwidth": 125000,
			  "type": "LORA",
			  "spreading": 12,
			  "coderate": "4/5"
			},
			"hardware": {
			  "status": 1,
			  "chain": 0,
			  "tmst": 1918574140,
			  "snr": -9.8,
			  "rssi": -97,
			  "channel": 3,
			  "gps": {
				"lat": -25.468379974365234,
				"lng": -49.30628967285156,
				"alt": 954
			  }
			},
			"time": 1669718927.7951567,
			"freq": 915.8,
			"size": 48
		  },
		  "counter_up": 6,
		  "lora": {
			"header": {
			  "class_b": false,
			  "confirmed": false,
			  "adr": true,
			  "ack": false,
			  "adr_ack_req": false,
			  "version": 0,
			  "type": 2
			},
			"mac_commands": []
		  },
		  "payload": "eyJUIjoyMC44LCJVIjo5My40LCJDIjoxMH0=",
		  "encrypted_payload": "THVy2YePaE57oNspHz8inGaoZ24EldRNtz8="
		},
		"meta": {
		  "network": "358b679dd88d47948f6e108677b4b79a",
		  "tags": ["jvtech"],
		  "packet_hash": "2d12d09d48669879a40f9779f12324a4",
		  "application": "6caabba5aa141e23",
		  "device_addr": "162a9521",
		  "time": 1669718927.862,
		  "device": "000516800011abf6",
		  "packet_id": "6e1735e722f0dd6835e9cd79bd8f740f",
		  "gateway": "b0fd0b7002640000"
		},
		"type": "uplink"
	  }	  
	`

	frameType, err := goeverynet.GetFrameType([]byte(packet))

	if err != nil {
		t.Fatal(err)
	}

	switch frameType {
	case goeverynet.UplinkEvent:
		var payload goeverynet.UplinkFrame

		err = json.Unmarshal([]byte(packet), &payload)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDownlinkFrame(t *testing.T) {
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

	frameType, err := goeverynet.GetFrameType([]byte(packet))

	if err != nil {
		t.Fatal(err)
	}

	switch frameType {
	case goeverynet.DownlinkEvent:
		var payload goeverynet.DownlinkFrame

		err = json.Unmarshal([]byte(packet), &payload)
		if err != nil {
			t.Error(err)
		}
	}
}
