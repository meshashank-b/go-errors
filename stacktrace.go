package apperror

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
)

// stackTrace represents a series of stack frames collected during an error.
type stackTrace struct {
	Frames []stackFrame `json:"frames"`
}

// stackFrame represents a single frame in a stack trace.
type stackFrame struct {
	Function        string  `json:"function"`
	File            string  `json:"file"`
	Line            int     `json:"line"`
	Pointer         uintptr `json:"pointer"`
	FunctionPointer uintptr `json:"function_pointer"`
}

// captureStackTrace collects the current stack trace, starting from the caller of this function.
func captureStackTrace() *stackTrace {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	frames := make([]stackFrame, 0, n) // Preallocate slice to reduce reallocations

	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			file, line := fn.FileLine(pc)
			frames = append(frames, stackFrame{
				Function:        fn.Name(),
				File:            file,
				Line:            line,
				Pointer:         pc,
				FunctionPointer: fn.Entry(),
			})
		} else {
			frames = append(frames, stackFrame{
				Function:        "unknown",
				File:            "unknown",
				Line:            0,
				Pointer:         pc,
				FunctionPointer: 0,
			})
		}
	}
	return &stackTrace{Frames: frames}
}

// String returns a formatted string representation of the stack trace.
func (st *stackTrace) String() string {
	var sb strings.Builder
	sb.WriteString("Stack Trace:\n---------------------------------------------------\n")
	for i, frame := range st.Frames {
		sb.WriteString("#" + strconv.Itoa(i+1) + " - Function: " + frame.Function + "\n")
		sb.WriteString("     Location: " + frame.File + ":" + strconv.Itoa(frame.Line) + "\n")
		sb.WriteString("     Pointer: 0x" + strconv.FormatUint(uint64(frame.Pointer), 16) + "\n")
		sb.WriteString("     Function Pointer: 0x" + strconv.FormatUint(uint64(frame.FunctionPointer), 16) + "\n")
		sb.WriteString("---------------------------------------------------\n")
	}
	return sb.String()
}

// enhanceWithCause merges another stack trace into the current one immutably.
func (st *stackTrace) enhanceWithCause(original *stackTrace) *stackTrace {
	newFrames := append(st.Frames, original.Frames...)
	return &stackTrace{Frames: newFrames}
}

// MarshalJSON provides a pretty JSON format for the stack trace.
func (st *stackTrace) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(st, "", "  ")
}
