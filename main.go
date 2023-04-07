package main

// #include <mach/mach.h>
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	var task C.mach_port_t
	if C.task_for_pid(C.mach_task_self_, C.int(os.Getpid()), &task) != C.KERN_SUCCESS {
		fmt.Println("Error: Unable to obtain the task port")
		return
	}
	var n_threads C.mach_msg_type_number_t
	var threads C.thread_act_array_t
	if C.task_threads(task, &threads, &n_threads) != C.KERN_SUCCESS {
		fmt.Println("Error: Unable to obtain the thread list")
		return
	}
	t := C.vm_address_t(uintptr(unsafe.Pointer(threads)))
	if C.vm_deallocate(C.mach_task_self_, t, C.vm_size_t(n_threads*C.sizeof_thread_act_t)) != C.KERN_SUCCESS {
		fmt.Println("Error: Unable to deallocate memory")
		return
	}
	fmt.Printf("Number of threads in the current process: %d\n", int(n_threads))
}
