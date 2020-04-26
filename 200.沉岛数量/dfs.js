// 方法一、深度优先
const dfs = (grid, r, c) => {
  const nr = grid.length
  const nc = grid[0].length
  if (r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] === '0') {
    return
  }
  grid[r][c] = '0'
  dfs(grid, r - 1, c)
  dfs(grid, r + 1, c)
  dfs(grid, r, c - 1)
  dfs(grid, r, c + 1)
}
const numIslands = grid => {
  if (!grid || !grid.length) {
    return 0
  }
  let sum = 0
  const nr = grid.length
  const nc = grid[0].length
  for (let i = 0; i < nr; i++) {
    for (let j = 0; j < nc; j++) {
      if (grid[i][j] === '1') {
        sum++
        dfs(grid, i, j)
      }
    }
  }
  return sum
}

// 下面的是别人的答案
/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function (grid) {
  let num = 0
  if (grid && grid.length) {
    const maxI = grid.length - 1,
      maxJ = grid[0].length - 1
    function overturn (i, j) {
      // 符合条件的翻转为'0'
      if (i < 0 || j < 0 || i > maxI || j > maxJ) return
      if (grid[i][j] === '1') {
        grid[i][j] = '0'
        overturn(i, j - 1)
        overturn(i - 1, j)
        overturn(i + 1, j)
        overturn(i, j + 1)
      }
    }
    for (let i = 0; i < grid.length; i++) {
      for (let j = 0; j < grid[i].length; j++) {
        if (grid[i][j] === '1') {
          num++
          overturn(i, j)
        }
      }
    }
  }
  return num
}

// 作者：tian-yi-19
// 链接：https://leetcode-cn.com/problems/number-of-islands/solution/chen-dao-ji-shu-by-tian-yi-19/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
