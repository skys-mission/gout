package displayw

import (
	"testing"
)

func TestGetAllDisplays(t *testing.T) {
	infos, err := GetAllDisplays()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for i, display := range infos {
		t.Logf("Display %d: %+v\n", i+1, display)
	}
}

func TestGetPrimaryDisplay(t *testing.T) {
	display, err := GetPrimaryDisplay()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("Display: %+v\n", display)

}
