package main

|\longremark{Include all the packages we need.}|
import (
	"errors"
	"fmt"
)

type Value int |\longremark{Declare a type for the value our list will contain;}|

type Node struct { |\longremark{declare a type for the each node in our list;}|
	Value
	prev, next *Node
}

type List struct {
	head, tail *Node
}

|\longremark{Mimic the interface of container/list.}|
func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Push(v Value) *List {
	n := &Node{Value: v} |\longremark{When pushing, create a new Node with the provided value;}|

	if l.head == nil { |\longremark{if the list is empty, put the new node at the head;}|
		l.head = n
	} else {
		l.tail.next = n |\longremark{otherwise put it at the tail;}|
		n.prev = l.tail |\longremark{make sure the new node points back to the previously existing one;}|
	}
	l.tail = n |\longremark{point tail to the newly inserted node.}|

	return l
}

var errEmpty = errors.New("List is empty")

func (l *List) Pop() (v Value, err error) {
	if l.tail == nil { |\longremark{When popping, return an error if the list is empty;}|
		err = errEmpty
	} else {
		v = l.tail.Value |\longremark{otherwise save the last value;}|
		l.tail = l.tail.prev |\longremark{discard the last node from the list;}|
		if l.tail == nil {
			l.head = nil |\longremark{and make sure the list is consistent if it becomes empty;}|
		}
	}

	return v, err
}

func main() {
	l := new(List)

	l.Push(1)
	l.Push(2)
	l.Push(4)

	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v\n", n.Value)
	}

        fmt.Println()

	for v, err := l.Pop(); err == nil; v, err = l.Pop() {
		fmt.Printf("%v\n", v)
	}
}
