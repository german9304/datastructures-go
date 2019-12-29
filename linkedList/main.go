package main 

import (
	ll "linkedList/main.go/linkedlist"
)


func main() {
	al := ll.LinkedList{}
	al.Init()
	al.Append(100)
	al.Append(200)
	al.Append(300)
	al.PrintItems()
}