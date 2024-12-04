// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package mbw

import (
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/skys-mission/gout/common"
)

const user32API = "user32.dll"

// MessageBox format constants
const (
	MB_OK                   = 0x00000000         // OK button only
	MB_OKCANCEL             = 0x00000001         // OK and Cancel buttons
	MB_ABORTRETRYIGNORE     = 0x00000002         // Abort, Retry, and Ignore buttons
	MB_YESNOCANCEL          = 0x00000003         // Yes, No, and Cancel buttons
	MB_YESNO                = 0x00000004         // Yes and No buttons
	MB_RETRYCANCEL          = 0x00000005         // Retry and Cancel buttons
	MB_CANCELTRYCONTINUE    = 0x00000006         // Cancel, Try Again, Continue buttons
	MB_ICONHAND             = 0x00000010         // Stop sign icon (Error)
	MB_ICONQUESTION         = 0x00000020         // Question mark icon (Query)
	MB_ICONEXCLAMATION      = 0x00000030         // Exclamation mark icon (Warning)
	MB_ICONASTERISK         = 0x00000040         // Asterisk icon (Information)
	MB_USERICON             = 0x00000080         // User-defined icon
	MB_ICONWARNING          = MB_ICONEXCLAMATION //Alias for the warning icon
	MB_ICONERROR            = MB_ICONHAND        //Alias for the error icon
	MB_ICONINFORMATION      = MB_ICONASTERISK    //Alias for the information icon
	MB_ICONSTOP             = MB_ICONHAND        //Alias for the stop icon
	MB_DEFBUTTON1           = 0x00000000         // First button is the default
	MB_DEFBUTTON2           = 0x00000100         // Second button is the default
	MB_DEFBUTTON3           = 0x00000200         // Third button is the default
	MB_DEFBUTTON4           = 0x00000300         // Fourth button is the default
	MB_APPLMODAL            = 0x00000000         // Application modal message box
	MB_SYSTEMMODAL          = 0x00001000         // System modal message box
	MB_TASKMODAL            = 0x00002000         // Task modal message box
	MB_HELP                 = 0x00004000         // Help button is included
	MB_SETFOREGROUND        = 0x00010000         // Bring the window to the foreground
	MB_DEFAULT_DESKTOP_ONLY = 0x00020000         // Only show on the default desktop
	MB_TOPMOST              = 0x00040000         // Message box is topmost
	MB_RIGHT                = 0x00080000         // Text is right-justified
	MB_RTLREADING           = 0x00100000         // Text is displayed with right-to-left reading order
	MB_SERVICE_NOTIFICATION = 0x00200000         // Service notification message box
)

// MessageBox button callback constants
const (
	ID_OK     = 1 // OK button identifier
	ID_CANCEL = 2 // Cancel button identifier
	ID_ABORT  = 3 // Abort button identifier
	ID_RETRY  = 4 // Retry button identifier
	ID_IGNORE = 5 // Ignore button identifier
	ID_YES    = 6 // Yes button identifier
	ID_NO     = 7 // No button identifier
)

// PopMsg function is used to display a message box on the Windows platform
// title: The title of the message box
// message: The message content of the message box
// The return value is of type error, indicating whether the function call was successful, nil for success, otherwise the reason for the error
func PopMsg(title, message string) (err error) {
	// Convert the title to UTF-16 encoding
	title16 := utf16.Encode([]rune(title))
	// Convert the message content to UTF-16 encoding
	message16 := utf16.Encode([]rune(message))
	// Call the Windows API function MessageBoxW to display the message box
	_, _, err = syscall.MustLoadDLL(user32API).MustFindProc("MessageBoxW").Call(
		0,                                      // Parent window handle, 0 indicates no parent window
		uintptr(unsafe.Pointer(&message16[0])), // Message string
		uintptr(unsafe.Pointer(&title16[0])),   // Title string
		MB_OK,                                  // Flags
	)
	// If the returned error is not nil and the error message is "The operation completed successfully.", set the error to nil
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return err
}

// PopErrMsg function is used to display an error message box on the Windows platform.
//
// Parameters:
// message string - The message content to be displayed in the error message box.
//
// Return value:
// err error - The error of the function call, nil for success, otherwise the reason for the error.
func PopErrMsg(message string) (err error) {
	// Convert the title "error" to UTF-16 encoding
	title16 := utf16.Encode([]rune("error"))
	// Convert the passed message content to UTF-16 encoding
	message16 := utf16.Encode([]rune(message))
	// Call the MessageBoxW function in user32.dll to display the message box
	_, _, err = syscall.MustLoadDLL(user32API).MustFindProc("MessageBoxW").Call(
		0,                                      // Parent window handle, 0 indicates no parent window
		uintptr(unsafe.Pointer(&message16[0])), // Message string
		uintptr(unsafe.Pointer(&title16[0])),   // Title string
		MB_OK|MB_ICONERROR,                     // Flags, indicating a message box with an OK button and error icon
	)
	// If err is not nil and the error message is "The operation completed successfully.", ignore the error
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return err
}

// PopWarningMsg displays a Windows message box.
//
// Parameters:
// message string - The content of the message to be displayed.
//
// Returns:
// err error - Returns an error if the function call fails, otherwise nil.
func PopWarningMsg(message string) (err error) {
	// Convert the title "error" to UTF-16 encoding
	title16 := utf16.Encode([]rune("warning"))
	// Convert the incoming message content to UTF-16 encoding
	message16 := utf16.Encode([]rune(message))
	// Call the MessageBoxW function in user32.dll to display the message box
	_, _, err = syscall.MustLoadDLL(user32API).MustFindProc("MessageBoxW").Call(
		0,                                      // Parent window handle, 0 indicates no parent window
		uintptr(unsafe.Pointer(&message16[0])), // Message string
		uintptr(unsafe.Pointer(&title16[0])),   // Title string
		MB_OK|MB_ICONWARNING,                   // Flags, indicating a message box with an OK button and error icon
	)
	// If err is not nil and the error message is "The operation completed successfully.", ignore the error
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return err

}

// PopCustomMsg function is used to display a custom message box on the Windows platform.
//
// Parameters:
// title: The title of the message box.
// message: The message content of the message box.
// format: The flag bits for the message box.
// hwnd: The parent window handle, 0 indicates no parent window.
//
// Return values:
// cb: The button number clicked by the user.
// err: Error information, nil if the operation is successful.
func PopCustomMsg(title, message string, format uintptr, hwnd uintptr) (cb uintptr, err error) {
	// Convert the title to UTF-16 encoding
	title16 := utf16.Encode([]rune(title))
	// Convert the message content to UTF-16 encoding
	message16 := utf16.Encode([]rune(message))
	// Call the Windows API function MessageBoxW to display the message box
	cb, _, err = syscall.MustLoadDLL(user32API).MustFindProc("MessageBoxW").Call(
		hwnd,                                   // Parent window handle, 0 indicates no parent window
		uintptr(unsafe.Pointer(&message16[0])), // Message string
		uintptr(unsafe.Pointer(&title16[0])),   // Title string
		format,                                 // Flags
	)
	// If the returned error is not nil and the error message is "The operation completed successfully.", set the error to nil
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return cb, err
}
