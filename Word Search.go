func search (i, row, col, m, n int, word string, board[][] byte) bool {
    if i == len(word) {
        return true
    }

    if row < 0 || col < 0 || row == m || col == n || board[row][col] != word[i] || board[row][col] == '!' {
        return false
    }

    c := board[row][col]
    board[row][col] = '!'

    up := search(i + 1, row - 1, col, m, n, word, board)
    down := search(i + 1, row + 1, col, m, n, word, board)
    left := search(i + 1, row, col - 1, m, n, word, board)
    right := search(i + 1, row, col + 1, m, n, word, board)

    board[row][col] = c

    return up || down || left || right
}

func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                if search(0, i, j, m, n, word, board) {
                    return true;
                }
            }
        }
    }

    return false;
}
