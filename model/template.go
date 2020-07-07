package model

// Template .
type Template struct {
	Type            TemplateType
	Value           string
	PackageNameFunc func(*Schema) string
}
