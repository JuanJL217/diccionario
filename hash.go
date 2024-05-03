package diccionario

import (
	"fmt"
)

const (
	_CASILLA_VACIA           int    = 0
	_CASILLA_OCUPADA         int    = 1
	_CLAVE_BORRADA           int    = 2
	_CAPACIDAD_INICIAL       int    = 20
	_DATO_INVALIDO                  = -1
	_PANIC_CLAVE_DICCIONARIO string = "La clave no pertenece al diccionario"
	PANIC_ITERADOR                  = "El iterador termino de iterar"
)

type tablaHash[K comparable, V any] struct { //Aquí se manejará el tema de la capacidad y su redimención
	capacidad int //cantidad de elementos agregados y borrados
	cantidad  int
	borrados  int
	tabla     []hashElemento[K, V] //en "tabla" será el slice de la estrutura hashElemento
}

type hashElemento[K comparable, V any] struct { //Esta estructura es la que tendrá cada celda que tengamos en crearTAbla
	clave  K //Cada celda tenda una clave, dato y estado
	dato   V
	estado int
}

type iterExternoTablaHash[K comparable, V any] struct {
	tabla    []hashElemento[K, V]
	contador int
}

func crearTabla[K comparable, V any](capacidad int) []hashElemento[K, V] {
	return make([]hashElemento[K, V], capacidad) //Al pasar la capacidad por parametro, vamos a poder "redimensionar" cuando haya más elementos en la tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] { //CrearHash inicializará la capacidad del vector y la tabla. Esa tabla es la que nos importa para la lógica
	return &tablaHash[K, V]{capacidad: _CAPACIDAD_INICIAL, tabla: crearTabla[K, V](_CAPACIDAD_INICIAL)}
}

func (th *tablaHash[K, V]) Guardar(clave K, dato V) {
	//pos := th.buscarClavePosicion(clave)
	//if th.tabla[pos].clave == clave {
	//	th.tabla[pos].dato = dato
	//} else {
	//	th.tabla[pos].clave = clave
	//	th.tabla[pos].dato = dato
	//	th.tabla[pos].estado = _CASILLA_OCUPADA
	//	th.cantidad++
	//}
	pos, estado := th.buscarEstado(clave)
	if estado == _CASILLA_OCUPADA {
		th.tabla[pos].dato = dato
	} else {
		th.tabla[pos].clave, th.tabla[pos].dato, th.tabla[pos].estado = clave, dato, _CASILLA_OCUPADA
		th.cantidad++
	}
}

func (th *tablaHash[K, V]) Pertenece(clave K) bool {
	//pos := th.hashingPosicion(clave)
	//return th.tabla[pos].estado == _CASILLA_OCUPADA && th.tabla[pos].clave == clave
	_, estado := th.buscarEstado(clave)
	return estado == _CASILLA_OCUPADA

}

func (th *tablaHash[K, V]) Obtener(clave K) V {
	//pos := th.buscarClavePosicion(clave)
	//if th.tabla[pos].estado == _CASILLA_VACIA {
	//	panic(_PANIC_CLAVE_DICCIONARIO)
	//}
	//return th.tabla[pos].dato
	pos, estado := th.buscarEstado(clave)
	if estado != _CASILLA_OCUPADA {
		panic(_PANIC_CLAVE_DICCIONARIO)
	}
	return th.tabla[pos].dato
}

func (th *tablaHash[K, V]) Borrar(clave K) V {
	//pos := th.buscarClavePosicion(clave)
	//if th.tabla[pos].estado == _CASILLA_VACIA {
	//	panic(_PANIC_CLAVE_DICCIONARIO)
	//}
	//if th.tabla[pos].clave == clave {
	//	th.tabla[pos].estado = _CLAVE_BORRADA
	//}
	//th.borrados++
	//return th.tabla[pos].dato
	pos, estado := th.buscarEstado(clave)
	if estado != _CASILLA_OCUPADA {
		panic(_PANIC_CLAVE_DICCIONARIO)
	}
	th.tabla[pos].estado = _CLAVE_BORRADA
	th.borrados++
	return th.tabla[pos].dato
}

func (th *tablaHash[K, V]) Cantidad() int {
	return th.cantidad
}

func (th *tablaHash[K, V]) Iterar(func(clave K, dato V) bool) {
	return
}

/*
****************************************************************
-----------------ITERADOR EXTERNO--------------------------
****************************************************************
*/

func (th tablaHash[K, V]) Iterador() IterDiccionario[K, V] {
	return &iterExternoTablaHash[K, V]{th.tabla, 0}
}

func (ieth *iterExternoTablaHash[K, V]) HaySiguiente() bool {
	estado := false
	for ieth.contador < len(ieth.tabla) && !estado {
		if ieth.tabla[ieth.contador].estado == _CASILLA_OCUPADA {
			estado = true
		}
		ieth.contador++
	}
	return estado
}

func (ieth iterExternoTablaHash[K, V]) VerActual() (K, V) {
	if !ieth.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	return ieth.tabla[ieth.contador].clave, ieth.tabla[ieth.contador].dato
}
func (ieth iterExternoTablaHash[K, V]) Siguiente() {
	if !ieth.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}

/*
****************************************************************
-----------------FUNCIONES AUXILIARES--------------------------
****************************************************************
*/

func (th tablaHash[K, V]) buscarClavePosicion(clave K) int {
	pos := th.hashingPosicion(clave)
	fmt.Println(clave)
	for i := pos; i < th.capacidad; i++ {
		if th.tabla[i].clave == clave && th.tabla[i].estado == _CASILLA_OCUPADA {
			return i
		} else if th.tabla[i].estado == _CASILLA_VACIA {
			pos = i
			break
		}
		if i == th.capacidad-1 {
			i = 0
		}
	}
	return pos
}

// Estaba haciendo una funcion de busqueda xd , igual podes modificarlo
func (th tablaHash[K, V]) buscarEstado(clave K) (int, int) {
	pos := th.hashingPosicion(clave)
	if th.tabla[pos].clave == clave && th.tabla[pos].estado == _CASILLA_OCUPADA {

		return pos, _CASILLA_OCUPADA
	}
	for i := pos + 1; i < th.capacidad && i != pos; i++ {
		if th.tabla[i].clave == clave && th.tabla[i].estado == _CASILLA_OCUPADA {
			return i, _CASILLA_OCUPADA
		}
		if th.tabla[i].estado == _CASILLA_VACIA {
			return i, _CASILLA_VACIA
		}
		if i == th.capacidad-1 {
			i = 0
		}
	}
	//Puse este dato pq si hay muchos borrados ,entonces es invalido la busqueda
	return _DATO_INVALIDO, _DATO_INVALIDO
}

// CODIGO FUENTE : https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
// USAMOS EL METODO FNV-HASHING
func (th tablaHash[K, V]) hashingPosicion(clave K) int {
	var (
		base uint32 = 2166136261
		dato uint32 = 16777619
	)
	arrayBytes := convertirABytes(clave)
	for _, valor := range arrayBytes {
		base ^= uint32(valor)
		base *= dato
	}
	return int(base) % th.capacidad

}

func convertirABytes[K comparable](clave K) []byte { //Función de la cátedra para obtener un array de bytes de la clave
	return []byte(fmt.Sprintf("%v", clave))
}
