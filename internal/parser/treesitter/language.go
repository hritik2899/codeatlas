package treesitter

import "fmt"

// Language identifies a source grammar supported by the ingestion pipeline.
// Keeping this registry small makes language support explicit and testable.
type Language string

const (
	LanguageGo Language = "go"
)

// LanguageForExtension maps repository file extensions to CodeAtlas languages.
func LanguageForExtension(extension string) (Language, error) {
	switch extension {
	case ".go":
		return LanguageGo, nil
	default:
		return "", fmt.Errorf("unsupported source extension %q", extension)
	}
}
