package useerr

import (
	"errors"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("平方根不能为负数")
	}
	return f * f, nil
}

func Use_my_sqrt() {
	f, err := sqrt(-1)
	if err != nil {
		log.Fatal("平方根不能为负数！")
	}
	fmt.Println(f)

}
