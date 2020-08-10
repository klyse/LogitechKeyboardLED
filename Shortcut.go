package main

import (
	"github.com/klyse/LogitechKeyboardLED/LogiKeyboardTypes"
	"github.com/moutend/go-hook/pkg/types"
)

type Shortcut struct {
	modifiers []types.VKCode
	keys      []LogiKeyboardTypes.Name
}

func (s *Shortcut) Create(modifiers []types.VKCode, keys []LogiKeyboardTypes.Name) *Shortcut {
	s.modifiers = modifiers
	s.keys = keys

	return s
}
