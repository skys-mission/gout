package processw

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/skys-mission/gout/common"
)

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

type DWORD uint32
type Handle uintptr

type LUID struct {
	LowPart  DWORD
	HighPart int32
}

//type TOKEN_PRIVILEGES struct {
//	PrivilegeCount DWORD
//	Privileges     [1]LUID_AND_ATTRIBUTES
//}
//type LUID_AND_ATTRIBUTES struct {
//	Luid       LUID
//	Attributes DWORD
//}

type ProcessInfo struct {
	processEntry32
	ExeFileName string
}

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

// SetPriorityClass sets the priority class for the specified process handle.
//
// Parameters:
// handle: HANDLE type, the handle of the specified process
// priorityClass: DWORD type, the priority class of the specified process
//
// Returns:
// error type, returns nil if the setting is successful, otherwise returns error information
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

// GetPriorityClass retrieves the priority class for the specified process handle.
//
// Parameters:
// handle: HANDLE type, the handle of the specified process
//
// Returns:
// DWORD type, representing the priority class of the process
// error type, returns nil if the retrieval is successful, otherwise returns error information
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

// SetProcessAffinityMask sets the affinity mask for the specified process handle.
//
// Parameters:
// handle: HANDLE type, the handle of the specified process
// processAffinityMask: DWORD type, the affinity mask for the specified process
//
// Returns:
// error type, returns nil if the setting is successful, otherwise returns error information
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

// GetProcessAffinityMask retrieves the process affinity mask and system affinity mask for the specified process handle.
//
// Parameters:
// handle: HANDLE type, the handle of the specified process
//
// Returns:
// DWORD type, representing the process affinity mask
// DWORD type, representing the system affinity mask
// error type, returns nil if the retrieval is successful, otherwise returns error information
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

// GetProcessList retrieves a list of all processes in the system.
//
// Returns:
// []*ProcessInfo - A slice of *ProcessInfo structures containing information for all processes
// error - If an error occurs during the function's execution, the corresponding error information is returned; otherwise, nil is returned
func GetProcessList() ([]*ProcessInfo, error) {
	// Call procCreateToolHelp32Snapshot to create a process snapshot and obtain a handle
	handle, _, err := procCreateToolHelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	if handle == uintptr(syscall.InvalidHandle) {
		return nil, err
	}
	defer syscall.CloseHandle(syscall.Handle(handle))

	// Check the returned error, if it's not an invalid handle error, return the error
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return nil, err
	}

	var processEntryList []processEntry32
	var entry processEntry32
	entry.Size = DWORD(unsafe.Sizeof(entry))

	// Call procProcess32First to get the first process information
	ret, _, err := procProcess32First.Call(handle, uintptr(unsafe.Pointer(&entry)))
	if ret == 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("ret == 0")
	}

	// Check the returned error, if it's not an invalid handle error, return the error
	if err != nil && err.Error() != common.WindowsCallInvalidError {
		return nil, err
	}

	// Loop through all process information
	for {
		processEntryList = append(processEntryList, entry)

		// Call procProcess32Next to get the next process information
		ret, _, err = procProcess32Next.Call(handle, uintptr(unsafe.Pointer(&entry)))
		if ret == 0 {
			break
		}

		// Check the returned error, if it's not an invalid handle error, return the error
		if err != nil && err.Error() != common.WindowsCallInvalidError {
			return nil, err
		}
	}

	// Convert process information to a slice of ProcessInfo structures
	var processList = make([]*ProcessInfo, len(processEntryList))
	for i, _ := range processEntryList {
		processList[i] = &ProcessInfo{
			processEntry32: processEntryList[i],
			// Convert the process name from UTF-16 to UTF-8 encoding
			ExeFileName: syscall.UTF16ToString(processEntryList[i].ExeFile[:]),
		}
	}

	return processList, nil
}
