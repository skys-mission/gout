// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package systemlw

import (
	"fmt"

	"github.com/skys-mission/gout/common"
	"golang.org/x/sys/windows"
)

// 定义了一系列语言代码常量，用于表示不同语言的IANA标准代码。
const (
	LanguageCodeEnglishIANA            = "en"
	LanguageCodeSimplifiedChineseIANA  = "zh-Hans"
	LanguageCodeTraditionalChineseIANA = "zh-Hant"
	LanguageCodeJapaneseIANA           = "ja"
	LanguageCodeKoreanIANA             = "ko"
	LanguageCodeFrenchIANA             = "fr"
	LanguageCodeRussianIANA            = "ru"
	LanguageCodeGermanIANA             = "de"
	LanguageCodeSpanishIANA            = "es"
)

// GetSystemLanguageCodeIANA 获取系统语言名称。
// 返回值是一个字符串，表示系统的默认语言名称，以及一个错误值，如果获取语言ID时发生错误。
func GetSystemLanguageCodeIANA() (string, error) {
	langID, err := getUserDefaultUILanguage()
	if err != nil {
		return "error", err
	}
	langCode := fmt.Sprintf("%04x", langID)

	switch langCode {
	case "0409":
		return LanguageCodeEnglishIANA, nil
	case "0804":
		return LanguageCodeSimplifiedChineseIANA, nil
	case "0404":
		return LanguageCodeTraditionalChineseIANA, nil
	case "0411":
		return LanguageCodeJapaneseIANA, nil
	case "0412":
		return LanguageCodeKoreanIANA, nil
	case "040C":
		return LanguageCodeFrenchIANA, nil
	case "0419":
		return LanguageCodeRussianIANA, nil
	case "0407":
		return LanguageCodeGermanIANA, nil
	case "0C0A":
		return LanguageCodeSpanishIANA, nil
	default:
		return "temporarily not recognized", nil
	}
}

// GetSystemLanguageCodeLCID 获取系统国家代码。
// 返回值是一个字符串，表示系统的默认语言ID的十六进制表示，以及一个错误值，如果获取语言ID时发生错误。
func GetSystemLanguageCodeLCID() (string, error) {
	langID, err := getUserDefaultUILanguage()
	return fmt.Sprintf("%04x", langID), err
}

// getUserDefaultUILanguage 获取用户的默认UI语言。
// 返回值是一个uint16，表示语言的LCID值，以及一个错误值，如果调用Windows API时发生错误。
func getUserDefaultUILanguage() (uint16, error) {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	procGetUserDefaultUILanguage := kernel32.NewProc("GetUserDefaultUILanguage")

	langID, _, err := procGetUserDefaultUILanguage.Call()
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return uint16(langID), err
}
