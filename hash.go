package diccionario

import (
	"fmt"
)

const (
	_CASILLA_VACIA                  int     = 0
	_CASILLA_OCUPADA                int     = 1
	_CLAVE_BORRADA                  int     = 2
	_CAPACIDAD_INICIAL              int     = 20
	_FACTOR_REDIMENCIONAR_CAPACIDAD int     = 2
	_FACTOR_CAPACIDAD_MAXIMA        float64 = 0.75
	_FACTOR_CAPACIDAD_MINIMA        float64 = 0.25
	_PANIC_CLAVE_DICCIONARIO        string  = "La clave no pertenece al diccionario"
	_PANIC_ITERADOR                         = "El iterador termino de iterar"
)

type tablaHash[K comparable, V any] struct {
	ocupados int
	borrados int
	tabla    []hashElemento[K, V]
}

type hashElemento[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type iterExternoTablaHash[K comparable, V any] struct {
	tabla    []hashElemento[K, V]
	contador int
}

func crearTabla[K comparable, V any](capacidad int) []hashElemento[K, V] {
	return make([]hashElemento[K, V], capacidad)
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &tablaHash[K, V]{tabla: crearTabla[K, V](_CAPACIDAD_INICIAL)}
}

func (th tablaHash[K, V]) capacidad() int {
	return cap(th.tabla)
}

func (th tablaHash[K, V]) cantidadMaximaRedimensionar() bool {
	return (th.ocupados + th.borrados) >= int((float64(th.capacidad()) * _FACTOR_CAPACIDAD_MAXIMA))
}

func (th tablaHash[K, V]) capacidadAceptadaParaAchicar() bool {
	return th.ocupados > _CAPACIDAD_INICIAL
}

func (th tablaHash[K, V]) unCuartoDeLaCapacidadActual() bool {
	return th.ocupados <= int(float64(th.capacidad())*_FACTOR_CAPACIDAD_MINIMA)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
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
	return int(base) % th.capacidad()
}

func (th *tablaHash[K, V]) redimensionar(aumentar bool) {
	capacidad := th.capacidad()
	tablaAnterior := th.tabla
	if aumentar {
		capacidad *= _FACTOR_REDIMENCIONAR_CAPACIDAD
	} else {
		capacidad /= _FACTOR_REDIMENCIONAR_CAPACIDAD
	}
	th.tabla = crearTabla[K, V](capacidad)
	th.borrados, th.ocupados = 0, 0
	for i := 0; i < len(tablaAnterior); i++ {
		if tablaAnterior[i].estado == _CASILLA_OCUPADA {
			th.Guardar(tablaAnterior[i].clave, tablaAnterior[i].dato)
		}
	}
}

func (th *tablaHash[K, V]) Guardar(clave K, dato V) {
	pos := th.buscarPosicion(clave)
	if th.tabla[pos].clave == clave && th.tabla[pos].estado == _CASILLA_OCUPADA {
		th.tabla[pos].dato = dato
	} else {
		th.tabla[pos].clave, th.tabla[pos].dato, th.tabla[pos].estado = clave, dato, _CASILLA_OCUPADA
		th.ocupados++
	}
	if th.cantidadMaximaRedimensionar() {
		th.redimensionar(true)
	}
}

func (th *tablaHash[K, V]) Pertenece(clave K) bool {
	pos := th.buscarPosicion(clave)
	return th.tabla[pos].estado == _CASILLA_OCUPADA && th.tabla[pos].clave == clave
}

func (th *tablaHash[K, V]) Obtener(clave K) V {
	pos := th.buscarPosicion(clave)
	if th.tabla[pos].estado == _CASILLA_VACIA {
		panic(_PANIC_CLAVE_DICCIONARIO)
	}
	return th.tabla[pos].dato
}

func (th *tablaHash[K, V]) Borrar(clave K) V {
	if th.capacidadAceptadaParaAchicar() && th.unCuartoDeLaCapacidadActual() {
		th.redimensionar(false)
	}
	pos := th.buscarPosicion(clave)
	if th.tabla[pos].estado == _CASILLA_VACIA {
		panic(_PANIC_CLAVE_DICCIONARIO)
	}
	if th.tabla[pos].clave == clave {
		th.tabla[pos].estado = _CLAVE_BORRADA
	}
	th.borrados++
	th.ocupados--
	return th.tabla[pos].dato
}

func (th *tablaHash[K, V]) Cantidad() int {
	return th.ocupados
}

func (th *tablaHash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	seguirIterando := true
	pos := 0
	for pos < th.capacidad() && seguirIterando {
		if th.tabla[pos].estado == _CASILLA_OCUPADA {
			if !visitar(th.tabla[pos].clave, th.tabla[pos].dato) {
				seguirIterando = false
			}
		}
		pos++
	}
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
		panic(_PANIC_ITERADOR)
	}
	return ieth.tabla[ieth.contador].clave, ieth.tabla[ieth.contador].dato
}
func (ieth iterExternoTablaHash[K, V]) Siguiente() {
	if !ieth.HaySiguiente() {
		panic(_PANIC_ITERADOR)
	}
}

/*
****************************************************************
-----------------FUNCIONES AUXILIARES--------------------------
****************************************************************
*/

func (th tablaHash[K, V]) buscarPosicion(clave K) int {
	pos := th.hashingPosicion(clave)
	for i := pos; i < th.capacidad(); i++ {
		if th.tabla[i].clave == clave && th.tabla[i].estado == _CASILLA_OCUPADA {
			return i
		} else if th.tabla[i].estado == _CASILLA_VACIA {
			pos = i
			break
		}
		if i == th.capacidad()-1 {
			i = 0
		}
	}
	return pos
}
