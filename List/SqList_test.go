package List

import (
	"fmt"
	"testing"
)

func TestInsertElem(t *testing.T) {

	// 初始化
	sq := InitSqList()

	// 判断是否为空，获取线性表的长度
	fmt.Println(sq.IsEmpty(), sq.GetLength())

	PrintSqList(sq)

	//插入
	for i := 0; i < 10; i++ {
		sq.InsertElem(1, ElemType(i+1000))
	}

	// 判断是否为空，获取线性表的长度
	fmt.Println(sq.IsEmpty(), sq.GetLength())

	PrintSqList(sq)

	// 获取元素的位置
	fmt.Print(">>>\t")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", sq.LocateElem(ElemType(i+1000)))
	}
	fmt.Print("<<<\n")

	// 删除元素
	var e ElemType
	for i := 0; i < 10; i++ {
		sq.RemoveElem(1, &e)
		fmt.Printf("删除的元素为: %d 删除后的线性表如下: \n", e)
		PrintSqList(sq)
	}

}
