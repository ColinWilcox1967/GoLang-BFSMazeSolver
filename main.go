package main

import (
	"fmt"
)

// Point represents a position in the maze
type Point struct {
	x, y int
}

// Directions to move in the maze: up, down, left, right
var directions = []Point{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

// BFS to find the shortest path in the maze
func bfs(maze [][]int, start, end Point) []Point {
	rows := len(maze)
	cols := len(maze[0])

	// Initialize the queue and visited set
	queue := []Point{start}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[start.x][start.y] = true

	// Initialize a map to track the path
	prev := make(map[Point]*Point)

	// BFS loop
	for len(queue) > 0 {
		// Get the current point
		current := queue[0]
		queue = queue[1:]

		// Check if we've reached the end
		if current == end {
			return reconstructPath(prev, start, end)
		}

		// Explore neighbors
		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}

			// Check if the next point is within bounds and not visited
			if next.x >= 0 && next.x < rows && next.y >= 0 && next.y < cols && !visited[next.x][next.y] && maze[next.x][next.y] == 0 {
				queue = append(queue, next)
				visited[next.x][next.y] = true
				prev[next] = &current
			}
		}
	}

	// No path found
	return nil
}

// Reconstructs the path from the BFS search
func reconstructPath(prev map[Point]*Point, start, end Point) []Point {
	var path []Point
	for at := &end; at != nil; at = prev[*at] {
		path = append([]Point{*at}, path...) // Prepend the point to the path
	}

	// If the start point is not at the beginning, return an empty path (no valid path)
	if path[0] != start {
		return nil
	}

	return path
}

func main() {
	// Example maze (0 is open space, 1 is a wall)
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	start := Point{0, 0} // Start position (top-left corner)
	end := Point{4, 4}   // End position (bottom-right corner)

	path := bfs(maze, start, end)

	if path != nil {
		fmt.Println("Path found:")
		for _, p := range path {
			fmt.Printf("(%d, %d) ", p.x, p.y)
		}
		fmt.Println()
	} else {
		fmt.Println("No path found!")
	}
}
