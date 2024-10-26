package dp

import (
	mdp "SharkLava/rl_grid_navigation/markov_decision_process"
	"time"
)

type DPAgent struct {
	MDP         *mdp.MDP
	Performance []float64
}

func (a *DPAgent) Train(episodes int) {
	for iter := 0; iter < episodes; iter++ {
		start := time.Now()
		a.MDP.EvaluatePolicy()
		a.MDP.ImprovePolicy()
		episodeTime := time.Since(start).Seconds()
		a.Performance = append(a.Performance, episodeTime)
	}
}

func (a *DPAgent) Benchmark() float64 {
	state := a.MDP.Grid.GetState()
	totalReward := 0.0
	for state != a.MDP.Grid.End {
		action := a.MDP.Policy[state]
		state = a.MDP.Grid.Move(state, action)
		totalReward += a.MDP.Grid.GetReward(state)
	}
	return totalReward
}
