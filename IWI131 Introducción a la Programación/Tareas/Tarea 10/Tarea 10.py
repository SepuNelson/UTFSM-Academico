def avistamientos_por_region(nombre_archivo):
    dic = {}
    i = 0
    arch = open(nombre_archivo, "r")
    for linea in arch:
        if i != 0:
            linea.replace("-", ";").replace("\n",";").split(";")
            year,month,region,avs,ovnis,aviones,otros = linea
            if region not in dic:
                dic[region] = (year,month,avs,ovnis,aviones,otros)
            dic[region].append((year,month,avs,ovnis,aviones,otros))
            suma = round(((ovnis*100)/avs),2)
            dic[region].append(suma,avs,year,month)
            dic[region].sort()
            dic[region].reverse()
        i += 1

    for reg in dic:
        archi = open(reg+".txt", "w")
        for n in dic[reg][0:3]:
            archi.write("En el mes {0} de {1} hubo {2}% de avistamientos confirmados de un total de {3}\n".format(n[3],n[2],n[0],n[1]))
        archi.close()
    arch.close()

    return dic


nombre_archivo = 'ovnis_grande.csv'