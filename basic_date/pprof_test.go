package basic_date

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func TestPprofCPU(t *testing.T) {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()
}
