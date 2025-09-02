package openroutergo

type OpenRouterOptionalConfig struct {
	SiteUrl  string
	SiteName string
}

type OpenRouterConfig struct {
	key string
	OpenRouterOptionalConfig
}

func (c *OpenRouterConfig) WithKey(key string) *OpenRouterConfig {
	c.key = key

	return c
}
