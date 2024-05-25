package systemlw

import (
	"fmt"

	"github.com/skys-mission/gout/common"

	"golang.org/x/sys/windows"
)

// Define supported language constants.
const (
	English            = "en"
	SimplifiedChinese  = "zh-CN"
	TraditionalChinese = "zh-TW"
	Japanese           = "ja"
	Korean             = "ko"
	French             = "fr"
	Russian            = "ru"
	German             = "de"
	Spanish            = "es"
)

// GetSystemLanguageName Return the name corresponding to the system default UI language.
func GetSystemLanguageName() (string, error) {
	langID, err := getUserDefaultUILanguage()
	if err != nil {
		return "error", err
	}
	langCode := fmt.Sprintf("%04x", langID)

	switch langCode {
	case "0409":
		return English, nil
	case "0804":
		return SimplifiedChinese, nil
	case "0404":
		return TraditionalChinese, nil
	case "0411":
		return Japanese, nil
	case "0412":
		return Korean, nil
	case "040C":
		return French, nil
	case "0419":
		return Russian, nil
	case "0407":
		return German, nil
	case "0C0A":
		return Spanish, nil
	default:
		return "other", nil
	}
}

// GetSystemCountryCode Return the country code for the system default language.
func GetSystemCountryCode() (string, error) {
	langID, err := getUserDefaultUILanguage()
	return fmt.Sprintf("%04x", langID), err
}

// getUserDefaultUILanguage Invoke the Windows API to obtain the system default user interface language.
func getUserDefaultUILanguage() (uint16, error) {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	procGetUserDefaultUILanguage := kernel32.NewProc("GetUserDefaultUILanguage")

	langID, _, err := procGetUserDefaultUILanguage.Call()
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return uint16(langID), err
}
