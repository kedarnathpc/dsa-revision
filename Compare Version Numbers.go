func compareVersion(v1 string, v2 string) int {
    i, j, m, n := 0, 0, len(v1), len(v2)

    for i < m || j < n {
        num1, num2 := 0, 0
        for i < m && v1[i] != '.' {
            num1 = num1 * 10 + int(v1[i]-'0')
            i++
        }
        for j < n && v2[j] != '.' {
            num2 = num2 * 10 + int(v2[j]-'0')
            j++
        }
        if num1 < num2 {
            return -1
        }
        if num1 > num2 {
            return 1
        }
        i++
        j++
    }

    return 0
}
