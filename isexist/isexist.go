package isexist

import (
	"os"
)

// Determines whether a file or directory exists
// 1 is existence
// 0 is nonexistence
// -1 is unkown error
func IsExist(path string) (exist int8) {
	_, err := os.Stat(path)
	if err == nil {
		exist = 1
	} else if os.IsNotExist(err) {
		exist = 0
	} else {
		exist = -1
	}
	return
}