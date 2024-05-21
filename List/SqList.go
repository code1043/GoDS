package List

const MAXSIZE = 100

type SqList struct {
	length int64
	data   [MAXSIZE]ElemType
}

func InitSqList() *SqList {
	return &SqList{}
}
