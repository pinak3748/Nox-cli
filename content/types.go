package content

import (
	"strings"
)

func GenerateTypesContent(name string) string {
	smallCasePageName := strings.ToLower(name)
	template := `export interface {smallCasePageName}Item {
  id: string;
  name: string;
  description: string;
}
`
	return strings.ReplaceAll(template, "{smallCasePageName}", smallCasePageName)
}
