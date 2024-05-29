package main

import (
	"fmt"

	"github.com/skys-mission/gout/util/win/privilegew"
	"github.com/skys-mission/gout/util/win/processw"
)

func main() {
	err := privilegew.EnableDebugPrivilege()
	if err != nil {
		panic(err)
	}
	list, err := processw.GetProcessList()
	if err != nil {
		panic(err)
	}
	//a := processw.ProcessInfo{}

	var audiodgPID processw.DWORD
	for _, p := range list {
		if p.ExeFileName == "audiodg.exe" {
			audiodgPID = p.ProcessID
			break
		}
	}
	if audiodgPID == 0 {
		panic("process not found")
	}
	handle, cf, err := processw.GetProcessHandle(audiodgPID)
	if err != nil {
		panic(err)
	}
	defer cf()

	priorityClass, err := processw.GetPriorityClass(handle)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current audiodg.exe priority class: 0x%X\n", priorityClass)

	err = processw.SetPriorityClass(handle, processw.HIGH_PRIORITY_CLASS)
	if err != nil {
		panic(err)
	}

	// Get current affinity mask
	processAffinityMask, systemAffinityMask, err := processw.GetProcessAffinityMask(handle)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current audiodg.exe affinity mask: 0x%X, system affinity mask: 0x%X\n", processAffinityMask, systemAffinityMask)
	// Set new affinity mask (e.g., use only the first CPU)
	err = processw.SetProcessAffinityMask(handle, 1<<22)
	if err != nil {
		panic(err)
	}
	fmt.Println("audiodg.exe affinity mask set to use only the first CPU")
}
