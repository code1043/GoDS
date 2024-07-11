package List

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {

	var e ElemType = 0

	LL := NewLinkList()
	fmt.Println("是否为空:", LL.IsEmpty())

	//插入
	for i := 1; i <= 10; i++ {
		LL.Insert(int64(i), ElemType(i+1000))
	}
	fmt.Println("是否为空:", LL.IsEmpty())

	PrintLinkList(LL)

	fmt.Println("查找位置为6的元素>>>>")
	LL.GetElem(6, &e)
	fmt.Println(e)

	fmt.Println("查找元素的位置： ", LL.Locate(1006), LL.Locate(100008))

	fmt.Println("删除:>>>>")

	for i := 10; i >= 1; i-- {
		LL.Delete(int64(i), &e)
		fmt.Printf("%d\t", e)
	}
	fmt.Println("\n是否为空:", LL.IsEmpty())

	PrintLinkList(LL)

}
