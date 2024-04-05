//Recursive Bubble Sort

func recBubbleSort(nums []int, n int) {
    if n == 1 {
        return
    }

    for i := 0; i < n-1; i++ {
        if nums[i] > nums[i+1] {
            nums[i], nums[i+1] = nums[i+1], nums[i]
        }
    }

    recBubbleSort(nums, n-1)
}
func sortArray(nums []int) []int {
    recBubbleSort(nums, len(nums))
    return nums
}
