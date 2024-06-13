func minMovesToSeat(seats []int, students []int) int {
    i, j, count := 0, 0, 0
    sort.Ints(seats)
    sort.Ints(students)

    for i < len(seats) && j < len(seats) {
        if seats[i] != students[i] {
            if seats[i] - students[i] > 0 {
                count += seats[i] - students[i]
            } else {
                count += students[i] - seats[i]
            }
        }
        i++
        j++
    }

    return count
}
