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

func reemplazante[K comparable, V any](nodoInvocado *nodoABB[K, V]) (*nodoABB[K, V], *nodoABB[K, V]) { //Haremos el más derecho del nodo izquierdo del que uqeremos borrar
	padre := nodoInvocado
	nodoActual := nodoInvocado.derecho
	if nodoActual.derecho == nil {
		return nodoActual, padre
	}
	return reemplazante(nodoActual)
}

func (arbol *abb[K, V]) borrarConDosHijos(diferencia int, nodoActual, padre *nodoABB[K, V]) {
	reemplazo, padreDelReemplazo := reemplazante(nodoActual.izquierdo)
	padreDelReemplazo.derecho = reemplazo.izquierdo
	if nodoActual == arbol.raiz {
		arbol.raiz = reemplazo
	} else {
		if diferencia < _VALOR_NULO {
			padre.izquierdo = reemplazo
		} else {
			padre.derecho = reemplazo
		}
	}
	reemplazo.izquierdo = nodoActual.izquierdo
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

func (arbol abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
	//ME FALTAAAAAAAAAAAAAAA
}

// --ITERADOR RANGO---//

type iteradorRangoABB[K comparable, V any] struct {
	arbolIterar *abb[K, V]
	pila        TDAPila.Pila[*nodoABB[K, V]]
	inicio      *K
	fin         *K
}

// --ITERADOR RANGO INTERNO--//
func (arbol abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	//ME FALTAAAAAAAAAAAAAAA
}

// -- ITERADOR RANGO EXTERNO
func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	if desde == nil && hasta == nil {
		desde = buscarMinimo(arbol.raiz)
		hasta = buscarMaximo(arbol.raiz)
	}
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	iteradorABB := &iteradorRangoABB[K, V]{arbol, pila, desde, hasta}
	apilarAIzquierda(iteradorABB, arbol.raiz)
	return iteradorABB
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

func apilarAIzquierda[K comparable, V any](iterABB *iteradorRangoABB[K, V], nodoActual *nodoABB[K, V]) {
	if nodoActual == nil {
		return
	}
	if iterABB.arbolIterar.cmp(nodoActual.clave, *iterABB.inicio) < 0 {
		apilarAIzquierda(iterABB, nodoActual.derecho)
		return
	}
	if iterABB.arbolIterar.cmp(nodoActual.clave, *iterABB.fin) > 0 {
		apilarAIzquierda(iterABB, nodoActual.izquierdo)
		return
	}
	iterABB.pila.Apilar(nodoActual)
	apilarAIzquierda(iterABB, nodoActual.izquierdo)
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
		apilarAIzquierda(iterABB, nodoActual.derecho)
	}
}

// VerActual implements IterDiccionario.
func (iterABB *iteradorRangoABB[K, V]) VerActual() (K, V) {
	if !iterABB.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	return iterABB.pila.VerTope().clave, iterABB.pila.VerTope().dato
}
