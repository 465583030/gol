package gol

// LogOption define which text to prefix to each log entry generated by the Logger.
type LogOption int8

const (
	// Bits or'ed together to control what's printed.
	// There is no control over the order they appear (the order listed
	// here) or the format they present (as described in the comments).
	// The prefix is followed by a colon only when Llongfile or Lshortfile
	// is specified.
	// For example, flags Ldate | Ltime (or LstdFlags) produce,
	//	2009/01/23 01:23:23 message
	// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         LogOption                    = 1 << iota // the date in the local time zone: 2009/01/23
	Ltime                                                  // the time in the local time zone: 01:23:23
	Lmicroseconds                                          // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                                              // full file name and line number: /a/b/c/d.go:23
	Lshortfile                                             // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                                                   // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime | Lshortfile             // initial values for the standard logger
)
