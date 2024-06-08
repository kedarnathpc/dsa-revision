func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	mp := make(map[int]int)
	sum := 0

	mp[0] = -1

	for i := 0; i < n; i++ {
		sum += nums[i]

		remainder := sum
		if k != 0 {
			remainder = sum % k
		}

		if val, exists := mp[remainder]; exists {
			if i - val > 1 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}

	return false
}
