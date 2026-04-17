package config

type OpenSearch struct {
	URL      string `env:"OPENSEARCH_URL"`
	Username string `env:"OPENSEARCH_USER"`
	Password string `env:"OPENSEARCH_PASS"`
}
