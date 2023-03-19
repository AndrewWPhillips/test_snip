package Windows

import (
	"log"
	"syscall"
	"testing"
	"unsafe"
)

func TestLogonUser(t *testing.T) {
	const (
		LOGON32_LOGON_INTERACTIVE = 2
		LOGON32_LOGON_NETWORK     = 3
	)
	advapi32, err := syscall.LoadLibrary("advapi32.dll")
	if err != nil {
		log.Printf("Error A %v\n", err)
	}
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Printf("Error B %v\n", err)
	}
	logonUser, err := syscall.GetProcAddress(advapi32, "LogonUserW")
	if err != nil {
		log.Printf("Error D %v\n", err)
	}
	closeHandle, err := syscall.GetProcAddress(kernel32, "CloseHandle")
	if err != nil {
		log.Printf("Error E %v\n", err)
	}

	name, err := syscall.UTF16PtrFromString("Andrew")
	if err != nil {
		log.Printf("Error J %v\n", err)
	}
	domain, err := syscall.UTF16PtrFromString("DOMAIN") // wrong domain? => An attempt was made to logon, but the network logon service was not started
	if err != nil {
		log.Printf("Error K %v\n", err)
	}
	password, err := syscall.UTF16PtrFromString("Password")
	if err != nil {
		log.Printf("Error L %v\n", err)
	}
	var handle uintptr

	ret, ret2, callErr := syscall.Syscall6(uintptr(logonUser),
		uintptr(6),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(domain)),
		uintptr(unsafe.Pointer(password)),
		uintptr(LOGON32_LOGON_NETWORK),
		uintptr(0),
		uintptr(unsafe.Pointer(&handle)),
	)
	log.Printf("%v %v %v\n", ret, ret2, callErr)
	// If unknown ID or wrong password:  0 0 The user name or password is incorrect.
	// If known ID and correct password: 1 0 The operation completed successfully.

	ret, ret2, callErr = syscall.Syscall(uintptr(closeHandle),
		uintptr(1),
		uintptr(unsafe.Pointer(handle)),
		0,
		0,
	)
	log.Printf("%v %v %v\n", ret, ret2, callErr)

	err = syscall.FreeLibrary(advapi32)
	if err != nil {
		log.Printf("Error Z %v\n", err)
	}
}
