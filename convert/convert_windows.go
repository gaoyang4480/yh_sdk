package convert

/*
#cgo CFLAGS: -I ./include
#cgo CPPFLAGS: -I ./include
#include "convert.h"
*/
import "C"

import (
	"errors"
	"github.com/gaoyang4480/yh_sdk/util"
	"syscall"
	"unsafe"
)

const (
	initSDKFuncName       = "YH_InitSDK"
	finalizeSDKFuncName   = "YH_FinalizeSDK"
	initAgentFuncName     = "YH_InitAgent"
	finalizeAgentFuncName = "YH_FinalizeAgent"
	deleteMemFuncName     = "YH_DeleteMem"
	officeToOFDFuncName   = "YH_OfficeToOFD"
)

var (
	globalDll *util.YHDll
)

// ConvertAgent 转换代理.
type ConvertAgent struct {
	agentPtr C.YH_CONVERT_AGENT
}

var _ IConvertAgent = (*ConvertAgent)(nil)

// LoadConvert 加载转换动态库.
func LoadConvert(convertDllFilePath string) error {
	if globalDll != nil {
		return nil
	}
	if len(convertDllFilePath) == 0 {
		return errors.New("convert dll file path is empty")
	}
	globalDll = &util.YHDll{}
	err := globalDll.LoadDll(convertDllFilePath)
	if err != nil {
		return errors.New("load convert dll error")
	}
	return nil
}

// UnloadConvert 卸载加载转换动态库.
func UnloadConvert() {
	if globalDll != nil {
		_ = globalDll.UnLoadDll()
	}
}

// InitSDK 初始化.
func (agent *ConvertAgent) InitSDK() error {
	if globalDll == nil {
		return errors.New("the convert dynamic library is not loaded")
	}
	proc, err := globalDll.GetProcAddress(initSDKFuncName)
	if err != nil {
		return err
	}
	ret, _, _ := syscall.Syscall(proc, uintptr(0), 0, 0, 0)
	if ret != C.YH_OK {
		return errors.New("init sdk error")
	}
	return nil
}

// FinalizeSDK 销毁.
func (agent *ConvertAgent) FinalizeSDK() error {
	if globalDll == nil {
		return errors.New("the convert dynamic library is not loaded")
	}
	proc, err := globalDll.GetProcAddress(finalizeSDKFuncName)
	if err != nil {
		return err
	}
	_, _, _ = syscall.Syscall(proc, uintptr(0), 0, 0, 0)
	return nil
}

// InitAgent 初始化代理.
func (agent *ConvertAgent) InitAgent(convertAgentType ConvertAgentType, convertServiceUrl string) (C.YH_CONVERT_AGENT, error) {
	if globalDll == nil {
		return nil, errors.New("the convert dynamic library is not loaded")
	}
	proc, err := globalDll.GetProcAddress(initAgentFuncName)
	if err != nil {
		return nil, err
	}
	var convertAgentTypeInt C.int
	convertAgentTypeInt = C.int(convertAgentType)
	convertServiceUrlPointer := C.CString(convertServiceUrl)
	defer C.free(unsafe.Pointer(convertServiceUrlPointer))
	ret, _, _ := syscall.Syscall(proc, uintptr(2), uintptr(convertAgentTypeInt), uintptr(unsafe.Pointer(convertServiceUrlPointer)), 0)
	if ret == 0 {
		return nil, errors.New("init agent error")
	}
	agent.agentPtr = C.YH_CONVERT_AGENT(ret)
	return agent.agentPtr, nil
}

// FinalizeAgent 销毁代理.
func (agent *ConvertAgent) FinalizeAgent(agentPtr C.YH_CONVERT_AGENT) error {
	if globalDll == nil {
		return errors.New("the convert dynamic library is not loaded")
	}
	proc, err := globalDll.GetProcAddress(finalizeAgentFuncName)
	if err != nil {
		return err
	}
	_, _, _ = syscall.Syscall(proc, uintptr(1), uintptr(unsafe.Pointer(agentPtr)), 0, 0)
	return nil
}

func (agent *ConvertAgent) Finalize() error {
	var err error
	if err = agent.FinalizeAgent(agent.agentPtr); err != nil {
		return err
	}
	return nil
}

// OfficeToOFD 将单个办公文件（Office文件如doc等、版式文件如pdf、xps、ceb等）转换为OFD文件并可附加元数据和语义树.
func (agent *ConvertAgent) OfficeToOFD(srcFilePath, outFilePath string, metaData MetaData, semantics Semantics) error {
	if globalDll == nil {
		return errors.New("the convert dynamic library is not loaded")
	}
	if agent.agentPtr == nil {
		return errors.New("the agent is not initialized")
	}
	proc, err := globalDll.GetProcAddress(officeToOFDFuncName)
	if err != nil {
		return err
	}
	srcFilePathPointer := C.CString(srcFilePath)
	defer C.free(unsafe.Pointer(srcFilePathPointer))
	outFilePathPointer := C.CString(outFilePath)
	defer C.free(unsafe.Pointer(outFilePathPointer))
	metaDataStr, err := encodeData(metaData)
	if err != nil {
		return err
	}
	var metaDataStrPointer *C.char
	if len(metaDataStr) > 0 {
		metaDataStrPointer = C.CString(metaDataStr)
		defer C.free(unsafe.Pointer(metaDataStrPointer))
	}
	semanticsStr, err := encodeData(semantics)
	if err != nil {
		return err
	}
	var semanticsStrPointer *C.char
	if len(semanticsStr) > 0 {
		semanticsStrPointer = C.CString(semanticsStr)
		defer C.free(unsafe.Pointer(semanticsStrPointer))
	}
	ret, _, _ := syscall.Syscall6(proc, uintptr(5), uintptr(unsafe.Pointer(agent.agentPtr)), uintptr(unsafe.Pointer(srcFilePathPointer)),
		uintptr(unsafe.Pointer(outFilePathPointer)), uintptr(unsafe.Pointer(metaDataStrPointer)), uintptr(unsafe.Pointer(semanticsStrPointer)), 0)
	if ret != C.YH_OK {
		return errors.New("init agent error")
	}
	return nil
}
