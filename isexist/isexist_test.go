package isexist

import "testing"

func TestIsExist(t *testing.T) {
	var (
		pathExist = "isexist.go"
		expectExist int8 = 1
	)

	exist := IsExist(pathExist)
	if exist == expectExist {
		t.Log("exist")
	} else if exist == 0 {
		t.Log("not exist")
	} else {
		t.Log("unkown error")
	}
}

func TestIsNotExist(t *testing.T) {
	var (
		pathNotExist = "notisexist.go"
		expectExist int8 = 1
	)

	exist := IsExist(pathNotExist)
	if exist == expectExist {
		t.Log("exist")
	} else if exist == 0 {
		t.Log("not exist")
	} else {
		t.Log("unkown error")
	}
}