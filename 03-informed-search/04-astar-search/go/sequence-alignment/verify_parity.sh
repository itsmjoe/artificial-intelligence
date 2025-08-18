#!/usr/bin/env bash
set -euo pipefail

# Ejecuta el programa y guarda la salida

go run . | tee salida_astar.txt

# Extrae secciones de alineamiento usando sed (más robusto)
# NW A: líneas entre "Alineamiento A:" y la siguiente etiqueta "Alineamiento B:"
nw_a=$(sed -n '/^Alineamiento A:$/,/^Alineamiento B:$/p' salida_astar.txt | sed '1d;$d' | tr -d '\n')
# NW B: entre "Alineamiento B:" y "Resumen:"
nw_b=$(sed -n '/^Alineamiento B:$/,/^Resumen:$/p' salida_astar.txt | sed '1d;$d' | tr -d '\n')

# A* A: entre "A* Alineamiento A:" y "A* Alineamiento B:"
astar_a=$(sed -n '/^A\* Alineamiento A:$/,/^A\* Alineamiento B:$/p' salida_astar.txt | sed '1d;$d' | tr -d '\n')
# A* B: entre "A* Alineamiento B:" y la línea que comienza "A* Coste total"
astar_b=$(sed -n '/^A\* Alineamiento B:$/,/^A\* Coste total/ p' salida_astar.txt | sed '1d;$d' | tr -d '\n')

# Normalizar (eliminar espacios/retornos) para comparación simple
norm(){ echo "$1" | tr -d ' \t\r\n' ; }

na=$(norm "$nw_a")
nb=$(norm "$nw_b")
aa=$(norm "$astar_a")
ab=$(norm "$astar_b")

if [[ "$na" == "$aa" && "$nb" == "$ab" ]]; then
  echo "OK: Ambos alineamientos coinciden."
  exit 0
else
  echo "ERROR: Diferencia entre alineamientos."
  echo "Needleman-Wunsch A: $nw_a"
  echo "A* A: $astar_a"
  echo "Needleman-Wunsch B: $nw_b"
  echo "A* B: $astar_b"
  exit 2
fi
