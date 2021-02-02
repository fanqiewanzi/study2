package calculate

//定义栈的数据结构
type SeqStack struct {
	top   int
	stack []interface{}
}

//入栈操作
func (seqstack *SeqStack) Push(elem interface{}) {
	seqstack.stack[seqstack.top] = elem
	seqstack.top++
}

//出栈操作
func (seqstack *SeqStack) Pop() interface{} {
	seqstack.top--
	return seqstack.stack[seqstack.top+1]
}

//取得栈顶并不出栈
func (seqstack *SeqStack) GetTop() interface{} {
	return seqstack.stack[seqstack.top]
}

//判断栈是否为空
func (seqstack *SeqStack) Empty() bool {
	if seqstack.top == -1 {
		return true
	}
	return false
}
