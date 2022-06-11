#ifndef YH_CONVERT_H
#define YH_CONVERT_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
//#include <dlfcn.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void*			HMODULE;
typedef HMODULE			YHHModule_t;
typedef int				YHStatus;
typedef HMODULE			YHHandleSym_t;

typedef void *YH_CONVERT_AGENT;

#define YH_OK 0
#define YH_ERROR 1
typedef int YH_STATUS;

extern YHHModule_t LoadConvert(const char *oesDllFilePath);

extern void UnloadConvert(YHHModule_t hModule);

extern YH_STATUS InitSDK();

extern void FinalizeSDK();

extern YH_CONVERT_AGENT InitAgent(int convertAgentType, const char *baseUrl);

extern void FinalizeAgent(YH_CONVERT_AGENT convertAgent);

extern YH_STATUS OfficeToOFD(YH_CONVERT_AGENT convertAgent, const char *srcFilePath, const char *outFilePath,
                          const char *metasStr, const char *semanticsStr);

#ifdef __cplusplus
}
#endif

#endif // OES_H