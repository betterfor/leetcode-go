#### 题目
<p>编写一个程序，通过已填充的空格来解决数独问题。</p>

<p>一个数独的解法需<strong>遵循如下规则</strong>：</p>

<ol>
	<li>数字&nbsp;<code>1-9</code>&nbsp;在每一行只能出现一次。</li>
	<li>数字&nbsp;<code>1-9</code>&nbsp;在每一列只能出现一次。</li>
	<li>数字&nbsp;<code>1-9</code>&nbsp;在每一个以粗实线分隔的&nbsp;<code>3x3</code>&nbsp;宫内只能出现一次。</li>
</ol>

<p>空白格用&nbsp;<code>&#39;.&#39;</code>&nbsp;表示。</p>

<p><img src="http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png"></p>

<p><small>一个数独。</small></p>

<p><img src="http://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png"></p>

<p><small>答案被标成红色。</small></p>

<p><strong>Note:</strong></p>

<ul>
	<li>给定的数独序列只包含数字&nbsp;<code>1-9</code>&nbsp;和字符&nbsp;<code>&#39;.&#39;</code>&nbsp;。</li>
	<li>你可以假设给定的数独只有唯一解。</li>
	<li>给定数独永远是&nbsp;<code>9x9</code>&nbsp;形式的。</li>
</ul>


 #### 题解
 1、暴力法
 将每种可能性都验证一遍，一个格子有9种可能，共81个格子，时间复杂度最多9^81^
 
 2、回溯法
 依次往空格中添加 *1-9* ，保证满足条件
 ```go
 func solveSudoku(board [][]byte) {
 	solveSudo(board,0)
 }
 
 func solveSudo(board [][]byte, index int) bool {
 	if index == 81 {
 		return true
 	}
 
 	row,col := index/9,index%9
 	if board[row][col] != '.' {
 		return solveSudo(board,index+1)
 	}
 
 	boxI,boxJ := (row/3)*3,(col/3)*3
 
 	isValid := func(b byte) bool { // 校验横行、纵行、9宫格的数字是否符合标准
 		for i := 0; i < 9; i++ {
 			if board[row][i] == b || board[i][col] == b || board[boxI+i/3][boxJ+i%3] == b {
 				return false
 			}
 		}
 		return true
 	}
 
 	for b := byte('1'); b <= '9'; b++ {
 		// 这里尝试填数字
 		if isValid(b) { // 如果有符合的数字，将该空赋值并校验下一个空
 			board[row][col] = b
 			if solveSudo(board,index+1) {
 				return true
 			}
 		}
 	}
 
 	// 沒有找到，置'.'
 	board[row][col] = '.'
 
 	return false
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-11/003702.png)