"""
===================================================================================
.                           F U N C I O N E S
===================================================================================
"""   
# Función que verifica si un número "x" en base "y" puede ser representado. // Recibe como parámetros 2 enteros (int).
def vrf(x, y):
    digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    x_string = str(x)
    for i in x_string:
        if int(digitos.index(i)) >= int(y):
            return 1
    return 0    

# Función que se utiliza cuando se debe cambiar desde y/o hacia una base mayor a 10. // Recibe como parámetros 2 enteros (int).
def conversor_masdiez(a,b):
    i = len(a)-1
    ind = i-1
    ret = ""
    ret_aux= ""
    digito = a
    while i > 0:
        dig = 0
        if digito[ind] == "1":
            if digito[i] == "0" and b > 10:
                while dig < len(digito)-1:
                    if dig < ind: 
                        ret_aux = ret_aux + digito[dig]
                    if dig == ind:
                        ret_aux = ret_aux + "A"
                    if dig > ind:
                        ret_aux = ret_aux + digito[dig+1] 
                    dig += 1
                ret = ret_aux
                ret_aux = ""    
                i = i-1
                ind = i -1    
            elif digito[i] == "1" and b > 11:
                while dig < len(digito)-1:
                    if dig < ind: 
                        ret_aux = ret_aux + digito[dig]
                    if dig == ind:
                        ret_aux = ret_aux + "B"
                    if dig > ind:
                        ret_aux = ret_aux + digito[dig+1]
                    dig += 1
                ret = ret_aux
                ret_aux = ""    
                i = i-1
                ind = i -1
            elif digito[i] == "2" and b > 12:
                while dig < len(digito)-1:
                    if dig < ind: 
                        ret_aux = ret_aux + digito[dig]
                    if dig == ind:
                        ret_aux = ret_aux + "C"
                    if dig > ind:
                        ret_aux = ret_aux + digito[dig+1]
                    dig += 1
                ret = ret_aux
                ret_aux = ""    
                i = i-1
                ind = i -1
            elif digito[i] == "3" and b > 13:
                while dig < len(digito)-1:
                    if dig < ind: 
                        ret_aux = ret_aux + digito[dig]
                    if dig == ind:
                        ret_aux = ret_aux + "D"
                    if dig > ind:
                        ret_aux = ret_aux + digito[dig+1]
                    dig += 1
                ret = ret_aux
                ret_aux = ""    
                i = i-1
                ind = i -1
            elif digito[i] == "4" and b > 14:
                while dig < len(digito)-1:
                    if dig < ind: 
                        ret_aux = ret_aux + digito[dig]
                    if dig == ind:
                        ret_aux = ret_aux + "E"
                    if dig > ind:
                       ret_aux = ret_aux + digito[dig+1]
                    dig += 1
                ret = ret_aux
                ret_aux = ""    
                i = i-1
                ind = i -1
            elif digito[i] == "5" and b > 15:
                while dig < len(digito)-1:
                    if dig < ind: 
                        ret_aux = ret_aux + digito[dig]
                    if dig == ind:
                        ret_aux = ret_aux + "F"
                    if dig > ind:
                        ret_aux = ret_aux + digito[dig+1]
                    dig += 1  
                ret = ret_aux
                ret_aux = ""                               
                i = i-1
                ind = i -1
        i = i-1
        ind = i-1
        #print(ret + " prim " + str(i) + " " + str(ind) )
        digito = ret
    copia = len(ret)-1
    while copia >= 0:
        ret_aux = ret_aux + ret[copia]
        copia -=1
        
    return ret

# Función que convierte un string a número.  // Recibe como parámetro un string (str).
def convert_stoi(a):
    if a == "A":
        a= "10" 
    elif a == "B":
        a= "11"
    elif a == "C":
        a = "12"
    elif a == "D":
        a = "13"
    elif a == "E":
        a = "14"
    elif a == "F":
        a = "15" 
    return a 
    
# Función que convierte un número "a" de base "b" a base "c". // Recibe como parámetros 3 enteros (int).
def convert_base(a,b,c):
    letras = "ABCDEF"
    decimal = 0
    for digit in str(a):
        if digit in letras:
            digit = convert_stoi(digit)
        decimal = decimal * int(b) + int(digit)
        
    result = ""
    while decimal > 0:
        digit = decimal % int(c)
        result = str(digit) + result
        decimal //= int(c) 
    
    if c > 10 and int(a)%c > 9:
        result = conversor_masdiez(result, c)
        
    return result
 
"""
===================================================================================
.                                 M A I N
===================================================================================
"""

flag = True
while (flag == True):
    bits = int(input("Ingrese número entre 0 y 32 incluídos: ")) 

    A = 0 # Contador de números en el Archivo
    B = 0 # Contador de números con error en la representación numérica
    C = 0 # Contador de números que no pueden ser representados con el valor ingresado
    D = 0 # Contador de sumas realizadas en complemento dos que provocan overflow en registros con el tamaño ingresado por el usuario

    #Primer Archivo ==> numeros.txt
    archivo1 = open("numeros.txt", "r")

    if(1 <= bits <= 32):             
        for linea in archivo1:
            #Cuando es la última línea
            if "\n" not in linea:
                linea = linea + "\n"
            #Separar datos de las líneas
            l = linea.replace("\n",";").replace("-",";").split(";")
            l.pop(-1)
            if len(l) == 2:
                b1, n1 = l
                A += 1
                B += vrf(n1,b1)
                if vrf(n1,b1) == 0: 
                    number = convert_base(n1,b1,2) 
                    if len(number) > bits:
                        C += 1
                
            elif len(l) == 4:
                b1, n1, b2, n2 = l
                A += 2
                B += vrf(n1,b1)
                B += vrf(n2,b2)
                if vrf(n1,b1) == 0: 
                    number = convert_base(n1,b1,2) 
                    if len(number) > bits:
                        C += 1
                
                if vrf(n2,b2) == 0: 
                    number = convert_base(n2,b2,2) 
                    if len(number) > bits:
                        C += 1

            #Segundo Archivo ==> resultados.txt
            archivo2 = open("resultados.txt", "w")
            l = str(A) + ";" + str(B) + ";" + str(C) + ";" + str(D) 
            archivo2.write(l)
            archivo2.close()

    elif (bits == 0):
        archivo2 = open("resultados.txt", "r")
        a = b = c = d = 0
        for linea in archivo2:
            if "\n" not in linea:
                linea = linea + "\n"
            l = linea.replace("\n",";").split(";")
            l.pop(-1)
            a,b,c,d = l
        archivo2.close()
        if int(a) >= (int(b) + int(c) + int(d)):
            flag = True

    else:
        print("Número fuera del rango")
    archivo1.close()