package List

import (
	"errors"
	"fmt"
)

type LNode struct {
	data ElemType
	next *LNode
}

type LinkList struct {
	length int64
	hNode  *LNode
}

func InitLinkList() *LinkList {
	return &LinkList{0, &LNode{}}
}

func PrintLinkList(l *LinkList) {
	fmt.Println("=============================>>>")

	p := (*l).hNode.next
	for p != nil {
		fmt.Printf("%d\t", (*p).data)
		p = (*p).next
	}
	fmt.Println("\nlength:", (*l).length)
	fmt.Println("<<<=============================")
}

func (l *LinkList) IsEmpty() bool {
	return l == nil || (*l).hNode.next == nil
}

func (l *LinkList) Length() int64 {
	return (*l).length
}

func (l *LinkList) Clear() error {

	if l == nil {
		return errors.New("list is nil")
	}

	p := (*l).hNode.next
	var node *LNode
	for p != nil {
		node = (*p).next
		p = node
		(*l).length -= 1
	}

	(*l).hNode.next = nil
	(*l).length = 0

	return nil
}

func (l *LinkList) Destory() error {

	err := (*l).Clear()
	if err != nil {
		l = nil
	} else {
		return errors.New("Clear is failure")
	}

	return nil
}

func (l *LinkList) Insert(i int64, e ElemType) error {

	var j int64 = 0

	p := (*l).hNode

	for p != nil && j < i-1 {
		p = (*p).next
		j += 1
	}

	if p == nil || j > i-1 {
		return errors.New("Insert is failure")
	}

	node := &LNode{
		data: e,
		next: (*p).next,
	}

	(*p).next = node
	(*l).length += 1

	return nil
}

func (l *LinkList) Delete(i int64, e *ElemType) error {
	var j int64 = 0

	p := (*l).hNode

	for p != nil && j < i-1 {
		p = (*p).next
		j += 1
	}

	if p == nil || j > i-1 {
		return errors.New("Delete is failure")
	}

	node := (*p).next
	*e = (*node).data
	(*p).next = (*node).next
	(*l).length -= 1

	return nil

}

func (l *LinkList) GetElem(i int64, e *ElemType) error {
	var j int64 = 1
	p := (*l).hNode.next

	for p != nil && j < i {
		p = (*p).next
		j++
	}

	if p == nil || j > i {
		return errors.New("Not found")
	}
	*e = (*p).data

	return nil
}

func (l *LinkList) Locate(e ElemType) int64 {
	var j int64 = 1
	p := (*l).hNode.next

	for p != nil && (*p).data != e {
		p = (*p).next
		j++
	}

	if p != nil {
		return j
	} else {
		return 0
	}
}
