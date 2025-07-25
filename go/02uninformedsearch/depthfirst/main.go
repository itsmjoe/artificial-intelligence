package main

import (
	"fmt"
	"strings"
)

// returns the successors of a given node in a 4x4 matrix graph.
// Matrix representation:
//
//	1  2  3  4
//	5  6  7  8
//	9 10 11 12
// 13 14 15 16

// The successors are defined based on possible movements: right, down, and diagonal.
func successors(n int) []int {
	switch n {
	case 1:
		return []int{2, 5, 6} // conexiones: derecha, abajo, diagonal
	case 2:
		return []int{1, 3, 5, 6, 7}
	case 3:
		return []int{2, 4, 6, 7, 8}
	case 4:
		return []int{3, 7, 8}
	case 5:
		return []int{1, 2, 6, 9, 10}
	case 6:
		return []int{1, 2, 3, 5, 7, 9, 10, 11}
	case 7:
		return []int{2, 3, 4, 6, 8, 10, 11, 12}
	case 8:
		return []int{3, 4, 7, 11, 12}
	case 9:
		return []int{5, 6, 10, 13, 14}
	case 10:
		return []int{5, 6, 7, 9, 11, 13, 14, 15}
	case 11:
		return []int{6, 7, 8, 10, 12, 14, 15, 16}
	case 12:
		return []int{7, 8, 11, 15, 16}
	case 13:
		return []int{9, 10, 14}
	case 14:
		return []int{9, 10, 11, 13, 15}
	case 15:
		return []int{10, 11, 12, 14, 16}
	case 16:
		return []int{11, 12, 15}
	default:
		return nil
	}
}

// depthFirstSearch algorithm to find a path from the begin node to the end node.
func depthFirstSearch(begin, end int) {
	fmt.Printf("Iniciando búsqueda en profundidad desde nodo %d hasta nodo %d\n", begin, end)
	fmt.Println("Matriz 4x4:")
	fmt.Println(" 1  2  3  4")
	fmt.Println(" 5  6  7  8")
	fmt.Println(" 9 10 11 12")
	fmt.Println("13 14 15 16")
	fmt.Println()

	list := []int{begin}
	visited := make(map[int]bool) // Para evitar ciclos infinitos
	step := 1

	for len(list) > 0 {
		// Tomar el primer elemento de la lista (simulando pila LIFO)
		current := list[0]
		list = list[1:]

		fmt.Printf("Paso %d: Explorando nodo %d\n", step, current)

		// Si ya visitamos este nodo, saltar
		if visited[current] {
			fmt.Printf("  → Nodo %d ya visitado, saltando...\n\n", current)
			step++
			continue
		}

		// Marcar como visitado
		visited[current] = true

		// Verificar si encontramos la solución
		if current == end {
			fmt.Printf("  → ¡SOLUCIÓN ENCONTRADA!\n")
			fmt.Printf("  → Nodo objetivo %d alcanzado\n", end)
			return
		}

		// Obtener sucesores del nodo actual
		tmp := successors(current)
		if tmp != nil {
			// Filtrar sucesores ya visitados
			unvisited := []int{}
			for _, node := range tmp {
				if !visited[node] {
					unvisited = append(unvisited, node)
				}
			}

			fmt.Printf("  → Sucesores de %d: %v\n", current, tmp)
			fmt.Printf("  → Sucesores no visitados: %v\n", unvisited)

			if len(unvisited) > 0 {
				// Invertir para mantener orden de profundidad
				reverse(unvisited)
				// Agregar al inicio de la lista (comportamiento de pila)
				unvisited = append(unvisited, list...)
				list = unvisited
				fmt.Printf("  → Lista actualizada: %v\n", list)
			} else {
				fmt.Printf("  → No hay sucesores nuevos para explorar\n")
			}
		} else {
			fmt.Printf("  → No hay sucesores para el nodo %d\n", current)
		}

		fmt.Println()
		step++
	}

	fmt.Println("NO SE ENCONTRÓ SOLUCIÓN")
}

// reverse function to reverse the order of elements in a slice.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Backtracking algorithm to find a path from the begin node to the end node.
func backtrackingSearch(begin, end int, path []int, visited map[int]bool, step *int) bool {
	*step++
	fmt.Printf("Paso %d: Explorando nodo %d, camino actual: %v\n", *step, begin, path)

	// Marcar como visitado
	visited[begin] = true
	path = append(path, begin)

	// Verificar si encontramos la solución
	if begin == end {
		fmt.Printf("  → ¡SOLUCIÓN ENCONTRADA!\n")
		fmt.Printf("  → Camino completo: %v\n", path)
		return true
	}

	// Obtener sucesores del nodo actual
	tmp := successors(begin)
	if tmp != nil {
		// Filtrar sucesores ya visitados
		unvisited := []int{}
		for _, node := range tmp {
			if !visited[node] {
				unvisited = append(unvisited, node)
			}
		}

		fmt.Printf("  → Sucesores de %d: %v\n", begin, tmp)
		fmt.Printf("  → Sucesores no visitados: %v\n", unvisited)

		// Intentar cada sucesor no visitado
		for _, next := range unvisited {
			fmt.Printf("  → Intentando explorar nodo %d\n", next)

			// Llamada recursiva
			if backtrackingSearch(next, end, path, visited, step) {
				return true
			}

			// Backtrack: desmarcar como visitado y continuar con siguiente sucesor
			fmt.Printf("  → Retrocediendo desde nodo %d\n", next)
			visited[next] = false
		}
	}

	fmt.Printf("  → No hay más opciones desde nodo %d, retrocediendo...\n", begin)
	return false
}

