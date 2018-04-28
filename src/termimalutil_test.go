package main

import "testing"

func TestGetTerminalDataRaw(t *testing.T) {
	data, err := determineTerminalSizeRaw()

	if len(data) <= 0 {
		//	t.Error("Got null length: " + data)
	}

	if err != nil {
		//t.Errorf("Got exception: %s\n", err.Error())
	}

}
