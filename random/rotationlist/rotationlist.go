package rotationlist

import (
	"fmt"
	"strings"
)

func TestRotationList() {
	l := &list{value: 1, next: &list{value: 2, next: &list{value: 3, next: &list{value: 4, next: &list{value: 5}}}}}
	fmt.Println(l)
	fmt.Println()
	fmt.Println(rotateList(l, 3))
	//l = &list{value: 1, next: &list{value: 2, next: &list{value: 3, next: &list{value: 4, next: &list{value: 5}}}}}
	fmt.Println(rotateList(l, 4))
	//l = &list{value: 1, next: &list{value: 2, next: &list{value: 3, next: &list{value: 4, next: &list{value: 5}}}}}
	fmt.Println(rotateList(l, 5))
	//l = &list{value: 1, next: &list{value: 2, next: &list{value: 3, next: &list{value: 4, next: &list{value: 5}}}}}
	fmt.Println(rotateList(l, 2))
	//l = &list{value: 1, next: &list{value: 2, next: &list{value: 3, next: &list{value: 4, next: &list{value: 5}}}}}
}


type list struct {
	value int
	next *list
}

func (l *list) String() string {
	arr := make([]string, 0)
	curr := l
	for curr != nil {
		arr = append(arr, fmt.Sprintf("%v", curr.value))
		curr = curr.next
	}

	return strings.Join(arr, " ")
}

func rotateList(l *list, k int) *list {
	if l.next == nil {
		return l
	}
	tmpList := l

	for i := 0; i < k; i++ {
		val := pop(tmpList)
		newList := *tmpList
		tmpList = &list{value:val, next: &newList}
	}

	*l = *tmpList
	return l
}

func pop(l *list) int {
	if l.next.next == nil {
		val := l.next.value
		l.next = nil
		return val
	}

	return pop(l.next)
}
