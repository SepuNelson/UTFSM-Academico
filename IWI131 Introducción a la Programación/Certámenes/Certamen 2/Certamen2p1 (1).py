def sismos_por_pais(archivo_sismos):
    dic = {}
    c = 0
    arch = open(archivo_sismos,"r")
    for line in arch:

        l = line.replace("\n"," ").split(";")
        tiempo,latitud,longitud,profundidad,magnitud,lugar = l
        if float(magnitud) >= 2:
            archi = open("estados_eeuu.txt","r")
            for linea in archi:
                l1 = linea.replace("\n"," ")
                if lugar == l1:
                    lugar = "EEUU "
                    c += 1
            if lugar not in dic:
                dic[lugar] = 1
            elif lugar in dic:
                dic[lugar] += 1
            archi.close()
    arch.close()
    list = []
    for x in dic:
        list.append((dic[x],x))
        list.sort()
        list.reverse()

    return list