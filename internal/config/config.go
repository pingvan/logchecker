package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// FileName is the name of the logchecker configuration file.
const FileName = ".logchecker.yml"

// Config represents the top-level logchecker configuration.
type Config struct {
	Rules RulesConfig `yaml:"rules"`
}

// RulesConfig holds per-rule configuration.
type RulesConfig struct {
	LowercaseLetter *RuleConfig          `yaml:"lowercase-letter"`
	EnglishLanguage *RuleConfig          `yaml:"english-language"`
	NoSpecialChars  *RuleConfig          `yaml:"no-special-chars"`
	NoSensitiveData *SensitiveDataConfig `yaml:"no-sensitive-data"`
}

// RuleConfig controls whether a rule is enabled.
type RuleConfig struct {
	Enabled *bool `yaml:"enabled"`
}

// SensitiveDataConfig extends RuleConfig with extra sensitive data patterns.
type SensitiveDataConfig struct {
	Enabled       *bool    `yaml:"enabled"`
	ExtraPatterns []string `yaml:"extra-patterns"`
}

// IsEnabled returns whether the rule is enabled (defaults to true).
func (rc *RuleConfig) IsEnabled() bool {
	if rc == nil || rc.Enabled == nil {
		return true
	}
	return *rc.Enabled
}

// IsEnabled returns whether the rule is enabled (defaults to true).
func (sc *SensitiveDataConfig) IsEnabled() bool {
	if sc == nil || sc.Enabled == nil {
		return true
	}
	return *sc.Enabled
}

// Default returns a config with all rules enabled and default settings.
func Default() *Config {
	return &Config{}
}

// Load reads the configuration file by walking up from dir. Returns Default() if not found.
func Load(dir string) (*Config, error) {
	path, found := findConfig(dir)
	if !found {
		return Default(), nil
	}

	path = filepath.Clean(path)
	data, err := os.ReadFile(path) //nolint:gosec // path is constructed internally via findConfig, not from user input
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
