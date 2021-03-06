package pytesting

import (
	"lime/3rdparty/libs/gopy/lib"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func test2() {
	l := py.NewLock()
	defer l.Unlock()
}

func test() {
	t := time.Now()
	for time.Since(t) < time.Second*2 {
		func() {
			l := py.NewLock()
			defer l.Unlock()
			test2()
			time.Sleep(time.Duration(float64(time.Millisecond) * (1 + rand.Float64())))
		}()
	}
}

func TestLock(t *testing.T) {
	l := py.InitAndLock()
	l.Unlock()
	defer func() {
		l.Lock()
		py.Finalize()
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())
	go test()
	go test()
	go test()
	test()
}
