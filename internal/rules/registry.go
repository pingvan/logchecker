package rules

// AllRules contains all available rules with default settings.
var AllRules = []Rule{
	&LowercaseLetterRule,
	&EnglishLanguageRule,
	&NoSpecialCharsRule,
	&NoSensitiveDataRule,
}
