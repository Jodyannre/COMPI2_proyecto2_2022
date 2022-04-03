package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
)

type DeclaracionConTipo struct {
	Id      string
	Mutable bool
	Publico bool
	Tipo    Ast.TipoRetornado
	Fila    int
	Columna int
	Stack   bool
}

func NewDeclaracionConTipo(id string, tipo Ast.TipoRetornado, mutable, publico bool,
	fila int, columna int) DeclaracionConTipo {
	nd := DeclaracionConTipo{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Tipo:    tipo,
		Fila:    fila,
		Columna: columna,
		Stack:   true,
	}
	return nd
}

func (d DeclaracionConTipo) Run(scope *Ast.Scope) interface{} {
	var nSimbolo Ast.Simbolo
	var obj3D Ast.O3D
	var direccion int
	var codigo3d string

	//Verificar que el id no exista
	existe := scope.Exist_actual(d.Id)

	//Verificar si ya existe
	if existe {
		//Ya existe y generar error semántico
		return errores.GenerarError(12, d, d, d.Id, "", "", scope)
	}

	//Verificar si el tipo es un acceso a un módulo
	if d.Tipo.Tipo == Ast.ACCESO_MODULO {
		//Traer el tipo y cambiar el tipo de la declaración
		nTipo := d.Tipo.Valor.(Ast.AccesosM).GetTipoFromAccesoModulo(d.Tipo, scope)
		if nTipo.Tipo == Ast.ERROR {
			return nTipo
		}
		if nTipo.Tipo == Ast.STRUCT_TEMPLATE {
			nTipo.Tipo = Ast.STRUCT
		}
		d.Tipo = nTipo
	}

	//No existe, entonces agregarla
	//Crear símbolo y agregarlo a la tabla del entorno actual

	nSimbolo = Ast.Simbolo{
		Identificador: d.Id,
		Valor:         Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL},
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          d.Tipo.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
		Entorno:       scope,
	}

	//Verificar si no es vector, array o struct para agregar el tipo especial
	if EsTipoEspecial(d.Tipo.Tipo) {
		nSimbolo.TipoEspecial = d.Tipo
	}

	/*Aquí esta todo lo de C3D*/
	if d.Stack {
		direccion = scope.Size
		nSimbolo.Direccion = direccion
		nSimbolo.TipoDireccion = Ast.STACK
		//Actualizar el obj3d
		obj3D = ValorPorDefecto(d.Tipo.Tipo, scope)
		nSimbolo.Valor = obj3D.Valor
	} else {
		direccion = Ast.GetH()
		nSimbolo.Direccion = direccion
		nSimbolo.TipoDireccion = Ast.HEAP
		/*Código del aumento del heap*/
		codigo3d = "H = H + 1;\n"
		//Actualizar el obj3d
		obj3D = ValorPorDefecto(d.Tipo.Tipo, scope)
		obj3D.Codigo = codigo3d
		nSimbolo.Valor = obj3D.Valor
	}

	scope.Add(nSimbolo)

	return Ast.TipoRetornado{
		Tipo:  Ast.DECLARACION,
		Valor: obj3D,
	}
}

func (op DeclaracionConTipo) GetFila() int {
	return op.Fila
}
func (op DeclaracionConTipo) GetColumna() int {
	return op.Columna
}

func (d DeclaracionConTipo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionConTipo) SetHeap() interface{} {
	Dcopia := d
	Dcopia.Stack = false
	return Dcopia
}
