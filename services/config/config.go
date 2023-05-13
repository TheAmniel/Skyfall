package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"skyfall/utils"
)

func Load() *Config {
	dir, _, err := utils.Executable()
	if err != nil {
		panic(err)
	}
	cfg, err := LoadFile(dir + utils.ConfigFile)
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadFile(fileName string) (*Config, error) {
	if fileName[len(fileName)-5:] != ".toml" {
		fileName += ".toml"
	}

	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	if data, err := os.ReadFile(fileName); err != nil {
		return nil, err
	} else {
		data = utils.ReplaceValues(data)
		var cfg Config
		if err = toml.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
		return &cfg, nil
	}
}
