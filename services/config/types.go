package config

type (
	Config struct {
		Server     *ServerConfig     `toml:"server"`
		Database   *DatabaseConfig   `toml:"database"`
		Middleware *MiddlewareConfig `toml:"middleware"`
	}

	ServerConfig struct {
		Host          string `toml:"host"`
		Port          string `toml:"port"`
		Secret        string `toml:"secret"`
		Limit         int    `toml:"limit"`
		Prefork       bool   `toml:"prefork"`
		StrictRouting bool   `toml:"strict-routing"`
		CaseSensitive bool   `toml:"case-sensitive"`
		UnescapePath  bool   `toml:"unescape-path"`
	}

	DatabaseConfig struct {
		Type                   string `toml:"type"`
		Name                   string `toml:"name"`
		PrepareStmt            bool   `toml:"prepare-stmt"`
		SkipDefaultTransaction bool   `toml:"skip-default-transaction"`
	}

	MiddlewareConfig struct {
		Cache    bool `toml:"cache"`
		Compress bool `toml:"compress"`
		Logger   bool `toml:"logger"`
		Recover  bool `toml:"recover"`
		Banned   bool `toml:"banned"`
	}
)
