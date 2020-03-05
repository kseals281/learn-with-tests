package arrays

func Sum(nums []int) (sum int) {
	for _, i := range nums {
		sum += i
	}
	return
}
func SumAll(lists ...[]int) (sums []int) {
	for _, list := range lists {
		sums = append(sums, Sum(list))
	}
	return
}

func SumAllTails(lists ...[]int) (sums []int) {
	for _, list := range lists {
		if len(list) <= 1 {
			sums = append(sums, 0)
			continue
		}
		list = list[1:]
		sums = append(sums, Sum(list))
	}
	return
}
