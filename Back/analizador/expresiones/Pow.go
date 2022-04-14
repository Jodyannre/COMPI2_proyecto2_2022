package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"math"
	"strconv"
)

type Pow struct {
	Tipo     Ast.TipoDato
	TipoOp   Ast.TipoDato
	Valor    interface{}
	Potencia interface{}
	Fila     int
	Columna  int
}

func NewPow(tipo Ast.TipoDato, tipoOp Ast.TipoDato, valor interface{}, potencia interface{}, fila, columna int) Pow {
	nP := Pow{
		Tipo:     tipo,
		TipoOp:   tipoOp,
		Valor:    valor,
		Potencia: potencia,
		Fila:     fila,
		Columna:  columna,
	}
	return nP
}

func (p Pow) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/************************************VARIABLES 3D***************************************/
	var obj3d, obj3dValor, obj3dPotencia Ast.O3D
	var codigo3d, referencia string
	/***************************************************************************************/

	valor := p.Valor.(Ast.Expresion).GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor
	potencia := p.Potencia.(Ast.Expresion).GetValue(scope)
	obj3dPotencia = potencia.Valor.(Ast.O3D)
	potencia = obj3dPotencia.Valor
	codigo3d += obj3dValor.Codigo
	codigo3d += obj3dPotencia.Codigo

	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if potencia.Tipo == Ast.ERROR {
		return potencia
	}

	if valor.Tipo != p.TipoOp {
		//Error, tipos diferentes en la operacion
		fila := p.Valor.(Ast.Abstracto).GetFila()
		columna := p.Valor.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected " + Ast.ValorTipoDato[p.TipoOp] + ", found " + Ast.ValorTipoDato[valor.Tipo] + "." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	if potencia.Tipo != p.TipoOp {
		//Error, tipos diferentes en la operacion
		fila := p.Potencia.(Ast.Abstracto).GetFila()
		columna := p.Potencia.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected " + Ast.ValorTipoDato[p.TipoOp] + ", found " + Ast.ValorTipoDato[potencia.Tipo] + "." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Todo bien, operar
	respuesta := Ast.TipoRetornado{
		Tipo:  p.TipoOp,
		Valor: true,
	}
	if p.TipoOp == Ast.I64 {
		respuesta.Valor = int(math.Pow(float64(valor.Valor.(int)), float64(potencia.Valor.(int))))
	} else {
		respuesta.Valor = math.Pow(valor.Valor.(float64), potencia.Valor.(float64))
	}

	/**************************GET CODIGO 3D DE OPERACION******************************/
	cod, ref := GetCod3DPowInt(obj3dValor.Referencia, obj3dPotencia.Referencia)
	codigo3d += cod
	referencia = ref
	/*********************************************************************************/

	obj3d.Valor = respuesta
	obj3d.Codigo = codigo3d
	obj3d.Referencia = referencia

	return Ast.TipoRetornado{
		Tipo:  respuesta.Tipo,
		Valor: obj3d,
	}
}

func (p Pow) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p Pow) GetFila() int {
	return p.Fila
}
func (p Pow) GetColumna() int {
	return p.Columna
}

func GetCod3DPowInt(val, pot string) (string, string) {
	var contador, resultado, valor, potencia, temp, temp2 string
	var lt, lt2, lt3, lf, lf2, lf3, salto, salto2, salto3 string
	var codigo3d string
	valor = Ast.GetTemp()
	potencia = Ast.GetTemp()
	contador = Ast.GetTemp()
	resultado = Ast.GetTemp()
	temp = Ast.GetTemp()
	temp2 = Ast.GetTemp()
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()
	lt2 = Ast.GetLabel()
	lf2 = Ast.GetLabel()
	salto2 = Ast.GetLabel()
	lt3 = Ast.GetLabel()
	lf3 = Ast.GetLabel()
	salto3 = Ast.GetLabel()
	codigo3d += "/*******************************POTENCIA DE INT*/\n"
	codigo3d += valor + " = " + val + "; //Guardar el valor \n"
	codigo3d += potencia + " = " + pot + "; //Guardar la potencia \n"
	codigo3d += contador + " = 1; //Iniciar contador \n"
	codigo3d += resultado + " = " + val + "; //Iniciar resultado \n"
	codigo3d += "/*********************VALIDACION POTENCIA 0 y 1*/\n"
	codigo3d += "if (" + potencia + " == 0) goto " + lt + "; \n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += resultado + " = 1; //Resultado de elevar el num a 0 \n"
	codigo3d += "goto " + salto + ";\n"

	codigo3d += lf + ":\n"
	codigo3d += "if (" + potencia + " == 1) goto " + lt2 + "; \n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += lt2 + ":\n"
	codigo3d += "goto " + salto2 + ";\n"

	codigo3d += "/***********************************************/\n"
	codigo3d += lf2 + ":\n"
	codigo3d += salto3 + ":\n"
	codigo3d += "if (" + contador + " < " + potencia + ") goto " + lt3 + "; \n"
	codigo3d += "goto " + lf3 + ";\n"
	codigo3d += lt3 + ": \n"
	codigo3d += temp + " = " + resultado + " * " + valor + "; //multiplicación por si mismo \n"
	codigo3d += resultado + " = " + temp + "; //Agregar a resultado \n"
	codigo3d += temp2 + " = " + contador + " + 1; \n"
	codigo3d += contador + " = " + temp2 + "; //Actualizar contador \n"
	codigo3d += "goto " + salto3 + ";\n"
	codigo3d += lf3 + ":\n"
	codigo3d += salto + ":\n"
	codigo3d += salto2 + ":\n"
	codigo3d += "/***********************************************/\n"
	return codigo3d, resultado
}
