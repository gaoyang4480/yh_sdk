package yh_sdk

import "yh_sdk/convert"

// IConvertAgent 转换代理接口.
type IConvertAgent interface {
	// OfficeToOFD 将单个办公文件（Office文件如doc等、版式文件如pdf、xps、ceb等）转换为OFD文件并可附加元数据和语义树.
	OfficeToOFD(srcFilePath, outFilePath string, metaData convert.MetaData, semantics convert.Semantics) error
	// FinalizeAgent 销毁代理.
	FinalizeAgent() error
}

// HttpAgent Http转换代理.
type HttpAgent struct {
	// 转换代理.
	convertAgent convert.IConvertAgent
}

var _ IConvertAgent = (*HttpAgent)(nil)

// InitSDK 初始化SDK.
func InitSDK(convertDllFilePath string) error {
	return convert.LoadConvert(convertDllFilePath)
}

// FinalizeSDK 销毁SDK.
func FinalizeSDK() {
	convert.UnloadConvert()
}

// NewHttpAgent 新建Http转换代理.
func NewHttpAgent(convertServiceUrl string) (IConvertAgent, error) {
	httpAgent := &HttpAgent{}
	var err error
	httpAgent.convertAgent, err = convert.NewConvertAgent(convert.AgentTypeHttp, convertServiceUrl)
	if err != nil {
		return nil, err
	}
	return httpAgent, nil
}

// OfficeToOFD 将单个办公文件（Office文件如doc等、版式文件如pdf、xps、ceb等）转换为OFD文件并可附加元数据和语义树.
func (agent *HttpAgent) OfficeToOFD(srcFilePath, outFilePath string, metaData convert.MetaData, semantics convert.Semantics) error {
	return agent.convertAgent.OfficeToOFD(srcFilePath, outFilePath, metaData, semantics)
}

// FinalizeAgent 销毁代理.
func (agent *HttpAgent) FinalizeAgent() error {
	if agent.convertAgent != nil {
		return agent.convertAgent.Finalize()
	}
	return nil
}
