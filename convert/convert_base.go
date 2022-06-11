// +build linux darwin

package convert

/*
#cgo linux LDFLAGS: -ldl
#cgo CFLAGS: -I ./include
#cgo CPPFLAGS: -I ./include
#include "convert.c"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// ConvertAgent 转换代理.
type ConvertAgent struct {
	agentPtr C.YH_CONVERT_AGENT
}

var (
	globalDll C.YHHModule_t
)

var _ IConvertAgent = (*ConvertAgent)(nil)

// LoadConvert 加载转换动态库.
func LoadConvert(convertDllFilePath string) error {
	if globalDll != nil {
		return nil
	}
	if len(convertDllFilePath) == 0 {
		return errors.New("convert dll file path is empty")
	}
	convertDllFilePathPointer := C.CString(convertDllFilePath)
	defer C.free(unsafe.Pointer(convertDllFilePathPointer))
	globalDll = C.LoadConvert((*C.char)(convertDllFilePathPointer))
	if globalDll == nil {
		return errors.New("load convert dll error")
	}
	return nil
}

// UnloadConvert 卸载加载转换动态库.
func UnloadConvert() {
	if globalDll != nil {
		C.UnloadConvert(globalDll)
	}
}

// InitSDK 初始化.
func (agent *ConvertAgent) InitSDK() error {
	if globalDll == nil {
		return errors.New("the convert dynamic library is not loaded")
	}
	ret := C.InitSDK(globalDll)
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
	C.FinalizeSDK(globalDll)
	return nil
}

// InitAgent 初始化代理.
func (agent *ConvertAgent) InitAgent(convertAgentType ConvertAgentType, convertServiceUrl string) (C.YH_CONVERT_AGENT, error) {
	if globalDll == nil {
		return nil, errors.New("the convert dynamic library is not loaded")
	}
	var convertAgentTypeInt C.int
	convertAgentTypeInt = C.int(convertAgentType)
	convertServiceUrlPointer := C.CString(convertServiceUrl)
	defer C.free(unsafe.Pointer(convertServiceUrlPointer))
	agentPtr := C.InitAgent(globalDll, convertAgentTypeInt, (*C.char)(convertServiceUrlPointer))
	if agentPtr == nil {
		return nil, errors.New("init agent error")
	}
	agent.agentPtr = agentPtr
	return agent.agentPtr, nil
}

// FinalizeAgent 销毁代理.
func (agent *ConvertAgent) FinalizeAgent(agentPtr C.YH_CONVERT_AGENT) error {
	if globalDll == nil {
		return errors.New("the convert dynamic library is not loaded")
	}
	C.FinalizeAgent(globalDll, agentPtr)
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
	ret := C.OfficeToOFD(globalDll, agent.agentPtr, srcFilePathPointer, outFilePathPointer,
		metaDataStrPointer, semanticsStrPointer)
	if ret != C.YH_OK {
		return errors.New("init agent error")
	}
	return nil
}
