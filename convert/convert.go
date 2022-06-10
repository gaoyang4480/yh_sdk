package convert

/*
#cgo CFLAGS: -I ./include
#cgo CPPFLAGS: -I ./include
#include "convert.h"
*/
import "C"

type ConvertAgentType int

const (
	AgentTypeHttp = iota
	AgentTypeMQ
)

// IConvertAgent 转换代理接口.
type IConvertAgent interface {
	// LoadConvert 加载转换动态库.
	LoadConvert(convertDllFilePath string) error
	// UnloadConvert 卸载加载转换动态库.
	UnloadConvert()
	// InitSDK 初始化.
	InitSDK() error
	// FinalizeSDK 销毁.
	FinalizeSDK() error
	// InitAgent 初始化代理.
	InitAgent(convertAgentType ConvertAgentType, convertServiceUrl string) (C.YH_CONVERT_AGENT, error)
	// FinalizeAgent 销毁代理.
	FinalizeAgent(agentPtr C.YH_CONVERT_AGENT) error
	// Finalize 销毁.
	Finalize() error
	// OfficeToOFD 将单个办公文件（Office文件如doc等、版式文件如pdf、xps、ceb等）转换为OFD文件并可附加元数据和语义树.
	OfficeToOFD(srcFilePath, outFilePath string, metaData MetaData, semantics Semantics) error
}

// NewConvertAgent 创建代理.
func NewConvertAgent(convertAgentType ConvertAgentType, convertServiceUrl, convertDllFilePath string) (IConvertAgent, error) {
	convertAgent := &ConvertAgent{}
	err := convertAgent.LoadConvert(convertDllFilePath)
	if err != nil {
		return nil, err
	}
	err = convertAgent.InitSDK()
	if err != nil {
		return nil, err
	}
	_, err = convertAgent.InitAgent(convertAgentType, convertServiceUrl)
	if err != nil {
		return nil, err
	}
	return convertAgent, nil
}
