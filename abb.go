package diccionario

type nodoABB[K comparable, V any] struct {
	izquierdo *nodoABB[K, V]
	derecho   *nodoABB[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      funcCmp[K] // porque devuelve menor, igual o mayor a 0
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccinarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{nil, nil, clave, dato}
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nuevoNodo := crearNodo(clave, dato)
	actual, padre := arbol.buscarNodo(clave, arbol.raiz, nil) // pongo nil, porque antes de la raiz, no hay nada
	if actual == nil && padre == nil {
		arbol.raiz = nuevoNodo
		arbol.cantidad++
	} else if actual != nil {
		actual.dato = dato
	} else {
		if arbol.cmp(clave, padre.clave) < 0 {
			padre.izquierdo = nuevoNodo
		} else {
			padre.derecho = nuevoNodo
		}
		arbol.cantidad++
	}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	actual, padre := arbol.buscarNodo(clave, arbol.raiz, nil)
	if padre != nil || (padre == nil && actual != nil) { //Si existe padre o si padre no existe y existe un actual (significa que apunta a la raiz)
		return true
	} // falta completar, yo me encargo
}

func (arbol *abb[K, V]) buscarNodo(clave K, nodoActual *nodoABB[K, V], padre *nodoABB[K, V]) (*nodoABB[K, V], *nodoABB[K, V]) {
	if nodoActual == nil {
		return nil, padre
	}
	if arbol.cmp(clave, nodoActual.clave) == 0 {
		return nodoActual, padre
	} else if arbol.cmp(clave, nodoActual.clave) < 0 {
		return arbol.buscarNodo(clave, nodoActual.izquierdo, nodoActual)
	} else {
		return arbol.buscarNodo(clave, nodoActual.derecho, nodoActual)
	}
}

// Ayudame con los iteradores :'v
