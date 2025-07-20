package main

import (
	"fmt"
	"strings"
	"time"
)

// State represents the complete state of the vacuum world
type State struct {
	Location string
	AState   string
	BState   string
}

// This function determines the action based on the current location and state
func reflexAgent(location, aState, bState string) string {
	// Priority: Clean dirty rooms first, then move to explore
	if location == "A" && aState == "DIRTY" {
		return "CLEAN"
	} else if location == "B" && bState == "DIRTY" {
		return "CLEAN"
	} else if location == "A" {
		return "RIGHT"
	} else if location == "B" {
		return "LEFT"
	}
	return ""
}

// Function to systematically visit all 8 states
func visitAllStates() {
	fmt.Println("=== SYSTEMATIC EXPLORATION OF ALL 8 STATES ===\n")

	// Define all possible states
	allStates := []State{
		{"A", "DIRTY", "DIRTY"},
		{"A", "CLEAN", "DIRTY"},
		{"A", "DIRTY", "CLEAN"},
		{"A", "CLEAN", "CLEAN"},
		{"B", "DIRTY", "DIRTY"},
		{"B", "CLEAN", "DIRTY"},
		{"B", "DIRTY", "CLEAN"},
		{"B", "CLEAN", "CLEAN"},
	}

	fmt.Println("All possible states in the vacuum world:")
	for i, state := range allStates {
		fmt.Printf("%d. (%s, %s, %s) - ", i+1, state.Location, state.AState, state.BState)

		// Determine what action the agent would take in this state
		action := reflexAgent(state.Location, state.AState, state.BState)
		fmt.Printf("Agent would choose: %s\n", action)

		// Show the resulting state after the action
		nextState := state
		if action == "CLEAN" {
			if state.Location == "A" {
				nextState.AState = "CLEAN"
			} else {
				nextState.BState = "CLEAN"
			}
		} else if action == "RIGHT" {
			nextState.Location = "B"
		} else if action == "LEFT" {
			nextState.Location = "A"
		}

		fmt.Printf("   → Next state: (%s, %s, %s)\n\n", nextState.Location, nextState.AState, nextState.BState)
		time.Sleep(1 * time.Second)
	}
}

// This function simulates the vacuum cleaner's operation and tracks visited states
func runLimitedSimulation(initialState State, maxSteps int) {
	visitedStates := make(map[string]bool)
	currentState := initialState
	step := 1

	fmt.Printf("\n=== LIMITED SIMULATION (Max %d steps) ===\n", maxSteps)
	fmt.Println("Starting from:", fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState))
	fmt.Println()

	for step <= maxSteps {
		// Create a string representation of the current state
		stateKey := fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState)

		// Check if we've visited this state before
		if !visitedStates[stateKey] {
			visitedStates[stateKey] = true
			fmt.Printf("Step %d - ✓ NEW STATE: %s\n", step, stateKey)
		} else {
			fmt.Printf("Step %d - ↻ Revisiting: %s\n", step, stateKey)
		}

		// Get action from reflex agent
		action := reflexAgent(currentState.Location, currentState.AState, currentState.BState)
		fmt.Printf("    Action: %s", action)

		// Apply the action and update state
		if action == "CLEAN" {
			if currentState.Location == "A" {
				currentState.AState = "CLEAN"
				fmt.Printf(" → Cleaned room A")
			} else if currentState.Location == "B" {
				currentState.BState = "CLEAN"
				fmt.Printf(" → Cleaned room B")
			}
		} else if action == "RIGHT" {
			currentState.Location = "B"
			fmt.Printf(" → Moved to room B")
		} else if action == "LEFT" {
			currentState.Location = "A"
			fmt.Printf(" → Moved to room A")
		}

		fmt.Printf("\n    Result: (%s, %s, %s)\n\n", currentState.Location, currentState.AState, currentState.BState)

		step++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("🎯 Total unique states visited: %d out of 8 possible states\n", len(visitedStates))
	fmt.Println("States visited:")
	for state := range visitedStates {
		fmt.Printf("  ✓ %s\n", state)
	}
}

// Model-based agent with memory to visit all states
type ModelBasedAgent struct {
	visitedStates map[string]bool
	targetStates  []State
}

