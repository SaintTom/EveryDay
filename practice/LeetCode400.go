package main

import (
	"fmt"
	"math"
)

//官
func findNthDigit(n int) int {
	d := 1
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10 //math.Pow底数,几次幂)
}

func main() {
	//fmt.Println(findNthDigit(499))
	//
	//fmt.Println(findNthDigit(500))
	//fmt.Println(findNthDigit(501))
	//fmt.Println(findNthDigit(502))
	//fmt.Println(findNthDigit(503))
	fmt.Println(findNthDigit(504))

}
