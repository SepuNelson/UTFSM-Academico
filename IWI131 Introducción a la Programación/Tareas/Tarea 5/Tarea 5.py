def ma(n1, n2, n3):
    ma = (n1+n2+n3)/3
    ma = round(ma,0)
    return ma

def mg(n1, n2, n3):
    mg = (n1*n2*n3)**(1/3)
    mg = round(mg, 0)
    return mg

def mv(n1, n2, n3):
    mv = (n3*(ma(n1, n2, n3))**2)**(1/3)
    mv = round(mv, 0)
    return mv

def x(ma, mg, mv):
    if ma > 55:
        return 1
    elif mg > 55:
        return 2
    elif mv > 55:
        return 3
    elif ma < 55 and mg < 55 and mv < 55:
        return 0

pedir_notas = True
while pedir_notas:
    ramo = input("Ingrese el nombre del ramo: ")

    if ramo == "salir":
        pedir_notas = False
        print("Fin del programa - Desarrollado por Kiwi :D!")
    else:
        n1 = int(input("Ingrese la 1era nota: "))
        n2 = int(input("Ingrese la 2era nota: "))
        n3 = int(input("Ingrese la 3era nota: "))

        nf_aritmetica = int(ma(n1, n2, n3))
        nf_geometrica = int(mg(n1, n2, n3))
        nf_vuelta = int(mv(n1, n2, n3))

        print("Su nota final según la Media Aritmética es:", nf_aritmetica)
        print("Su nota final según la Media Geométrica es:", nf_geometrica)
        print("Su nota final según la Media Vuelta es:", nf_vuelta)

        nf_aprobacion = x(nf_aritmetica, nf_vuelta, nf_geometrica)

        if nf_aprobacion == 0:
            print("Lamentablemente no puedes aprobar con ninguna de las fórmulas :'c")
        elif nf_aprobacion == 1:
            print("Si la NF del ramo se calcula usando la Media Aritmética, entonces apruebas", ramo, ":D")
        elif nf_aprobacion == 2:
            print("Si la NF del ramo se calcula usando la Media Geométrica, entonces apruebas", ramo, ":D")
        elif nf_aprobacion == 3:
            print("Si la NF del ramo se calcula usando la Media Vuelta, entonces apruebas", ramo, ":D")
