package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboard"
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/klyse/LogitechKeyboardLED/Shortcuts"
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

	logiKeyboard.SetTargetDevice(LogiKeyboardTypes.LogiDeviceTypePerKeyRgb)
	defaultLightning(false)

	var shortcuts = []Shortcuts.Shortcut{
		Shortcuts.CreateWithKey([]types.VKCode{types.VK_LSHIFT}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F6, 100, 0, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F9, 0, 0, 100),
		}),

		Shortcuts.CreateWithKey([]types.VKCode{types.VK_LCONTROL}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F2, 100, 0, 0),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.C, 100, 0, 100),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.V, 100, 0, 100),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.X, 100, 0, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.Y, 100, 0, 0),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.Z, 100, 0, 0),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F, 0, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.R, 100, 100, 0),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.B, 0, 100, 0),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.NUM_SLASH, 100, 0, 100),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.HOME, 0, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.END, 0, 100, 0),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.W, 50, 0, 100),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F4, 100, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F8, 0, 0, 100),
		}),

		Shortcuts.CreateWithKey([]types.VKCode{types.VK_LCONTROL, types.VK_LSHIFT, types.VK_LMENU}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.T, 50, 0, 50),
		}),

		Shortcuts.CreateWithKey([]types.VKCode{types.VK_LCONTROL, types.VK_LMENU}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.B, 0, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.L, 0, 100, 0),
		}),

		Shortcuts.CreateWithKey([]types.VKCode{types.VK_LCONTROL, types.VK_LSHIFT}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.Z, 0, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.B, 0, 100, 0),

			Shortcuts.CreateKeyColorEffect(LogiKeyboardTypes.F, 0, 100, 0, Shortcuts.Blinking),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.R, 100, 100, 0),
		}),

		Shortcuts.CreateWithKey([]types.VKCode{types.VK_RMENU, types.VK_LCONTROL}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.SEVEN, 50, 50, 100),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.ZERO, 50, 50, 100),

			Shortcuts.CreateKeyColor(LogiKeyboardTypes.EIGHT, 50, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.NINE, 50, 100, 0),

			// todo: check what this is
			Shortcuts.CreateKeyColor(0x66, 50, 100, 0),
		}),

		Shortcuts.CreateWithKey([]types.VKCode{types.VK_LMENU}, []Shortcuts.ShortcutKey{
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.F4, 50, 50, 100),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.ONE, 0, 100, 0),
			Shortcuts.CreateKeyColor(LogiKeyboardTypes.FIVE, 0, 0, 100),
		}),
	}

	if err := run(shortcuts); err != nil {
		log.Fatal(err)
	}
}

func defaultLightning(preHotkey bool) {
	fmt.Printf("defaultLighting %v\n", preHotkey)
	logiKeyboard.StopEffects()
	logiKeyboard.SetLightning(100, 100, 100)

	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.LEFT_CONTROL, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.LEFT_SHIFT, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.LEFT_ALT, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.RIGHT_ALT, 100, 0, 0)
	logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.RIGHT_CONTROL, 100, 0, 0)

	if !preHotkey {
		logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.F9, 0, 0, 100)
		logiKeyboard.SetLightingForKeyWithKeyName(LogiKeyboardTypes.ESC, 0, 100, 0)

		//logiKeyboard.SetFlashSingleKey(LogiKeyboardTypes.F, 0, 100, 0, 0, 500)
		//logiKeyboard.SetPulseSingleKey(LogiKeyboardTypes.F, 0, 0, 0, 100, 100, 100, 500, 1)
	}
}

func run(shortcuts []Shortcuts.Shortcut) error {
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

	// get all relevant keys (modifiers)
	var relevantKeys []types.VKCode
	linq.From(shortcuts).
		SelectManyT(func(c Shortcuts.Shortcut) linq.Query {
			return linq.From(c.Modifiers())
		}).
		GroupByT(func(modifier types.VKCode) types.VKCode {
			return modifier
		}, func(modifier types.VKCode) types.VKCode {
			return modifier
		}).
		SelectT(func(c linq.Group) types.VKCode {
			return c.Key.(types.VKCode)
		}).
		ToSlice(&relevantKeys)

	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case k := <-keyboardChan:
			var prevState = currentlyPressedKeys[k.VKCode]
			currentlyPressedKeys[k.VKCode] = k.Message == types.WM_KEYDOWN || k.Message == types.WM_SYSKEYDOWN

			// if the key remained in the same state: ignore
			if prevState == currentlyPressedKeys[k.VKCode] {
				continue
			}

			// if the key is no modifier: ignore
			if !linq.From(relevantKeys).
				Contains(k.VKCode) {
				continue
			}

			fmt.Printf("Received %v %v\n", k.Message, k.VKCode)

			shortCut := linq.From(shortcuts).
				WhereT(func(c Shortcuts.Shortcut) bool {
					found := linq.From(c.Modifiers()).
						AllT(func(y types.VKCode) bool {
							return currentlyPressedKeys[y]
						})

					if !found {
						return false
					}

					found = linq.From(currentlyPressedKeys).
						WhereT(func(y linq.KeyValue) bool {
							return y.Value.(bool)
						}).
						AnyWithT(func(y linq.KeyValue) bool {
							return !linq.From(c.Modifiers()).
								Contains(y.Key.(types.VKCode))
						})

					return !found
				}).
				First()

			if shortCut != nil {
				defaultLightning(true)
				for _, logiKey := range shortCut.(Shortcuts.Shortcut).Keys() {
					switch logiKey.Effect() {
					case Shortcuts.Fixed:
						logiKeyboard.SetLightingForKeyWithKeyName(logiKey.Key(), logiKey.Red(), logiKey.Green(), logiKey.Blue())
					case Shortcuts.Blinking:
						logiKeyboard.SetFlashSingleKey(logiKey.Key(), logiKey.Red(), logiKey.Green(), logiKey.Blue(), LogiKeyboardTypes.LogiLedDurationInfinite, 500)
					}
				}
			} else {
				defaultLightning(false)
			}
		}

		continue
	}
}
