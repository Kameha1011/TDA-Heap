package cola_prioridad

const (
	_TAMINICIAL         = 10
	_FACTOR_REDIMENSION = 2
	_FACTORDISMINUCION  = 4
)

type funcionComp[T any] func(T, T) int

type colaPrioridad[T any] struct {
	datos    []T
	cantidad int
	cmp      funcionComp[T]
}

func swap[T any](a, b int, arr []T) {
	arr[a], arr[b] = arr[b], arr[a]
}

func (heap *colaPrioridad[T]) panicEstaVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func upHeap[T any](indice int, heap *colaPrioridad[T]) {
	indicePadre := (indice - 1) / 2
	for indice >= 0 && heap.cmp(heap.datos[indicePadre], heap.datos[indice]) < 0 {
		swap(indicePadre, indice, heap.datos)
		indice = indicePadre
		indicePadre = (indice - 1) / 2
	}
}

func indiceHijoMayor[T any](indicePadre int, arr []T, cmp func(a, b T) int, cantidad int) int {
	hijoIzq, hijoDer := (2*indicePadre)+1, (2*indicePadre)+2
	if hijoDer < cantidad && cmp(arr[hijoIzq], arr[hijoDer]) < 0 {
		return hijoDer
	}
	return hijoIzq
}

func downHeap[T any](indice int, cmp funcionComp[T], arr []T, cantidad int) {
	hijoMayor := indiceHijoMayor(indice, arr, cmp, cantidad)
	for hijoMayor < cantidad && cmp(arr[indice], arr[hijoMayor]) < 0 {
		swap(indice, hijoMayor, arr)
		indice = hijoMayor
		hijoMayor = indiceHijoMayor(indice, arr, cmp, cantidad)
	}
}

func (heap *colaPrioridad[T]) redimension(cant int) {
	arrNuevo := make([]T, cant)
	copy(arrNuevo, heap.datos)
	heap.datos = arrNuevo
}

func heapify[T any](arr []T, cmp funcionComp[T]) {
	for i := range arr {
		downHeap(len(arr)-1-i, cmp, arr, len(arr))
	}
}

func crearHeap[T any](cmp funcionComp[T], tam int) *colaPrioridad[T] {
	heap := new(colaPrioridad[T])
	heap.datos = make([]T, tam)
	heap.cantidad = 0
	heap.cmp = cmp
	return heap
}

func CrearHeap[T any](cmp funcionComp[T]) ColaPrioridad[T] {
	return crearHeap(cmp, _TAMINICIAL)
}

func CrearHeapArr[T any](cmp funcionComp[T], arr []T) ColaPrioridad[T] {
	heap := crearHeap(cmp, max(len(arr), _TAMINICIAL))
	heap.cantidad = len(arr)
	copy(heap.datos, arr)
	heapify(heap.datos, cmp)
	return heap
}

func (heap *colaPrioridad[T]) EstaVacia() bool {
	return heap.cantidad <= 0
}

func (heap *colaPrioridad[T]) Encolar(dato T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimension(len(heap.datos) * _FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = dato
	upHeap(heap.cantidad, heap)
	heap.cantidad++
}

func (heap *colaPrioridad[T]) VerMax() T {
	heap.panicEstaVacia()
	return heap.datos[0]
}

func (heap *colaPrioridad[T]) Desencolar() T {
	heap.panicEstaVacia()
	dato := heap.datos[0]
	swap(0, heap.cantidad-1, heap.datos)
	heap.cantidad--
	downHeap(0, heap.cmp, heap.datos, heap.cantidad)
	if heap.cantidad <= len(heap.datos)/_FACTORDISMINUCION && len(heap.datos) > _TAMINICIAL {
		heap.redimension(len(heap.datos) / _FACTOR_REDIMENSION)
	}
	return dato
}

func (heap *colaPrioridad[T]) Cantidad() int {
	return heap.cantidad
}

func HeapSort[T any](arr []T, cmp funcionComp[T]) {
	heapify(arr, cmp)
	for i := len(arr) - 1; i >= 0; i-- {
		swap(0, i, arr)
		downHeap(0, cmp, arr, i)
	}
}
