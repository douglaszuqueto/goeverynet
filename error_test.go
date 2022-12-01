package goeverynet_test

import (
	"testing"

	"github.com/douglaszuqueto/goeverynet"
)

func TestError(t *testing.T) {
	packet := `
	{
		"params": {
			"message": "Bad message format. Received data: {\"success\": true}"
		},
		"meta": {
			"time": 1669718933.089572
		},
		"type": "error"
	}
	`

	frame, err := goeverynet.NewErrorFrame([]byte(packet))
	if err != nil {
		t.Fatal(err)
	}

	if frame.Type != "error" {
		t.Errorf("frame type is not error")
	}
}
