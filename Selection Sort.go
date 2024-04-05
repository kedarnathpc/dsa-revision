// Selection Sort
func sortArray(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        mini := i
        for j := i + 1; j < len(nums); j++ {
            if nums[j] < nums[mini] {
                mini = j
            }
        }
        nums[i], nums[mini] = nums[mini], nums[i]
    }

    return nums
}
