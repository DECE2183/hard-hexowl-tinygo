//go:build gc.conservative && lib
// +build gc.conservative,lib

package runtime

import (
	"unsafe"
	"internal/task"
)

type freeRTOSTaskTCB struct {
	pxTopOfStack uintptr
	useless      [48]byte
	pxStack      uintptr
	taskName     [16]byte
	xCoreID      int32
	pxEndOfStack uintptr
}

// markStack marks all root pointers found on the stack.
//
// This implementation is conservative and relies on the stack top (provided by
// the linker) and getting the current stack pointer from a register. Also, it
// assumes a descending stack. Thus, it is not very portable.
func markStack() {
	vTaskSuspendAll()

	// Scan the current stack, and all current registers.
	scanCurrentStack()

	if !task.OnSystemStack() {
		// Mark system stack.
		markRoots(getSystemStackPointer(), stackTop)
	}

	xTaskResumeAll()
}

//go:export tinygo_scanCurrentStack
func scanCurrentStack()

//go:export tinygo_scanstack
func scanstack(sp uintptr) {
	// Mark current stack.
	// This function is called by scanCurrentStack, after pushing all registers onto the stack.
	// Callee-saved registers have been pushed onto stack by tinygo_localscan, so this will scan them too.
	markRoot(0, sp)
	tcb := xTaskGetCurrentTaskHandle()
	// println("stack from:", tcb.pxTopOfStack, "to:", tcb.pxEndOfStack)
	markRoots(tcb.pxStack, tcb.pxEndOfStack)
}

//go:extern pxTaskGetStackStart
//go:export pxTaskGetStackStart
func pxTaskGetStackStart(task unsafe.Pointer) uintptr

//go:extern xTaskGetCurrentTaskHandle
//go:export xTaskGetCurrentTaskHandle
func xTaskGetCurrentTaskHandle() *freeRTOSTaskTCB

//go:extern vTaskSuspendAll
//go:export vTaskSuspendAll
func vTaskSuspendAll()

//go:extern xTaskResumeAll
//go:export xTaskResumeAll
func xTaskResumeAll() uintptr
