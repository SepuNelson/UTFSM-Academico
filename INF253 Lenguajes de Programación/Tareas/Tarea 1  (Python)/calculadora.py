"""
=========================================================
                B I B L I O T E C A S                    
=========================================================
"""
import re

"""
=========================================================
        E X P R E SI O N E S   R E G U L A R E S                     
=========================================================
"""

digito = r'[1-9]'
digito_o_zero = r'[0-9]'
entero = fr'(?:{digito}{digito_o_zero}*)|0'
espacio = r'(?:\s*)'
clave = fr'(?:ANS|CUPON\s*\(\s*(?:{entero}|ANS)(?:\s*,\s*(?:{entero}|ANS))?\s*\))'
operador = fr'{espacio}(?:\+|-|\*|//){espacio}'
operacion = fr'(?:{clave}|{entero}){operador}(?:{clave}|{entero})'
sentencia = fr'{operacion}(?:{operador}(?:{entero}|{clave}))*'

"""
=========================================================
                   F U N C I O N E S                    
=========================================================
"""
def listas(linea, resultado_actual):
    '''
    ***
    * linea : String
    * resultado_actual : Int
    ***
    La función recibe la linea en la cual se va leyendo el archivo.txt, para ir separando los datos que necesito en 2 listas, num_list y
    op_list, hace filtros si encuentra entero, ANS o CUPON en sus 2 variantes estas 2 ultimas se calcula el valor numerico ANS = resultado_actual y
    CUPON dependiendo de su formula para retornar las 2 listas en una sola lista list, donde num_list = list[0] y op_list = list[1] en un futuro.
    '''
    lista = []
    num_list = []
    op_list = []
    inicio = 0
    final = 0
    descuento = 0
    while final < len(linea):
        if linea[final] == " ":
            descuento += 1
        if linea[final] == "+" or linea[final] == "-" or linea[final] == "*" or linea[final] == "/" or linea[final] == "":
            texto = linea[inicio:final-descuento]
            descuento = 0
            if re.fullmatch(entero, texto) != None:
                num_list.append(int(texto))
                if linea[final] == "/":
                    op_list.append("//")
                else:
                    op_list.append(linea[final])
            elif re.fullmatch(clave, texto) != None:
                if texto == "ANS":
                    if resultado_actual == None:
                        return "Error"
                    else:    
                        num_list.append(resultado_actual)
                else:
                    if "," in texto:
                        xy = re.findall(entero,texto)
                        if calcular_cupon_xy(int(xy[0]),int(xy[1])) == "Error":
                            archivo2.write(linea + " = Error\n")
                            return []
                        else:
                            num_list.append(calcular_cupon_xy(int(xy[0]),int(xy[1])))
                    else:
                        x = re.findall(entero,texto)
                        num_list.append(calcular_cupon_x(int(x[0])))

                if linea[final] == "/":
                    op_list.append("//")
                else:
                    op_list.append(linea[final])
            inicio = final + 1
            while linea[inicio] == " ":
                final += 1
                inicio += 1

        elif final == (len(linea) - 1):
                while linea[inicio] == " ":
                    inicio += 1
                texto = linea[inicio:]
                if re.fullmatch(entero, texto) != None:
                    num_list.append(int(texto))
                elif re.fullmatch(clave, texto) != None:
                    if texto == "ANS":
                        num_list.append(resultado_actual)
                    else:
                        if "," in texto:
                            xy = re.findall(entero,texto)
                            if calcular_cupon_xy(int(xy[0]),int(xy[1])) == "Error":
                                archivo2.write(linea + " = Error\n")
                                return []
                            else:
                                num_list.append(calcular_cupon_xy(int(xy[0]),int(xy[1])))
                        else:
                            x = re.findall(entero,texto)
                            num_list.append(calcular_cupon_x(int(x[0])))
        final += 1
    lista.append(num_list)
    lista.append(op_list)
    return lista

