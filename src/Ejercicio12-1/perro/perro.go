// Package perro, es un paquete con funciones para el ejercicio 12-1 del curso en GO.
package perro

// EdadPerro es una función exportada que recibe como argumento los años de un humano (edadHumano) y arroja
// el resultado de multiplicarlos por 7. Así se obtiene un int de la conversión de esos años a edad de perro.
func EdadPerro(edadHumano int) int {
	return edadHumano * 7
}
