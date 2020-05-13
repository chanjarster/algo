package heap

// 比较两个元素，判断a是否应该在b前面
type less func(a, b int) bool

// 原地建堆
// 数据的下标从0开始
// 当前元素下标i，左节点下标i*2+1，右节点下标i*2+2
func buildHeap(data []int, less less) {
	n := len(data)
	firstLeafI := n/2 - 1 // 第一个非叶子节点
	// 先从第一个非叶子节点开始，因为叶子节点没有子节点，没什么好比的。
	// 从下面开始堆化，这样每个节点就是和自己的左右子节点比，把最小/最大数放到堆顶
	for i := firstLeafI; i >= 0; i-- {
		//
		heapify(data, n, i, less)
	}
}

// 自上而下堆化
// params:
//  data: 数据
//  n: data的操作边界，只操作data[0:n]的区间
//  i: 从data[i]开始堆化
//  less: 一个函数，用来控制两个数字比较的时候，哪个应该放在前面
func heapify(data []int, n, i int, less less) {

	for {
		minI := i
		leftI := i*2 + 1
		rightI := i*2 + 2
		if leftI < n && less(data[leftI], data[minI]) {
			minI = leftI
		}
		if rightI < n && less(data[rightI], data[minI]) {
			minI = rightI
		}
		if i == minI {
			// 自己就是最小的
			break
		}
		swap(data, i, minI)
		i = minI
	}

}
