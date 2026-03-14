package rules


import "github.com/pingvan/logchecker/internal/config"

// AllRules contains all available rules with default settings.
var AllRules = []Rule{
	&LowercaseLetterRule,
	&EnglishLanguageRule,
	&NoSpecialCharsRule,
	&NoSensitiveDataRule,
}

func FromConfig(cfg *config.Config) []Rule {
	if cfg == nil {
		return AllRules
	}

	var result []Rule

	if cfg.Rules.LowercaseLetter.IsEnabled() {
		result = append(result, &LowercaseLetterRule)
	}
	if cfg.Rules.EnglishLanguage.IsEnabled() {
		result = append(result, &EnglishLanguageRule)
	}
	if cfg.Rules.NoSpecialChars.IsEnabled() {
		result = append(result, &NoSpecialCharsRule)
	}
	if cfg.Rules.NoSensitiveData.IsEnabled() {
		extra := cfg.Rules.NoSensitiveData.ExtraPatterns
		if len(extra) > 0 {
			result = append(result, NewNoSensitiveDataRule(extra))
		} else {
			result = append(result, &NoSensitiveDataRule)
		}
	}

	return result
}
