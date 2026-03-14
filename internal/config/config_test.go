package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixtureDir(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok, "failed to resolve test file path")
	return filepath.Join(filepath.Dir(filename), "testdata")
}

func prepareConfig(t *testing.T, fixtureName string) string {
	t.Helper()

	src := filepath.Join(fixtureDir(t), fixtureName)
	data, err := os.ReadFile(src)
	require.NoError(t, err)

	dir := t.TempDir()
	err = os.WriteFile(filepath.Join(dir, FileName), data, 0o644)
	require.NoError(t, err)

	return dir
}

func TestDefault(t *testing.T) {
	cfg := Default()
	require.NotNil(t, cfg)

	assert.True(t, cfg.Rules.LowercaseLetter.IsEnabled())
	assert.True(t, cfg.Rules.EnglishLanguage.IsEnabled())
	assert.True(t, cfg.Rules.NoSpecialChars.IsEnabled())
	assert.True(t, cfg.Rules.NoSensitiveData.IsEnabled())
}

func TestLoadNoFile(t *testing.T) {
	dir := t.TempDir()

	cfg, err := Load(dir)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	assert.True(t, cfg.Rules.LowercaseLetter.IsEnabled())
}

func TestLoadDisableRules(t *testing.T) {
	dir := prepareConfig(t, "disable_rules.yml")

	cfg, err := Load(dir)
	require.NoError(t, err)

	assert.False(t, cfg.Rules.LowercaseLetter.IsEnabled())
	assert.False(t, cfg.Rules.EnglishLanguage.IsEnabled())
	assert.True(t, cfg.Rules.NoSpecialChars.IsEnabled())
	assert.True(t, cfg.Rules.NoSensitiveData.IsEnabled())
}

func TestLoadAllDisabled(t *testing.T) {
	dir := prepareConfig(t, "all_disabled.yml")

	cfg, err := Load(dir)
	require.NoError(t, err)

	assert.False(t, cfg.Rules.LowercaseLetter.IsEnabled())
	assert.False(t, cfg.Rules.EnglishLanguage.IsEnabled())
	assert.False(t, cfg.Rules.NoSpecialChars.IsEnabled())
	assert.False(t, cfg.Rules.NoSensitiveData.IsEnabled())
}

func TestLoadExtraPatterns(t *testing.T) {
	dir := prepareConfig(t, "extra_patterns.yml")

	cfg, err := Load(dir)
	require.NoError(t, err)

	require.Len(t, cfg.Rules.NoSensitiveData.ExtraPatterns, 2)
	assert.Equal(t, "ssn", cfg.Rules.NoSensitiveData.ExtraPatterns[0])
	assert.Equal(t, "credit_card", cfg.Rules.NoSensitiveData.ExtraPatterns[1])
	assert.True(t, cfg.Rules.NoSensitiveData.IsEnabled())
}

func TestLoadInvalidYAML(t *testing.T) {
	dir := prepareConfig(t, "invalid.yml")

	_, err := Load(dir)
	assert.Error(t, err)
}

func TestFindConfigWalksUp(t *testing.T) {
	root := t.TempDir()
	child := filepath.Join(root, "a", "b", "c")
	require.NoError(t, os.MkdirAll(child, 0o755))

	cfgPath := filepath.Join(root, FileName)
	require.NoError(t, os.WriteFile(cfgPath, []byte("rules:\n"), 0o644))

	found, ok := findConfig(child)
	assert.True(t, ok)
	assert.Equal(t, cfgPath, found)
}
