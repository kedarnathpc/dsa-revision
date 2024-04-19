type Pair struct {
	first, second int
}

func bfs(row, col, n, m int, grid [][]byte, visited [][]int) {
	visited[row][col] = 1
	q := []Pair{{row, col}}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for len(q) > 0 {
		r, c := q[0].first, q[0].second
		q = q[1:]

		for i := 0; i < 4; i++ {
			newRow := r + dx[i]
			newCol := c + dy[i]

			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < m && grid[newRow][newCol] == '1' && visited[newRow][newCol] != 1 {
				visited[newRow][newCol] = 1
				q = append(q, Pair{newRow, newCol})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	count := 0
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if visited[row][col] == 0 && grid[row][col] == '1' {
				count++
				bfs(row, col, n, m, grid, visited)
			}
		}
	}

	return count
}
