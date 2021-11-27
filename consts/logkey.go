package consts

type LogKey string

const (
	LogId     LogKey = "log_id"
	LogLevel  LogKey = "log_level"
	SourcePsm LogKey = "source_psm"
	Psm       LogKey = "psm"
	PodName   LogKey = "pod_name"
)
