package diccionario

import (
	TDAPila "tdas/pila"
)

const _VALOR_NULO int = 0

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
	cmp      funcCmp[K]
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

func (arbol *abb[K, V]) Obtener(clave K) V {
	actual, _ := arbol.buscarNodo(clave, arbol.raiz, nil)
	if actual == nil {
		panic(PANIC_CLAVE_DICCIONARIO)
	}
	return actual.dato
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	actual, padre := arbol.buscarNodo(clave, arbol.raiz, nil)
	if actual == nil {
		panic(PANIC_CLAVE_DICCIONARIO)
	}
	if actual == arbol.raiz {
		padre = arbol.raiz
	}
	dato := actual.dato
	valorComparativo := arbol.cmp(actual.clave, padre.clave) // Me dice si está a la izquierda o derecha del padre
	if actual.izquierdo == nil && actual.derecho == nil {
		arbol.borrarConCeroHijos(valorComparativo, actual, padre)
	} else if actual.izquierdo != nil && actual.derecho == nil {
		arbol.borrarConUnHijo(valorComparativo, true, actual, padre)
	} else if actual.izquierdo == nil && actual.derecho != nil {
		arbol.borrarConUnHijo(valorComparativo, false, actual, padre)
	} else {
		arbol.borrarConDosHijos(valorComparativo, actual, padre)
	}
	arbol.cantidad--
	return dato
}

