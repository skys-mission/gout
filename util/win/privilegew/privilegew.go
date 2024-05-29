package privilegew

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/skys-mission/gout/common"
)

type handle uintptr
type dword uint32

const (
	se_DEBUG_NAME = "SeDebugPrivilege"

	se_PRIVILEGE_ENABLED = 0x00000002
)

var (
	advapi32 = syscall.NewLazyDLL("advapi32.dll")

	procOpenProcessToken      = advapi32.NewProc("OpenProcessToken")
	procLookupPrivilegeValue  = advapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges = advapi32.NewProc("AdjustTokenPrivileges")
)

type luid struct {
	LowPart  dword
	HighPart int32
}

type token_PRIVILEGES struct {
	PrivilegeCount dword
	Privileges     [1]luid_AND_ATTRIBUTES
}

type luid_AND_ATTRIBUTES struct {
	Luid       luid
	Attributes dword
}

// EnableDebugPrivilege Enables debugging privileges for the current process.
// To execute this code, administrator privileges are required.
func EnableDebugPrivilege() error {
	var hToken handle
	// Get a handle to the current process
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return err
	}

	// Call OpenProcessToken to open the access token of the current process
	ret, _, err := procOpenProcessToken.Call(uintptr(handle), syscall.TOKEN_ADJUST_PRIVILEGES|syscall.TOKEN_QUERY, uintptr(unsafe.Pointer(&hToken)))
	if ret == 0 {
		return fmt.Errorf("OpenProcessToken: %w", err)
	}
	// Close the handle in a deferred manner
	defer syscall.CloseHandle(syscall.Handle(hToken))

	var l luid
	// Convert the string to a UTF-16 encoded pointer
	debugNameUINT16, err := syscall.UTF16PtrFromString(se_DEBUG_NAME)
	if err != nil {
		return err
	}
	// Call LookupPrivilegeValue to find the privilege value
	ret, _, err = procLookupPrivilegeValue.Call(0, uintptr(unsafe.Pointer(debugNameUINT16)), uintptr(unsafe.Pointer(&l)))
	if ret == 0 {
		return fmt.Errorf("LookupPrivilegeValue: %w", err)
	}

	// Set the token_PRIVILEGES structure
	tp := token_PRIVILEGES{
		PrivilegeCount: 1,
		Privileges: [1]luid_AND_ATTRIBUTES{
			{
				// Privilege identifier
				Luid: l,
				// Set the privilege to the enabled state
				Attributes: se_PRIVILEGE_ENABLED,
			},
		},
	}

	// Call AdjustTokenPrivileges to adjust the privileges of the access token
	ret, _, err = procAdjustTokenPrivileges.Call(uintptr(hToken), 0, uintptr(unsafe.Pointer(&tp)), 0, 0, 0)
	if ret == 0 {
		return fmt.Errorf("AdjustTokenPrivileges: %w", err)
	}
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return err
}
