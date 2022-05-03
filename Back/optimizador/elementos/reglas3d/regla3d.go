package reglas3d

import (
	"Back/optimizador/elementos/bloque3d"
	"Back/optimizador/elementos/interfaces3d"
	"strings"
)

func ComprobarReglas(bloque bloque3d.Bloque3d) string {
	mapStrings := make(map[int]string)
	//contador := 0
	instrucciones := bloque.Instrucciones
	//var tipoActual interfaces3d.TipoCD3
	//var tipoActual2 interfaces3d.TipoCD3
	//var tipoActual3 interfaces3d.TipoCD3
	salida := ""
	operacionesActual := ""
	operacionesActual2 := ""
	operacionesActual3 := ""
	temporalActual := ""
	temporalActual2 := ""
	elementoAmodificar := ""
	var elementoActual interfaces3d.Expresion3D
	//var elementoActual2 interfaces3d.Expresion3D
	//var elementoActual3 interfaces3d.Expresion3D
	//var regla2Encontrada bool = false
	/*PRIMERO TODOS A TEXTO*/
	for i := 0; i < instrucciones.Len(); i++ {
		elementoActual = instrucciones.GetValue(i).(interfaces3d.Expresion3D)
		mapStrings[i] = elementoActual.GetValor()
	}

	/*SEGUNDO VERIFICAR REGLA 1*/
	for i := 0; i < instrucciones.Len(); i++ {
		elemento := mapStrings[i]
		arr1 := strings.Split(elemento, "=")
		if !strings.Contains(elemento, "=") {
			continue
		}
		temporalActual = strings.TrimSpace(arr1[0])
		operacionesActual = arr1[1]
		if strings.Contains(operacionesActual, "P") ||
			strings.Contains(operacionesActual, "H") ||
			strings.Contains(operacionesActual, "stack") ||
			strings.Contains(operacionesActual, "heap") ||
			strings.Contains(temporalActual, "P") ||
			strings.Contains(temporalActual, "H") {
			continue
		}
		for j := i + 1; j < instrucciones.Len(); j++ {
			if j == i {
				continue
			}
			elemento2 := mapStrings[j]
			if !strings.Contains(elemento2, "=") {
				continue
			}
			arr2 := strings.Split(elemento2, "=")
			temporalActual2 = strings.TrimSpace(arr2[0])
			operacionesActual2 = arr2[1]
			if operacionesActual == operacionesActual2 {
				if strings.Contains(temporalActual, "stack") ||
					strings.Contains(temporalActual, "heap") ||
					strings.Contains(temporalActual, "P") ||
					strings.Contains(temporalActual, "H") ||
					strings.Contains(temporalActual2, "stack") ||
					strings.Contains(temporalActual2, "heap") ||
					strings.Contains(temporalActual2, "P") ||
					strings.Contains(temporalActual2, "H") {
					continue
				}
				elementoAmodificar = mapStrings[j]
				elementoAmodificar = strings.Replace(elementoAmodificar, operacionesActual2, " "+temporalActual+";\n", -1)
				mapStrings[j] = elementoAmodificar
				/*VERIFICAR REGLA 2*/
				for k := j + 1; k < instrucciones.Len(); k++ {
					elemento3 := mapStrings[k]
					if !strings.Contains(elemento3, "=") {
						continue
					}
					arr3 := strings.Split(elemento3, "=")
					operacionesActual3 = arr3[1]
					if strings.Contains(operacionesActual3, temporalActual2) {
						elementoAmodificar = mapStrings[k]
						elementoAmodificar = strings.Replace(elementoAmodificar, temporalActual2, temporalActual, -1)
						mapStrings[k] = elementoAmodificar
						/*VERIFICAR REGLA 3*/
						mapStrings[j] = ""
					}
				}
			}

		}

	}

	/*RECUPERAR TODAS LAS INSTRUCCIONES DESDE EL MAP*/
	for i := 0; i < instrucciones.Len(); i++ {
		salida += mapStrings[i]
	}

	return salida
}
