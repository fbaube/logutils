// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logutils

// TODO: If "app" is default, can be omitted from log msgs.
// if is category, and filterable, then use for go pkg name,
// or for input file name.

// Level describes the level of a log message.
type Level int

// RFC5424 log message levels.
const (
	LevelPanic    Level = iota + 2
	LevelError          // 3
	LevelWarning        // 4
	LevelOkay           // 5
	LevelInfo           // 6
	LevelProgress       // 7
	LevelDbg            // misspelled cos 8 != RFC5424 "7"
	//
	// Utility levels for fancier console output
	GreenBG
)

// LevelNames maps log levels to user-frenly names
var LevelNames = map[Level]string{
	LevelDbg:      "?Dbg",
	LevelProgress: "Progress",
	LevelInfo:     "Info",
	LevelOkay:     "Okay",
	LevelWarning:  "Warning",
	LevelError:    "Error",
	LevelPanic:    "PANIC",
}

// String returns the string representation of the log level
func (l Level) String() string {
	if name, ok := LevelNames[l]; ok {
		return name
	}
	return "Unknown_WTH"
}
