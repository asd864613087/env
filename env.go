package env

import (
	"context"
	"github.com/asd864613087/env/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

/*
	获取环境变量
 */
func GetPsm() string {
	return os.Getenv(consts.ENV_PSM)
}

func GetNamespace() string {
	return os.Getenv(consts.ENV_NAMESPACE)
}

func GetPodName() string {
	return os.Getenv(consts.ENV_PODNAME)
}

/*
	从Ctx中获取变量
*/

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

/*
	设置请求下游时Header的KV，ginex的rawcall使用
*/

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

/*
	从Header中获取参数，ginex Log中间件使用
 */
func GetHeaderSourcePsmFromGinCtx(c *gin.Context) string{
	v := c.GetHeader(consts.HEADER_SOURCE_PSM)
	return v
}

func GetHeaderLogLevelFromGinCtx(c *gin.Context) string{
	v := c.GetHeader(consts.HEADER_LOG_LEVEL)
	return v
}

func GetHeaderLogIdFromGinCtx(c *gin.Context) string{
	v := c.GetHeader(consts.HEADER_LOG_ID)
	return v
}


/*
	往gin.Context中设置参数，ginex的log中间件使用
 */

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

/*
	Metrics 的 constLabel
 */

func GetConstLabels() map[string]string {
	return map[string]string{
		consts.METRICS_PSM: GetPsm(),
		consts.METRICS_NAMESPACE: GetNamespace(),
		consts.METRICS_POD_NAME: GetPodName(),
	}
}
