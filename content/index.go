package content

import (
	"fmt"
)

func GenerateIndexContent(name string) string {
	return fmt.Sprintf(`import %s from './modules';

export default %s;
`, name, name)
}
