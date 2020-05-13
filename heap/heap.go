// 堆（Heap）
// 是一个完全二叉树，每个节点>=子节点（大顶堆），或者，每个节点<=子节点（小顶堆）
// 极客时间: https://time.geekbang.org/column/article/69913
package heap

// 小顶堆
type Heap interface {
	Len() int   // 返回元素数量
	Push(n int) // 添加元素
	Pop() int   // 移除堆顶元素
	Peek() int  // 返回堆顶元素
}

// 第一个元素的下标是1
const headI = 1

// 这是一个小顶堆
type heapImpl struct {
	data []int // 下标0一直空着，元素从下标1插起
	big  bool  // 是否大顶堆
}

func NewHeap(big bool) Heap {
	return &heapImpl{
		data: make([]int, 1),
		big:  big,
	}
}

func (h *heapImpl) Len() int {
	return len(h.data) - 1
}

// 判断a是否应该放在b前面
func (h *heapImpl) less(a, b int) bool {
	if h.big {
		return a > b
	}
	return a < b
}

// 添加新元素
func (h *heapImpl) Push(n int) {

	// 新元素放到数组最后，然后自下而上构建堆
	h.data = append(h.data, n)

	i := len(h.data) - 1

	for i != headI {
		p := i / 2 // 父节点
		if h.less(h.data[i], h.data[p]) {
			swap(h.data, i, p)
			i = p
		} else {
			break
		}
	}

}

// 移除堆顶元素
func (h *heapImpl) Pop() int {
	if h.Len() <= 0 {
		return 0
	}

	v := h.data[headI]

	// 把最后的元素放到堆顶，然后长度-1
	lastI := len(h.data) - 1 // 最后一个元素
	h.data[headI] = h.data[lastI]
	h.data[lastI] = 0
	h.data = h.data[:lastI]

	// 从上而下构建堆
	i := headI
	for {
		minI := i
		le := minI * 2   // 左节点
		ri := minI*2 + 1 // 右节点

		// 看自己，左节点，右节点哪个最小，就把它放上去
		if le < lastI && h.less(h.data[le], h.data[minI]) {
			minI = le
		}
		if ri < lastI && h.less(h.data[ri], h.data[minI]) {
			minI = ri
		}
		// 还是自己最小
		if minI == i {
			break
		}
		swap(h.data, i, minI)
		i = minI
	}
	return v

}

func (h *heapImpl) Peek() int {
	if h.Len() == 0 {
		return 0
	}
	return h.data[headI]
}

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
