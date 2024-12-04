// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package privilegew

// 设置中文注释

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/skys-mission/gout/common"
)

// handle 是 Windows API 中进程句柄的别名。
type handle uintptr

// dword 是 Windows API 中常用的无符号32位整型的别名。
type dword uint32

// seDebugName 定义了调试权限的名称。
const (
	seDebugName = "SeDebugPrivilege"
	// sePrivilegeEnabled 表示启用权限的标志。
	sePrivilegeEnabled = 0x00000002
)

// advapi32 是用于惰性加载的 advapi32.dll 库。
var (
	advapi32 = syscall.NewLazyDLL("advapi32.dll")

	// procOpenProcessToken 是 OpenProcessToken 函数的句柄。
	procOpenProcessToken = advapi32.NewProc("OpenProcessToken")
	// procLookupPrivilegeValue 是 LookupPrivilegeValueW 函数的句柄。
	procLookupPrivilegeValue = advapi32.NewProc("LookupPrivilegeValueW")
	// procAdjustTokenPrivileges 是 AdjustTokenPrivileges 函数的句柄。
	procAdjustTokenPrivileges = advapi32.NewProc("AdjustTokenPrivileges")
)

// luID 代表本地唯一标识符（LUID）。
type luID struct {
	LowPart  dword
	HighPart int32
}

// tokenPrivileges 用于描述要调整的令牌权限。
type tokenPrivileges struct {
	PrivilegeCount dword
	Privileges     [1]luIDAndAttributes
}

// luIDAndAttributes 结合了 LUID 和其相应的属性。
type luIDAndAttributes struct {
	LuID       luID
	Attributes dword
}

// EnableDebugPrivilege 启用当前进程的调试权限。
// 这个函数通过调用 Windows API 来提升进程的权限，以便进行调试操作。
func EnableDebugPrivilege() error {
	var hToken handle
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return err
	}

	ret, _, err := procOpenProcessToken.Call(uintptr(handle), syscall.TOKEN_ADJUST_PRIVILEGES|syscall.TOKEN_QUERY, uintptr(unsafe.Pointer(&hToken)))
	if ret == 0 {
		return fmt.Errorf("OpenProcessToken: %w", err)
	}
	defer syscall.CloseHandle(syscall.Handle(hToken))

	var l luID
	debugNameUINT16, err := syscall.UTF16PtrFromString(seDebugName)
	if err != nil {
		return err
	}
	ret, _, err = procLookupPrivilegeValue.Call(0, uintptr(unsafe.Pointer(debugNameUINT16)), uintptr(unsafe.Pointer(&l)))
	if ret == 0 {
		return fmt.Errorf("LookupPrivilegeValue: %w", err)
	}

	tp := tokenPrivileges{
		PrivilegeCount: 1,
		Privileges: [1]luIDAndAttributes{
			{
				LuID:       l,
				Attributes: sePrivilegeEnabled,
			},
		},
	}

	ret, _, err = procAdjustTokenPrivileges.Call(uintptr(hToken), 0, uintptr(unsafe.Pointer(&tp)), 0, 0, 0)
	if ret == 0 {
		return fmt.Errorf("AdjustTokenPrivileges: %w", err)
	}
	if err != nil && err.Error() == common.WindowsCallInvalidError {
		err = nil
	}
	return err
}
