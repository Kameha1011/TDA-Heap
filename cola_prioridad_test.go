package cola_prioridad_test

import (
	"math/rand/v2"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN_CHICO  = 1000
	_VOLUMEN_GRANDE = 100000
)

var funcionCmpInts = func(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

var funcionCmpStrings = func(a, b string) int {
	return strings.Compare(a, b)
}

func cmpArrInts(a, b []int) int {
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return 1
		}
		if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func poblarArr(arr []int, max int) {
	hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for {
			n := rand.IntN(max)
			if _, ok := hash[n]; !ok {
				arr[i] = n
				hash[n] = 1
				break
			}
		}
	}
}

func ordenarArr(arr []int) {
	if len(arr) <= 1 {
		return
	}
	medio := len(arr) / 2
	izq := make([]int, medio)
	der := make([]int, len(arr)-medio)
	copy(izq, arr[:medio])
	copy(der, arr[medio:])
	ordenarArr(izq)
	ordenarArr(der)
	i, j, k := 0, 0, 0
	for i < len(izq) && j < len(der) {
		if izq[i] < der[j] {
			arr[k] = izq[i]
			i++
		} else {
			arr[k] = der[j]
			j++
		}
		k++
	}
	for i < len(izq) {
		arr[k] = izq[i]
		i++
		k++
	}
	for j < len(der) {
		arr[k] = der[j]
		j++
		k++
	}
}

func TestHeapArrInts(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[[]int](cmpArrInts)
	heap.Encolar([]int{1, 2, 3})
	heap.Encolar([]int{4, 5, 6})
	heap.Encolar([]int{7, 8, 9})
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, []int{7, 8, 9}, heap.VerMax())
	require.Equal(t, []int{7, 8, 9}, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.False(t, heap.EstaVacia())

}

func TestVacia(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.Equal(t, 0, heap.Cantidad())
	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
}

func TestEncolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	heap.Encolar(10)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
	heap.Encolar(20)
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())
	heap.Encolar(5)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())
	heap.Encolar(15)
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())
}

func TestDesencolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)
	heap.Encolar(15)
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 15, heap.Desencolar())
	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
}

func TestHeapStrings(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[string](funcionCmpStrings)
	heap.Encolar("hola")
	heap.Encolar("chau")
	heap.Encolar("como")
	heap.Encolar("estas")
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, "hola", heap.VerMax())
	require.Equal(t, "hola", heap.Desencolar())
	require.Equal(t, "estas", heap.Desencolar())
	require.Equal(t, "como", heap.Desencolar())
	require.Equal(t, "chau", heap.Desencolar())
}

func TestVolumen(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	arr := make([]int, _VOLUMEN_CHICO)
	poblarArr(arr, _VOLUMEN_CHICO)
	for i := 0; i < len(arr); i++ {
		heap.Encolar(arr[i])
	}
	require.Equal(t, _VOLUMEN_CHICO, heap.Cantidad())
	ordenarArr(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		require.Equal(t, arr[i], heap.Desencolar())
	}
}

func TestVolumenGrande(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	arr := make([]int, _VOLUMEN_GRANDE)
	poblarArr(arr, _VOLUMEN_GRANDE)
	for i := 0; i < len(arr); i++ {
		heap.Encolar(arr[i])
	}
	require.Equal(t, _VOLUMEN_GRANDE, heap.Cantidad())
	ordenarArr(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		require.Equal(t, arr[i], heap.Desencolar())
	}
}

func TestVaciar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	arr := make([]int, _VOLUMEN_CHICO)
	poblarArr(arr, _VOLUMEN_CHICO)
	for i := 0; i < len(arr); i++ {
		heap.Encolar(arr[i])
		heap.Desencolar()
	}
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapArr(t *testing.T) {
	arr := make([]int, _VOLUMEN_CHICO)
	poblarArr(arr, _VOLUMEN_CHICO)
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, funcionCmpInts)
	require.Equal(t, _VOLUMEN_CHICO, heap.Cantidad())
	ordenarArr(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		require.Equal(t, arr[i], heap.Desencolar())
	}
}

func TestHeapArrVacio(t *testing.T) {
	arr := make([]int, 0)
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, funcionCmpInts)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
}

