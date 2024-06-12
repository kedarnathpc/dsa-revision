func sortColors(nums []int)  {
    i, j, k := 0, len(nums)-1, 0

    for k <= j {
        if nums[k] == 0 {
            nums[i], nums[k] = nums[k], nums[i]
            i++
            k++
        } else if nums[k] == 1 {
            k++
        } else {
            nums[k], nums[j] = nums[j], nums[k]
            j--
        }
    }
}
