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

/* NOTE also
fromm the new package log/slog:
const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)
We wanted to make it easy to use levels to specify logger verbosity.
Since a larger level means a more severe event, a logger that accepts
events with smaller (or more negative) level means a more verbose log-
ger. Logger verbosity is thus the negation of event severity, and the
default verbosity of 0 accepts all events at least as severe as INFO.

Also we wanted some room between levels to accommodate schemes with
named levels between ours. For example, Google Cloud Logging defines
a Notice level between Info and Warn. Since there are only a few of
these intermediate levels, the gap between the numbers need not be
large. Our gap of 4 matches OpenTelemetry's mapping. Subtracting 9
from an OpenTelemetry level in the DEBUG, INFO, WARN and ERROR rang-
es converts it to the corresponding slog Level range. OpenTelemetry
also has the names TRACE and FATAL, which slog does not. But those
OpenTelemetry levels can still be represented as slog Levels by
using the appropriate integers.
*/

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
