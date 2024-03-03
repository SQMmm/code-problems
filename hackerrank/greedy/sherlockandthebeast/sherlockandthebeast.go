package sherlockandthebeast

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/sherlock-and-the-beast/problem
 */

func TestDecentNumber() {
	decentNumber(1)
	decentNumber(2)
	decentNumber(3)
	decentNumber(4)
	decentNumber(5)
	decentNumber(6)
	decentNumber(7)
	decentNumber(8)
	decentNumber(9)
	decentNumber(10)

}

func decentNumber(n int32) {
	if n < 3 || n == 4 {
		fmt.Println(-1)
		return
	}
	// Write your code here
	c := n/3


	for c >= 0 {
		//fmt.Println(c, (n-c*3)%5, c*3)
		if (n-c*3)%5 == 0 {
			for i := int32(0); i < c*3; i++ {
				fmt.Print("5")
			}
			for i := int32(0); i < (n-c*3); i++ {
				fmt.Print("3")
			}
			fmt.Println()
			return
		}
		c--
	}

	fmt.Println(-1)
}

