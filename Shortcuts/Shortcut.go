package Shortcuts

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/types"
)

type Shortcut interface {
	Modifiers() []types.VKCode
	Keys() []ShortcutKey
}

type shortcut struct {
	modifiers []types.VKCode
	keys      []ShortcutKey
}

func (s shortcut) Modifiers() []types.VKCode {
	return s.modifiers
}

func (s shortcut) Keys() []ShortcutKey {
	return s.keys
}

func Create(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name) Shortcut {
	s := shortcut{modifiers: modifiers}

	var k = make([]ShortcutKey, len(keys))
	s.keys = k

	for _, key := range keys {
		sck := CreateKey(key)
		s.keys = append(s.keys, sck)
	}

	return s
}

func CreateColor(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name, red, green, blue int) Shortcut {
	s := shortcut{modifiers: modifiers}

	var k = make([]ShortcutKey, len(keys))
	s.keys = k

	for _, key := range keys {
		sck := CreateKeyColor(key, red, green, blue)
		s.keys = append(s.keys, sck)
	}

	return s
}

func CreateWithKey(modifiers []types.VKCode, keys []ShortcutKey) Shortcut {
	s := shortcut{modifiers, keys}

	return s
}
