package qlog

//ConsoleColor ...
type ConsoleColor string

const (
	//ConsoleBlack black ansi.
	ConsoleBlack ConsoleColor = "[30m"

	//ConsoleRed red ansi.
	ConsoleRed ConsoleColor = "[31m"

	//ConsoleGreen green ansi.
	ConsoleGreen ConsoleColor = "[32m"

	//ConsoleYellow yellow ansi.
	ConsoleYellow ConsoleColor = "[33m"

	//ConsoleBlue blue ansi.
	ConsoleBlue ConsoleColor = "[34m"

	//ConsoleMagenta magenta ansi.
	ConsoleMagenta ConsoleColor = "[35m"

	//ConsoleCyan cyan ansi.
	ConsoleCyan ConsoleColor = "[36m"

	//ConsoleLightGray light gray ansi.
	ConsoleLightGray ConsoleColor = "[37m"

	//ConsoleDarkGray dark gray ansi.
	ConsoleDarkGray ConsoleColor = "[90m"

	//ConsoleLightRed light red ansi.
	ConsoleLightRed ConsoleColor = "[91m"

	//ConsoleLightGreen light green ansi.
	ConsoleLightGreen ConsoleColor = "[92m"

	//ConsoleLightYellow light yellow ansi.
	ConsoleLightYellow ConsoleColor = "[93m"

	//ConsoleLightBlue light blue ansi.
	ConsoleLightBlue ConsoleColor = "[94m"

	//ConsoleLightMagenta light magenta ansi.
	ConsoleLightMagenta ConsoleColor = "[95m"

	//ConsoleLightCyan light cyan ansi.
	ConsoleLightCyan ConsoleColor = "[96m"

	//ConsoleWhite light white ansi.
	ConsoleWhite ConsoleColor = "[97m"
)