func countStudents(students []int, sandwiches []int) int {
    queue := make([]int, 0)
    i, unchanged := 0, 0
    
    for _, s := range students {
        queue = append(queue, s)
    }

    for i < len(sandwiches) && unchanged < len(queue) {
        if queue[0] == sandwiches[i] {
            queue = queue[1:]
            i++
            unchanged = 0
        } else {
            queue = append(queue, queue[0])
            queue = queue[1:]
            unchanged++
        }
    }

    return len(queue)
}
