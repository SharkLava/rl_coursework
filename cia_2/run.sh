#!/bin/bash

# Define the parameters
GRID_SIZE=100
NUM_OBSTACLES=10
DISCOUNT_FACTOR=0.95
LEARNING_RATE=0.2
EPSILON=0.1
OUTPUT_FILE="results.csv"

# Loop through the episode values from 200 to 5000 with a step of 100
for EPISODES in $(seq 200 100 5000)
do
  while true; do
    # Run the program in the background and get its process ID (PID)
    ./rl_grid_navigation -grid_size=$GRID_SIZE -num_obstacles=$NUM_OBSTACLES -discount_factor=$DISCOUNT_FACTOR -learning_rate=$LEARNING_RATE -epsilon=$EPSILON -episodes=$EPISODES -output=$OUTPUT_FILE &
    PID=$!

    # Wait for up to 45 seconds for the process to complete
    SECONDS=0
    while kill -0 $PID 2>/dev/null && [ $SECONDS -lt 45 ]; do
      sleep 1
    done

    # Check if the process is still running after 45 seconds
    if kill -0 $PID 2>/dev/null; then
      # If it is still running, kill it and retry
      echo "Episodes: $EPISODES took too long (more than 45 seconds), killing process and retrying..."
      kill -9 $PID
      wait $PID 2>/dev/null # Clean up the terminated process
    else
      # If it finished within 45 seconds, break out of the retry loop
      echo "Episodes: $EPISODES completed, Time: $(date +%s)"
      break
    fi
  done
done
