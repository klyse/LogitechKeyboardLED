package LogiKeyboard

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"golang.org/x/sys/windows"
	"log"
)

type LogiKeyboard struct {
	dll                             *windows.LazyDLL
	ledInit                         *windows.LazyProc
	ledSetTargetDevice              *windows.LazyProc
	ledSetLighting                  *windows.LazyProc
	ledSetLightingForKeyWithKeyName *windows.LazyProc
	ledSetLightingForTargetZone     *windows.LazyProc
	ledFlashSingleKey               *windows.LazyProc
	ledPulseSingleKey               *windows.LazyProc
	ledStopEffects                  *windows.LazyProc

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
	p.ledFlashSingleKey = p.dll.NewProc("LogiLedFlashSingleKey")
	p.ledPulseSingleKey = p.dll.NewProc("LogiLedPulseSingleKey")
	p.ledStopEffects = p.dll.NewProc("LogiLedStopEffects")

	p.ledShutdown = p.dll.NewProc("LogiLedShutdown")

	return p
}

func (v LogiKeyboard) Init() {
	ret, _, _ := v.ledInit.Call()

	if ret != 1 {
		log.Fatal("Already initialized")
	}
}

func (v LogiKeyboard) SetTargetDevice(device int) {
	_, _, _ = v.ledSetTargetDevice.Call(uintptr(device))
}

func (v LogiKeyboard) StopEffects() {
	_, _, _ = v.ledStopEffects.Call()
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

func (v LogiKeyboard) SetFlashSingleKey(name LogiKeyboardTypes.Name, red, green, blue, msDuration, msInterval int) {
	_, _, _ = v.ledFlashSingleKey.Call(uintptr(name), uintptr(red), uintptr(green), uintptr(blue), uintptr(msDuration), uintptr(msInterval))
}

func (v LogiKeyboard) SetPulseSingleKey(name LogiKeyboardTypes.Name, red, green, blue, finishRedPercentage, finishGreenPercentage, finishBluePercentage, msDuration, isInfinite int) {
	_, _, _ = v.ledPulseSingleKey.Call(uintptr(name), uintptr(red), uintptr(green), uintptr(blue), uintptr(finishRedPercentage), uintptr(finishGreenPercentage), uintptr(finishBluePercentage), uintptr(msDuration), uintptr(isInfinite))
}

func (v LogiKeyboard) Shutdown() {
	_, _, _ = v.ledShutdown.Call()
}
