package Daily26

func Remove_dups(nums []int) (length int) {
	for ptr := 1; ptr < len(nums); ptr++ {
		if nums[length] != nums[ptr] {
			length++
			nums[length], nums[ptr] = nums[ptr], nums[length]
		}
	}
	return length + 1
}
