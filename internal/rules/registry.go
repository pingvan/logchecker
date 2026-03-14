package rules

var AllRules = []Rule{
	&LowercaseLetterRule,
	&EnglishLanguageRule,
	&NoSpecialCharsRule,
	&NoSensitiveDataRule,
}
