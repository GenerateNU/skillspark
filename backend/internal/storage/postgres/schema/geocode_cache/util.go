package geocode_cache

import "embed"

//go:embed sql/*.sql
var SqlGeocodeCacheFiles embed.FS
