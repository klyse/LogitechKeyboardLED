package main

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/types"
)

type ShortcutKey struct {
	Key   LogiKeyboardTypes.Name
	Red   int
	Green int
	Blue  int
}

func (s *ShortcutKey) Create(key LogiKeyboardTypes.Name) *ShortcutKey {
	s.Key = key
	s.Blue = 0
	s.Green = 0
	s.Red = 100

	return s
}

func (s *ShortcutKey) CreateColor(key LogiKeyboardTypes.Name, red, green, blue int) *ShortcutKey {
	s.Key = key
	s.Blue = blue
	s.Green = green
	s.Red = red

	return s
}

type Shortcut struct {
	Modifiers []types.VKCode
	Keys      []ShortcutKey
}

func (s *Shortcut) Create(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name) *Shortcut {
	s.Modifiers = modifiers

	var k = make([]ShortcutKey, len(keys))
	s.Keys = k

	for _, key := range keys {
		sck := new(ShortcutKey).Create(key)
		s.Keys = append(s.Keys, *sck)
	}

	return s
}

func (s *Shortcut) CreateColor(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name, red, green, blue int) *Shortcut {
	s.Modifiers = modifiers

	var k = make([]ShortcutKey, len(keys))
	s.Keys = k

	for _, key := range keys {
		sck := new(ShortcutKey).CreateColor(key, red, green, blue)
		s.Keys = append(s.Keys, *sck)
	}

	return s
}

func (s *Shortcut) CreateWithKey(modifiers []types.VKCode, keys []ShortcutKey) *Shortcut {
	s.Modifiers = modifiers

	s.Keys = keys

	return s
}
