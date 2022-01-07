/* Golang BEINNER.ТЕОРИЯ
Тема: массивы, срезы, карты
Задача № 4 (программa которая находит самый наименьший элемент в списке)

package main

import "fmt"

/*func main () {
	x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,
	}
	min := x[0]
	for i:=1; i < len(x); i ++ {
	if	min > x[i] {
	min = x[i]
	}
	}
	fmt.Println(min)
	}
*/

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	min := x[0]
	for i := range x {
		if min > x[i] {
			min = x[i]
		}
	}
	fmt.Println(min)
}