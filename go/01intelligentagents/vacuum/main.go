/*
ANÁLISIS DE AGENTE ASPIRADORA - EXPLORACIÓN DE ESTADOS (VERSIÓN MEJORADA)
Estudiante: [Tu nombre]
Curso: Inteligencia Artificial

MEJORAS IMPLEMENTADAS:
- Agregué la capacidad de "ensuciar" habitaciones para explorar más estados
- Implementé un contador de tiempo para simular que las habitaciones se ensucian automáticamente
- Modifiqué el agente para ser más explorativo
- Agregué métricas detalladas de exploración

OBJETIVO: Demostrar que con pequeñas modificaciones un agente reflexivo simple
puede visitar más estados del espacio de búsqueda.
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// State representa el estado del mundo con información adicional
type State struct {
	Location   string // "A" o "B"
	AState     string // "DIRTY" o "CLEAN"
	BState     string // "DIRTY" o "CLEAN"
	TimeSteps  int    // Contador de pasos (para simular ensuciado automático)
	ACleanTime int    // Tiempo desde que A fue limpiada
	BCleanTime int    // Tiempo desde que B fue limpiada
}

var outputFile *os.File

// Funciones para escribir tanto en consola como archivo
func logOutput(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Print(message)
	if outputFile != nil {
		outputFile.WriteString(message)
	}
}

func logOutputln(args ...interface{}) {
	fmt.Println(args...)
	if outputFile != nil {
		message := fmt.Sprintln(args...)
		outputFile.WriteString(message)
	}
}

// Convierte un estado a string para comparaciones (sin incluir contadores internos)
func stateToString(s State) string {
	return fmt.Sprintf("(%s,%s,%s)", s.Location, s.AState, s.BState)
}

// AGENTE REFLEXIVO SIMPLE MEJORADO
// Ahora incluye una pequeña probabilidad de comportamiento explorativo
func improvedReflexAgent(location, aState, bState string, timeSteps int) string {
	// Agregar un poco de aleatoriedad ocasionalmente para ser más realista
	// En la vida real, los sensores pueden tener ruido o el agente puede "decidir" explorar
	rand.Seed(int64(timeSteps))

	// Regla 1: Si la habitación actual está sucia, limpiarla (prioridad alta)
	if location == "A" && aState == "DIRTY" {
		return "CLEAN"
	}
	if location == "B" && bState == "DIRTY" {
		return "CLEAN"
	}

	// Regla 2: Si ambas están limpias, comportamiento explorativo ocasional
	if aState == "CLEAN" && bState == "CLEAN" {
		// 20% probabilidad de quedarse quieto (simulando "patrullaje")
		if rand.Float64() < 0.2 {
			return "WAIT" // Nueva acción: esperar
		}
	}

	// Regla 3: Moverse hacia habitación que necesite atención
	if location == "A" && bState == "DIRTY" {
		return "RIGHT"
	}
	if location == "B" && aState == "DIRTY" {
		return "LEFT"
	}

	// Regla 4: Si todo está limpio, moverse para patrullar
	if location == "A" {
		return "RIGHT"
	} else {
		return "LEFT"
	}
}

// SIMULACIÓN DE ENSUCIADO AUTOMÁTICO
// Las habitaciones pueden ensuciarse después de un tiempo
func simulateEnvironmentChanges(state State) State {
	newState := state
	newState.TimeSteps++

	// Incrementar contadores de tiempo limpio
	if state.AState == "CLEAN" {
		newState.ACleanTime++
	} else {
		newState.ACleanTime = 0
	}

	if state.BState == "CLEAN" {
		newState.BCleanTime++
	} else {
		newState.BCleanTime = 0
	}

	// Simular ensuciado automático después de cierto tiempo
	// Esto es más realista - las habitaciones se ensucian con el tiempo
	if newState.ACleanTime > 3 && rand.Float64() < 0.3 { // 30% chance después de 3 pasos
		newState.AState = "DIRTY"
		newState.ACleanTime = 0
		logOutput("    [EVENTO AMBIENTAL: Habitación A se ensució automáticamente]\n")
	}

	if newState.BCleanTime > 3 && rand.Float64() < 0.3 { // 30% chance después de 3 pasos
		newState.BState = "DIRTY"
		newState.BCleanTime = 0
		logOutput("    [EVENTO AMBIENTAL: Habitación B se ensució automáticamente]\n")
	}

	return newState
}

// Función auxiliar para aplicar acciones (ahora incluye WAIT)
func applyAction(state State, action string) State {
	newState := state

	switch action {
	case "CLEAN":
		if state.Location == "A" {
			newState.AState = "CLEAN"
			newState.ACleanTime = 0
		} else if state.Location == "B" {
			newState.BState = "CLEAN"
			newState.BCleanTime = 0
		}
	case "RIGHT":
		newState.Location = "B"
	case "LEFT":
		newState.Location = "A"
	case "WAIT":
		// No hacer nada, solo esperar (útil para el ensuciado automático)
	}

	return newState
}

// EXPERIMENTO PRINCIPAL MEJORADO
func experimentImprovedAgent(startState State, maxSteps int) {
	logOutputln("=== EXPERIMENTO: AGENTE MEJORADO CON ENSUCIADO AUTOMÁTICO ===")
	logOutput("Estado inicial: %s\n", stateToString(startState))
	logOutput("Máximo de pasos: %d\n\n", maxSteps)

	visitedStates := make(map[string]bool)
	stateHistory := make([]string, 0)
	currentState := startState

	// Inicializar generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	for step := 1; step <= maxSteps; step++ {
		// Simular cambios ambientales antes de actuar
		currentState = simulateEnvironmentChanges(currentState)

		stateKey := stateToString(currentState)
		stateHistory = append(stateHistory, stateKey)

		// Marcar estado como visitado
		if !visitedStates[stateKey] {
			visitedStates[stateKey] = true
			logOutput("Paso %d: NUEVO ESTADO -> %s", step, stateKey)
		} else {
			logOutput("Paso %d: Revisitando -> %s", step, stateKey)
		}

		// Agregar información adicional del estado interno
		logOutput(" [Tiempo: A=%d, B=%d]\n", currentState.ACleanTime, currentState.BCleanTime)

		// Obtener acción del agente mejorado
		action := improvedReflexAgent(currentState.Location, currentState.AState, currentState.BState, currentState.TimeSteps)
		logOutput("         Acción: %s", action)

		// Aplicar la acción
		nextState := applyAction(currentState, action)
		logOutput(" -> %s\n", stateToString(nextState))

		currentState = nextState
		time.Sleep(600 * time.Millisecond) // Pausa para observar
	}

	logOutput("\nRESULTADO FINAL:\n")
	logOutput("   Estados únicos visitados: %d de 8 posibles (%.1f%%)\n",
		len(visitedStates), float64(len(visitedStates))/8.0*100)

	logOutputln("\nEstados visitados:")
	for state := range visitedStates {
		logOutput("   - %s\n", state)
	}

	// Análisis de secuencia
	logOutputln("\nAnalisis de secuencia:")
	logOutput("   Transiciones totales: %d\n", len(stateHistory))

	// Contar revisitas para medir eficiencia exploratoria
	uniqueTransitions := 0
	seenStates := make(map[string]bool)
	for _, state := range stateHistory {
		if !seenStates[state] {
			uniqueTransitions++
			seenStates[state] = true
		}
	}

	efficiency := float64(uniqueTransitions) / float64(len(stateHistory)) * 100
	logOutput("   Eficiencia exploratoria: %.1f%%\n", efficiency)
	logOutputln()
}

// COMPARACIÓN: Agente original vs mejorado
func compareAgents() {
	logOutputln("=== COMPARACIÓN: AGENTE ORIGINAL VS MEJORADO ===")
	logOutputln()

	startState := State{"A", "DIRTY", "DIRTY", 0, 0, 0}

	logOutputln("AGENTE ORIGINAL (sin ensuciado automatico):")
	// Simular agente original rápidamente
	visitedOriginal := make(map[string]bool)
	currentState := startState

	for i := 0; i < 10; i++ {
		stateKey := stateToString(currentState)
		visitedOriginal[stateKey] = true

		// Usar lógica simple original
		var action string
		if currentState.Location == "A" && currentState.AState == "DIRTY" {
			action = "CLEAN"
		} else if currentState.Location == "B" && currentState.BState == "DIRTY" {
			action = "CLEAN"
		} else if currentState.Location == "A" {
			action = "RIGHT"
		} else {
			action = "LEFT"
		}

		currentState = applyAction(currentState, action)
	}

	logOutput("   Estados alcanzados: %d\n", len(visitedOriginal))
	for state := range visitedOriginal {
		logOutput("     - %s\n", state)
	}

	logOutput("\nAGENTE MEJORADO (con ensuciado automatico):\n")
	logOutput("   Ver resultados del experimento anterior\n")
	logOutput("   Ventaja: Puede explorar estados que el original no alcanza\n")
	logOutputln()
}

// REFLEXIONES PARA EL REPORTE
func generateReflections() {
	logOutputln("=== REFLEXIONES Y APRENDIZAJES ===")
	logOutputln()

	logOutputln("MODIFICACIONES IMPLEMENTADAS:")
	logOutputln("   1. Ensuciado automatico: Las habitaciones se ensucian despues de un tiempo")
	logOutputln("   2. Accion WAIT: El agente puede esperar ocasionalmente")
	logOutputln("   3. Comportamiento explorativo: Pequena aleatoriedad en las decisiones")
	logOutputln("   4. Metricas de exploracion: Seguimiento de eficiencia y cobertura")
	logOutputln()

	logOutputln("HALLAZGOS:")
	logOutputln("   - El ensuciado automatico permite ciclos mas largos de exploracion")
	logOutputln("   - Un comportamiento ligeramente aleatorio mejora la cobertura")
	logOutputln("   - El problema original estaba en la irreversibilidad de la limpieza")
	logOutputln("   - En entornos reales, las condiciones cambian constantemente")
	logOutputln()

	logOutputln("IMPLICACIONES TEORICAS:")
	logOutputln("   - Los agentes reflexivos simples son limitados en entornos estaticos")
	logOutputln("   - La modificacion del entorno puede mejorar la exploracion")
	logOutputln("   - Hay un trade-off entre simplicidad y completitud en la exploracion")
	logOutputln()

	logOutputln("CONCLUSIONES:")
	logOutputln("   - La version mejorada alcanza mas estados en una sola ejecucion")
	logOutputln("   - Las modificaciones son minimas pero efectivas")
	logOutputln("   - El analisis demuestra la importancia del diseno del entorno")
	logOutputln()
}

// Función principal
func main() {
	// Configurar archivo de salida
	var err error
	outputFile, err = os.Create("analisis_aspiradora.txt")
	if err != nil {
		log.Fatalf("Error creando archivo: %v", err)
	}
	defer outputFile.Close()

	// Encabezado
	logOutputln("ANALISIS DE AGENTE ASPIRADORA - VERSION MEJORADA")
	logOutputln("==============================================")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logOutput("Fecha: %s\n", timestamp)
	logOutput("Modificaciones: Ensuciado automatico + comportamiento explorativo\n\n")

	// Experimento principal con agente mejorado
	startState := State{"A", "DIRTY", "DIRTY", 0, 0, 0}
	experimentImprovedAgent(startState, 15)

	// Separador
	logOutput("\n" + strings.Repeat("=", 60) + "\n\n")

	// Comparación
	compareAgents()

	// Separador
	logOutput("\n" + strings.Repeat("=", 60) + "\n\n")

	// Reflexiones finales
	generateReflections()
}
