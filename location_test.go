package goeverynet_test

import (
	"testing"

	"github.com/douglaszuqueto/goeverynet"
)

func TestLocation(t *testing.T) {
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

	frame, err := goeverynet.NewLocationFrame([]byte(packet))
	if err != nil {
		t.Fatal(err)
	}

	if frame.Type != "location" {
		t.Errorf("frame type is not location")
	}
}
