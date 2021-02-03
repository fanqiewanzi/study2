package calculate

//定义栈的数据结构
type SeqStack struct {
	top   int
	stack []interface{}
}

const defaultTop = -1

//创建栈
func NewSeqStack() *SeqStack {
	stack := make([]interface{}, 10)
	return &SeqStack{defaultTop, stack}
}

//入栈操作
func (seqStack *SeqStack) Push(elem interface{}) {
	seqStack.top++
	seqStack.stack[seqStack.top] = elem
}

//出栈操作
func (seqStack *SeqStack) Pop() interface{} {
	seqStack.top--
	return seqStack.stack[seqStack.top+1]
}

//取得栈顶并不出栈
func (seqStack *SeqStack) GetTop() interface{} {
	return seqStack.stack[seqStack.top]
}

//判断栈是否为空
func (seqStack *SeqStack) Empty() bool {
	if seqStack.top == -1 {
		return true
	}
	return false
}
