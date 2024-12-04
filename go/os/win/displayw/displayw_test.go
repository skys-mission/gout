// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package displayw

import (
	"testing"
)

func TestGetAllDisplayInfo(t *testing.T) {
	infos, err := GetAllDisplayInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for i, display := range infos {
		t.Logf("Display %d: %+v\n", i+1, display)
	}
}

func TestGetPrimaryDisplayInfo(t *testing.T) {
	display, err := GetPrimaryDisplayInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("Display: %+v\n", display)

}
