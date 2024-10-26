package main

import (
	"SharkLava/rl_grid_navigation/agent"
	"SharkLava/rl_grid_navigation/grid"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type AgentResult struct {
	Name      string
	TrainTime float64
	BenchTime float64
	Reward    float64
}

func trainAndBenchmark(agent agent.Agent, episodes int, name string, resultChan chan<- AgentResult) {
	start := time.Now()
	agent.Train(episodes)
	trainTime := time.Since(start).Seconds()

	start = time.Now()
	reward := agent.Benchmark()
	benchTime := time.Since(start).Seconds()

	resultChan <- AgentResult{Name: name, TrainTime: trainTime, BenchTime: benchTime, Reward: reward}
}

func main() {
	gridSize := flag.Int("grid_size", 5, "Size of the grid")
	numObstacles := flag.Int("num_obstacles", 10, "Number of obstacles in the grid")
	discountFactor := flag.Float64("discount_factor", 0.9, "Discount factor for the agent")
	learningRate := flag.Float64("learning_rate", 0.1, "Learning rate for Q-Learning and SARSA agents")
	epsilon := flag.Float64("epsilon", 0.1, "Exploration rate for Q-Learning and SARSA agents")
	episodes := flag.Int("episodes", 1000, "Number of training episodes")
	outputFile := flag.String("output", "results.csv", "Output CSV file for results")

	flag.Parse()

	grid := grid.NewGrid(*gridSize, *numObstacles)

	dpAgent := agent.NewDPAgent(grid, *discountFactor)
	qAgent := agent.NewQLearningAgent(grid, *discountFactor, *learningRate, *epsilon)
	sarsaAgent := agent.NewSarsaAgent(grid, *discountFactor, *learningRate, *epsilon)

	resultChan := make(chan AgentResult)

	var wg sync.WaitGroup
	wg.Add(3)

	// Start training and benchmarking in separate goroutines
	go func() {
		defer wg.Done()
		trainAndBenchmark(dpAgent, *episodes, "DP Agent", resultChan)
		// fmt.Println("MDP")
	}()

	go func() {
		defer wg.Done()
		trainAndBenchmark(qAgent, *episodes, "Q-Learning Agent", resultChan)
		// fmt.Println("Q-Learning")
	}()

	go func() {
		defer wg.Done()
		trainAndBenchmark(sarsaAgent, *episodes, "SARSA Agent", resultChan)
		// fmt.Println("SARSA")
	}()

	results := make([]AgentResult, 3)
	go func() {
		for i := 0; i < 3; i++ {
			results[i] = <-resultChan
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	// Open the CSV file in append mode, or create it if it doesnt exist
	file, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open output file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Check if the file is empty to write a header if needed
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	// Write the header only if the file is empty
	if fileInfo.Size() == 0 {
		err = writer.Write([]string{"Algorithm", "Episodes", "Reward"})
		if err != nil {
			log.Fatalf("Failed to write CSV header: %v", err)
		}
	}

	// Write the results
	for _, result := range results {
		fmt.Printf("%s %d %.2f\n", result.Name, *episodes, result.Reward)
		err := writer.Write([]string{
			result.Name,
			strconv.Itoa(*episodes),
			fmt.Sprintf("%.2f", result.Reward),
		})
		if err != nil {
			log.Fatalf("Failed to write CSV record: %v", err)
		}
	}
	fmt.Printf("Results have been appended to %s\n", *outputFile)
}