func (arbol abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

// ---METODOS AUXILIARES INTERNOS---//
func (arbol *abb[K, V]) borrarConCeroHijos(diferencia int, nodoActual, padre *nodoABB[K, V]) {
	if nodoActual == arbol.raiz {
		arbol.raiz = nil
	} else {
		if diferencia < _VALOR_NULO {
			padre.izquierdo = nil
		} else {
			padre.derecho = nil
		}
	}
}

func (arbol *abb[K, V]) borrarConUnHijo(diferencia int, ladoIzquierdo bool, nodoActual, padre *nodoABB[K, V]) {
	if ladoIzquierdo {
		if nodoActual == arbol.raiz {
			arbol.raiz = nodoActual.izquierdo
		} else {
			if diferencia < _VALOR_NULO {
				padre.izquierdo = nodoActual.izquierdo
			} else {
				padre.derecho = nodoActual.izquierdo
			}
		}
	} else {
		if nodoActual == arbol.raiz {
			arbol.raiz = nodoActual.derecho
		} else {
			if diferencia < _VALOR_NULO {
				padre.izquierdo = nodoActual.derecho
			} else {
				padre.derecho = nodoActual.derecho
			}
		}
	}
}

func reemplazante[K comparable, V any](nodoInvocado *nodoABB[K, V]) (*nodoABB[K, V], *nodoABB[K, V]) {
	nodoHijo := nodoInvocado.derecho
	if nodoHijo == nil {
		return nodoInvocado, nil
	}
	if nodoHijo.derecho == nil {
		return nodoHijo, nodoInvocado
	}
	return reemplazante(nodoHijo)
}

func (arbol *abb[K, V]) borrarConDosHijos(diferencia int, nodoActual, padre *nodoABB[K, V]) {
	reemplazo, padreDelReemplazo := reemplazante(nodoActual.izquierdo)
	if padreDelReemplazo != nil {
		padreDelReemplazo.derecho = reemplazo.izquierdo
	}
	if nodoActual == arbol.raiz {
		arbol.raiz = reemplazo
	} else {
		if diferencia < _VALOR_NULO {
			padre.izquierdo = reemplazo
		} else {
			padre.derecho = reemplazo
		}
	}
	if padreDelReemplazo != nil {
		reemplazo.izquierdo = nodoActual.izquierdo
	} else {
		reemplazo.izquierdo = nodoActual.izquierdo.izquierdo
	}
	reemplazo.derecho = nodoActual.derecho
}

func (arbol *abb[K, V]) buscarNodo(clave K, nodoActual, padre *nodoABB[K, V]) (*nodoABB[K, V], *nodoABB[K, V]) {
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

// --ITERADORES---//

type iteradorRangoABB[K comparable, V any] struct {
	arbolIterar *abb[K, V]
	pila        TDAPila.Pila[*nodoABB[K, V]]
	inicio      *K
	fin         *K
}

func (iterABB *iteradorRangoABB[K, V]) HaySiguiente() bool {
	return !iterABB.pila.EstaVacia()
}

// Siguiente implements IterDiccionario.
func (iterABB *iteradorRangoABB[K, V]) Siguiente() {
	if !iterABB.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	nodoActual := iterABB.pila.Desapilar()
	if nodoActual.derecho != nil {
		iterABB.apilarElementos(nodoActual.derecho)
	}
}

// VerActual implements IterDiccionario.
func (iterABB *iteradorRangoABB[K, V]) VerActual() (K, V) {
	if !iterABB.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	return iterABB.pila.VerTope().clave, iterABB.pila.VerTope().dato
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

// --ITERADOR RANGO INTERNO--//
func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if desde == nil {
		desde = buscarMinimo(arbol.raiz)
	}
	if hasta == nil {
		hasta = buscarMaximo(arbol.raiz)
	}
	iter := &iteradorRangoABB[K, V]{arbol, TDAPila.CrearPilaDinamica[*nodoABB[K, V]](), desde, hasta}
	iter.visitarElementos(arbol.raiz, visitar)

}

// -- ITERADOR RANGO EXTERNO
func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	if desde == nil {
		desde = buscarMinimo(arbol.raiz)
	}
	if hasta == nil {
		hasta = buscarMaximo(arbol.raiz)
	}
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	iteradorABB := &iteradorRangoABB[K, V]{arbol, pila, desde, hasta}
	iteradorABB.apilarElementos(arbol.raiz)
	return iteradorABB
}

func (iterABB *iteradorRangoABB[K, V]) apilarElementos(nodoActual *nodoABB[K, V]) {
	if nodoActual == nil {
		return
	}
	if iterABB.arbolIterar.cmp(nodoActual.clave, *iterABB.inicio) < 0 {
		iterABB.apilarElementos(nodoActual.derecho)
		return
	}
	if iterABB.arbolIterar.cmp(nodoActual.clave, *iterABB.fin) > 0 {
		iterABB.apilarElementos(nodoActual.izquierdo)
		return
	}
	iterABB.pila.Apilar(nodoActual)
	iterABB.apilarElementos(nodoActual.izquierdo)
}

func (iterABB *iteradorRangoABB[K, V]) visitarElementos(nodo *nodoABB[K, V], visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	if iterABB.arbolIterar.cmp(nodo.clave, *iterABB.inicio) >= 0 {
		iterABB.visitarElementos(nodo.izquierdo, visitar)
	}
	if iterABB.arbolIterar.cmp(nodo.clave, *iterABB.inicio) >= 0 && iterABB.arbolIterar.cmp(nodo.clave, *iterABB.fin) <= 0 {
		visitar(nodo.clave, nodo.dato)
	}
	if iterABB.arbolIterar.cmp(nodo.clave, *iterABB.fin) <= 0 {
		iterABB.visitarElementos(nodo.derecho, visitar)
	}
}

func buscarMinimo[K comparable, V any](nodo *nodoABB[K, V]) *K {
	if nodo == nil {
		return nil
	}
	if nodo.izquierdo == nil {
		return &nodo.clave
	}
	return buscarMinimo(nodo.izquierdo)
}

func buscarMaximo[K comparable, V any](nodo *nodoABB[K, V]) *K {
	if nodo == nil {
		return nil
	}
	if nodo.derecho == nil {
		return &nodo.clave
	}
	return buscarMaximo(nodo.derecho)
}
