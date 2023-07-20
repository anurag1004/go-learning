package main

import "fmt"

func dfs(node int, visited []bool, graph [][]int) {
	fmt.Printf("%v,", node)
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, visited, graph)
		}
	}
}
func dfsIter(start int, graph [][]int) {
	stack := make([]int, 0, len(graph))
	visited := make([]bool, len(graph))
	stack = append(stack, start)
	for len(stack) != 0 {
		node := stack[len(stack)-1]  // get the top
		stack = stack[:len(stack)-1] // pop

		if !visited[node] {
			fmt.Printf("%v,", node)
			visited[node] = true
		} else {
			continue
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
}
func bfs(start int, graph [][]int) {
	queue := make([]int, 0, len(graph))
	visited := make([]bool, len(graph))

	queue = append(queue, start)
	visited[start] = true

	for len(queue) != 0 {
		node := queue[0]  // peek
		queue = queue[1:] // dequeue
		fmt.Printf("%v,", node)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}

	}
}
func main() {
	graph := [][]int{{1, 2, 5}, {0, 3}, {0, 4}, {1, 5}, {2, 5}, {0, 3, 4}}
	dfs(0, make([]bool, len(graph)), graph)
	fmt.Println()
	dfsIter(0, graph)
	fmt.Println()
	bfs(0, graph)
}
