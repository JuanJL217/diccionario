package diccionario

import "fmt"

const (
	_CASILLA_VACIA           int    = 0
	_CASILLA_OCUPADA         int    = 1
	_CLAVE_BORRADA           int    = 2
	_CAPACIDAD_INICIAL       int    = 20
	_PANIC_CLAVE_DICCIONARIO string = "La clave no pertenece al diccionario"
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

func crearTabla[K comparable, V any](capacidad int) []hashElemento[K, V] {
	return make([]hashElemento[K, V], capacidad) //Al pasar la capacidad por parametro, vamos a poder "redimensionar" cuando haya más elementos en la tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] { //CrearHash inicializará la capacidad del vector y la tabla. Esa tabla es la que nos importa para la lógica
	return &tablaHash[K, V]{capacidad: _CAPACIDAD_INICIAL, tabla: []hashElemento[K, V]{}}
}

func (th *tablaHash[K, V]) Guardar(clave K, dato V) {
	

	return
}

func (hash *tablaHash[K, V]) Pertenece(clave K) bool {
	return true
}

func (hash *tablaHash[K, V]) Obtener(clave K) V {
	return
}

func (hash *tablaHash[K, V]) Borrar(clave K) V {
	return
}

func (hash *tablaHash[K, V]) Cantidad() int {
	return
}

func (hash *tablaHash[K, V]) Iterar(func(clave K, dato V) bool) {
	return
}

func (hash *tablaHash[K, V]) Iterador() IterDiccionario[K, V] {
	return
}

/*
****************************************************************
-----------------FUNCIONES AUXILIARES--------------------------
****************************************************************
*/

func (th tablaHash[K, V]) buscarEstado(pos int) int {
	if th.tabla[pos].estado == _CASILLA_VACIA || th.tabla[pos].estado  == _CLAVE_BORRADA{
		return pos
	}
	for pos := pos + 1 ; pos < len(th.tabla) ; pos ++ {
		if pos
	}

}

// CODIGO FUENTE : https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
// USAMOS EL METODO FNV-HASHING
func hashing[K comparable](clave K, largo int) int {
	var (
		base uint32 = 2166136261
		dato uint32 = 16777619
	)
	arrayBytes := convertirABytes(clave)
	for _, valor := range arrayBytes {
		base ^= uint32(valor)
		base *= dato
	}
	return int(base) % largo

}

func convertirABytes[K comparable](clave K) []byte { //Función de la cátedra para obtener un array de bytes de la clave
	return []byte(fmt.Sprintf("%v", clave))
}
