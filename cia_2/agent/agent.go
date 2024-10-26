package agent

import (
	"SharkLava/rl_grid_navigation/dp"
	"SharkLava/rl_grid_navigation/grid"
	mdp "SharkLava/rl_grid_navigation/markov_decision_process"
	"SharkLava/rl_grid_navigation/qlearning"
	"SharkLava/rl_grid_navigation/sarsa"
)

type Agent interface {
	Train(int)
	Benchmark() float64
}

func NewDPAgent(grid *grid.Grid, discountFactor float64) *dp.DPAgent {
	return &dp.DPAgent{
		MDP: mdp.NewMDP(grid, discountFactor),
	}
}

func NewQLearningAgent(grid *grid.Grid, discountFactor float64, learningRate float64, epsilon float64) *qlearning.QLearningAgent {
	return &qlearning.QLearningAgent{
		Grid:           grid,
		DiscountFactor: discountFactor,
		LearningRate:   learningRate,
		Epsilon:        epsilon,
		QTable:         make(map[[2]int]map[string]float64),
	}
}

func NewSarsaAgent(grid *grid.Grid, discountFactor float64, learningRate float64, epsilon float64) *sarsa.SarsaAgent {
	return &sarsa.SarsaAgent{
		Grid:           grid,
		DiscountFactor: discountFactor,
		LearningRate:   learningRate,
		Epsilon:        epsilon,
		QTable:         make(map[[2]int]map[string]float64),
	}
}
