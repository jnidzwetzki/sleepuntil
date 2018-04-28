package terminalutil

import "testing"

func TestGetTerminalDataRaw(t *testing.T) {
	data, err := determineSize()

	if len(data) <= 0 {
		//t.Error("Got null length")
	}

	if err != nil {
		//t.Errorf("Got exception: %s\n", err.Error())
	}

}
