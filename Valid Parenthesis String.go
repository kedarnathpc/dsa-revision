func checkValidString(s string) bool {
    mini, maxi := 0, 0

    for _, i := range s {
        if i == '(' {
            mini++
            maxi++
        } else if i == ')' {
            mini--
            maxi--
        } else {
            mini--
            maxi++
        }
        if maxi < 0 {
            return false
        }
        if mini < 0 {
            mini = 0
        }
    }

    return mini == 0
}
