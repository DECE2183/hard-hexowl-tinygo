//go:build lib
// +build lib

package runtime

//go:linkname sleep time.Sleep
func sleep(duration int64) {
	if duration <= 0 {
		return
	}

	sleepTicks(nanosecondsToTicks(duration))
}

// run is called by the program entry point to execute the go program.
// With the "none" scheduler, init and the main function are invoked directly.
//export gorun
func run(heapSize uintptr) {
	initHeap(heapSize)
	initAll()
}

const hasScheduler = false
