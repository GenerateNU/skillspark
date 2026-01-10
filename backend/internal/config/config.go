package config

// Config holds the entire application configuration
type Config struct {
	Application Application
	DB          DB
	Supabase    Supabase
}
