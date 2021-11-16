package env

import (
	"context"
	"github.com/asd864613087/env/consts"
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
