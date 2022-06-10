// +build windows

package util

import (
	"syscall"
)

type YHDll struct {
	Handle syscall.Handle
}

// 加载Dll.
func (dll *YHDll) LoadDll(dllPath string) error {
	var err error
	dll.Handle, err = syscall.LoadLibrary(dllPath)
	if err != nil {
		return err
	}

	//dll.Handle = syscall.NewLazyDLL(dllPath)
	return nil
}

// 释放Dll.
func (dll *YHDll) UnLoadDll() error {
	err := syscall.FreeLibrary(dll.Handle)
	if err != nil {
		return err
	}

	return nil
}

// 获取dll函数地址.
func (dll *YHDll) GetProcAddress(funcName string) (uintptr, error) {
	proc, err := syscall.GetProcAddress(dll.Handle, funcName)
	if err != nil {
		return 0, err
	}

	//proc := dll.Handle.NewProc(funcName)

	return proc, nil
}
