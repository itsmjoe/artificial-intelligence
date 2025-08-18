grep "A\* Coste total" salida_astar.txt
# Alineamiento de secuencias (Needleman–Wunsch + A*)

Esta carpeta contiene dos implementaciones para el alineamiento pareado de secuencias:

- `main.go` — Implementación de Needleman–Wunsch (`stringalign`) y una llamada a la implementación A* para comparación.
- `astar.go` — Implementación de búsqueda A* para alineamiento de secuencias (función `astarAlign`).
- `go.mod` — módulo para ejecutar el programa dentro de esta carpeta.
- `salida_astar.txt` — salida de consola guardada durante una ejecución de ejemplo.

Cómo ejecutar

1. Asegúrate de tener Go instalado (probado con Go 1.23).
2. Desde esta carpeta ejecuta:

```bash
cd 03-informed-search/04-astar-search/go/sequence-alignment
go run .
```

Qué esperar

- El programa imprime primero el alineamiento calculado por Needleman–Wunsch y su matriz de costes, y después imprime el resultado de A* y un resumen.
- Para las secuencias de ejemplo (`ATCGTACGTA` y `ATGGTCGTA`) ambos métodos producen el mismo alineamiento y coste total `-6.00`.
- La salida de la ejecución de ejemplo se guardó en `salida_astar.txt`.

Verificación rápida

Puedes confirmar rápidamente que ambas implementaciones producen el mismo alineamiento revisando que `salida_astar.txt` contenga las dos secciones y las mismas líneas de alineamiento. Por ejemplo:

```bash
# ejecutar y guardar salida
go run . | tee salida_astar.txt

# buscar el coste total encontrado por A*
grep "A\* Coste total" salida_astar.txt

# o abrir el archivo para inspección
less salida_astar.txt
```
