package yh_sdk

import (
	"fmt"
	"testing"
	"yh_sdk/convert"
)

const (
	convertDllFilePath = ""
	convertServiceUrl  = "http://1.202.28.218:9000/v1/"
	srcFilePath        = ""
	outFilePath        = ""
)

func init() {
	if err := InitSDK(convertDllFilePath); err != nil {
		panic(err)
	}
}

func TestOfficeToOFD(t *testing.T) {
	agent, err := NewHttpAgent(convertServiceUrl)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = agent.FinalizeAgent()
	}()
	err = agent.OfficeToOFD(srcFilePath, outFilePath, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("convert office file to ofd successfully")
}

func TestOfficeToOFDWithMetaData(t *testing.T) {
	agent, err := NewHttpAgent(convertServiceUrl)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = agent.FinalizeAgent()
	}()
	// 元数据.
	metaData := make(convert.MetaData)
	metaData["title"] = "测试标题"
	err = agent.OfficeToOFD(srcFilePath, outFilePath, metaData, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("convert office file to ofd with metadata successfully")
}
