package queue

// Queue An FIFO queue
type Queue []int

// Push 推送方法
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop 弹出方法
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty 如果slice为空的情况
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
