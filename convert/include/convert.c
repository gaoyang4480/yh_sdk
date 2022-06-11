#include "convert.h"
#include <dlfcn.h>

///////////////////////////////////////////////////////////////////////////////////
//
// 动态库C接口函数名
//
///////////////////////////////////////////////////////////////////////////////////
#define YH_INITSDK_FUNC_NAME "YH_InitSDK"
#define YH_FINALIZESDK_FUNC_NAME "YH_FinalizeSDK"
#define YH_INITAGENT_FUNC_NAME "YH_InitAgent"
#define YH_FINALIZEAGENT_FUNC_NAME "YH_FinalizeAgent"
#define YH_OFFICETOOFD_FUNC_NAME "YH_OfficeToOFD"

///////////////////////////////////////////////////////////////////////////////////
//
// 动态库C接口函数声明
//
///////////////////////////////////////////////////////////////////////////////////
typedef YH_STATUS (*YH_INITSDK_FUNC)();

typedef void (*YH_FINALIZESDK_FUNC)();

typedef YH_CONVERT_AGENT (*YH_INITAGENT_FUNC)(int convertAgentType, const char *baseUrl);

typedef void (*YH_FINALIZEAGENT_FUNC)(YH_CONVERT_AGENT convertAgent);

typedef YH_STATUS (*YH_OFFICETOOFD_FUNC)(YH_CONVERT_AGENT convertAgent, const char *srcFilePath, const char *outFilePath,
                   const char *metasStr, const char *semanticsStr);

YHStatus GetFuncAddress(YHHModule_t hModule, const char *procName, YHHandleSym_t *ressym) {
    char* error = NULL;
    *ressym = (YHHandleSym_t)dlsym(hModule, procName);
    if((error = dlerror()) != NULL)
    {
        return -1;
    }
    return 0;
}

YHHModule_t LoadConvert(const char *oesDllFilePath) {
    hModule = dlopen(oesDllFilePath, RTLD_LAZY);
    if (!hModule) {
        return NULL;
    }
    return hModule;
}

void UnloadConvert(YHHModule_t hModule) {
    if (NULL != hModule) {
        dlclose(hModule);
    }
}

YH_STATUS InitSDK(YHHModule_t hModule) {
    YH_STATUS status = YH_OK;
    YHHandleSym_t ressym = NULL;
    YH_INITSDK_FUNC initSDKFunc = NULL;

    do {
        if (NULL == hModule) {
            status = OES_CANCEL;
            break;
        }

        if (GetFuncAddress(hModule, YH_INITSDK_FUNC_NAME, &ressym)) {
            status = OES_NO_FUNC;
            break;
        }

        initSDKFunc = (YH_INITSDK_FUNC) ressym;
        if (initSDKFunc == NULL) {
            status = OES_NO_FUNC;
            break;
        }

        status = initSDKFunc();
        if (status != YH_OK) {
            break;
        }
    } while (0);

    return status;
}

void FinalizeSDK(YHHModule_t hModule) {
    YHHandleSym_t ressym = NULL;
    YH_FINALIZESDK_FUNC finalizeSDKFunc = NULL;

    do {
        if (NULL == hModule) {
            break;
        }

        if (GetFuncAddress(hModule, YH_FINALIZESDK_FUNC_NAME, &ressym)) {
            break;
        }

        finalizeSDKFunc = (YH_FINALIZESDK_FUNC) ressym;
        if (finalizeSDKFunc == NULL) {
            break;
        }

        finalizeSDKFunc();
    } while (0);
}

YH_CONVERT_AGENT InitAgent(YHHModule_t hModule, int convertAgentType, const char *baseUrl) {
    YH_CONVERT_AGENT agentPtr = NULL;
    YHHandleSym_t ressym = NULL;
    YH_INITAGENT_FUNC initAgentFunc = NULL;

    do {
        if (NULL == hModule) {
            break;
        }

        if (GetFuncAddress(hModule, YH_INITAGENT_FUNC_NAME, &ressym)) {
            break;
        }

        initAgentFunc = (YH_INITAGENT_FUNC) ressym;
        if (initAgentFunc == NULL) {
            break;
        }

        agentPtr = initAgentFunc();
    } while (0);

    return agentPtr;
}

void FinalizeAgent(YHHModule_t hModule, YH_CONVERT_AGENT convertAgent) {
    YHHandleSym_t ressym = NULL;
    YH_FINALIZEAGENT_FUNC finalizeAgentFunc = NULL;

    do {
        if (NULL == hModule) {
            break;
        }

        if (GetFuncAddress(hModule, YH_FINALIZEAGENT_FUNC_NAME, &ressym)) {
            break;
        }

        finalizeAgentFunc = (YH_FINALIZEAGENT_FUNC) ressym;
        if (finalizeAgentFunc == NULL) {
            break;
        }

        finalizeAgentFunc(convertAgent);
    } while (0);
}

YH_STATUS OfficeToOFD(YHHModule_t hModule, YH_CONVERT_AGENT convertAgent, const char *srcFilePath, const char *outFilePath,
                          const char *metasStr, const char *semanticsStr) {
    YH_STATUS status = YH_OK;
    YHHandleSym_t ressym = NULL;
    YH_OFFICETOOFD_FUNC officeToOFDFunc = NULL;

    do {
        if (NULL == hModule) {
            status = YH_ERROR;
            break;
        }

        if (GetFuncAddress(hModule, YH_OFFICETOOFD_FUNC_NAME, &ressym)) {
            status = YH_ERROR;
            break;
        }

        officeToOFDFunc = (YH_OFFICETOOFD_FUNC) ressym;
        if (officeToOFDFunc == NULL) {
            status = YH_ERROR;
            break;
        }

        status = officeToOFDFunc(convertAgent, srcFilePath, outFilePath, metasStr, semanticsStr);
        if (status != YH_OK) {
            break;
        }
    } while (0);

    return status;
}