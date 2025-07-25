/*
SIMULACIÓN DE AGENTE ASPIRADORA - ANÁLISIS DE ESTADOS

Este programa simula una aspiradora inteligente que opera en un mundo de dos habitaciones (A y B).
Demuestra las diferencias entre diferentes tipos de agentes de inteligencia artificial:

1. AGENTE REFLEXIVO SIMPLE:
   - Solo considera el estado actual (ubicación + estado de habitaciones)
   - No tiene memoria de estados anteriores
   - Usa reglas simples condición-acción
   - Comportamiento determinístico

2. AGENTE BASADO EN MODELOS:
   - Mantiene memoria de estados visitados
   - Puede planificar acciones para alcanzar objetivos
   - Más sofisticado que el agente reflexivo simple

ESPACIO DE ESTADOS:
El mundo tiene 8 estados posibles:
- (A,DIRTY,DIRTY), (A,CLEAN,DIRTY), (A,DIRTY,CLEAN), (A,CLEAN,CLEAN)
- (B,DIRTY,DIRTY), (B,CLEAN,DIRTY), (B,DIRTY,CLEAN), (B,CLEAN,CLEAN)

LIMITACIÓN CLAVE:
Un agente reflexivo simple no puede visitar todos los estados en una sola ejecución
porque no puede ensuciar habitaciones que ya están limpias.
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

// State representa el estado completo del mundo de la aspiradora
// Incluye la ubicación de la aspiradora y el estado de ambas habitaciones
type State struct {
	Location string // "A" o "B" - dónde está la aspiradora
	AState   string // "DIRTY" o "CLEAN" - estado de la habitación A
	BState   string // "DIRTY" o "CLEAN" - estado de la habitación B
}

// Esta función determina la acción basándose en la ubicación actual y el estado de las habitaciones
// Es un agente reflexivo simple porque solo considera el estado presente sin memoria
func reflexAgent(location, aState, bState string) string {
	// Prioridad: Limpiar habitaciones sucias primero, luego moverse para explorar
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

// Función que muestra sistemáticamente todos los 8 estados posibles
// Esto es útil para entender el espacio completo de estados del mundo de la aspiradora
func visitAllStates() {
	fmt.Println("=== EXPLORACIÓN SISTEMÁTICA DE LOS 8 ESTADOS ===\n")

	// Definir todos los estados posibles en el mundo de la aspiradora
	// Cada estado se representa como (Ubicación, Estado_A, Estado_B)
	allStates := []State{
		{"A", "DIRTY", "DIRTY"}, // Aspiradora en A, ambas habitaciones sucias
		{"A", "CLEAN", "DIRTY"}, // Aspiradora en A, A limpia, B sucia
		{"A", "DIRTY", "CLEAN"}, // Aspiradora en A, A sucia, B limpia
		{"A", "CLEAN", "CLEAN"}, // Aspiradora en A, ambas habitaciones limpias
		{"B", "DIRTY", "DIRTY"}, // Aspiradora en B, ambas habitaciones sucias
		{"B", "CLEAN", "DIRTY"}, // Aspiradora en B, A limpia, B sucia
		{"B", "DIRTY", "CLEAN"}, // Aspiradora en B, A sucia, B limpia
		{"B", "CLEAN", "CLEAN"}, // Aspiradora en B, ambas habitaciones limpias
	}

	fmt.Println("Todos los estados posibles en el mundo de la aspiradora:")
	for i, state := range allStates {
		fmt.Printf("%d. (%s, %s, %s) - ", i+1, state.Location, state.AState, state.BState)

		// Determinar qué acción tomaría el agente en este estado
		action := reflexAgent(state.Location, state.AState, state.BState)
		fmt.Printf("El agente eligiría: %s\n", action)

		// Mostrar el estado resultante después de la acción
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

		fmt.Printf("   → Siguiente estado: (%s, %s, %s)\n\n", nextState.Location, nextState.AState, nextState.BState)
		time.Sleep(1 * time.Second)
	}
}

// Esta función simula la operación de la aspiradora y rastrea los estados visitados
// Demuestra las limitaciones del agente reflexivo simple en una ejecución continua
func runLimitedSimulation(initialState State, maxSteps int) {
	// Mapa para rastrear qué estados ya hemos visitado
	visitedStates := make(map[string]bool)
	currentState := initialState
	step := 1

	fmt.Printf("\n=== SIMULACIÓN LIMITADA (Máximo %d pasos) ===\n", maxSteps)
	fmt.Println("Comenzando desde:", fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState))
	fmt.Println()

	for step <= maxSteps {
		// Crear representación en cadena del estado actual
		stateKey := fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState)

		// Verificar si ya hemos visitado este estado antes
		if !visitedStates[stateKey] {
			visitedStates[stateKey] = true
			fmt.Printf("Paso %d - NUEVO ESTADO: %s\n", step, stateKey)
		} else {
			fmt.Printf("Paso %d - Revisitando: %s\n", step, stateKey)
		}

		// Obtener acción del agente reflexivo
		action := reflexAgent(currentState.Location, currentState.AState, currentState.BState)
		fmt.Printf("    Acción: %s", action)

		// Aplicar la acción y actualizar el estado
		if action == "CLEAN" {
			if currentState.Location == "A" {
				currentState.AState = "CLEAN"
				fmt.Printf(" → Se limpió la habitación A")
			} else if currentState.Location == "B" {
				currentState.BState = "CLEAN"
				fmt.Printf(" → Se limpió la habitación B")
			}
		} else if action == "RIGHT" {
			currentState.Location = "B"
			fmt.Printf(" → Se movió a la habitación B")
		} else if action == "LEFT" {
			currentState.Location = "A"
			fmt.Printf(" → Se movió a la habitación A")
		}

		fmt.Printf("\n    Resultado: (%s, %s, %s)\n\n", currentState.Location, currentState.AState, currentState.BState)

		step++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Total de estados únicos visitados: %d de 8 estados posibles\n", len(visitedStates))
	fmt.Println("Estados visitados:")
	for state := range visitedStates {
		fmt.Printf("  %s\n", state)
	}
}

// Agente basado en modelos con memoria para intentar visitar todos los estados
// A diferencia del agente reflexivo simple, este agente puede recordar estados visitados
type ModelBasedAgent struct {
	visitedStates map[string]bool // Memoria de estados visitados
	targetStates  []State         // Lista de estados objetivo a alcanzar
}

// Función que elige la acción basándose en el estado actual y la memoria
func (agent *ModelBasedAgent) chooseAction(currentState State) string {
	stateKey := fmt.Sprintf("(%s, %s, %s)", currentState.Location, currentState.AState, currentState.BState)

	// Marcar el estado actual como visitado en la memoria
	if !agent.visitedStates[stateKey] {
		agent.visitedStates[stateKey] = true
		fmt.Printf("NUEVO ESTADO DESCUBIERTO: %s\n", stateKey)
	}

	// Si hemos visitado todos los estados, continuar con comportamiento reflexivo simple
	if len(agent.visitedStates) >= 8 {
		fmt.Printf("TODOS LOS ESTADOS VISITADOS! Usando comportamiento reflexivo simple.\n")
		return reflexAgent(currentState.Location, currentState.AState, currentState.BState)
	}

	// Intentar alcanzar el siguiente estado no visitado
	for _, targetState := range agent.targetStates {
		targetKey := fmt.Sprintf("(%s, %s, %s)", targetState.Location, targetState.AState, targetState.BState)
		if !agent.visitedStates[targetKey] {
			fmt.Printf("Intentando alcanzar: %s\n", targetKey)
			return agent.planToReachState(currentState, targetState)
		}
	}

	// Respaldo: usar agente reflexivo simple
	return reflexAgent(currentState.Location, currentState.AState, currentState.BState)
}

// Función de planificación simple para intentar alcanzar un estado objetivo
func (agent *ModelBasedAgent) planToReachState(current, target State) string {
	// Lógica de planificación simple para alcanzar el estado objetivo

	// Si necesitamos estar en una ubicación diferente, moverse allí primero
	if current.Location != target.Location {
		if current.Location == "A" {
			return "RIGHT"
		} else {
			return "LEFT"
		}
	}

	// Si estamos en la ubicación correcta, ajustar los estados de las habitaciones
	if current.Location == "A" {
		if target.AState == "CLEAN" && current.AState == "DIRTY" {
			return "CLEAN"
		}
		// Si necesitamos la habitación A sucia pero está limpia, no podemos ensuciarla
		// Así que nos movemos para explorar otras posibilidades
		return "RIGHT"
	} else { // location == "B"
		if target.BState == "CLEAN" && current.BState == "DIRTY" {
			return "CLEAN"
		}
		// Si necesitamos la habitación B sucia pero está limpia, no podemos ensuciarla
		// Así que nos movemos para explorar otras posibilidades
		return "LEFT"
	}
}

// Simulación que puede visitar todos los estados usando diferentes condiciones iniciales
// Esta función demuestra que se necesitan múltiples escenarios para explorar completamente
func simulateAllStates() {
	fmt.Println("EXPLORACIÓN COMPLETA DE ESTADOS")
	fmt.Println("===============================")
	fmt.Println("Nota: Para visitar TODOS los estados, simularemos diferentes condiciones iniciales\n")

	// Todos los posibles estados iniciales que incluyen habitaciones sucias
	// (necesitamos habitaciones sucias para poder limpiarlas y crear transiciones)
	startingStates := []State{
		{"A", "DIRTY", "DIRTY"}, // Empezar en A con ambas sucias
		{"A", "DIRTY", "CLEAN"}, // Empezar en A con A sucia, B limpia
		{"B", "DIRTY", "DIRTY"}, // Empezar en B con ambas sucias
		{"B", "DIRTY", "CLEAN"}, // Empezar en B con A sucia, B limpia
	}

	// Mapa global para rastrear todos los estados visitados en todos los escenarios
	allVisitedStates := make(map[string]bool)

	for i, startState := range startingStates {
		fmt.Printf("Escenario %d: Comenzando desde (%s, %s, %s)\n",
			i+1, startState.Location, startState.AState, startState.BState)

		// Crear agente basado en modelos para este escenario
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
			fmt.Printf("  Paso %d: %s → Acción: %s", step, stateKey, action)

			// Aplicar acción
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

	fmt.Printf("TOTAL DE ESTADOS ÚNICOS VISITADOS: %d/8\n", len(allVisitedStates))
	fmt.Println("Estados visitados en todos los escenarios:")
	for state := range allVisitedStates {
		fmt.Printf("  %s\n", state)
	}
}

// Exploración exhaustiva que visita TODOS los 8 estados
// Esta función demuestra conceptualmente cómo se puede alcanzar cada estado
func completeStateExploration() {
	fmt.Println("\nEXPLORACIÓN COMPLETA - TODOS LOS 8 ESTADOS")
	fmt.Println("=========================================")
	fmt.Println("Estrategia: Visitar manualmente cada estado para demostrar todas las posibilidades\n")

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

	fmt.Println("Visitando todos los 8 estados posibles:")
	for i, state := range allStates {
		fmt.Printf("%d. Estado: (%s, %s, %s)\n", i+1, state.Location, state.AState, state.BState)

		// Mostrar qué haría el agente reflexivo simple
		action := reflexAgent(state.Location, state.AState, state.BState)
		fmt.Printf("   El Agente Reflexivo Simple elegiría: %s\n", action)

		// Mostrar la transición
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
		fmt.Printf("   → Transicionaría a: (%s, %s, %s)\n", nextState.Location, nextState.AState, nextState.BState)

		// Mostrar cómo se puede alcanzar este estado
		fmt.Printf("   Cómo alcanzar este estado:\n")
		if state.Location == "A" && state.AState == "DIRTY" && state.BState == "DIRTY" {
			fmt.Printf("      - Estado inicial (ambas habitaciones sucias, empezar en A)\n")
		} else if state.Location == "A" && state.AState == "CLEAN" && state.BState == "DIRTY" {
			fmt.Printf("      - Desde (A,DIRTY,DIRTY): acción CLEAN\n")
		} else if state.Location == "A" && state.AState == "DIRTY" && state.BState == "CLEAN" {
			fmt.Printf("      - Empezar con A sucia, B limpia (estado inicial alternativo)\n")
		} else if state.Location == "A" && state.AState == "CLEAN" && state.BState == "CLEAN" {
			fmt.Printf("      - Desde (A,DIRTY,CLEAN): acción CLEAN, o desde (B,CLEAN,CLEAN): acción LEFT\n")
		} else if state.Location == "B" && state.AState == "DIRTY" && state.BState == "DIRTY" {
			fmt.Printf("      - Empezar con ambas sucias en B, o desde (A,DIRTY,DIRTY): acción RIGHT\n")
		} else if state.Location == "B" && state.AState == "CLEAN" && state.BState == "DIRTY" {
			fmt.Printf("      - Desde (A,CLEAN,DIRTY): acción RIGHT\n")
		} else if state.Location == "B" && state.AState == "DIRTY" && state.BState == "CLEAN" {
			fmt.Printf("      - Desde (B,DIRTY,DIRTY): acción CLEAN\n")
		} else if state.Location == "B" && state.AState == "CLEAN" && state.BState == "CLEAN" {
			fmt.Printf("      - Desde (B,CLEAN,DIRTY): acción CLEAN, o desde (A,CLEAN,CLEAN): acción RIGHT\n")
		}

		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("TODOS LOS 8 ESTADOS DEMOSTRADOS!")
	fmt.Println("\nCONCLUSIONES CLAVE:")
	fmt.Println("• Un agente reflexivo simple PUEDE visitar todos los estados, pero no en una sola ejecución continua")
	fmt.Println("• Se necesitan diferentes condiciones iniciales para alcanzar ciertos estados")
	fmt.Println("• La limitación es que las habitaciones no pueden volver a ensuciarse una vez limpias")
	fmt.Println("• Para visitar todos los estados sistemáticamente, necesitamos:")
	fmt.Println("  - Múltiples condiciones iniciales, O")
	fmt.Println("  - Un agente que pueda 'reiniciar' el entorno, O")
	fmt.Println("  - Un agente más sofisticado con capacidades de planificación")
}

// La función principal inicializa los estados y comienza la simulación de la aspiradora
// Demuestra las diferencias entre agentes reflexivos simples y agentes más sofisticados
func main() {
	fmt.Println("SIMULACIÓN DE AGENTE ASPIRADORA")
	fmt.Println("===============================")

	// Mostrar exploración sistemática de todos los estados
	visitAllStates()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("LIMITACIONES DEL AGENTE REFLEXIVO SIMPLE")
	fmt.Println("Nota: Un agente reflexivo simple no puede visitar todos los 8 estados")
	fmt.Println("en una sola ejecución porque no puede ensuciar las habitaciones nuevamente.")

	// Comenzar simulación limitada para mostrar estados alcanzables
	initialState := State{Location: "A", AState: "DIRTY", BState: "DIRTY"}
	runLimitedSimulation(initialState, 8)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Mostrar cómo visitar todos los estados con diferentes enfoques
	simulateAllStates()

	// Exploración completa mostrando todos los estados
	completeStateExploration()
}
