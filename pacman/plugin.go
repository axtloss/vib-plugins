// Copyright 2023 - 2023, axtlos <axtlos@disroot.org>
// SPDX-License-Identifier: GPL-3.0-ONLY

package main

import (
	"C"
	"encoding/json"
	"fmt"
	"strings"
)

type PacmanModule struct {
	Name string `json:"name"`
	Type string `json:"type"`

	ExtraFlags []string
	Packages   []string
}

//export PlugInfo
func PlugInfo() *C.char {
	plugininfo := &api.PluginInfo{Name: "pacman", Type: api.BuildPlugin}
	pluginjson, err := json.Marshal(plugininfo)
	if err != nil {
		return C.CString(fmt.Sprintf("ERROR: %s", err.Error()))
	}
	return C.CString(string(pluginjson))
}

func convertToCString(s string) *C.char {
	return C.CString(s)
}

func convertToGoString(s *C.char) string {
	return C.GoString(s)
}

//export BuildModule
func BuildModule(moduleInterface *C.char, _ *C.char) *C.char {
	var module PacmanModule
	err := json.Unmarshal([]byte(C.GoString(moduleInterface)), &module)
	if err != nil {
		return C.CString(fmt.Sprintf("ERROR: %s", err.Error()))
	}

	cmd := fmt.Sprintf("pacman -S --noconfirm %s %s", strings.Join(module.ExtraFlags, " "), strings.Join(module.Packages, " "))

	return C.CString(cmd)
}

func main() {}
