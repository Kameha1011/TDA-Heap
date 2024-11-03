package cola_prioridad

const (
	_TAMINICIAL         = 10
	_FACTOR_REDIMENSION = 2
	_FACTORDISMINUCION  = 4
)

type funcionComp[T any] func(a, b T) int

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

func upHeap[T any](posicion int, heap *colaPrioridad[T]) {
	if posicion == 0 {
		return
	}
	posicionPadre := (posicion - 1) / 2
	padre := heap.datos[posicionPadre]
	if heap.cmp(padre, heap.datos[posicion]) < 0 {
		swap(posicionPadre, posicion, heap.datos)
		upHeap(posicionPadre, heap)
	}
}

func obtenerMasGrande[T any](posicionIzq, posicionDer int, arr []T, cmp funcionComp[T]) int {
	if cmp(arr[posicionIzq], arr[posicionDer]) > 0 {
		return posicionIzq
	}
	return posicionDer
}

func downHeap[T any](posicion int, cmp funcionComp[T], arr []T, cantidad int) {
	if posicion >= cantidad-1 {
		return
	}

	poscIzq := posicion*2 + 1
	poscDer := posicion*2 + 2

	if poscIzq > cantidad-1 {
		return
	}

	mayor := poscIzq

	if poscDer <= cantidad-1 {
		mayor = obtenerMasGrande(poscIzq, poscDer, arr, cmp)
	}

	if cmp(arr[posicion], arr[mayor]) < 0 {
		swap(posicion, mayor, arr)
		downHeap(mayor, cmp, arr, cantidad)
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
