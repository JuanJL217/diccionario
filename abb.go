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

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{nil, nil, clave, dato}
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nuevoNodo := crearNodo(clave, dato)
	actual, padre := arbol.buscarNodo(clave, arbol.raiz, nil)
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
	actual, _ := arbol.buscarNodo(clave, arbol.raiz, nil)
	return actual != nil
}

func (arbol abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	actual, _ := arbol.buscarNodo(clave, arbol.raiz, nil)
	if actual == nil {
		panic(PANIC_CLAVE_DICCIONARIO)
	}
	return actual.dato
}
func (arbol *abb[K, V]) Borrar(clave K) V {

}

func (arbol abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {

}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
}

// --ITERADOR RANGO---//
func (arbol abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (arbol abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}

// ---METODOS AUXILIARES INTERNOS---//
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