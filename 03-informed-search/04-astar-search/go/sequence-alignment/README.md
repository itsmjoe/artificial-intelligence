# Alineamiento de secuencias (Needleman–Wunsch + A*)

Esta carpeta contiene dos implementaciones para el alineamiento pareado de secuencias:

- `main.go` — Implementación de Needleman–Wunsch (`stringalign`) y una llamada a la implementación A* para comparación.
- `astar.go` — Implementación de búsqueda A* para alineamiento de secuencias (función `astarAlign`).
- `go.mod` — módulo para ejecutar el programa dentro de esta carpeta.
- `salida_astar.txt` — salida de consola guardada durante una ejecución de ejemplo.

Cómo ejecutar

1. Asegúrate de tener Go instalado (probado con Go 1.23).
2. Desde la raíz del repo ejecuta:

```bash
cd 03-informed-search/04-astar-search/go/sequence-alignment
go run .
```

Qué esperar

- El programa imprime primero el alineamiento calculado por Needleman–Wunsch y su matriz de costes, y después imprime el resultado de A* y un resumen.
- Para las secuencias de ejemplo (`ATCGTACGTA` y `ATGGTCGTA`) ambos métodos producen el mismo alineamiento y coste total `-6.00`.
- La salida de la ejecución de ejemplo se guardó en `salida_astar.txt`.

Verificación rápida (manual)

```bash
# ejecutar y guardar salida
go run . | tee salida_astar.txt

# buscar el coste total encontrado por A*
grep "A\* Coste total" salida_astar.txt

# o abrir el archivo para inspección
less salida_astar.txt
```

Verificación automática

Hay un script `verify_parity.sh` que:

- ejecuta `go run .` y guarda la salida en `salida_astar.txt`;
- extrae las líneas de alineamiento generadas por Needleman–Wunsch y por A*;
- compara las dos alineaciones y devuelve `OK` si coinciden o imprime las diferencias en caso contrario.

Para usarlo:

```bash
chmod +x verify_parity.sh
./verify_parity.sh
```

Salida esperada del script:

- "OK: Ambos alineamientos coinciden." si NW y A* devuelven el mismo alineamiento.
- En caso de diferencias, el script mostrará las dos alineaciones y terminará con código de error.

Si quieres, puedo agregar una prueba en Go que ejecute esta verificación para varios casos automáticamente.
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
