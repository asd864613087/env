package env

import (
	"context"
	"github.com/asd864613087/env/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GetPsm() string {
	return os.Getenv(consts.ENV_PSM)
}

func GetNamespace() string {
	return os.Getenv(consts.ENV_NAMESPACE)
}

func GetPodName() string {
	return os.Getenv(consts.ENV_PODNAME)
}

func GetCtxSourcePsm(ctx context.Context) string {
	v := ctx.Value(consts.CTX_SOURCE_PSM)
	return v.(string)
}

func GetCtxPsm(ctx context.Context) string {
	v := ctx.Value(consts.CTX_PSM)
	return v.(string)
}

func GetCtxLogId(ctx context.Context) string {
	v := ctx.Value(consts.CTX_LOG_ID)
	return v.(string)
}

func GetCtxLogLevel(ctx context.Context) string {
	v := ctx.Value(consts.CTX_LOG_LEVEL)
	return v.(string)
}

func SetHeaderSourcePsm(req *http.Request, sourcePsm string) *http.Request {
	req.Header.Set(consts.HEADER_SOURCE_PSM, sourcePsm)
	return req
}

func SetHeaderPsm(req *http.Request, Psm string) *http.Request {
	req.Header.Set(consts.HEADER_PSM, Psm)
	return req
}

func SetHeaderDestPsm(req *http.Request, destPsm string) *http.Request {
	req.Header.Set(consts.HEADER_DEST_PSM, destPsm)
	return req
}

func SetHeaderLogId(req *http.Request, logId string) *http.Request {
	req.Header.Set(consts.HEADER_LOG_ID, logId)
	return req
}

func SetHeaderLogLevel(req *http.Request, logLevel string) *http.Request {
	req.Header.Set(consts.HEADER_LOG_LEVEL, logLevel)
	return req
}

func SetHeaderNamespace(req *http.Request, namespace string) *http.Request {
	req.Header.Set(consts.HEADER_NAMESPACE, namespace)
	return req
}

func SetHeaderPodName(req *http.Request, podName string) *http.Request {
	req.Header.Set(consts.HEADER_POD_NAME, podName)
	return req
}

func SetGinCtxPsm(c *gin.Context, psm string) *gin.Context {
	c.Set(consts.CTX_PSM, psm)
	return c
}

func SetGinCtxSourcePsm(c *gin.Context, sourcePsm string ) *gin.Context {
	c.Set(consts.CTX_SOURCE_PSM, sourcePsm)
	return c
}

func SetGinCtxLogId(c *gin.Context, logId string) *gin.Context {
	c.Set(consts.CTX_LOG_ID, logId)
	return c
}

func SetGinCtxLogLevel(c *gin.Context, logLevel string) *gin.Context {
	c.Set(consts.CTX_LOG_LEVEL, logLevel)
	return c
}

