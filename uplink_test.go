package goeverynet_test

import (
	"encoding/json"
	"testing"

	"github.com/douglaszuqueto/goeverynet"
)

func TestUplink(t *testing.T) {
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

	var payload goeverynet.UplinkFrame

	err := json.Unmarshal([]byte(packet), &payload)
	if err != nil {
		t.Error(err)
	}

	t.Log(payload.Type)
}
