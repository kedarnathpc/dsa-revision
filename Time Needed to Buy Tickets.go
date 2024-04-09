func timeRequiredToBuy(tickets []int, k int) int {
    count := 0

    for i := range tickets {
        if i <= k {
            count += min(tickets[i], tickets[k])
        } else {
            count += min(tickets[i], tickets[k]-1)
        }
    }

    return count
}