func (agent *ModelBasedAgent) chooseAction(currentState State) string {
	stateKey := fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState)

	// Mark current state as visited
	if !agent.visitedStates[stateKey] {
		agent.visitedStates[stateKey] = true
		fmt.Printf("🆕 NEW STATE DISCOVERED: %s\n", stateKey)
	}

	// If we've visited all states, continue with simple reflex behavior
	if len(agent.visitedStates) >= 8 {
		fmt.Printf("✅ ALL STATES VISITED! Using simple reflex behavior now.\n")
		return reflexAgent(currentState.Location, currentState.AState, currentState.BState)
	}

	// Try to reach next unvisited state
	for _, targetState := range agent.targetStates {
		targetKey := fmt.Sprintf("(%s, %s, %s)", targetState.Location, targetState.AState, targetState.BState)
		if !agent.visitedStates[targetKey] {
			fmt.Printf("🎯 Trying to reach: %s\n", targetKey)
			return agent.planToReachState(currentState, targetState)
		}
	}

	// Fallback to simple reflex
	return reflexAgent(currentState.Location, currentState.AState, currentState.BState)
}

func (agent *ModelBasedAgent) planToReachState(current, target State) string {
	// Simple planning logic to reach target state

	// If we need to be in a different location, move there first
	if current.Location != target.Location {
		if current.Location == "A" {
			return "RIGHT"
		} else {
			return "LEFT"
		}
	}

	// If we're in the right location, adjust room states
	if current.Location == "A" {
		if target.AState == "CLEAN" && current.AState == "DIRTY" {
			return "CLEAN"
		}
		// If we need room A dirty but it's clean, we can't make it dirty
		// So move to explore other possibilities
		return "RIGHT"
	} else { // location == "B"
		if target.BState == "CLEAN" && current.BState == "DIRTY" {
			return "CLEAN"
		}
		// If we need room B dirty but it's clean, we can't make it dirty
		// So move to explore other possibilities
		return "LEFT"
	}
}

// Simulation that can visit all states by "resetting" the environment
func simulateAllStates() {
	fmt.Println("🌟 COMPLETE STATE EXPLORATION")
	fmt.Println("===============================")
	fmt.Println("Note: To visit ALL states, we'll simulate different starting conditions\n")

	// All possible starting states
	startingStates := []State{
		{"A", "DIRTY", "DIRTY"},
		{"A", "DIRTY", "CLEAN"},
		{"B", "DIRTY", "DIRTY"},
		{"B", "DIRTY", "CLEAN"},
	}

	allVisitedStates := make(map[string]bool)

	for i, startState := range startingStates {
		fmt.Printf("🚀 Scenario %d: Starting from (%s, %s, %s)\n",
			i+1, startState.Location, startState.AState, startState.BState)

		// Create model-based agent
		agent := &ModelBasedAgent{
			visitedStates: make(map[string]bool),
			targetStates: []State{
				{"A", "DIRTY", "DIRTY"},
				{"A", "CLEAN", "DIRTY"},
				{"A", "DIRTY", "CLEAN"},
				{"A", "CLEAN", "CLEAN"},
				{"B", "DIRTY", "DIRTY"},
				{"B", "CLEAN", "DIRTY"},
				{"B", "DIRTY", "CLEAN"},
				{"B", "CLEAN", "CLEAN"},
			},
		}

		currentState := startState
		for step := 1; step <= 8; step++ {
			stateKey := fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState)
			allVisitedStates[stateKey] = true

			action := agent.chooseAction(currentState)
			fmt.Printf("  Step %d: %s → Action: %s", step, stateKey, action)

			// Apply action
			if action == "CLEAN" {
				if currentState.Location == "A" {
					currentState.AState = "CLEAN"
				} else {
					currentState.BState = "CLEAN"
				}
			} else if action == "RIGHT" {
				currentState.Location = "B"
			} else if action == "LEFT" {
				currentState.Location = "A"
			}

			fmt.Printf(" → (%s, %s, %s)\n", currentState.Location, currentState.AState, currentState.BState)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()
	}

	fmt.Printf("🎉 TOTAL UNIQUE STATES VISITED: %d/8\n", len(allVisitedStates))
	fmt.Println("States visited across all scenarios:")
	for state := range allVisitedStates {
		fmt.Printf("  ✅ %s\n", state)
	}
}

// The main function initializes the states and starts the vacuum cleaner simulation
func main() {
	fmt.Println("🤖 VACUUM CLEANER AGENT SIMULATION")
	fmt.Println("==================================")

	// Show systematic exploration of all states
	visitAllStates()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🔄 SIMPLE REFLEX AGENT LIMITATIONS")
	fmt.Println("Note: A simple reflex agent cannot visit all 8 states")
	fmt.Println("from a single run because it cannot make rooms dirty again.")

	// Start limited simulation to show reachable states
	initialState := State{Location: "A", AState: "DIRTY", BState: "DIRTY"}
	runLimitedSimulation(initialState, 8)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Show how to visit all states with different approaches
	simulateAllStates()
}
