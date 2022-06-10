// +build linux darwin

package convert

/*
#cgo linux LDFLAGS: -ldl
#cgo CFLAGS: -I ./include
#cgo CPPFLAGS: -I ./include
#include "convert.h"
*/
import "C"

// ConvertAgent 转换代理.
type ConvertAgent struct {
	hModule C.YHHModule_t
}
