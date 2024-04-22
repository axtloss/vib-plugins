// Copyright 2023 - 2023, axtlos <axtlos@disroot.org>
// SPDX-License-Identifier: GPL-3.0-ONLY

package main

import (
	"encoding/json"
	"testing"
)

type testModule struct {
	Name       string
	Type       string
	ExtraFlags []string
	Packages   []string
}

type testCases struct {
	module   interface{}
	expected string
}

var test = []testCases{
	{testModule{"Single Package, Single Flag", "pacman", []string{"--overwrite=\"*\""}, []string{"bash"}}, "pacman -S --noconfirm --overwrite=\"*\" bash"},
	{testModule{"Single Package, No Flag", "pacman", []string{""}, []string{"bash"}}, "pacman -S --noconfirm  bash"},
	{testModule{"Multiple Packages, No Flag", "pacman", []string{""}, []string{"bash", "fish"}}, "pacman -S --noconfirm  bash fish"},
	{testModule{"Multiple Packages, Multiple Flags", "pacman", []string{"--overwrite=\"*\"", "--verbose"}, []string{"bash", "fish"}}, "pacman -S --noconfirm --overwrite=\"*\" --verbose bash fish"},
}

func TestBuildModule(t *testing.T) {
	for _, testCase := range test {
		moduleInterface, err := json.Marshal(testCase.module)
		if err != nil {
			t.Errorf("Error in json %s", err.Error())
		}
		if output := BuildModule(convertToCString(string(moduleInterface)), convertToCString("")); convertToGoString(output) != testCase.expected {
			t.Errorf("Output %s not equivalent to expected %s", convertToGoString(output), testCase.expected)
		}
	}

}
