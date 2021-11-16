package env

import (
	"env/consts"
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
