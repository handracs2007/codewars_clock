// https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a/train/go

package main

import "fmt"

func Past(h, m, s int) int {
	return h * 60 * 60 * 1000 + m * 60 * 1000 + s * 1000
}

func main(){
	fmt.Println(Past(0,1,1))
}