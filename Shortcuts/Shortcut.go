package Shortcuts

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/types"
)

type Shortcut struct {
	Modifiers []types.VKCode
	Keys      []ShortcutKey
}

func (s *Shortcut) Create(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name) *Shortcut {
	s.Modifiers = modifiers

	var k = make([]ShortcutKey, len(keys))
	s.Keys = k

	for _, key := range keys {
		sck := CreateKey(key)
		s.Keys = append(s.Keys, sck)
	}

	return s
}

func (s *Shortcut) CreateColor(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name, red, green, blue int) *Shortcut {
	s.Modifiers = modifiers

	var k = make([]ShortcutKey, len(keys))
	s.Keys = k

	for _, key := range keys {
		sck := CreateKeyColor(key, red, green, blue)
		s.Keys = append(s.Keys, sck)
	}

	return s
}

func (s *Shortcut) CreateWithKey(modifiers []types.VKCode, keys []ShortcutKey) *Shortcut {
	s.Modifiers = modifiers

	s.Keys = keys

	return s
}
