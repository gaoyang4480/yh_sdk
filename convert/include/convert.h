#ifndef YH_CONVERT_H
#define YH_CONVERT_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

///////////////////////////////////////////////////////////////////////////////////
//
// 通用宏定义
//
///////////////////////////////////////////////////////////////////////////////////
typedef void*			HMODULE;
typedef HMODULE			YHHModule_t;
typedef int				YHStatus;
typedef HMODULE			YHHandleSym_t;

typedef void *YH_CONVERT_AGENT;

#define YH_OK 0
#define YH_ERROR 1
typedef int YH_STATUS;

///////////////////////////////////////////////////////////////////////////////////
//
// 接口声明
//
///////////////////////////////////////////////////////////////////////////////////
/// 加载转换动态库.
extern YHHModule_t LoadConvert(const char *oesDllFilePath);

/// 卸载转换动态库.
extern void UnloadConvert(YHHModule_t hModule);

/// 初始化.
extern YH_STATUS InitSDK(YHHModule_t hModule);

/// 销毁.
extern void FinalizeSDK(YHHModule_t hModule);

/// 初始化代理.
extern YH_CONVERT_AGENT InitAgent(YHHModule_t hModule, int convertAgentType, const char *baseUrl);

/// 销毁代理.
extern void FinalizeAgent(YHHModule_t hModule, YH_CONVERT_AGENT convertAgent);

/// 将单个办公文件（Office文件如doc等、版式文件如pdf、xps、ceb等）转换为OFD文件并可附加元数据和语义树.
extern YH_STATUS OfficeToOFD(YHHModule_t hModule, YH_CONVERT_AGENT convertAgent, const char *srcFilePath, const char *outFilePath,
                          const char *metasStr, const char *semanticsStr);

#ifdef __cplusplus
}
#endif

#endif // YH_CONVERT_H