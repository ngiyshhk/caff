package utils

import (
	"regexp"
	"strings"
)

// UpperInitial .
func UpperInitial(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// ToSnakeCase .
func ToSnakeCase(s string) string {
	result := s
	for _, pattern := range []string{"(.)([A-Z][^A-Z_]+)", "([^A-Z_])([A-Z])"} {
		rgp := regexp.MustCompile(pattern)
		result = rgp.ReplaceAllString(result, "${1}_${2}")
	}
	return strings.ToLower(result)
}

// ToCamelCase .
func ToCamelCase(s string) string {
	rgp := regexp.MustCompile("(.+?)(_+|$)")
	return rgp.ReplaceAllStringFunc(s, func(t string) string {
		tmp := strings.ReplaceAll(t, "_", "")
		switch tmp {
		case "acl", "api", "ascii", "cpu", "css", "dns", "eof", "guid", "html", "http",
			"https", "id", "ip", "json", "lhs", "qps", "ram", "rhs", "rpc", "sla", "smtp",
			"sql", "ssh", "tcp", "tls", "ttl", "udp", "ui", "uid", "uuid", "uri", "url",
			"utf8", "vm", "xml", "xmpp", "xsrf", "xss":
			return strings.ToUpper(tmp)
		default:
			return UpperInitial(strings.ReplaceAll(t, "_", ""))
		}
	})
}
