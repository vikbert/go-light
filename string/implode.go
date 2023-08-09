package string

import (
	"fmt"
	"strings"
)

// Implode Join array elements with a separator string
func Implode(elements []interface{}, separator string) string {
	size := len(elements)
	stringArray := make([]string, size)

	for i := 0; i < size; i++ {
		stringArray[i] = fmt.Sprintf("%v", elements[i])
	}

	return strings.Join(stringArray, separator)
}
