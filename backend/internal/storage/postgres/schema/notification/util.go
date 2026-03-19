package notification

import (
	"embed"
)

//go:embed sql/*.sql
var SqlNotificationFiles embed.FS

