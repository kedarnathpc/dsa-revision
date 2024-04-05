// Recursive Insertion Sort

func recInsertionSort(nums []int, i, n int) {
    if i == n {
        return
    }

    j := i
    for j > 0 && nums[j-1] > nums[j] {
        nums[j-1], nums[j] = nums[j], nums[j-1]
        j--
    }

    recInsertionSort(nums, i+1, n)
}

func sortArray(nums []int) []int {
    recInsertionSort(nums, 0, len(nums))
    return nums
}
