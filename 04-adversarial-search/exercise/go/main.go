package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var idCounter int = 0
var dotLines []string

func inc() int {
	idCounter++
	return idCounter
}

// Nodo representa un nodo en el árbol: [id, heurística]
type Nodo struct {
	id         int
	heuristica int
}

func addNode(id int, label string, extra string) {
	// label más descriptivo: mostrar heurística y id
	labelWithID := fmt.Sprintf("%s\\n(id:%d)", label, id)
	line := fmt.Sprintf("    %d [label=\"%s\"%s];", id, labelWithID, extra)
	dotLines = append(dotLines, line)
}

func addEdge(from, to int, label string) {
	line := fmt.Sprintf("    %d -- %d [label=\"%s\"];", from, to, label)
	dotLines = append(dotLines, line)
}

func succ(node Nodo, depth int, bfactor int, arrow string) []Nodo {
	var children []Nodo
	for i := 0; i < bfactor; i++ {
		childID := inc()
		heuristic := 0
		if depth == 1 {
			heuristic = rand.Intn(41) - 10 // [-10, 30]
		}
		addNode(childID, strconv.Itoa(heuristic), "")
		addEdge(node.id, childID, arrow)
		children = append(children, Nodo{id: childID, heuristica: heuristic})
	}
	return children
}

// alphabeta implements minimax with alpha-beta pruning.
// It logs when a branch is pruned and marks pruned child nodes in the DOT output.
func alphabeta(node Nodo, depth int, maximizing bool, bfactor int, alpha int, beta int, path []int) Nodo {
	if depth == 0 {
		return node
	}
	// Debug: show entering node with alpha/beta
	curPath := append(path, node.id)
	fmt.Printf("Entrando en nodo %d depth=%d maximizing=%v alpha=%d beta=%d path=%v\n", node.id, depth, maximizing, alpha, beta, curPath)

	arrow := "max"
	if !maximizing {
		arrow = "min"
	}

	children := succ(node, depth, bfactor, arrow)

	var best Nodo
	if maximizing {
		best = Nodo{id: -1, heuristica: -999}
		for i, child := range children {
			result := alphabeta(child, depth-1, false, bfactor, alpha, beta, curPath)
			if result.heuristica > best.heuristica {
				best = result
			}
			// Debug: show child result and alpha update
			fmt.Printf("Nodo %d -> hijo %d heur=%d (mejor=%d)\n", node.id, child.id, result.heuristica, best.heuristica)
			if result.heuristica > alpha {
				alpha = result.heuristica
				fmt.Printf("  Actualizado alpha=%d\n", alpha)
			}
			// poda beta
			if alpha >= beta {
				// mark remaining children as pruned
				for j := i + 1; j < len(children); j++ {
					pr := children[j]
					// mark node visually as pruned (bright red fill, white font, dashed)
					addNode(pr.id, strconv.Itoa(pr.heuristica), ", style=filled, fillcolor=red, fontcolor=white, style=dashed")
					// imprimir ruta completa hasta el pruned child
					prPath := append(curPath, pr.id)
					fmt.Printf("Poda alfa-beta: nodo padre=%d, se poda la rama hija %d, ruta=%v (desde índice %d)\n", node.id, pr.id, prPath, j)
				}
				break
			}
		}
	} else {
		best = Nodo{id: -1, heuristica: 999}
		for i, child := range children {
			result := alphabeta(child, depth-1, true, bfactor, alpha, beta, curPath)
			if result.heuristica < best.heuristica {
				best = result
			}
			// Debug: show child result and beta update
			fmt.Printf("Nodo %d -> hijo %d heur=%d (mejor=%d)\n", node.id, child.id, result.heuristica, best.heuristica)
			if result.heuristica < beta {
				beta = result.heuristica
				fmt.Printf("  Actualizado beta=%d\n", beta)
			}
			// poda alfa
			if alpha >= beta {
				for j := i + 1; j < len(children); j++ {
					pr := children[j]
					addNode(pr.id, strconv.Itoa(pr.heuristica), ", style=filled, fillcolor=red, fontcolor=white, style=dashed")
					prPath := append(curPath, pr.id)
					fmt.Printf("Poda alfa-beta: nodo padre=%d, se poda la rama hija %d, ruta=%v (desde índice %d)\n", node.id, pr.id, prPath, j)
				}
				break
			}
		}
	}

	node.heuristica = best.heuristica
	addNode(best.id, strconv.Itoa(best.heuristica), ", shape=doublecircle")
	addNode(node.id, strconv.Itoa(node.heuristica), "")
	return node
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Iniciar DOT
	dotLines = append(dotLines, "graph G {")
	dotLines = append(dotLines, "    rankdir=TB;")
	dotLines = append(dotLines, "    node [shape=circle, fontname=\"Arial\", color=black, fontcolor=black, fontsize=12, width=0.9];")
	dotLines = append(dotLines, "    edge [color=black];")

	rootID := inc()
	addNode(rootID, "0", ", color=red, fontcolor=red, shape=doublecircle")

	root := Nodo{id: rootID, heuristica: 0}
	// Usar alphabeta en lugar de minimax (minimax no estaba definido)
	// parámetros: nodo raíz, profundidad, maximizing, branching factor, alpha, beta
	result := alphabeta(root, 3, true, 2, -999, 999, []int{})

	// Close graph (legend removed to save espacio)
	dotLines = append(dotLines, "}")
	fmt.Println("Resultado final:", result)

	// Guardar DOT en archivo
	dotContent := strings.Join(dotLines, "\n") + "\n"
	dotPath := "tree.dot"
	if err := os.WriteFile(dotPath, []byte(dotContent), 0644); err != nil {
		fmt.Println("Error escribiendo archivo DOT:", err)
		return
	}
	fmt.Println("DOT guardado en:", dotPath)

	// Intentar generar imagen PNG usando Graphviz 'dot' si está disponible
	pngPath := "tree.png"
	cmd := exec.Command("dot", "-Tpng", "-o", pngPath, dotPath)
	if err := cmd.Run(); err != nil {
		fmt.Println("No se pudo generar PNG con 'dot' (¿Graphviz instalado?). Error:", err)
		fmt.Println("Puedes generar la imagen manualmente: dot -Tpng -o tree.png tree.dot")
	} else {
		fmt.Println("Imagen de grafico generada en:", pngPath)
	}
}
