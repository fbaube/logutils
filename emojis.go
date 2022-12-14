package logutils

// Levels (3,4,5) (Error,Warning,Okay) (3,4,5) (R,Y,G) are
// intended as summary items for execution checkpoints. Grn,
// Ylw, Red (currently disabled) are calm B/G indicator lights.
// .
const (
	// NOTIFICATION / SUMMARY
	EmojiPanic   = "â"  // 2 X
	EmojiError   = "â"  // 3 R
	EmojiWarning = "đ¨"  // 4 Y
	EmojOkay     = "đŠ"  // 5 G
	EmojiInfo    = "âšī¸" // 6 I
	// TRANSIENT
	EmojiProgress = "âĢī¸" // 7
	EmojiDbg      = "â"  // misspelled cos 8 != RFC5424 "7"
	/* STATE INDICATORS
	Red = "đ´"
	Ylw = "đĄ"
	Grn = "đĸ"
	*/
)

/*  This    RFC5424
0   -      Emergency (system is unusable)
1   -      Alert (take action ASAP)
2  Panic   Critical
3  Error   Error
4  Warning Warning
5  Okay    Notice (normal but significant condition)
6  Info    Informational
7  Debug   Debug
*/

func EmojiOfLevel(L Level) string {
	switch L {
	case 0, 1, 2:
		return "âđâ"
	case 3:
		return "â"
	case 4:
		return "đ¨"
	case 5:
		return "đŠ"
	case 6:
		return "đŦ"
	case 7:
		return "ã°ī¸"
	case 8:
		return "â"
	}
	return "?!?!"
}

/*
Stockpile of useful emojis:
â­ â â â
đ´ đ  đĄ đĸ đĩ đŖ đ¤ âĢ âĒ
đĨ đ§ đ¨ đŠ đĻ đĒ đĢ âŦ âŦ âž âŊ
đļ đˇ đ¸ đš đē đģ đ  đ đŗ đ˛
*/
