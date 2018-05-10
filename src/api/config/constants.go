package config

const (
	GENERIC_FAILURE_MESSAGE = "Something went wrong..."

	APP_NAME = "jwt-go-api"

	LOG_FORMAT = `${time_custom} → id=${id} ` +
		`status=${status} ` +
		`host=${host} path=${path} ` +
		`latency=${latency_human} ` +
		"\n\n"
)