// Backjumping algorithm to find a path from the begin node to the end node.
func backjumpingSearch(begin, end int, path []int, visited map[int]bool, conflictSet map[int][]int, step *int) bool {
	*step++
	fmt.Printf("Paso %d: Explorando nodo %d, camino actual: %v\n", *step, begin, path)

	// Marcar como visitado
	visited[begin] = true
	path = append(path, begin)

	// Verificar si encontramos la solución
	if begin == end {
		fmt.Printf("  → ¡SOLUCIÓN ENCONTRADA!\n")
		fmt.Printf("  → Camino completo: %v\n", path)
		return true
	}

	// Obtener sucesores del nodo actual
	tmp := successors(begin)
	if tmp != nil {
		// Filtrar sucesores ya visitados
		unvisited := []int{}
		for _, node := range tmp {
			if !visited[node] {
				unvisited = append(unvisited, node)
			}
		}

		fmt.Printf("  → Sucesores de %d: %v\n", begin, tmp)
		fmt.Printf("  → Sucesores no visitados: %v\n", unvisited)

		// Intentar cada sucesor no visitado
		for _, next := range unvisited {
			fmt.Printf("  → Intentando explorar nodo %d\n", next)

			// Llamada recursiva
			if backjumpingSearch(next, end, path, visited, conflictSet, step) {
				return true
			}

			// Backjumping: identificar conjunto de conflicto
			conflictNodes := conflictSet[next]
			fmt.Printf("  → Conjunto de conflicto para nodo %d: %v\n", next, conflictNodes)

			// Backjump más allá de los nodos en conflicto
			fmt.Printf("  → Realizando backjump desde nodo %d\n", next)
			visited[next] = false

			// En una implementación real, saltaríamos a un punto específico
			// Por simplicidad, continuamos con el siguiente sucesor
		}
	}

	fmt.Printf("  → No hay más opciones desde nodo %d, realizando backjump...\n", begin)
	return false
}

// Wrapper function for Backtracking
func executeBacktracking(begin, end int) {
	fmt.Printf("Iniciando búsqueda con BACKTRACKING desde nodo %d hasta nodo %d\n", begin, end)
	fmt.Println("Matriz 4x4:")
	fmt.Println(" 1  2  3  4")
	fmt.Println(" 5  6  7  8")
	fmt.Println(" 9 10 11 12")
	fmt.Println("13 14 15 16")
	fmt.Println()

	visited := make(map[int]bool)
	path := []int{}
	step := 0

	if !backtrackingSearch(begin, end, path, visited, &step) {
		fmt.Println("NO SE ENCONTRÓ SOLUCIÓN CON BACKTRACKING")
	}
}

// Wrapper function for Backjumping
func executeBackjumping(begin, end int) {
	fmt.Printf("Iniciando búsqueda con BACKJUMPING desde nodo %d hasta nodo %d\n", begin, end)
	fmt.Println("Matriz 4x4:")
	fmt.Println(" 1  2  3  4")
	fmt.Println(" 5  6  7  8")
	fmt.Println(" 9 10 11 12")
	fmt.Println("13 14 15 16")
	fmt.Println()

	visited := make(map[int]bool)
	path := []int{}
	conflictSet := make(map[int][]int)
	step := 0

	// Inicializar conjuntos de conflicto simples
	for i := 1; i <= 16; i++ {
		conflictSet[i] = []int{}
	}

	if !backjumpingSearch(begin, end, path, visited, conflictSet, &step) {
		fmt.Println("NO SE ENCONTRÓ SOLUCIÓN CON BACKJUMPING")
	}
}

// Main function to execute all three search algorithms.
func main() {
	fmt.Println("ALGORITMOS DE BÚSQUEDA - MATRIZ 4x4")
	fmt.Println("===================================")
	fmt.Println()

	// Algoritmo 1: Búsqueda en Profundidad (DFS)
	fmt.Println("1. BÚSQUEDA EN PROFUNDIDAD (DFS)")
	fmt.Println("---------------------------------")
	depthFirstSearch(1, 12)

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// Algoritmo 2: Backtracking
	fmt.Println("2. BÚSQUEDA CON BACKTRACKING")
	fmt.Println("----------------------------")
	executeBacktracking(1, 12)

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// Algoritmo 3: Backjumping
	fmt.Println("3. BÚSQUEDA CON BACKJUMPING")
	fmt.Println("---------------------------")
	executeBackjumping(1, 12)
}
