package logutils

// 2024.04 Remove level Progress. This aligns this file
// better with the RFC, and it leaves these levels:
// - Panic/Abort
// - R/Y/G Error/Warning/Okay
// - Info (the workhorse)
// - Debug (changed back from "Dbg" 
//   cos now the numbers align again)  

// Levels (3,4,5) (Error,Warning,Okay) (3,4,5) (R,Y,G) are
// intended as summary items for execution checkpoints. Grn,
// Ylw, Red (currently disabled) are calm B/G indicator lights.
// .
const (
	// NOTIFICATION / SUMMARY
	EmojiPanic   = "❌"  // 2 X
	EmojiError   = "❌"  // 3 R
	EmojiWarning = "🟨"  // 4 Y
	EmojOkay     = "🟩"  // 5 G
	EmojiInfo    = "ℹ️" // 6 I
	// TRANSIENT
	// EmojiProgress = "▫️" // 7
	// EmojiDbg = "❓"  // misspelled cos 8 != RFC5424 "7"
	EmojiDebug = "💠"
	/* STATE INDICATORS
	Red = "🔴"
	Ylw = "🟡"
	Grn = "🟢"
	*/
)

/*  This   RFC5424
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
		return "❌💀❌"
	case 3:
		return "❌"
	case 4:
		return "🟨"
	case 5:
		return "🟩"
	case 6:
		return "ℹ️ " //  "💬"
	case 7: /* 
		return "〰️"
	case 8: */ 
		return "💠" //  "❓"
	}
	return "?!?!"
}

/* Stockpile of useful emojis:
⭕ ✅ ❌ ❎
🔴 🟠 🟡 🟢 🔵 🟣 🟤 ⚫ ⚪
🟥 🟧 🟨 🟩 🟦 🟪 🟫 ⬛ ⬜ ◾ ◽
🔶 🔷 🔸 🔹 🔺 🔻 💠 🔘 🔳 🔲
ⓘ ⓘ  🛈 ℹ️   
*/
