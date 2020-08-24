package Shortcuts

import "github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"

type Effect int

const (
	Fixed    Effect = iota
	Blinking        = iota
)

type ShortcutKey interface {
	Key() LogiKeyboardTypes.Name
	Effect() Effect
	Red() int
	Green() int
	Blue() int
}

type shortcutKey struct {
	key    LogiKeyboardTypes.Name
	effect Effect
	red    int
	green  int
	blue   int
}

func (c shortcutKey) Key() LogiKeyboardTypes.Name {
	return c.key
}

func (c shortcutKey) Effect() Effect {
	return c.effect
}

func (c shortcutKey) Red() int {
	return c.red
}

func (c shortcutKey) Green() int {
	return c.green
}

func (c shortcutKey) Blue() int {
	return c.blue
}

func CreateKey(key LogiKeyboardTypes.Name) ShortcutKey {
	return shortcutKey{key, Fixed, 0, 0, 0}
}

func CreateKeyColor(key LogiKeyboardTypes.Name, red, green, blue int) ShortcutKey {
	return shortcutKey{key, Fixed, red, green, blue}
}

func CreateKeyColorEffect(key LogiKeyboardTypes.Name, red, green, blue int, effect Effect) ShortcutKey {
	return shortcutKey{key, effect, red, green, blue}
}
