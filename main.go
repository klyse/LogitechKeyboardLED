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
		new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LSHIFT}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F6, 0, 100, 0),
		}),
	}

	if err := run(shortcuts); err != nil {
		log.Fatal(err)
	}
}

func defaultLightning() {
	fmt.Println("defaultLighting")
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

			found := false
			for _, shortcut := range shortcuts {
				found = true
				for _, key := range shortcut.modifiers {
					if !keyMap[key] {
						found = false
						break
					}
				}

				if found {
					defaultLightning()
					for _, logiKey := range shortcut.keys {
						logiKeyboard.SetLightingForKeyWithKeyName(logiKey.key, logiKey.red, logiKey.green, logiKey.blue)
					}
					break
				}
			}

			if !found {
				defaultLightning()
			}
		}

		continue
	}
}
