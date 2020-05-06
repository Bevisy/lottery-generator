package generater

import "testing"

// 测试isExist()
func TestExist(t *testing.T) {
	var tests = []struct {
		nums []int
		num  int
	}{
		{[]int{1, 2, 3}, 3},
	}

	for _, test := range tests {
		if !isExist(test.nums, test.num) {
			t.Errorf("%d is existed in %v", test.num, test.nums)
		}
	}
}
