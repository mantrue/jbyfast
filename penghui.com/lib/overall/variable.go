package overall

import "runtime"

var (
	Filepath string
	Line     int
)

func init() {
	_, Filepath, Line, _ = runtime.Caller(0)
}
