package common

import (
	"sync"
)

type Lock struct {
	DeviceLock *sync.Mutex
	// Timer *time.Timer
	// flag bool
}

func (dl *Lock) Lock() {
	dl.DeviceLock.Lock()
	// dl.Timer = time.AfterFunc(time.Second, func() {
	// dl.DeviceLock.Unlock()
	// // need to skip critical section
	// dl.flag = true
	// })
}

func (dl *Lock) Unlock() {
	//dl.Timer.Stop()
	//if !dl.flag{
	//	dl.DeviceLock.Unlock()
	//}
	dl.DeviceLock.Unlock()
}