def evaluar(op, num_1, num_2):
    '''
    ***
    * op : String
    * num_1 : Int
    * num_2 : Int
    ***
    La función recibe 3 parámetros, el primero es un simbolo de +,-,* o / , seguido de un numero 1 y un numero 2, los que se sumarán, restarán, 
    multiplicarán o dividirán dependiendo de la comparación que reciba de los ifs/elifs, retornando el resultado de esa operación.
    '''
    if op == '+':
        return num_1 + num_2
    elif op == '-':
        return num_1 - num_2
    elif op == '*':
        return num_1 * num_2
    elif op == '//':
        if num_2 == 0:
            return "Error"
        else:
            return num_1 // num_2

def calcular_cupon_x(x):
    '''
    ***
    * x : Int
    ***
    Retorna el valor del número x con el 20% de descuento.
    '''
    resultado = int(x * 0.20)
    return resultado

def calcular_cupon_xy(x, y):
    '''
    ***
    * x : Int
    * y : Int
    ***
    Retorna el valor del número x con el descuento y.
    '''
    if 0 <= y <= 100:
        resultado = int(x * y / 100)
        return resultado
    else:
        return "Error"

def resolver(num_list, op_list, resultado_actual):
    '''
    ***
    * num_list : Lista de Números y ANSs
    * op_list : Lista de Operadores
    * resultado_actual : Int
    ***
    La función recibe 3 parámetros, el primero es una lista de Números y ANSs, el segundo es una lista de Operadores y el tercero es el resultado de las lineas,
    la función se encarga de ir buscando y leyendo las operaciones por el orden en que se deben resolver, retornando el resultado final.
    '''
    while op_list != []:
        if "*" in op_list or "//" in op_list:
            for i in op_list:
                if i == "*" or i == "//":
                    pos = op_list.index(i)
                    if re.findall(clave,str(num_list[pos])) != []:
                        num_list[pos] = resultado_actual
                        if re.findall(clave,str(num_list[pos + 1])) != []:
                            num_list[pos + 1] = resultado_actual
                    else:
                        if re.findall(clave,str(num_list[pos + 1])) != []:
                            num_list[pos + 1] = resultado_actual
                    if evaluar(i,int(num_list[pos]),int(num_list[pos + 1])) == "Error":
                        return "Error"
                    else:
                        result = evaluar(i,int(num_list[pos]),int(num_list[pos + 1]))
                        num_list[pos] = result
                        num_list.pop(pos + 1)
                        op_list.pop(pos)

        elif "+" in op_list or "-" in op_list:
            for i in op_list:
                if i == "+" or i == "-":
                    pos = op_list.index(i)
                    if re.findall(clave,str(num_list[pos])) != []:
                        num_list[pos] = resultado_actual
                        if re.findall(clave,str(num_list[pos + 1])) != []:
                            num_list[pos + 1] = resultado_actual
                    else:
                        if re.findall(clave,str(num_list[pos + 1])) != []:
                            num_list[pos + 1] = resultado_actual
                    result = evaluar(i,int(num_list[pos]),int(num_list[pos + 1]))
                    if result < 0:
                        result = 0
                    num_list[pos] = result
                    num_list.pop(pos + 1)
                    op_list.pop(pos)
    return num_list[0]

"""
=========================================================
                    A R C H I V O S                    
=========================================================
"""

archivo1 = open("problemas.txt", "r")
archivo2 = open("desarrollos.txt", "w")
resultado_actual = None
flag = True
for linea in archivo1:
    linea = linea[:-1]
    if re.fullmatch(sentencia,linea) != None:
        list = listas(linea, resultado_actual)
        if list != [] and list != "Error":
            num_list = list[0]
            op_list = list[1]
            if resolver(num_list, op_list, resultado_actual) == "Error":
                archivo2.write(linea + " = Error\n")
                flag = False
            else:
                if flag == True:
                    resp = resolver(num_list, op_list, resultado_actual)
                    resultado_actual = resp
                    archivo2.write(linea + " = " + str(resp) + "\n")

                elif flag == False:
                    archivo2.write(linea + " = Sin Resolver\n")
        elif list == "Error":
            archivo2.write(linea + " = Sin Resolver\n")
            flag = False
        else:
            flag = False

    elif re.fullmatch(sentencia,linea) == None and linea != "":
        flag = False
        archivo2.write(linea + " = Error\n")  

    elif linea == "":
        flag = True
        resultado_actual = None
        archivo2.write(linea + "\n")

archivo1.close()
archivo2.close()