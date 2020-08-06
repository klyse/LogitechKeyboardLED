package main

import "C"
import (
	"fmt"
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboard"
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	k := LogiKeyboard.Create()

	defer k.Shutdown()

	k.Init()

	k.SetTargetDevice(LogiKeyboardTypes.LogiDeviceTypeAll)

	if err := run(k); err != nil {
		log.Fatal(err)
	}
}

func run(ka *LogiKeyboard.LogiKeyboard) error {
	// Buffer size is depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")

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

			switch key := k.VKCode; key {
			case types.VK_LWIN:

				var red = rand.Intn(100)
				var green = rand.Intn(100)
				var blue = rand.Intn(100)
				for x := LogiKeyboardTypes.ESC; x < LogiKeyboardTypes.END; x++ {
					red = (red + 1) % 100
					green = (green + 1) % 100
					blue = (blue + 1) % 100
					ka.SetLightingForKeyWithKeyName(x, red, green, blue)
				}
			}

			continue
		}
	}
}
