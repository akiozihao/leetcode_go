package sanye

type NumArray []int

func Constructor(nums []int) NumArray {
    n := len(nums)
    seg := make(NumArray, n*4)
    seg.build(nums, 0, 0, n-1)
    return seg
}

func (seg NumArray) build(nums []int, node, s, e int) {
    if s == e {
        seg[node] = nums[s]
        return
    }
    m := s + (e-s)/2
    seg.build(nums, node*2+1, s, m)
    seg.build(nums, node*2+2, m+1, e)
    seg[node] = seg[node*2+1] + seg[node*2+2]
}

func (seg NumArray) change(index, val, node, s, e int) {
    if s == e {
        seg[node] = val
        return
    }
    m := s + (e-s)/2
    if index <= m {
        seg.change(index, val, node*2+1, s, m)
    } else {
        seg.change(index, val, node*2+2, m+1, e)
    }
    seg[node] = seg[node*2+1] + seg[node*2+2]
}

func (seg NumArray) range_(left, right, node, s, e int) int {
    if left == s && right == e {
        return seg[node]
    }
    m := s + (e-s)/2
    if right <= m {
        return seg.range_(left, right, node*2+1, s, m)
    }
    if left > m {
        return seg.range_(left, right, node*2+2, m+1, e)
    }
    return seg.range_(left, m, node*2+1, s, m) + seg.range_(m+1, right, node*2+2, m+1, e)
}

func (seg NumArray) Update(index, val int) {
    seg.change(index, val, 0, 0, len(seg)/4-1)
}

func (seg NumArray) SumRange(left, right int) int {
    return seg.range_(left, right, 0, 0, len(seg)/4-1)
}
