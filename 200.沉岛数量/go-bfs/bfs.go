type point struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	ret := 0
	directs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	if len(grid) == 0 {
		return 0
	}

	xLen := len(grid)
	yLen := len(grid[0])

	walk := func(x int, y int) bool {
		// 为什么把 '0' 改成 0 就会失败呢？
		if x >= xLen || x < 0 || y < 0 || y >= yLen || grid[x][y] == '0' {
			return false
		}
		grid[x][y] = '0'
		return true
	}

	for {
		var queue []point
		for i := 0; i < xLen; i++ {
			for j := 0; j < yLen; j++ {
				if walk(i, j) {
					queue = append(queue, point{i, j})
					break
				}
			}
			if len(queue) != 0 {
				break
			}
		}

		if len(queue) == 0 {
			break
		}

		for len(queue) > 0 {
			tmp := queue[0]
			for _, direct := range directs {
				px := tmp.x + direct[0]
				py := tmp.y + direct[1]
				if walk(px, py) {
					queue = append(queue, point{px, py})
				}
			}
			queue = queue[1:]
		}
		ret++
	}

	return ret
}


// 作者：haozi007
// 链接：https://leetcode-cn.com/problems/number-of-islands/solution/golang-bfsshi-xian-by-haozi007/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。