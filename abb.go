package diccionario

type nodoABB[K comparable, V any] struct {
	izquierdo *nodoABB[K, V]
	derecho   *nodoABB[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      funcCmp[K] // porque devuelve menor, igual o mayor a 0
}

type nodoAuxiliar[K comparable, V any] struct { //La idea de este nodo es saber de donde vengo, o sea, tener un nodo actual, donde estoy parado, mas que odo para borra
	actual   *nodoABB[K, V] // Padre
	anterior *nodoABB[K, V] // el Padre del actual, de donde vengo
}

type funcCmp[K comparable] func(K, K) int

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccinarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{nil, nil, clave, dato}
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	// AYUDAAAAA
	if arbol.raiz == nil {
		arbol.raiz.clave, arbol.raiz.dato = clave, dato
	} else {
		arbol.Guardar(clave, dato)
	}
	arbol.cantidad++
}
