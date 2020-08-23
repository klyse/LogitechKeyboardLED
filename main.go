package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
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
	defaultLightning(false)

	var shortcuts = []Shortcut{
		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LSHIFT}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F6, 100, 0, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F9, 0, 0, 100),
		}),

		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LCONTROL}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F2, 100, 0, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.C, 100, 0, 100),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.V, 100, 0, 100),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.X, 100, 0, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.Y, 100, 0, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.Z, 100, 0, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F, 0, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.R, 100, 100, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.B, 0, 100, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.NUM_SLASH, 100, 0, 100),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.HOME, 0, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.END, 0, 100, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.W, 50, 0, 100),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F4, 100, 100, 0),
		}),

		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LCONTROL, types.VK_LSHIFT, types.VK_LMENU}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.T, 50, 0, 50),
		}),

		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LCONTROL, types.VK_LMENU}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.B, 0, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.L, 0, 100, 0),
		}),

		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LCONTROL, types.VK_LSHIFT}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.Z, 0, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.B, 0, 100, 0),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F, 0, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.R, 100, 100, 0),
		}),

		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_RMENU, types.VK_LCONTROL}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.SEVEN, 50, 50, 100),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.ZERO, 50, 50, 100),

			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.EIGHT, 50, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.NINE, 50, 100, 0),

			*new(ShortcutKey).CreateColor(0x66, 50, 100, 0),
		}),

		//*new(Shortcut).CreateColor([]types.VKCode{types.VK_LWIN}, []LogiKeyboardTypes.Name{LogiKeyboardTypes.TAB, LogiKeyboardTypes.ONE, LogiKeyboardTypes.TWO, LogiKeyboardTypes.THREE, LogiKeyboardTypes.FOUR, LogiKeyboardTypes.FIVE, LogiKeyboardTypes.SIX, LogiKeyboardTypes.SEVEN, LogiKeyboardTypes.EIGHT, LogiKeyboardTypes.NINE, LogiKeyboardTypes.ZERO}, 0, 100, 0),

		*new(Shortcut).CreateWithKey([]types.VKCode{types.VK_LMENU}, []ShortcutKey{
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.F4, 50, 50, 100),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.ONE, 0, 100, 0),
			*new(ShortcutKey).CreateColor(LogiKeyboardTypes.FIVE, 0, 0, 100),
		}),
	}

	if err := run(shortcuts); err != nil {
		log.Fatal(err)
	}
}

func defaultLightning(preHotkey bool) {
	fmt.Printf("defaultLighting %v\n", preHotkey)
	logiKeyboard.SetLightning(100, 100, 100)

	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.LEFT_CONTROL, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.LEFT_SHIFT, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.LEFT_ALT, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.RIGHT_ALT, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.RIGHT_CONTROL, 100, 0, 0)

	if !preHotkey {
		logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.F9, 0, 0, 100)
		logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.ESC, 0, 100, 0)
	}
}

func run(shortcuts []Shortcut) error {
	// Buffer size is depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")

	currentlyPressedKeys := make(map[types.VKCode]bool)

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

			currentlyPressedKeys[k.VKCode] = k.Message == types.WM_KEYDOWN || k.Message == types.WM_SYSKEYDOWN

			shortCut := linq.From(shortcuts).Where(func(c interface{}) bool {
				found := linq.From(c.(Shortcut).Modifiers).All(func(y interface{}) bool {
					return currentlyPressedKeys[y.(types.VKCode)]
				})

				if !found {
					return false
				}

				found = linq.From(currentlyPressedKeys).Where(func(y interface{}) bool {
					return y.(linq.KeyValue).Value.(bool)
				}).AnyWith(func(y interface{}) bool {
					return !linq.From(c.(Shortcut).Modifiers).Contains(y.(linq.KeyValue).Key.(types.VKCode))
				})

				return !found
			}).First()

			if shortCut != nil {
				defaultLightning(true)
				for _, logiKey := range shortCut.(Shortcut).Keys {
					logiKeyboard.SetLightingForKeyWithKeyName(logiKey.Key, logiKey.Red, logiKey.Green, logiKey.Blue)
				}
			} else {
				defaultLightning(false)
			}
		}

		continue
	}
}
