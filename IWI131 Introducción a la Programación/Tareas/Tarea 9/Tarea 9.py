day = int(input("Día: "))
month = int(input("Mes: "))
year = int(input("Año: "))
x = "s"
while x == "s":
    cond = 0
    rut = input("Rut: ")
    fecha = year, month, day
    # Ya existe rut en dic dosis
    if rut in dosis:
        for t in vacunas:
            for w in vacunas[t]:
                if w == rut:
                    cond = 1
                    print("Segunda dosis. Paciente debe ser inoculado con: ",t)

        for i in dosis:
            if i == rut:
                dosis[i].append(fecha)

    #No existe rut en dic dosis
    if rut not in dosis:
        years = int(input("Edad: "))
        l_edad_fecha = [years,fecha]
        dosis[rut] = l_edad_fecha

    if cond == 0:
        vacuna = input("Tipo vacuna: ")
        #Ya existe la vacuna en dic vacunas
        if vacuna in vacunas:
            for i in vacunas:
                if i == vacuna:
                    vacunas[i].append(rut)

        # No existe la vacuna en dic vacunas
        if vacuna not in vacunas:
            vacunas[vacuna] = rut
    x = input("¿Desea continuar? (s / n): ")

edades_cantidad = {}
for i in dosis:
    if len(dosis[i]) == 3:
        edad_f = dosis[i][0]

        if edad_f not in edades_cantidad:
            edades_cantidad[edad_f] = 1

        elif edad_f in edades_cantidad:
            for n in edades_cantidad:
                if n == edad_f:
                    edades_cantidad[n] = edades_cantidad[n] + 1
print(" ")
print("Edades con más personas con esquema de inoculación completo: ")
primero = 0
segundo = 0
for a in edades_cantidad:
    if edades_cantidad[a] > primero:
        primero = edades_cantidad[a]
    if primero > edades_cantidad[a] > segundo:
        segundo = edades_cantidad[a]
for q in edades_cantidad:
    if edades_cantidad[q] == primero:
        print(q,"Años: ",edades_cantidad[q],"personas")
for b in edades_cantidad:
    if primero != segundo:
        if edades_cantidad[b] == segundo:
            print(b,"Años: ",edades_cantidad[b],"personas")
