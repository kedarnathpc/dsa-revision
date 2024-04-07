func reverseSlice(nums []int, i, j int){
    for i <= j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}

func rotate(nums []int, k int)  {
    k = k % len(nums)
    reverseSlice(nums, 0, len(nums)-k-1)
    reverseSlice(nums, len(nums)-k, len(nums)-1)
    reverseSlice(nums, 0, len(nums)-1)
}
