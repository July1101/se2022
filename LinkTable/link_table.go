package LinkTable

//package main

import (
	"fmt"
	"strconv"
)

type LinkNode struct {
	Data int
	Next *LinkNode
}

type LinkTable struct {
	Head *LinkNode
	Tail *LinkNode
}

func NewLinkTable() LinkTable {
	newHead := new(LinkNode)
	newHead.Next = nil
	newHead.Data = 0 //头节点的data用作记录节点个数
	return LinkTable{newHead, newHead}
}

//打印结点
func (t *LinkTable) PrintLinkTable() {
	fmt.Println("the table has " + strconv.Itoa(t.Head.Data) + " nodes")
	a := t.Head.Next
	for i := 0; i < t.Head.Data; i++ {
		fmt.Print(a.Data)
		if a != t.Tail {
			fmt.Print(" -> ")
		}
		a = a.Next
	}
	fmt.Println()
}

//尾插法插入
func (t *LinkTable) InsertAtTail(val int) {
	newNode := new(LinkNode)
	newNode.Data = val
	newNode.Next = nil

	t.Tail.Next = newNode
	t.Tail = newNode
	t.Head.Data++
}

//头插法插入
func (t *LinkTable) InsertAtHead(val int) {
	newNode := new(LinkNode)
	newNode.Data = val
	newNode.Next = t.Head.Next
	t.Head.Next = newNode
	t.Head.Data++
	if t.Tail == t.Head {
		t.Tail = newNode
	}
}

//根据下标插入值
func (t *LinkTable) InsertUseIndex(index, val int) {
	if index > t.Head.Data || index < 0 {
		fmt.Println("index error , insert fail")
		return
	}

	newNode := new(LinkNode)
	newNode.Data = val
	p := t.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	newNode.Next = p.Next
	p.Next = newNode

	t.Head.Data++
	//在最后一个节点插入 处理尾指针改变的情形
	if index == t.Head.Data-1 {
		t.Tail = newNode
	}
}

//根据值删除节点
func (t *LinkTable) RemoveUseValue(value int) {
	a := t.Head
	for a.Next != nil {
		if a.Next.Data == value {
			if a.Next == t.Tail {
				t.Tail = a
			}
			a.Next = a.Next.Next
			t.Head.Data--
		} else {
			a = a.Next
		}
	}
}

//根据下标删除节点
func (t *LinkTable) RemoveUseIndex(index int) {
	if index >= t.Head.Data || index < 0 {
		fmt.Println("index error , delete fail")
		return
	}

	a := t.Head
	for i := 0; i < index; i++ {
		a = a.Next
	}

	//处理删除最后一个节点的情况
	if index == t.Head.Data-1 {
		t.Tail = a
	}

	t.Head.Data--
	a.Next = a.Next.Next
}

//func main() {
//	t := NewLinkTable()
//
//	t.InsertAtHead(1)
//	t.InsertAtHead(2)
//	t.InsertAtHead(3)
//
//	t.InsertAtTail(4)
//	t.InsertAtTail(5)
//	t.InsertAtTail(6)
//
//	t.PrintLinkTable()
//
//	t.RemoveUseIndex(0)
//
//	t.PrintLinkTable()
//
//	t.RemoveUseValue(1)
//
//	t.PrintLinkTable()
//}
