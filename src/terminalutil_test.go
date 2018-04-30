package main

import "testing"

func TestGetTerminalDataRaw1(t *testing.T) {

	width, height, err := parseTerminalSizeRaw("10 20")

	if err != nil {
		t.Errorf("Got exception: %s\n", err.Error())
	}

	if width != 20 {
		t.Errorf("Got wrong length: %d", width)
	}

	if height != 10 {
		t.Errorf("Got wrong height: %d", height)
	}

}

func TestGetTerminalDataRaw2(t *testing.T) {

	width, height, err := parseTerminalSizeRaw("10 20\n")

	if err != nil {
		t.Errorf("Got exception: %s\n", err.Error())
	}

	if width != 20 {
		t.Errorf("Got wrong length: %d", width)
	}

	if height != 10 {
		t.Errorf("Got wrong height: %d", height)
	}

}
