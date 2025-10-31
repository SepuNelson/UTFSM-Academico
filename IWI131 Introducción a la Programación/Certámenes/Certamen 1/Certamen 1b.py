def votos_partido(votos, partido):
    c = 0
    i = 0
    f = 0
    votos = votos + "$"
    while f < len (votos):
        if votos[f] == "$":
            v = votos[i:f]
            i = f + 1
            if partido == v:
                c += 1
        f += 1
    return c

coaliciones = input("Ingrese coaliciones: ")
votos = input("Ingrese votos por partido: ")
i1 = 0
f1 = 0
contf = 0
conti = 0
suma = 0
suma_mayor = 0
coaliciones += ";"
while contf < len(coaliciones):
    if coaliciones[contf] == ";":
        coalicionesc = coaliciones[conti:contf] + "-"
        conti = contf + 1
        if suma > suma_mayor:
            suma_mayor = suma
            c_ganadora = coalicion
        i1 = 0
        f1 = 0
        suma = 0
        while f1 < len(coalicionesc):
            if coalicionesc[f1] == ":":
                coalicion = coalicionesc[i1:f1]
                i1 = f1 + 1
                print("Coalicion: ", coalicion)
            elif coalicionesc[f1] == "-":
                partido = coalicionesc[i1:f1]
                i1 = f1 + 1
                votos_x_partido = votos_partido(votos, partido)
                suma += votos_x_partido
                print(partido, votos_x_partido)
            f1 += 1
        print("Total coalicion",coalicion, ":", suma)
    contf += 1
print("La coalición ganadora es", c_ganadora, "con", suma_mayor, "votos")
