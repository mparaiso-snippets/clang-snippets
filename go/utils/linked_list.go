package utils

type Any = interface{}

type Node struct {
	Item Any
	Next *Node
}

type LinkedList struct {
	First *Node
}
