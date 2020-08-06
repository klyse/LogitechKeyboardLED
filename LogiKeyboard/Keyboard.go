package LogiKeyboard

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"syscall"
)

type LogiKeyboard struct {
	dll                             *syscall.LazyDLL
	ledInit                         *syscall.LazyProc
	ledSetTargetDevice              *syscall.LazyProc
	ledSetLighting                  *syscall.LazyProc
	ledSetLightingForKeyWithKeyName *syscall.LazyProc
	ledSetLightingForTargetZone     *syscall.LazyProc

	ledShutdown *syscall.LazyProc
}

func Create() *LogiKeyboard {
	p := new(LogiKeyboard)

	//https://github.com/golang/go/wiki/WindowsDLLs
	p.dll = syscall.NewLazyDLL("LogitechLedEnginesWrapper.dll")
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
	_, _, _ = v.ledSetLightingForKeyWithKeyName.Call(uintptr(deviceType), uintptr(zoneId), uintptr(red), uintptr(green), uintptr(blue))
}

func (v LogiKeyboard) Shutdown() {
	_, _, _ = v.ledShutdown.Call()
}
