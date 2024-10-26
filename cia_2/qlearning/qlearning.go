package qlearning

import (
	"SharkLava/rl_grid_navigation/grid"
	"SharkLava/rl_grid_navigation/utils"
	"math"
	"math/rand"
	"time"
)

type QLearningAgent struct {
	Grid           *grid.Grid
	DiscountFactor float64
	LearningRate   float64
	Epsilon        float64
	QTable         map[[2]int]map[string]float64
	Performance    []float64
}

func (a *QLearningAgent) Train(episodes int) {
	for episode := 0; episode < episodes; episode++ {
		state := a.Grid.GetState()
		start := time.Now()
		for state != a.Grid.End {
			action := a.chooseAction(state)
			nextState := a.Grid.Move(state, action)
			reward := a.Grid.GetReward(nextState)
			maxQ := a.getMaxQ(nextState)
			a.updateQTable(state, action, reward, maxQ)
			state = nextState
		}
		episodeTime := time.Since(start).Seconds()
		a.Performance = append(a.Performance, episodeTime)
		a.Epsilon = utils.Max(0.001, a.Epsilon*0.9995) // Decay epsilon slowly
	}
}

func (a *QLearningAgent) Benchmark() float64 {
	state := a.Grid.GetState()
	totalReward := 0.0
	for state != a.Grid.End {
		action := a.chooseAction(state)
		state = a.Grid.Move(state, action)
		totalReward += a.Grid.GetReward(state)
	}
	return totalReward
}

func (a *QLearningAgent) chooseAction(state [2]int) string {
	if rand.Float64() < a.Epsilon {
		return a.randomAction()
	}
	return a.getBestAction(state)
}

func (a *QLearningAgent) randomAction() string {
	actions := []string{"up", "down", "left", "right"}
	return actions[rand.Intn(len(actions))]
}

func (a *QLearningAgent) getBestAction(state [2]int) string {
	bestAction := "up"
	bestQ := -math.Inf(1)
	for action, q := range a.getQValues(state) {
		if q > bestQ {
			bestQ = q
			bestAction = action
		}
	}
	return bestAction
}

func (a *QLearningAgent) getMaxQ(state [2]int) float64 {
	maxQ := -math.Inf(1)
	for _, q := range a.getQValues(state) {
		maxQ = math.Max(maxQ, q)
	}
	return maxQ
}

func (a *QLearningAgent) getQValues(state [2]int) map[string]float64 {
	if q, ok := a.QTable[state]; ok {
		return q
	}
	q := make(map[string]float64)
	for _, action := range []string{"up", "down", "left", "right"} {
		q[action] = rand.Float64() * 0.01 // Small random value
	}
	a.QTable[state] = q
	return q
}

func (a *QLearningAgent) updateQTable(state [2]int, action string, reward float64, maxQ float64) {
	qValues := a.getQValues(state)
	qValues[action] = qValues[action] + a.LearningRate*(reward+a.DiscountFactor*maxQ-qValues[action])
}
