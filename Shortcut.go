package main

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/types"
)

type ShortcutKey struct {
	key   LogiKeyboardTypes.Name
	red   int
	green int
	blue  int
}

func (s *ShortcutKey) Create(key LogiKeyboardTypes.Name) *ShortcutKey {
	s.key = key
	s.blue = 0
	s.green = 0
	s.red = 100

	return s
}

func (s *ShortcutKey) CreateColor(key LogiKeyboardTypes.Name, red, green, blue int) *ShortcutKey {
	s.key = key
	s.blue = blue
	s.green = green
	s.red = red

	return s
}

type Shortcut struct {
	modifiers []types.VKCode
	keys      []ShortcutKey
}

func (s *Shortcut) Create(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name) *Shortcut {
	s.modifiers = modifiers

	var k = make([]ShortcutKey, len(keys))
	s.keys = k

	for _, key := range keys {
		sck := new(ShortcutKey).Create(key)
		s.keys = append(s.keys, *sck)
	}

	return s
}

func (s *Shortcut) CreateColor(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name, red, green, blue int) *Shortcut {
	s.modifiers = modifiers

	var k = make([]ShortcutKey, len(keys))
	s.keys = k

	for _, key := range keys {
		sck := new(ShortcutKey).CreateColor(key, red, green, blue)
		s.keys = append(s.keys, *sck)
	}

	return s
}

func (s *Shortcut) CreateWithKey(modifiers []types.VKCode, keys []ShortcutKey) *Shortcut {
	s.modifiers = modifiers

	s.keys = keys

	return s
}
