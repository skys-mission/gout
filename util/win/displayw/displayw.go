package displayw

import (
	"errors"
	"syscall"
	"unsafe"
)

type DisplayInfo struct {
	Name        string
	Width       int
	Height      int
	RefreshRate int
	// BitDepth is information in the display adapter driver and is unrelated to the display settings in Windows.
	BitDepth    int
	ColorFormat string
	//ColorSpace  string
	IsPrimary bool
}

const (
	ENUM_CURRENT_SETTINGS = 0xFFFFFFFF
	CCHDEVICENAME         = 32
)

var (
	user32                   = syscall.NewLazyDLL("user32.dll")
	gdi32                    = syscall.NewLazyDLL("gdi32.dll")
	procEnumDisplayDevicesW  = user32.NewProc("EnumDisplayDevicesW")
	procEnumDisplaySettingsW = user32.NewProc("EnumDisplaySettingsW")
)

type DISPLAY_DEVICE struct {
	cb           uint32
	DeviceName   [CCHDEVICENAME]uint16
	DeviceString [128]uint16
	StateFlags   uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

type DEVMODE struct {
	dmDeviceName         [CCHDEVICENAME]uint16
	dmSpecVersion        uint16
	dmDriverVersion      uint16
	dmSize               uint16
	dmDriverExtra        uint16
	dmFields             uint32
	dmPositionX          int32
	dmPositionY          int32
	dmDisplayOrientation uint32
	dmDisplayFixedOutput uint32
	dmColor              uint16
	dmDuplex             uint16
	dmYResolution        uint16
	dmTTOption           uint16
	dmCollate            uint16
	dmFormName           [32]uint16
	dmLogPixels          uint16
	dmBitsPerPel         uint32
	dmPelsWidth          uint32
	dmPelsHeight         uint32
	dmDisplayFlags       uint32
	dmDisplayFrequency   uint32
	dmICMMethod          uint32
	dmICMIntent          uint32
	dmMediaType          uint32
	dmDitherType         uint32
	dmReserved1          uint32
	dmReserved2          uint32
	dmPanningWidth       uint32
	dmPanningHeight      uint32
}

//func getColorSpace() string {
//	return "sRGB"
//}
//

// enumDisplayDevices function enumerates the display devices of the specified device
// Parameter i represents the index of the device, starting from 0
// The return value DISPLAY_DEVICE contains the display information of the device, the bool value indicates whether the enumeration is successful
func enumDisplayDevices(i uint32) (DISPLAY_DEVICE, bool) {
	// Define a DISPLAY_DEVICE structure variable dd
	var dd DISPLAY_DEVICE
	// Set the cb field of dd to the size of dd
	dd.cb = uint32(unsafe.Sizeof(dd))
	// Call the procEnumDisplayDevicesW function, passing in four parameters
	// The first parameter is a nil pointer, indicating enumeration of all display devices
	// The second parameter is i, representing the index of the device
	// The third parameter is the address of dd, used to store the returned device information
	// The fourth parameter is 0, indicating the use of default flags
	ret, _, _ := procEnumDisplayDevicesW.Call(
		uintptr(unsafe.Pointer(nil)),
		uintptr(i),
		uintptr(unsafe.Pointer(&dd)),
		0,
	)
	// Return dd and whether ret is not equal to 0, used to determine if the enumeration is successful
	return dd, ret != 0
}

// enumDisplaySettings function enumerates the display modes of the specified display device
// Parameter deviceName is a pointer to the device name
// Parameter modeNum is the number of the display mode to be enumerated, starting from 0
// The return value is the enumerated display mode information and whether the enumeration is successful
func enumDisplaySettings(deviceName *uint16, modeNum uint32) (DEVMODE, bool) {
	// Define a DEVMODE type variable dm
	var dm DEVMODE
	// Set the dmSize field of dm to the size of dm
	dm.dmSize = uint16(unsafe.Sizeof(dm))
	// Call the procEnumDisplaySettingsW function, passing in three parameters
	// The first parameter is a pointer to the device name
	// The second parameter is the number of the display mode to be enumerated
	// The third parameter is the address of dm
	ret, _, _ := procEnumDisplaySettingsW.Call(
		uintptr(unsafe.Pointer(deviceName)),
		uintptr(modeNum),
		uintptr(unsafe.Pointer(&dm)),
	)
	// Return the enumerated display mode information and the result of whether the enumeration is successful
	return dm, ret != 0
}

// getColorFormat function gets the color format string based on the bit depth
// Parameters:
//
//	bitDepth uint32 - The bit depth, 32 indicates RGBA format, 24 indicates RGB format, other values indicate unknown format
//
// Return value:
//
//	string - The color format string, which is "RGBA", "RGB", or "Unknown"
func getColorFormat(bitDepth uint32) string {
	switch bitDepth {
	case 32:
		// If the bit depth is 32, return "RGBA"
		return "RGBA"
	case 24:
		// If the bit depth is 24, return "RGB"
		return "RGB"
	default:
		// In other cases, return "Unknown"
		return "Unknown"
	}
}

// GetAllDisplays function is used to retrieve information about all monitors.
// Return values:
//
//	[]*DisplayInfo: A slice containing information about all displays.
//	error: If no monitors are found, an error message is returned.
func GetAllDisplays() ([]*DisplayInfo, error) {
	var displays []*DisplayInfo
	for i := uint32(0); ; i++ {
		// Call enumDisplayDevices function to get the display device information.
		dd, ok := enumDisplayDevices(i)
		if !ok {
			break
		}
		// If the state flag of the monitor is not STATE_ATTACHED, skip it.
		if dd.StateFlags&1 == 0 {
			continue
		}
		// Call enumDisplaySettings function to get the current settings of the monitor.
		dm, ok := enumDisplaySettings(&dd.DeviceName[0], ENUM_CURRENT_SETTINGS)
		if !ok {
			continue
		}
		// Add the monitor information to the slice.
		displays = append(displays, &DisplayInfo{
			// Monitor name
			Name: syscall.UTF16ToString(dd.DeviceName[:]),
			// Monitor width
			Width: int(dm.dmPelsWidth),
			// Monitor height
			Height: int(dm.dmPelsHeight),
			// Monitor refresh rate
			RefreshRate: int(dm.dmDisplayFrequency),
			// Monitor bit depth
			BitDepth: int(dm.dmBitsPerPel),
			// Monitor color format
			ColorFormat: getColorFormat(dm.dmBitsPerPel),
			// // Monitor color space (this line is commented out)
			// ColorSpace:  getColorSpace(),
			// Whether it is the primary monitor
			IsPrimary: dd.StateFlags&4 != 0,
		})
	}
	// If no monitors are found, return an error message.
	if len(displays) == 0 {
		return nil, errors.New("no display found")
	}
	// Return the slice of monitor information and nil error.
	return displays, nil
}

// GetPrimaryDisplay function is used to retrieve information about the primary monitor.
// Return values:
//
//	*DisplayInfo: A pointer to a structure containing information about the primary monitor.
//	error: If the primary monitor is not found, an error message is returned.
func GetPrimaryDisplay() (*DisplayInfo, error) {
	// Get information about all monitors.
	displays, err := GetAllDisplays()
	if err != nil {
		// If the retrieval fails, return an error message.
		return nil, err
	}
	// Iterate through all monitor information.
	for _, display := range displays {
		// If the primary monitor is found
		if display.IsPrimary {
			// Return the information of the primary monitor.
			return display, nil
		}
	}
	// If the primary monitor is not found, return an error message.
	return nil, errors.New("no primary display found")
}
