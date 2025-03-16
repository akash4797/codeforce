package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Direction vectors for king's movement (8 possible moves)
var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

// Check if a position is under attack by the queen
func isUnderAttack(x, y, qx, qy int) bool {
	// Same row or column
	if x == qx || y == qy {
		return true
	}
	// Same diagonal
	if math.Abs(float64(x-qx)) == math.Abs(float64(y-qy)) {
		return true
	}
	return false
}

// Check if position is within board boundaries
func isValid(x, y, n int) bool {
	return x >= 1 && x <= n && y >= 1 && y <= n
}

func canKingEscape(n, ax, ay, bx, by, cx, cy int) bool {
	// If target is under attack, return false
	if isUnderAttack(cx, cy, ax, ay) {
		return false
	}

	// BFS queue
	type pos struct{ x, y int }
	queue := []pos{{bx, by}}
	visited := make(map[pos]bool)
	visited[pos{bx, by}] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If we reached the target
		if curr.x == cx && curr.y == cy {
			return true
		}

		// Try all possible moves
		for i := 0; i < 8; i++ {
			nx := curr.x + dx[i]
			ny := curr.y + dy[i]
			newPos := pos{nx, ny}

			if isValid(nx, ny, n) && !visited[newPos] && !isUnderAttack(nx, ny, ax, ay) {
				visited[newPos] = true
				queue = append(queue, newPos)
			}
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, ax, ay, bx, by, cx, cy int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &ax, &ay)
	fmt.Fscan(reader, &bx, &by)
	fmt.Fscan(reader, &cx, &cy)

	if canKingEscape(n, ax, ay, bx, by, cx, cy) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
