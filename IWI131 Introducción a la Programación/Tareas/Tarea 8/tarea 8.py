def obtener_valor_caracteristica(caracteristicas, buscada):
    i = 0
    while i < len(caracteristicas):
        carac,canti = caracteristicas[i]
        if carac == buscada:
            return canti
        i += 1

def puntaje_amigo(amigo, caracteristicas):
    nombre,caracs = amigo
    i1 = 0
    total = 0
    while i1 < len(caracs):
        c = caracs[i1]
        #Conteo
        i = 0
        while i < len(caracteristicas):
            carac, canti = caracteristicas[i]
            if c == carac:
                total += canti
            i += 1
        i1 += 1
    return total

caracteristicas = [
('kawaii',10), ('leal',20), ('acusete',-10), ('avaro',-15), ('respetuoso',20),
('otaku',25),('lolero',25),('furro',-50),('vtuver',25),('mechero',-30)
]

amigos = [('Sneki',('leal','kawaii','vtuver')),
          ('Otaku-taku',('otaku','avaro','lolero','leal')),
          ('Maiga',('paciente','otaku','leal')),
          ('Mojojojo',('mechero','kawaii','Furro','lolero')),
          ('Seiya',('leal','acusete')),
          ('Vegeta',('otaku','avaro')),
          ('Kalila',('lolero','kawaii')),
          ('Grogu',('avaro','kawaii','lolero','otaku')),
          ('Freezer',('acusete','furro','otaku','lolero'))
]

n = 0
puntos_primero = 0
puntos_segundo = 0
primero = ""
segundo = ""
while n < len(amigos):
    nom,cts = amigos[n]
    n1 = 0
    while n1 < len(cts):
        if 'lolero' == cts[n1]:
            amigo = amigos[n]
            puntos = puntaje_amigo(amigo, caracteristicas)
            if puntos > puntos_primero:
                puntos_primero = puntos
                primero = nom + ","
            if puntos_segundo < puntos < puntos_primero:
                puntos_segundo = puntos
                segundo = nom + ","
        n1 += 1
    n += 1


#Pregunta 1
print(obtener_valor_caracteristica(caracteristicas, "vtuver"))
print(" ")

#Pregunta 2
print(puntaje_amigo(('Vegeta',('otaku','avaro')), caracteristicas))
print(" ")

#Pregunta 3
print("Equipo seleccionado: ")
print(primero, puntos_primero, "puntos")
print(segundo, puntos_segundo, "puntos")