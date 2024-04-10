func deckRevealedIncreasing(deck []int) []int {
    sort.Ints(deck)
    n, i := len(deck), 0
    queue, result := make([]int, 0), make([]int, n)

    for i < n {
        queue = append(queue, i)
        i++
    }

    i = 0
    for i < n {
        result[queue[0]] = deck[i]
        queue = queue[1:]
        if len(queue) > 0 {
            queue = append(queue, queue[0])
            queue = queue[1:]
        }
        i++
    }

    return result
}