func TestHeapArrEncolar(t *testing.T) {
	arr := make([]int, _VOLUMEN_CHICO)
	poblarArr(arr, _VOLUMEN_CHICO)
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, funcionCmpInts)
	heap.Encolar(_VOLUMEN_CHICO * 2)
	require.Equal(t, _VOLUMEN_CHICO+1, heap.Cantidad())
	require.Equal(t, _VOLUMEN_CHICO*2, heap.VerMax())
	heap.Encolar(_VOLUMEN_CHICO * 3)
	require.Equal(t, _VOLUMEN_CHICO+2, heap.Cantidad())
	require.Equal(t, _VOLUMEN_CHICO*3, heap.VerMax())
}

func TestComportamiento(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](funcionCmpInts)
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)
	heap.Encolar(15)
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.Equal(t, 15, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	heap.Encolar(30)
	heap.Encolar(25)

	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 30, heap.VerMax())
	require.Equal(t, 30, heap.Desencolar())
	require.Equal(t, 25, heap.Desencolar())

	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestComportamientoStrings(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[string](funcionCmpStrings)
	heap.Encolar("hola")
	heap.Encolar("chau")
	heap.Encolar("como")
	heap.Encolar("estas")
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, "hola", heap.VerMax())
	require.Equal(t, "hola", heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.Equal(t, "estas", heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.Equal(t, "como", heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	heap.Encolar("hola")
	heap.Encolar("chau")

	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, "hola", heap.VerMax())
	require.Equal(t, "hola", heap.Desencolar())
	require.Equal(t, "chau", heap.Desencolar())

	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "chau", heap.VerMax())
	require.Equal(t, "chau", heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestHeapArrComportamiento(t *testing.T) {
	arr := []int{10, 20, 5, 15, 30, 25, 5, 90, 72, 1, 0, 100, 200, 300, 400, 500, 600, 700, 800, 900}
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, funcionCmpInts)
	require.Equal(t, 20, heap.Cantidad())
	require.Equal(t, 900, heap.VerMax())

	require.Equal(t, 900, heap.Desencolar())
	require.Equal(t, 19, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.Equal(t, 800, heap.Desencolar())
	require.Equal(t, 18, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	heap.Encolar(1000)
	heap.Encolar(2000)
	heap.Encolar(3000)

	require.Equal(t, 21, heap.Cantidad())
	require.Equal(t, 3000, heap.VerMax())
	require.Equal(t, 3000, heap.Desencolar())

	for !heap.EstaVacia() {
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapArrComportamientoStrings(t *testing.T) {
	arr := []string{"goku", "vegetta", "gohan", "trunks", "bulma", "krilin", "piccolo", "bills"}
	heap := TDAColaPrioridad.CrearHeapArr[string](arr, funcionCmpStrings)
	require.Equal(t, 8, heap.Cantidad())
	require.Equal(t, "vegetta", heap.VerMax())

	require.Equal(t, "vegetta", heap.Desencolar())
	require.Equal(t, 7, heap.Cantidad())

	require.Equal(t, "trunks", heap.Desencolar())
	require.Equal(t, 6, heap.Cantidad())

	heap.Encolar("z")
	heap.Encolar("zzz")
	heap.Encolar("zzzzz")

	require.Equal(t, 9, heap.Cantidad())
	require.Equal(t, "zzzzz", heap.VerMax())
	require.Equal(t, "zzzzz", heap.Desencolar())

	require.Equal(t, 8, heap.Cantidad())
	require.Equal(t, "zzz", heap.VerMax())
	require.Equal(t, "zzz", heap.Desencolar())

	require.Equal(t, 7, heap.Cantidad())
	require.Equal(t, "z", heap.VerMax())
	require.Equal(t, "z", heap.Desencolar())

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, "piccolo", heap.VerMax())
	require.Equal(t, "piccolo", heap.Desencolar())

	for !heap.EstaVacia() {
		heap.Desencolar()
	}
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSort(t *testing.T) {
	arr := []int{10, 7, 5, 3, 2, 1, 8, 9, 4, 6}
	TDAColaPrioridad.HeapSort(arr, funcionCmpInts)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
}
