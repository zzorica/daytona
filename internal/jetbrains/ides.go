// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package jetbrains

var ides = map[Id]Ide{
	CLion:    {"CL", "CLion"},
	IntelliJ: {"IIU", "IntelliJ IDEA Ultimate"},
	GoLand:   {"GO", "GoLand"},
	PyCharm:  {"PCP", "PyCharm Professional"},
	PhpStorm: {"PS", "PhpStorm"},
	WebStorm: {"WS", "WebStorm"},
	Rider:    {"RD", "Rider"},
	RubyMine: {"RM", "RubyMine"},
}

func GetIdes() map[Id]Ide {
	return ides
}

func GetIde(id Id) (Ide, bool) {
	ide, ok := ides[id]
	return ide, ok
}
