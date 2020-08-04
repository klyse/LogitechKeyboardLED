package Keyboard

import "syscall"

type Keyboard struct {
	dll                             *syscall.LazyDLL
	ledInit                         *syscall.LazyProc
	ledSetTargetDevice              *syscall.LazyProc
	ledSetLighting                  *syscall.LazyProc
	ledSetLightingForKeyWithKeyName *syscall.LazyProc
	ledShutdown                     *syscall.LazyProc
}

func Create() *Keyboard {
	p := new(Keyboard)

	//https://github.com/golang/go/wiki/WindowsDLLs
	p.dll = syscall.NewLazyDLL("LogitechLedEnginesWrapper.dll")
	p.ledInit = p.dll.NewProc("LogiLedInit")
	p.ledSetTargetDevice = p.dll.NewProc("LogiLedSetTargetDevice")
	p.ledSetLighting = p.dll.NewProc("LogiLedSetLighting")

	p.ledShutdown = p.dll.NewProc("LogiLedShutdown")

	return p
}

func (v Keyboard) Init() {
	_, _, _ = v.ledInit.Call()
}

func (v Keyboard) SetTargetDevice(device int) {
	_, _, _ = v.ledSetTargetDevice.Call(uintptr(device))
}

func (v Keyboard) SetLightning(red int, green int, blue int) {
	_, _, _ = v.ledSetLighting.Call(uintptr(red), uintptr(green), uintptr(blue))
}

func (v Keyboard) Shutdown() {
	_, _, _ = v.ledShutdown.Call()
}
