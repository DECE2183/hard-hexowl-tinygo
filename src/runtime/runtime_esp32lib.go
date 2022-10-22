//go:build esp32 && lib
// +build esp32,lib

package runtime

import (
	"device"
)

//go:extern _bss_start
var _sbss [0]byte

//go:extern _bss_end
var _ebss [0]byte

//go:linkname main main.main
func main() {
	for {
		sleep(1000)
	}
}

func abort() {
	for {
		device.Asm("waiti 0")
	}
}

// interruptInit initialize the interrupt controller and called from runtime once.
func interruptInit() {
	device.Asm(`
		movi    a3, 0				
    	xsr     a3, INTENABLE      
    	rsync						
    	or      a2, a3, a2         
    	wsr     a2, INTENABLE      
    	rsync
	`)
}