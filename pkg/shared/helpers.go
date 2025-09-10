package shared

import "regexp"

// Obfuscate credentials in URLs of the form //user:password@
func SanitizeConnectionString(str string) string {
	re := regexp.MustCompile(`(//[^:/@]+:)[^@]+(@)`)
	return re.ReplaceAllString(str, `$1***$2`)
}
