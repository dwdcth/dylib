package dylib

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"unsafe"
)

type LazyDLL struct {
	mu     sync.Mutex
	handle uintptr
	err    error
	Name   string
	// 加了异常处理的
	mySyscall *LazyProc
}

func NewLazyDLL(name string) *LazyDLL {
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (l *LazyDLL) libFullPath(name string) string {
	if runtime.GOOS == "darwin" {
		file, _ := exec.LookPath(os.Args[0])
		libPath := filepath.Dir(file) + "/" + name
		if fileExists(libPath) {
			return libPath
		}
	}
// +build !windows
return name
}

func (l *LazyDLL) Load() error {
	return l.err
}

func (l *LazyDLL) NewProc(name string) *LazyProc {
	p := new(LazyProc)
	p.Name = name
	p.lzdll = l
	return p
}

func (l *LazyDLL) Close() {

}

func (d *LazyDLL) call(proc *LazyProc, a ...uintptr) (r1, r2 uintptr, lastErr error) {
	return 0, 0, errors.New("not supported")
}

type LazyProc struct {
	mu    sync.Mutex
	p     uintptr
	Name  string
	lzdll *LazyDLL
}

func (p *LazyProc) Addr() uintptr {
	err := p.Find()
	if err != nil {
		fmt.Println(err)
	}
	return p.p
}

func (p *LazyProc) Find() error {
	return errors.New("not supported")
}

func toPtr(a uintptr) unsafe.Pointer {
	return unsafe.Pointer(a)
}

func (p *LazyProc) Call(a ...uintptr) (uintptr, uintptr, error) {
	return 0, 0, errors.New("not supported")
}

func (p *LazyProc) CallOriginal(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	return 0, 0, errors.New("not supported")
}
