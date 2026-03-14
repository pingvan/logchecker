package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const FileName = ".logchecker.yml"

type Config struct {
	Rules RulesConfig `yaml:"rules"`
}

type RulesConfig struct {
	LowercaseLetter *RuleConfig         `yaml:"lowercase-letter"`
	EnglishLanguage *RuleConfig         `yaml:"english-language"`
	NoSpecialChars  *RuleConfig         `yaml:"no-special-chars"`
	NoSensitiveData *SensitiveDataConfig `yaml:"no-sensitive-data"`
}

type RuleConfig struct {
	Enabled *bool `yaml:"enabled"`
}

type SensitiveDataConfig struct {
	Enabled       *bool    `yaml:"enabled"`
	ExtraPatterns []string `yaml:"extra-patterns"`
}

func (rc *RuleConfig) IsEnabled() bool {
	if rc == nil || rc.Enabled == nil {
		return true
	}
	return *rc.Enabled
}

func (sc *SensitiveDataConfig) IsEnabled() bool {
	if sc == nil || sc.Enabled == nil {
		return true
	}
	return *sc.Enabled
}

func Default() *Config {
	return &Config{}
}

func Load(dir string) (*Config, error) {
	path, found := findConfig(dir)
	if !found {
		return Default(), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}

	return &cfg, nil
}

func findConfig(dir string) (string, bool) {
	for {
		path := filepath.Join(dir, FileName)
		if _, err := os.Stat(path); err == nil {
			return path, true
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}
