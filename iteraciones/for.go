package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 10; i > 0; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
