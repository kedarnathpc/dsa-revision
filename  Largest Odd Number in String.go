func largestOddNumber(num string) string {
    res := ""
    flag := -1
    i := len(num)-1

    for ; i >= 0; i-- {
        if (num[i]-'0') % 2 == 1 {
            flag = i
            break
        }
    }

    if flag == -1 {
        return res
    }
    res = num[0:i+1]
    return res
}
