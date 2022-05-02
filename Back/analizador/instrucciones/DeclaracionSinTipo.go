package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionSinTipo struct {
	Id      string
	Mutable bool
	Publico bool
	Valor   interface{}
	Fila    int
	Columna int
	Stack   bool
}

func NewDeclaracionSinTipo(id string, valor interface{}, mutable, publico bool,
	fila int, columna int) DeclaracionSinTipo {
	nd := DeclaracionSinTipo{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
		Stack:   true,
	}
	return nd
}

func (d DeclaracionSinTipo) Run(scope *Ast.Scope) interface{} {
	var nSimbolo Ast.Simbolo
	var cod3D string = ""
	var obj3DValor, obj3D Ast.O3D
	var temp string
	var direccion int

	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	//Primero verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else if tipoIn == Ast.FUNCION {
		preValor = Ast.TipoRetornado{
			Valor: d.Valor,
			Tipo:  Ast.FUNCION,
		}
	} else {
		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	obj3DValor = preValor.(Ast.TipoRetornado).Valor.(Ast.O3D)
	valor := obj3DValor.Valor
	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		/////////////////////////////ERROR/////////////////////////////////////
		return valor
	}
	/*Error ya existe*/
	if existe {
		/////////////////////////////ERROR/////////////////////////////////////
		return errores.GenerarError(12, d, d, d.Id, "", "", scope)
	}

	//No existe, entonces agregarla
	//Verificar la mutabilidad para cambiarla si es un vector o un array o un struct
	//Agregar el tipo al simbolo en el apartado de tipo especial para despues
	if valor.Tipo == Ast.VECTOR {
		elemento := valor.Valor.(expresiones.Vector)
		elemento.Mutable = d.Mutable
		valor.Valor = elemento
	} else if valor.Tipo == Ast.ARRAY {
		elemento := valor.Valor.(expresiones.Array)
		elemento.Mutable = d.Mutable
		valor.Valor = elemento
	} else if valor.Tipo == Ast.STRUCT {
		nValor := valor.Valor.(Ast.Structs).SetMutabilidad(d.Mutable)
		valor.Valor = nValor
	} else if valor.Tipo == Ast.DIMENSION_ARRAY {
		elemento := valor.Valor.(expresiones.Array)
		elemento.Mutable = d.Mutable
		valor.Tipo = Ast.ARRAY
		valor.Valor = elemento
	}
	//Crear símbolo y agregarlo a la tabla del entorno actual
	nSimbolo = Ast.Simbolo{
		Identificador: d.Id,
		Valor:         valor,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          valor.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
		Entorno:       scope,
	}

	//Puede que sea un tipo especial

	if EsTipoEspecial(valor.Tipo) {
		//Agregar el tipo al simbolo en el apartado de tipo especial para despues
		if valor.Tipo == Ast.VECTOR {
			elemento := valor.Valor.(expresiones.Vector)
			nSimbolo.TipoEspecial = elemento.TipoVector
		} else if valor.Tipo == Ast.ARRAY {
			elemento := valor.Valor.(expresiones.Array)
			nSimbolo.TipoEspecial = elemento.TipoDelArray
		} else if valor.Tipo == Ast.STRUCT {
			elemento := valor.Valor.(Ast.Structs).GetPlantilla(scope)
			nSimbolo.TipoEspecial = Ast.TipoRetornado{
				Valor: elemento,
				Tipo:  Ast.STRUCT,
			}
		}
	}

	//Desde aquí trabajar el C3D

	//Verificar si la variable se va a guardar en el heap o en el stack

	if d.Stack {
		//Agregar el código anterior
		cod3D += obj3DValor.Codigo

		if valor.Tipo == Ast.BOOLEAN {
			if obj3DValor.Lt != "" {
				salto := Ast.GetLabel()
				cod3D += "/*************************CAMBIO VALOR BOOLEANO*/ \n"
				cod3D += obj3DValor.Lt + ": \n"
				cod3D += obj3DValor.Referencia + " = 1;\n"
				cod3D += "goto " + salto + ";\n"
				cod3D += obj3DValor.Lf + ": \n"
				cod3D += obj3DValor.Referencia + " = 0;\n"
				cod3D += salto + ":\n"
				cod3D += "/***********************************************/ \n"
			}
		}

		cod3D += "/***********************DECLARACIÓN DE VARIABLE*/ \n"
		/*Get el nuevo temporal*/
		temp = Ast.GetTemp()
		/*Get la dirección donde se guardara */
		direccion = scope.ContadorDeclaracion
		//Conseguir la dirección del stack donde se va a guardar la nueva variable
		cod3D += temp + " = " + " P + " + strconv.Itoa(direccion) + ";\n"
		//Se guarda la nueva variable en el stack
		cod3D += "stack[(int)" + temp + "] = " + obj3DValor.Referencia + ";\n"
		//Aumentar el SP
		scope.ContadorDeclaracion++
		cod3D += "/***********************************************/\n"
		//Agregar la dirección al símbolo
		nSimbolo.Direccion = direccion
		nSimbolo.TipoDireccion = Ast.STACK
		//Actualizar el obj3d
		obj3D.Codigo = cod3D
		obj3D.Valor = valor
	} else {

		//Agregar el código anterior
		cod3D += obj3DValor.Codigo

		if valor.Tipo == Ast.BOOLEAN {
			if obj3DValor.Lt != "" {
				salto := Ast.GetLabel()
				cod3D += "/*************************CAMBIO VALOR BOOLEANO*/ \n"
				cod3D += obj3DValor.Lt + ": \n"
				cod3D += obj3DValor.Referencia + " = 1;\n"
				cod3D += "goto " + salto + ";\n"
				cod3D += obj3DValor.Lf + ": \n"
				cod3D += obj3DValor.Referencia + " = 0;\n"
				cod3D += salto + ":\n"
				cod3D += "/***********************************************/ \n"
			}
		}

		cod3D += "/***********************DECLARACIÓN DE VARIABLE*/ \n"
		/*Get el nuevo temporal*/
		temp = Ast.GetTemp()
		/*Get la dirección donde se guardara */
		direccion = Ast.GetH()
		//Conseguir la dirección del stack donde se va a guardar la nueva variable
		cod3D += temp + " = " + " H;\n"
		//Se guarda la nueva variable en el stack
		cod3D += "heap[(int)" + temp + "] = " + obj3DValor.Referencia + ";\n"
		//Aumentar H
		cod3D += "H = H + 1;\n"
		cod3D += "/***********************************************/\n"

		//Aumentar el scope
		scope.ContadorDeclaracion++
		//Agregar la dirección al símbolo
		nSimbolo.Direccion = direccion
		nSimbolo.TipoDireccion = Ast.HEAP
		//Actualizar el obj3d
		obj3D.Codigo = cod3D
		obj3D.Valor = valor
	}

	scope.Add(nSimbolo)

	return Ast.TipoRetornado{
		Tipo:  Ast.DECLARACION,
		Valor: obj3D,
	}
}

func (op DeclaracionSinTipo) GetFila() int {
	return op.Fila
}
func (op DeclaracionSinTipo) GetColumna() int {
	return op.Columna
}

func (d DeclaracionSinTipo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionSinTipo) SetHeap() interface{} {
	Dcopia := d
	Dcopia.Stack = false
	return Dcopia
}
