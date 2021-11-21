package consts

const (
	// Standard For Env Name
	ENV_PSM       = "PSM"
	ENV_NAMESPACE = "NAMESPACE"
	ENV_PODNAME   = "HOSTNAME"
	ENV_PODIP     = "POD_IP"
	ENV_NODENAME  = "POD_NODE_NAME"

	// Standard For Http Params Header Name
	HEADER_SOURCE_PSM = "Paas-Source-Psm"
	HEADER_PSM        = "Paas-Psm"
	HEADER_DEST_PSM   = "Paas-Dest-Psm"
	HEADER_NAMESPACE  = "Paas-Name-Space"
	HEADER_POD_NAME   = "Paas-Pod-Name"
	HEADER_LOG_ID     = "Paas-Log-Id"
	HEADER_LOG_LEVEL  = "Paas-Log-Level"

	// Standard For Context Key Name
	CTX_SOURCE_PSM = "source_psm"
	CTX_PSM        = "psm"
	CTX_DEST_PSM   = "dest_psm"
	CTX_LOG_ID     = "log_id"
	CTX_LOG_LEVEL  = "log_level"

	// Standard For Metrics Label Name
	METRICS_PSM       = "psm"
	METRICS_POD_NAME  = "pod_name"
	METRICS_NAMESPACE = "namespace"
)
