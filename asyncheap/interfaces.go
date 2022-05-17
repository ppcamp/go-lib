package asyncheap

type AsyncHeap interface {
	Push(x any)
	Pop() any
	Len() int
}
