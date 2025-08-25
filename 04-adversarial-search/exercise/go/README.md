# Adversarial search - Minimax con poda alfa-beta

Este pequeño programa genera un árbol de búsqueda aleatorio y ejecuta una versión de minimax usando poda alfa-beta.

Archivos
- `main.go` - código fuente con implementación de `alphabeta`, generación de DOT y creación de imagen `tree.png`.
- `tree.dot` - archivo DOT generado (si se ejecuta el programa).
- `tree.png` - imagen del grafo generada por Graphviz (si está instalado).
- `run_output.txt` - salida de consola de la última ejecución (si fue guardada).

Uso

1) Ejecutar el programa desde esta carpeta:

```bash
cd 04-adversarial-search/exercise/go
go run main.go
```

2) Ficheros generados:
- `tree.dot` contiene la representación Graphviz del árbol y una leyenda.
- `tree.png` es la imagen (si `dot` está instalado el programa intentará generarla automáticamente).

Notas
- El programa imprime mensajes de depuración mostrando alpha/beta al entrar en cada nodo y las actualizaciones tras evaluar cada hijo.
- Los nodos podados se muestran con color `lightcoral` y texto en blanco; el nodo seleccionado (mejor) aparece como doble círculo.
- Si quieres cambiar la profundidad o el factor de ramificación, edita la llamada en `main.go`.
