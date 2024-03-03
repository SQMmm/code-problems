package stackmaxelement

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/maximum-element/problem
 */

type stack struct {
	values []int32
}

func newStack() *stack {
	return &stack{
		values: make([]int32, 0),
	}
}

func (s *stack) AddVal(val int32) {
	s.values = append(s.values, val)
}

func (s *stack) DeleteVal() {
	if len(s.values) == 0 {
		return
	}
	s.values = s.values[:len(s.values) - 1]
}

func (s *stack) GetMax() int32 {
	max := int32(0)
	for _, v := range s.values {
		if v > max {
			max = v
		}
	}

	return max
}


func getMax(operations []string) []int32 {
	s := &stack{}
	res := make([]int32, 0)
	for _, op := range operations {
		strs := strings.Split(op, " ")
		if len(strs) == 0 {
			log.Fatalln("no operations")
		}
		switch strs[0] {
		case "1":
			if len(strs) < 2 {
				log.Fatal("failed")
			}
			num, err := strconv.Atoi(strs[1])
			if err != nil {
				log.Fatal(err)
			}
			s.AddVal(int32(num))
		case "2":
			s.DeleteVal()
		case "3":
			res = append(res, s.GetMax())
		default :
			log.Fatal("no failed")
		}
	}

	return res
}


func TestGetMax() {
	strs := strings.Split(tt, "\n")
	fmt.Println(getMax(strs))
}