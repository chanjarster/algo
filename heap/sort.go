package heap

// 堆排序
// 原理：
// 1. 把数组堆化（大顶堆）
// 2. data[0]是堆顶，和data[tail]交换，那么data[tail]就是最大值了
// 3. 对data[0:tail-1]自上而下堆化
// 4. 重复2-3，直到所有元素都弄完为止
func Sort(data []int) {
	if len(data) <= 1 {
		return
	}

	bigger := func(a, b int) bool {
		return a > b
	}
	buildHeap(data, bigger)

	tail := len(data) - 1
	for tail > 0 {
		// data[0]就是堆顶，它肯定是最小/最大的，直接和data[tail]交换
		swap(data, tail, 0)
		// 对data[0:tail-1]的堆化，因为data[0]（堆顶）元素变了，因此需要再次自上而下堆化一把
		heapify(data, tail, 0, bigger)
		tail--
	}

}
