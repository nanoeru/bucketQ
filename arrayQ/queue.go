package arrayQ

//	最大保持数指定
func NewQueue(size int) *Queue {
	return &Queue{
		buff: make([]interface{}, size),
	}
}

//	配列を用いたキュー(スレッドセーフではない)
type Queue struct {
	front, rear, count int
	buff               []interface{}
}

func (q *Queue) Enqueue(a ...interface{}) bool {
	for _, v := range a {
		if q.count == len(q.buff) {
			return false
		} else {
			q.buff[q.rear] = v
			q.rear++
			q.count++
			//	ループ
			if q.rear >= len(q.buff) {
				q.rear = 0
			}
			//			return true
		}
	}
	return true
}

func (q *Queue) Clear() {
	q.front, q.rear, q.count = 0, 0, 0
}

func (q *Queue) Length() int {
	return q.count
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue) IsFull() bool {
	return q.count != 0
}

func (q *Queue) Dequeue() (a interface{}, result bool) {
	//	Enqueueの実装により以下の場合は起こりえないはず
	//	q.rear == q.front {
	if q.count == 0 {
		return
	}
	a = q.buff[q.front]
	q.front++
	q.count--
	//	ループ
	if q.front >= len(q.buff) {
		q.front = 0
	}
	result = true
	return
}
