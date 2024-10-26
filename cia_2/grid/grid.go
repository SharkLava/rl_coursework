package grid

import (
	"math/rand"
	"time"
)

type Grid struct {
	Size       int
	Grid       [][]int
	Start, End [2]int
}

func NewGrid(size int, numObstacles int) *Grid {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	rand.Seed(time.Now().UnixNano())

	// Place obstacles
	for i := 0; i < numObstacles; i++ {
		x, y := rand.Intn(size), rand.Intn(size)
		for grid[x][y] == -1 || (x == 0 && y == 0) || (x == size-1 && y == size-1) {
			x, y = rand.Intn(size), rand.Intn(size)
		}
		grid[x][y] = -1
	}

	// Place start and end points
	start := [2]int{0, 0}
	end := [2]int{size - 1, size - 1}
	for grid[end[0]][end[1]] == -1 {
		end = [2]int{rand.Intn(size), rand.Intn(size)}
	}

	grid[start[0]][start[1]] = 1
	grid[end[0]][end[1]] = 2

	return &Grid{
		Size:  size,
		Grid:  grid,
		Start: start,
		End:   end,
	}
}

func (g *Grid) GetState() [2]int {
	return g.Start
}

func (g *Grid) GetReward(state [2]int) float64 {
	if state == g.End {
		return 1000.0 // High reward for reaching the goal
	} else if g.Grid[state[0]][state[1]] == 0 {
		return -1.0 // Small penalty for valid moves
	}
	return -5.0 // Larger penalty for obstacles or poor areas

}

func (g *Grid) IsValid(state [2]int) bool {
	x, y := state[0], state[1]
	return x >= 0 && x < g.Size && y >= 0 && y < g.Size && g.Grid[x][y] != -1
}

func (g *Grid) Move(state [2]int, action string) [2]int {
	x, y := state[0], state[1]
	switch action {
	case "up":
		x--
	case "down":
		x++
	case "left":
		y--
	case "right":
		y++
	}
	if g.IsValid([2]int{x, y}) {
		return [2]int{x, y}
	}
	return state
}
