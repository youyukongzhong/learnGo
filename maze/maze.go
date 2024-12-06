package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readMaze 读取迷宫数据
func readMaze(filename string) [][]int {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err) // 如果文件无法打开,则退出程序
	}

	// 使用 bufio.Scanner 逐行读取文件内容
	//scanner.Scan()：调用这个方法时，bufio.Scanner 会从输入文件中读取下一行的内容，直到遇到换行符（\n）为止。
	//如果文件的第一行是 6 5，scanner.Scan() 就会将 6 5 这一行读取到内部缓冲中。
	scanner := bufio.NewScanner(file)
	scanner.Scan()                           // 读取第一行
	header := strings.Fields(scanner.Text()) // 分割行内容
	row, _ := strconv.Atoi(header[0])        // 提取行数
	col, _ := strconv.Atoi(header[1])        // 提取列数

	// 初始化二维数组来存储迷宫
	// make 函数用于创建切片、映射（map）和通道（channel）。它接受三个参数：类型、长度和容量（可选）。
	maze := make([][]int, row)
	for i := 0; i < row; i++ {
		scanner.Scan()                         // 读取每一行
		line := strings.Fields(scanner.Text()) // 分割行内容为字符串数组
		maze[i] = make([]int, col)             // 将 maze 中的每一行初始化为一个切片，确保每一行有 col 列，并且所有的值是 0（默认值）
		for j := 0; j < col; j++ {
			maze[i][j], _ = strconv.Atoi(line[j]) // 转换为整数
		}
	}
	return maze // 返回二维数组
}

// 定义点结构体, 表示二维坐标
type point struct {
	i, j int // i 表示行索引，j 表示列索引
}

// 定义方向数组,表示上左下右四个移动方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}, // 上、左、下、右
}

// `add` 方法用于将当前点和另一个点的坐标相加，返回新点
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// `at` 方法检查某点是否在二维数组范围内，并返回该点的值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) { // 检查行是否越界
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) { // 检查列是否越界
		return 0, false
	}
	return grid[p.i][p.j], true // 返回点值和有效性
}

// 广度优先搜索算法，计算从起点到终点的最短路径步数
func walk(maze [][]int, start, end point) [][]int {
	// 初始化步数数组，记录从起点到每个点的步数
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i])) // 每行分配内存
	}

	// BFS 队列，初始只包含起点
	Q := []point{start}

	// BFS 主循环
	for len(Q) > 0 {
		cur := Q[0] // 获取队列中的第一个点
		Q = Q[1:]   // 将该点从队列中移除

		// 发现终点,退出程序
		if cur == end {
			break
		}

		// 遍历四个方向
		for _, dir := range dirs {
			next := cur.add(dir) //计算移动后的新点

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			// 检查新点是否在迷宫范围内且不是障碍物
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			// 检查新点是否已经被访问过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			// 起点本身不能作为下一个点
			if next == start {
				continue
			}

			// 当前点的步数加一，记录到新点
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			// 将新点加入队列
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
