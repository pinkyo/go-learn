package main

func main() {
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		// 统计每个顶点的入度
		indegree[edge[0]]++
		// 构建邻接表
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	res := make([]int, 0)
	for len(queue) > 0 {
		// 出队列
		top := queue[0]
		queue = queue[1:]
		res = append(res, top)
		// 遍历顶点的邻接表
		for _, v := range graph[top] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}

	return nil
}
