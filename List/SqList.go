package List

import "fmt"

const MAXSIZE = 100

type SqList struct {
	length int64
	data   [MAXSIZE]ElemType
}

func InitSqList() *SqList {
	return &SqList{}
}

func PrintSqList(sq *SqList) {
	fmt.Println("=============================>>>")
	for i := int64(0); i < (*sq).length; i++ {
		fmt.Printf("%d\t", (*sq).data[i])
	}
	fmt.Println("\n<<<=============================")
}

func (sq *SqList) ClearList() {
	(*sq).length = 0
}

func (sq *SqList) GetLength() int64 {
	return (*sq).length
}

func (sq *SqList) IsEmpty() bool {
	return (*sq).length == 0
}

func (sq *SqList) LocateElem(e ElemType) int64 {
	var i int64 = 0
	for i < (*sq).length && (*sq).data[i] != e {
		i++
	}

	if i < (*sq).length {
		return i + 1
	}

	return 0
}

func (sq *SqList) InsertElem(i int64, e ElemType) bool {
	if i < 1 || i > (*sq).length+1 {
		return false
	}

	if (*sq).length == MAXSIZE {
		return false
	}

	for j := (*sq).length - 1; j >= i-1; j-- {
		(*sq).data[j+1] = (*sq).data[j]
	}

	(*sq).data[i-1] = e
	(*sq).length += 1

	return true
}

func (sq *SqList) RemoveElem(i int64, e *ElemType) bool {
	if i < 1 || i > (*sq).length {
		return false
	}

	*e = (*sq).data[i-1]

	for j := i; j < (*sq).length; j++ {
		(*sq).data[j-1] = (*sq).data[j]
	}

	(*sq).length -= 1

	return true
}
