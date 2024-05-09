package diccionario

import (
	"fmt"
)

const (
	_CASILLA_VACIA                  = 0
	_CASILLA_OCUPADA                = 1
	_CLAVE_BORRADA                  = 2
	_CAPACIDAD_INICIAL              = 19
	_FACTOR_REDIMENCIONAR_CAPACIDAD = 2
	_FACTOR_CAPACIDAD_MAXIMA        = 0.75
	_FACTOR_CAPACIDAD_MINIMA        = 0.25
	PANIC_ITERADOR                  = "El iterador termino de iterar"
	PANIC_CLAVE_DICCIONARIO         = "La clave no pertenece al diccionario"
)

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	ocupados int
	borrados int
	tam      int
}

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type iterTablaHash[K comparable, V any] struct {
	tablaIterar      *hashCerrado[K, V]
	posicion         int
	contadorOcupados int
}

/*
****************************************************************
-----------------FUNCION DE CREACION DEL TDA--------------------
****************************************************************
*/

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return crearTabla[K, V](_CAPACIDAD_INICIAL)
}
func crearTabla[K comparable, V any](capacidad int) *hashCerrado[K, V] {
	// Crea la tabla con dicho tamanio
	return &hashCerrado[K, V]{tabla: make([]celdaHash[K, V], capacidad), tam: capacidad}
}

/*
****************************************************************
-----------------PRIMITIVAS DEL TDA-----------------------------
****************************************************************
*/

func (th *hashCerrado[K, V]) Guardar(clave K, dato V) {
	pos, estado := th.buscarPosicion(clave)
	if estado == _CASILLA_OCUPADA {
		th.tabla[pos].clave, th.tabla[pos].dato = clave, dato
	} else if estado == _CASILLA_VACIA {
		th.tabla[pos].clave, th.tabla[pos].dato, th.tabla[pos].estado = clave, dato, _CASILLA_OCUPADA
		th.ocupados++
	}
	if th.verCapacidadAumentar() {
		th.nuevaTabla(th.tam * _FACTOR_REDIMENCIONAR_CAPACIDAD)
	}
}

func (th *hashCerrado[K, V]) Pertenece(clave K) bool {
	pos, _ := th.buscarPosicion(clave)
	return th.tabla[pos].estado == _CASILLA_OCUPADA && th.tabla[pos].clave == clave
}

func (th *hashCerrado[K, V]) Obtener(clave K) V {
	pos, estado := th.buscarPosicion(clave)
	if estado == _CASILLA_VACIA {
		panic(PANIC_CLAVE_DICCIONARIO)
	}
	return th.tabla[pos].dato
}

func (th *hashCerrado[K, V]) Borrar(clave K) V {
	if th.verCapacidadDisminuir() {
		th.nuevaTabla(th.tam / _FACTOR_REDIMENCIONAR_CAPACIDAD)
	}
	pos, estado := th.buscarPosicion(clave)
	if estado == _CASILLA_VACIA {
		panic(PANIC_CLAVE_DICCIONARIO)
	}
	if th.tabla[pos].clave == clave {
		th.tabla[pos].estado = _CLAVE_BORRADA
	}
	th.borrados++
	th.ocupados--
	return th.tabla[pos].dato
}

func (th hashCerrado[K, V]) Cantidad() int {
	return th.ocupados
}

/*
****************************************************************
-----------------ITERADOR INTERNO-------------------------------
****************************************************************
*/

func (th hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iter := new(iterTablaHash[K, V])
	iter.tablaIterar = &th
	var estado bool
	for iter.buscarOcupados() && !estado {
		if !visitar(th.tabla[iter.posicion].clave, th.tabla[iter.posicion].dato) {
			estado = true
		} else {
			iter.contadorOcupados++
			iter.posicion++
		}
	}

}

/*
****************************************************************
-----------------ITERADOR EXTERNO-------------------------------
****************************************************************
*/

func (th *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterTablaHash[K, V])
	iter.tablaIterar = th
	iter.buscarOcupados()
	return iter
}

func (i *iterTablaHash[K, V]) HaySiguiente() bool {
	return i.contadorOcupados < i.tablaIterar.ocupados
}

func (i *iterTablaHash[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	return i.tablaIterar.tabla[i.posicion].clave, i.tablaIterar.tabla[i.posicion].dato
}

func (i *iterTablaHash[K, V]) Siguiente() {
	if !i.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	i.posicion++
	i.contadorOcupados++
	i.buscarOcupados()
}

/*
****************************************************************
-----------------METODOS AUXILIARES INTERNOS--------------------
****************************************************************
*/

func (th hashCerrado[K, V]) buscarPosicion(clave K) (int, int) {
	// Busca la casilla de la clave para las primitivas
	// Guardar,Borrar,Pertenece,Obtener
	// Devuelve la posicion y el estado de la casilla
	pos := hashing(clave, th.tam)
	for pos < th.tam {
		if th.tabla[pos].clave == clave && th.tabla[pos].estado == _CASILLA_OCUPADA {
			return pos, _CASILLA_OCUPADA
		}
		if th.tabla[pos].estado == _CASILLA_VACIA {
			return pos, _CASILLA_VACIA
		}
		if pos == th.tam-1 {
			pos = 0
		}
		pos++

	}
	return pos, _CLAVE_BORRADA
}

func (i *iterTablaHash[K, V]) buscarOcupados() bool {
	// Buscamos la celda que tenga un estado "ocupado"
	var estado bool
	for i.posicion < i.tablaIterar.tam && i.contadorOcupados < i.tablaIterar.ocupados && !estado {
		if i.tablaIterar.tabla[i.posicion].estado == _CASILLA_OCUPADA {
			estado = true
		} else {
			i.posicion++
		}
	}
	return estado
}

func (th *hashCerrado[K, V]) verCapacidadAumentar() bool {
	// Si los elementos (ocupados + borrados) ocupan un 75% de la tabla, Aumentamos el tamanio de la tabla
	return (th.ocupados + th.borrados) >= int(float64(th.tam)*_FACTOR_CAPACIDAD_MAXIMA)
}

func (th *hashCerrado[K, V]) verCapacidadDisminuir() bool {
	// Si los elementos ocupados ocupan un 25% de la tabla, Disminuimos la tabla siempre y cuando los ocupados sean mayores a la CAPACIDAD_INICIAL
	return th.ocupados <= int(float64(th.tam)*_FACTOR_CAPACIDAD_MINIMA) && th.ocupados > _CAPACIDAD_INICIAL
}
func (th *hashCerrado[K, V]) nuevaTabla(tamanio int) {
	// Creamos una nueva tabla para la redimension
	nuevo := crearTabla[K, V](tamanio)
	iter := new(iterTablaHash[K, V])
	iter.tablaIterar = th
	for iter.buscarOcupados() {
		nuevo.Guardar(th.tabla[iter.posicion].clave, th.tabla[iter.posicion].dato)
		iter.posicion++
		iter.contadorOcupados++
	}
	th.tabla = nuevo.tabla
	th.ocupados = nuevo.ocupados
	th.borrados = nuevo.borrados
	th.tam = nuevo.tam

}

/*
****************************************************************
-----------------FUNCIONES AUXILIARES---------------------------
****************************************************************
*/
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// CODIGO FUENTE : https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
// USAMOS EL METODO FNV-HASHING
func hashing[K comparable](clave K, tamanio int) int {
	var (
		base uint32 = 2166136261
		dato uint32 = 16777619
	)
	arrayBytes := convertirABytes(clave)
	for _, valor := range arrayBytes {
		base ^= uint32(valor)
		base *= dato
	}
	return int(base) % tamanio
}
