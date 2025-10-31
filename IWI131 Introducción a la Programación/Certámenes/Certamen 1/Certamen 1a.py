def votos_partido(votos, partido): #    pollo$carne$fideos$pollo$arroz
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

votos = input("Votos: ")
