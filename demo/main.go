package main

import (
	"flag"
	"fmt"
	"yh_sdk"
)

var (
	convertDllFilePath string
	convertServiceUrl  string
	srcFilePath        string
	outFilePath        string
)

func init() {
	flag.StringVar(&convertDllFilePath, "d", "", "convert dll file path")
	flag.StringVar(&convertServiceUrl, "u", "", "convert service url")
	flag.StringVar(&srcFilePath, "i", "", "source file path")
	flag.StringVar(&outFilePath, "o", "", "output file path")
}

func main() {
	flag.Parse()
	agent, err := yh_sdk.NewHttpAgent(convertDllFilePath, convertServiceUrl)
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
	fmt.Println("convert successfully")
}
