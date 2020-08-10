package main

import "C"
import (
	"fmt"
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboard"
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"log"
	"os"
	"os/signal"
	"time"
)

var logiKeyboard *LogiKeyboard.LogiKeyboard

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	logiKeyboard = LogiKeyboard.Create()

	defer logiKeyboard.Shutdown()

	logiKeyboard.Init()

	logiKeyboard.SetTargetDevice(LogiKeyboardTypes.LogiDeviceTypeAll)
	defaultLightning()

	var shortcuts = []*Shortcut{
		new(Shortcut).Create([]types.VKCode{types.VK_LCONTROL, types.VK_LSHIFT}, []LogiKeyboardTypes.Name{LogiKeyboardTypes.T, LogiKeyboardTypes.R}),
		new(Shortcut).Create([]types.VKCode{types.VK_LWIN}, []LogiKeyboardTypes.Name{LogiKeyboardTypes.ONE, LogiKeyboardTypes.TWO, LogiKeyboardTypes.THREE}),
		new(Shortcut).Create([]types.VKCode{types.VK_LCONTROL}, []LogiKeyboardTypes.Name{LogiKeyboardTypes.C, LogiKeyboardTypes.Z, LogiKeyboardTypes.Y}),
		new(Shortcut).Create([]types.VKCode{types.VK_LSHIFT}, []LogiKeyboardTypes.Name{LogiKeyboardTypes.F6}),
	}

	if err := run(shortcuts); err != nil {
		log.Fatal(err)
	}
}

func defaultLightning() {
	fmt.Print("defaultLighting")
	logiKeyboard.SetLightning(100, 100, 100)
}

func run(shortcuts []*Shortcut) error {
	// Buffer size is depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")

	keyMap := make(map[types.VKCode]bool)

	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case k := <-keyboardChan:
			fmt.Printf("Received %v %v\n", k.Message, k.VKCode)

			keyMap[k.VKCode] = k.Message == types.WM_KEYDOWN
			fmt.Printf("Value of ctrl %v\n", keyMap[types.VK_LCONTROL])

			for _, shortcut := range shortcuts {
				found := true
				for _, key := range shortcut.modifiers {
					if !keyMap[key] {
						found = false
						break
					}
				}

				if found {
					defaultLightning()
					for _, logiKey := range shortcut.keys {
						logiKeyboard.SetLightingForKeyWithKeyName(logiKey, 100, 0, 0)
					}
					break
				}
			}

			if k.Message == types.WM_KEYUP &&
				(k.VKCode == types.VK_LWIN ||
					k.VKCode == types.VK_RWIN ||
					k.VKCode == types.VK_RCONTROL ||
					k.VKCode == types.VK_LCONTROL ||
					k.VKCode == types.VK_LSHIFT ||
					k.VKCode == types.VK_RSHIFT ||
					k.VKCode == types.VK_RMENU ||
					k.VKCode == types.VK_LMENU) {
				defaultLightning()
			}
		}

		continue
	}
}
