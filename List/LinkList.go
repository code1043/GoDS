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
	next   *LNode
}

func InitLinkList() *LinkList {
	return &LinkList{}
}

func PrintLinkList(l *LinkList) {
	fmt.Println("=============================>>>")

	p := (*l).next
	for p != nil {
		fmt.Printf("%d\t", (*p).data)
		p = (*p).next
	}
	fmt.Println("\nlength:", (*l).length)
	fmt.Println("<<<=============================")
}

func (l *LinkList) IsEmpty() bool {
	return (*l).next == nil
}

func (l *LinkList) Length() int64 {
	return (*l).length
}

func (l *LinkList) Clear() error {

	if l == nil {
		return errors.New("list is nil")
	}

	p := (*l).next
	var q *LNode
	for p != nil {
		q = (*p).next
		p = q
		(*l).length -= 1
	}

	(*l).next = nil
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

	if i < 1 || i > l.length+1 {
		return errors.New("index out of range")
	}

	var j int64 = 1
	p := (*l).next
	for p != nil && j < i-1 {
		p = (*p).next
		j++
	}

	node := &LNode{
		data: e,
		next: nil,
	}

	if i == 1 {
		(*node).next = (*l).next
		(*l).next = node
	} else {
		(*node).next = (*p).next
		(*p).next = node
	}
	(*l).length += 1
	return nil
}

func (l *LinkList) Delete(i int64, e *ElemType) error {
	if i < 1 || i > l.length {
		return errors.New("index out of range")
	}

	if i == 1 {
		*e = (*l).next.data
		(*l).next = (*l).next.next
	} else {
		var j int64 = 1
		p := (*l).next
		for p != nil && j < i-1 {
			p = (*p).next
			j++
		}

		*e = (*p).next.data
		(*p).next = (*p).next.next
	}

	(*l).length -= 1
	return nil

}

func (l *LinkList) GetElem(i int64, e *ElemType) error {
	var j int64 = 1
	p := (*l).next

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
	p := (*l).next

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
