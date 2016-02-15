package filename

import (
	"strings"
)

func Getfix(a string) string {

	n := strings.LastIndex(a, ".")
	if n != 0 && n != -1 {

		return a[strings.LastIndex(a, ".")+1:]

	}
	return ""

}
