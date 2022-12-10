package logutils

// LogMessage is a log level + string message pair.
// It should be accumulated into a slice by code that
// does not have access to system logging facilities.
//
// The string can be (for example) the output of
// a call to fmt.Sprintf(..) that would normally
// be passed to a logger. It could therefore be
// created using the "%w" format specifier.
// .
type LogMessage struct {
	Level
	string
}

type LogMsgList []LogMessage

func (llm LogMsgList) MaxLevel(maxLvl Level) {
}
