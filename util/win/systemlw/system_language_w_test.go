package systemlw

import (
	"fmt"
	"testing"
)

func TestGetSystemCountryCode(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf(fmt.Sprintf("panic: %v", err))
		}
	}()
	code, err := GetSystemCountryCode()
	if err != nil {
		t.Errorf(fmt.Sprintf("GetSystemCountryCode: %v", err))
		return
	}
	t.Logf("Country Code:%v", code)

}

func TestGetSystemLanguage(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf(fmt.Sprintf("panic: %v", err))
		}
	}()
	name, err := GetSystemLanguageName()
	if err != nil {
		t.Errorf(fmt.Sprintf("GetSystemLanguageName: %v", err))
		return
	}
	t.Logf("Language Name:%s", name)
}
