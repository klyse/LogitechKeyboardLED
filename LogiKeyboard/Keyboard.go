package LogiKeyboard

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"golang.org/x/sys/windows"
)

type LogiKeyboard struct {
	dll                             *windows.LazyDLL
	ledInit                         *windows.LazyProc
	ledSetTargetDevice              *windows.LazyProc
	ledSetLighting                  *windows.LazyProc
	ledSetLightingForKeyWithKeyName *windows.LazyProc
	ledSetLightingForTargetZone     *windows.LazyProc

	ledShutdown *windows.LazyProc
}

func Create() *LogiKeyboard {
	p := new(LogiKeyboard)

	//https://github.com/golang/go/wiki/WindowsDLLs
	p.dll = windows.NewLazyDLL("LogitechLedEnginesWrapper.dll")
	p.ledInit = p.dll.NewProc("LogiLedInit")
	p.ledSetTargetDevice = p.dll.NewProc("LogiLedSetTargetDevice")
	p.ledSetLighting = p.dll.NewProc("LogiLedSetLighting")
	p.ledSetLightingForKeyWithKeyName = p.dll.NewProc("LogiLedSetLightingForKeyWithKeyName")
	p.ledSetLightingForTargetZone = p.dll.NewProc("LogiLedSetLightingForTargetZone")

	p.ledShutdown = p.dll.NewProc("LogiLedShutdown")

	return p
}

func (v LogiKeyboard) Init() {
	_, _, _ = v.ledInit.Call()
}

func (v LogiKeyboard) SetTargetDevice(device int) {
	_, _, _ = v.ledSetTargetDevice.Call(uintptr(device))
}

func (v LogiKeyboard) SetLightning(red int, green int, blue int) {
	_, _, _ = v.ledSetLighting.Call(uintptr(red), uintptr(green), uintptr(blue))
}

func (v LogiKeyboard) SetLightingForKeyWithKeyName(name LogiKeyboardTypes.Name, red int, green int, blue int) {
	_, _, _ = v.ledSetLightingForKeyWithKeyName.Call(uintptr(name), uintptr(red), uintptr(green), uintptr(blue))
}

func (v LogiKeyboard) SetLightingForTargetZone(deviceType LogiKeyboardTypes.DeviceType, zoneId int, red int, green int, blue int) {
	_, _, _ = v.ledSetLightingForTargetZone.Call(uintptr(deviceType), uintptr(zoneId), uintptr(red), uintptr(green), uintptr(blue))
}

func (v LogiKeyboard) Shutdown() {
	_, _, _ = v.ledShutdown.Call()
}
