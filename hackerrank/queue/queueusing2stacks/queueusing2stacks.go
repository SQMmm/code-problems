package queueusing2stacks

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/queue-using-two-stacks/problem
 */

type stack struct {
	values []int
}

func newStack() *stack {
	return &stack{values: make([]int, 0)}
}

func (s *stack) Push(num int) {
	s.values = append(s.values, num)
}

func (s *stack) Pop() int {
	if len(s.values) == 0 {
		return -1
	}
	val := s.values[len(s.values) - 1]
	s.values = s.values[:len(s.values) - 1]
	return val
}
func (s *stack) Length() int {
	return len(s.values)
}

func (s *stack) Get() int {
	if len(s.values) == 0 {
		return 0
	}

	return s.values[len(s.values) - 1]
}

type queue struct {
	st1 *stack
	st2 *stack
}

func newQueue() *queue {
	return &queue{
		st1: newStack(),
		st2: newStack(),
	}
}

func (q *queue) Enqueue(num int) {
	q.reverseSt2()
	q.st1.Push(num)
}

func (q *queue) Dequeue() int {
	q.reverseSt1()
	return q.st2.Pop()
}

func (q *queue) Get() int {
	q.reverseSt1()
	return q.st2.Get()
}

func (q *queue) reverseSt1() {
	if q.st2.Length() > 0 {
		return
	}

	for q.st1.Length() > 0 {
		val := q.st1.Pop()
		q.st2.Push(val)
	}
}

func (q *queue) reverseSt2() {
	if q.st1.Length() > 0 {
		return
	}

	for q.st2.Length() > 0 {
		val := q.st2.Pop()
		q.st1.Push(val)
	}
}

func process(queries []string) {
	q := newQueue()
	res := make([]string, 0)
	for _, query := range queries {
		strs := strings.Split(query, " ")
		if len(strs) == 0 {
			return
		}
		switch  strs[0] {
		case "1":
			if len(strs) < 2 {
				return
			}
			n, _ := strconv.Atoi(strs[1])
			q.Enqueue(n)
		case "2":
			q.Dequeue()
		case "3":
			res = append(res, strconv.Itoa(q.Get()))
		default:
			return
		}
	}

	fmt.Println(strings.Join(res, "\n"))
}


func TestQueueProcess() {
	process([]string{
		"1 42",
		"2",
		"1 14",
		"3",
		"1 28",
		"3",
		"1 60",
		"1 78",
		"2",
		"2",
	})
}