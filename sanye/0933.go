package sanye

type RecentCounter struct {
	history []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		history: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.history = append(this.history, t)
	for t-this.history[0] > 3000 {
		this.history = this.history[1:]
	}
	return len(this.history)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
