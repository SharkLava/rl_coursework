package sarsa

import (
	"SharkLava/rl_grid_navigation/grid"
	"SharkLava/rl_grid_navigation/utils"
	"math"
	"math/rand"
	"time"
)

type SarsaAgent struct {
	Grid           *grid.Grid
	DiscountFactor float64
	LearningRate   float64
	Epsilon        float64
	QTable         map[[2]int]map[string]float64
	Performance    []float64
}

func (a *SarsaAgent) Train(episodes int) {
	for episode := 0; episode < episodes; episode++ {
		state := a.Grid.GetState()
		action := a.chooseAction(state)
		start := time.Now()
		for state != a.Grid.End {
			nextState := a.Grid.Move(state, action)
			nextAction := a.chooseAction(nextState)
			reward := a.Grid.GetReward(nextState)
			a.updateQTable(state, action, reward, nextState, nextAction)
			state = nextState
			action = nextAction
		}
		episodeTime := time.Since(start).Seconds()
		a.Performance = append(a.Performance, episodeTime)
		a.Epsilon = utils.Max(0.01, a.Epsilon*0.9995) // Decay epsilon slowly
	}
}

func (a *SarsaAgent) Benchmark() float64 {
	state := a.Grid.GetState()
	totalReward := 0.0
	for state != a.Grid.End {
		action := a.chooseAction(state)
		state = a.Grid.Move(state, action)
		totalReward += a.Grid.GetReward(state)
	}
	return totalReward
}

func (a *SarsaAgent) chooseAction(state [2]int) string {
	if rand.Float64() < a.Epsilon {
		return a.randomAction()
	}
	return a.getBestAction(state)
}

func (a *SarsaAgent) randomAction() string {
	actions := []string{"up", "down", "left", "right"}
	return actions[rand.Intn(len(actions))]
}

func (a *SarsaAgent) getBestAction(state [2]int) string {
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

func (a *SarsaAgent) getQValues(state [2]int) map[string]float64 {
	if q, ok := a.QTable[state]; ok {
		return q
	}
	q := make(map[string]float64)
	for _, action := range []string{"up", "down", "left", "right"} {
		q[action] = 0.0
	}
	a.QTable[state] = q
	return q
}

func (a *SarsaAgent) updateQTable(state [2]int, action string, reward float64, nextState [2]int, nextAction string) {
	qValues := a.getQValues(state)
	qValues[action] = qValues[action] + a.LearningRate*(reward+a.DiscountFactor*a.getQValues(nextState)[nextAction]-qValues[action])
}
