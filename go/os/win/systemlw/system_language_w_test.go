package systemlw

import (
	"fmt"
	"testing"
)

func TestGetSystemLanguageCodeIANA(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf(fmt.Sprintf("panic: %v", err))
		}
	}()
	code, err := GetSystemLanguageCodeIANA()
	if err != nil {
		t.Errorf(fmt.Sprintf("GetSystemCountryCode: %v", err))
		return
	}
	t.Logf("Country Code:%v", code)

}

func TesGetSystemLanguageCodeLCID(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf(fmt.Sprintf("panic: %v", err))
		}
	}()
	name, err := GetSystemLanguageCodeLCID()
	if err != nil {
		t.Errorf(fmt.Sprintf("GetSystemLanguageName: %v", err))
		return
	}
	t.Logf("Language Name:%s", name)
}
