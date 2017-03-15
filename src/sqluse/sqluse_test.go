package sqluse

import (
	"testing"
)

func Test_findByPk(t *testing.T) {
	num := FindByUser(1)
	t.Log(num)
}
