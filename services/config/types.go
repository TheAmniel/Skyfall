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
		TimeZone      string `toml:"timeZone"`
		Limit         int    `toml:"limit"`
		Prefork       bool   `toml:"prefork"`
		StrictRouting bool   `toml:"strictRouting"`
		CaseSensitive bool   `toml:"caseSensitive"`
		UnescapePath  bool   `toml:"unescapePath"`
	}

	DatabaseConfig struct {
		Type                   string `toml:"type"`
		Name                   string `toml:"name"`
		PrepareStmt            bool   `toml:"prepareStmt"`
		SkipDefaultTransaction bool   `toml:"skipDefaultTransaction"`
	}

	MiddlewareConfig struct {
		Cache     bool `toml:"cache"`
		Compress  bool `toml:"compress"`
		Logger    bool `toml:"logger"`
		Recover   bool `toml:"recover"`
		Shortener bool `toml:"shortener"`
		Banned    bool `toml:"banned"`
		Traffic   bool `toml:"traffic"`
	}
)
