package tracex

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var (
	//version string
	incrNum uint64
	//pid     = os.Getpid()
)

type (
	TraceIDCtx struct{}
)

// NewTraceID New tracex id
func NewTraceID() string {
	return fmt.Sprintf("%d%s%d",
		os.Getpid(),
		time.Now().Format("20060102150405999"),
		atomic.AddUint64(&incrNum, 1))
}
