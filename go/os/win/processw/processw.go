// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package processw

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/skys-mission/gout/common"
)

// 定义进程访问权限常量
const (
	process_ALL_ACCESS = 0x1F0FFF
	TH32CS_SNAPPROCESS = 0x00000002

	IDLE_PRIORITY_CLASS         = 0x40
	BELOW_NORMAL_PRIORITY_CLASS = 0x4000
	NORMAL_PRIORITY_CLASS       = 0x20
	ABOVE_NORMAL_PRIORITY_CLASS = 0x8000
	HIGH_PRIORITY_CLASS         = 0x80
	REALTIME_PRIORITY_CLASS     = 0x100
)

// 定义kernel32.dll相关变量和函数指针
var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procSetPriorityClass         = kernel32.NewProc("SetPriorityClass")
	procGetPriorityClass         = kernel32.NewProc("GetPriorityClass")
	procSetProcessAffinityMask   = kernel32.NewProc("SetProcessAffinityMask")
	procGetProcessAffinityMask   = kernel32.NewProc("GetProcessAffinityMask")
	procCreateToolHelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = kernel32.NewProc("Process32FirstW")
	procProcess32Next            = kernel32.NewProc("Process32NextW")
)

// DWORD 和 Handle 类型定义
type DWORD uint32
type Handle uintptr

// LUID 结构体定义
type LUID struct {
	LowPart  DWORD
	HighPart int32
}

// ProcessInfo 结构体定义，包含进程信息和可执行文件名
type ProcessInfo struct {
	processEntry32
	ExeFileName string
}

// processEntry32 结构体定义，包含详细的进程信息
type processEntry32 struct {
	Size              DWORD
	CntUsage          DWORD
	ProcessID         DWORD
	DefaultHeapID     uintptr
	ModuleID          DWORD
	CntThreads        DWORD
	ParentProcessID   DWORD
	PriorityClassBase DWORD
	Flags             DWORD
	ExeFile           [260]uint16
}

// GetProcessHandle retrieves the process handle for the specified process ID.
//
// pid: Process ID
//
// Returns:
// handle: Process handle
// cf: Function to close the handle, call to release handle resources
// err: Error information, returns a non-nil error if opening the process fails
func GetProcessHandle(pid DWORD) (handle Handle, cf func() error, err error) {
	handleS, err := syscall.OpenProcess(process_ALL_ACCESS, false, uint32(pid))
	if err != nil {
		return 0, nil, err
	}

	return Handle(handleS), func() error {
		return syscall.CloseHandle(syscall.Handle(handle))
	}, nil
}

// SetPriorityClass 设置进程的优先级类
func SetPriorityClass(handle Handle, priorityClass DWORD) error {
	ret, _, err := procSetPriorityClass.Call(uintptr(handle), uintptr(priorityClass))
	if ret == 0 {
		if err != nil {
			return err
		}
		return errors.New("ret == 0")
	}
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return err
	}
	return nil
}

// GetPriorityClass 获取进程的优先级类
func GetPriorityClass(handle Handle) (DWORD, error) {
	ret, _, err := procGetPriorityClass.Call(uintptr(handle))
	if ret == 0 {
		if err != nil {
			return 0, err
		}
		return 0, errors.New("ret == 0")
	}
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return 0, err
	}
	return DWORD(ret), nil
}

// SetProcessAffinityMask 设置进程的亲和性掩码
func SetProcessAffinityMask(handle Handle, processAffinityMask DWORD) error {
	ret, _, err := procSetProcessAffinityMask.Call(uintptr(handle), uintptr(processAffinityMask))
	if ret == 0 {
		if err != nil {
			return err
		}
		return errors.New("ret == 0")
	}
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return err
	}
	return nil
}

// GetProcessAffinityMask 获取进程的亲和性掩码和系统的亲和性掩码
func GetProcessAffinityMask(handle Handle) (processAffinityMask DWORD, systemAffinityMask DWORD, err error) {
	ret, _, err := procGetProcessAffinityMask.Call(uintptr(handle), uintptr(unsafe.Pointer(&processAffinityMask)), uintptr(unsafe.Pointer(&systemAffinityMask)))

	if ret == 0 {
		if err != nil {
			return 0, 0, err
		}
		return 0, 0, errors.New("ret == 0")
	}
	return processAffinityMask, systemAffinityMask, nil
}

// GetProcessList 获取系统中的进程列表
func GetProcessList() ([]*ProcessInfo, error) {
	handle, _, err := procCreateToolHelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	if handle == uintptr(syscall.InvalidHandle) {
		return nil, err
	}
	defer syscall.CloseHandle(syscall.Handle(handle))
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return nil, err
	}

	var processEntryList []processEntry32
	var entry processEntry32
	entry.Size = DWORD(unsafe.Sizeof(entry))
	ret, _, err := procProcess32First.Call(handle, uintptr(unsafe.Pointer(&entry)))
	if ret == 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("ret == 0")
	}
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return nil, err
	}
	for {
		processEntryList = append(processEntryList, entry)
		ret, _, err = procProcess32Next.Call(handle, uintptr(unsafe.Pointer(&entry)))
		if ret == 0 {
			break
		}
		if err != nil && err.Error() != common.WindowsCallInvalidError {
			return nil, err
		}
	}

	var processList = make([]*ProcessInfo, len(processEntryList))
	for i, _ := range processEntryList {
		processList[i] = &ProcessInfo{
			processEntry32: processEntryList[i],
			ExeFileName:    syscall.UTF16ToString(processEntryList[i].ExeFile[:]),
		}
	}

	return processList, nil
}
