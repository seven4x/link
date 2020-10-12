package learn

import (
	"strconv"
	"testing"
)

func Test_Validate(t *testing.T) {

	for i := 0; i < 10; i++ {

		Validate(strconv.Itoa(i))
	}
}
