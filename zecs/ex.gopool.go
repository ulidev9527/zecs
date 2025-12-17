package zecs

import (
	"fmt"
	"sync"

	"github.com/panjf2000/ants/v2"
)

// 扩展 go 对象池

const (
	GoPool_Max = 100000 // 最大协程数 10W
	GoPool_Min = 1000   // 最小协程数 1K
)

var (
	goPool, _             = ants.NewPool(GoPool_Min)
	goPool_mu             sync.Mutex
	goPool_maxErrOutputCD = 0
)

func GO(f func()) {
	goPool_mu.Lock()
	defer goPool_mu.Unlock()
	if e := goPool.Submit(f); e != nil {
		fmt.Print("ERR: zserbice.GO :", e)
	}
	if goPool.Free() < 100 {

		if goPool.Cap() < GoPool_Max {
			goPool.Tune(goPool.Cap() + GoPool_Min)
			fmt.Print("Warn: zserbice.GO TUNE:", goPool.Cap())
		} else if goPool_maxErrOutputCD <= 0 {
			fmt.Print("Err: zserbice.GO TUNE MAX", goPool.Cap())
			goPool_maxErrOutputCD = GoPool_Max
		}
		goPool_maxErrOutputCD--
	}
}
