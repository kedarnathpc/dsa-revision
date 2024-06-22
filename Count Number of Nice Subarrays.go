func solve (nums []int, k int) int {
    count, ans, i, j, n := 0, 0, 0, 0, len(nums)

    for j < n {
        if nums[j] & 1 == 1 {
            count++
        }

        for count > k {
            if nums[i] & 1 == 1 {
                count--
            }
            i++
        }

        ans += (j - i + 1)
        j++
    }

    return ans
}

func numberOfSubarrays(nums []int, k int) int {
    return solve(nums, k) - solve(nums, k - 1)
}
