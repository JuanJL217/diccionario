package diccionario_test

import (
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const TAMANIO = 10000

var arrVolumen = make([]int, TAMANIO)

func compararCadenas(a, b string) int {
	return strings.Compare(a, b)
}

func compararNumeros(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func numeroAleatorio(n int) int {
	return rand.Intn(n)
}

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestABBClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compararNumeros)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElementoABB(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](compararCadenas)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoABB(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestABBBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorradosABB(t *testing.T) {
	t.Log("Prueba de reutilizacion de borrados")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericasABB(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararNumeros)
	clave := 10
	valor := "Gatito"
	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVaciaABB(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloABB(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](compararCadenas)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}
func TestIteradorInternoSinRangoVacio(t *testing.T) {
	t.Log("Iterar sobre un diccionario vacio usando el iterador interno sin rango")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Iterar(func(c, v int) bool {
		require.Equal(t, 0, dic.Cantidad())
		return true
	})
	require.False(t, dic.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dic.Borrar(1)
		dic.Obtener(1)
	})
}

func TestIteradorInternoSinRangoCompleto(t *testing.T) {
	t.Log("Iterar todos los elementos sin condicion de corte usando el iterador interno sin rango")
	claves := []string{"H", "A", "B", "Z", "S", "T", "R"}
	valor := []int{1, 2, 3, 4, 5, 6, 7}
	dic := TDADiccionario.CrearABB[string, int](compararCadenas)
	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valor[i])
	}
	require.Equal(t, 7, dic.Cantidad())
	contador := 0
	dic.Iterar(func(c string, v int) bool {
		require.Equal(t, valor[contador], dic.Obtener(claves[contador]))
		contador++
		return true
	})
	require.Equal(t, contador, dic.Cantidad(), "El contador deberia coincidir con la cantidad que hay en el diccionario")
}

func TestIteradorInternoSinRangoConCorte(t *testing.T) {
	t.Log("Iterar sobre un diccionario con condicion de corte usando el iterador interno sin rango")
	elementos := []int{6, 1, 4, 15, 10, 16, 8, 13, 11, 14}
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	for _, valor := range elementos {
		dic.Guardar(valor, valor)
	}
	contador := 0
	dic.Iterar(func(c, v int) bool {
		if contador == 5 {
			return false
		}
		contador++
		return true
	})
	require.Equal(t, elementos[contador], dic.Obtener(elementos[contador]), "obtener el elemento en la posicion del contador deben coincidir ")
}
func TestIteradorInternoRangoCorte(t *testing.T) {
	t.Log("Iteramos un diccionario con elementos usando el iterador interno con limites")
	datos := []int{100, 50, 150, 25, 75, 125, 200, 10, 30, 60, 80, 110, 130, 190, 300}
	minimo := 60
	maximo := 120
	solucion := 0
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	for _, valor := range datos {
		dic.Guardar(valor, valor)
		if valor >= minimo && valor <= maximo {
			solucion += valor
		}
	}

	suma := 0
	dic.IterarRango(&minimo, &maximo, func(c int, v int) bool {
		//sumamos todas las claves que esten compredidos en el rango
		suma += c
		return true
	})
	require.Equal(t, solucion, suma)

}

func TestIteradorInternoConRangoCompleto(t *testing.T) {
	t.Log("Iteramos todos los elementos usando el iterador interno con rango , el comportamiento tiene que ser como el iterador interno sin rango")
	var (
		minimo *string
		maximo *string
	)
	datos := []string{"H", "I", "G", "F", "J", "E", "K"}
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	for _, valor := range datos {
		dic.Guardar(valor, valor)
	}
	contador := 0
	dic.IterarRango(minimo, maximo, func(c string, v string) bool {
		contador++
		return true
	})
	require.Equal(t, contador, dic.Cantidad())
}

func TestIteradorExternoSinRangoVacioABB(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](compararCadenas)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func buscarClave(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorExternoSinRangoCompleto(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarClave(primero, claves))

	iter.Siguiente()
	segundo, segundoValor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarClave(segundo, claves))
	require.EqualValues(t, valores[buscarClave(segundo, claves)], segundoValor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarClave(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoSinRango(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	claves := []string{"B", "A", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarClave(primero, claves))
	require.NotEqualValues(t, -1, buscarClave(segundo, claves))
	require.NotEqualValues(t, -1, buscarClave(tercero, claves))
}

func TestVolumenABB(t *testing.T) {
	t.Log("Prueba de volumen")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	for i := 0; i < len(arrVolumen); i++ {
		n := numeroAleatorio(TAMANIO)
		arrVolumen[i] = n
		dic.Guardar(n, n)
	}
	for i := 0; i < len(arrVolumen); i++ {
		require.True(t, dic.Pertenece(arrVolumen[i]))
	}
}
