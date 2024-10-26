package markov_decision_process

import (
	"SharkLava/rl_grid_navigation/grid"
	"math"
	"sync"
)

type MDP struct {
	Grid           *grid.Grid
	DiscountFactor float64
	ValueFunc      map[[2]int]float64
	Policy         map[[2]int]string
	mu             sync.Mutex
}

func NewMDP(grid *grid.Grid, discountFactor float64) *MDP {
	valueFunc := make(map[[2]int]float64)
	policy := make(map[[2]int]string)
	for x := 0; x < grid.Size; x++ {
		for y := 0; y < grid.Size; y++ {
			state := [2]int{x, y}
			valueFunc[state] = 0.0
			policy[state] = "up"
		}
	}
	return &MDP{
		Grid:           grid,
		DiscountFactor: discountFactor,
		ValueFunc:      valueFunc,
		Policy:         policy,
	}
}

func (m *MDP) EvaluatePolicy() {
	for iter := 0; iter < 1000; iter++ {
		delta := 0.0
		for x := 0; x < m.Grid.Size; x++ {
			for y := 0; y < m.Grid.Size; y++ {
				state := [2]int{x, y}
				oldValue := m.ValueFunc[state]
				action := m.Policy[state]
				nextState := m.Grid.Move(state, action)
				reward := m.Grid.GetReward(nextState)
				m.ValueFunc[state] = reward + m.DiscountFactor*m.ValueFunc[nextState]
				delta = math.Max(delta, math.Abs(oldValue-m.ValueFunc[state]))
			}
		}
		if delta < 0.001 {
			break
		}
	}
}

func (m *MDP) ImprovePolicy() {
	for x := 0; x < m.Grid.Size; x++ {
		for y := 0; y < m.Grid.Size; y++ {
			state := [2]int{x, y}
			bestValue := -math.Inf(1)
			bestAction := "up"
			for _, action := range []string{"up", "down", "left", "right"} {
				nextState := m.Grid.Move(state, action)
				reward := m.Grid.GetReward(nextState)
				value := reward + m.DiscountFactor*m.ValueFunc[nextState]
				if value > bestValue {
					bestValue = value
					bestAction = action
				}
			}
			m.Policy[state] = bestAction
		}
	}
}

func (m *MDP) Train() {
	for iter := 0; iter < 100; iter++ {
		m.EvaluatePolicy()
		m.ImprovePolicy()
	}
}

