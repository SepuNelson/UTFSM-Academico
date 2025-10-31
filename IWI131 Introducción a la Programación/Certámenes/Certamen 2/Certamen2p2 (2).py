def categorizar_sismos(archivo_sismos):
    dic = {}
    c = 0
    arch = open(archivo_sismos, "r")
    for line in arch:
        ll = line.replace("\n", " ").split(";")
        tiempo, latitud, longitud, profundidad, magnitud, lugar = ll
        fecha = tiempo[0:10]
        hora = tiempo[11:16]
        if 2 <= float(magnitud) < 3:
            str_menor = "2"
            if str_menor not in dic:
                dic[str_menor] = [[magnitud, fecha, hora, lugar]]
            elif str_menor in dic:
                dic[str_menor].append([magnitud, fecha, hora, lugar])
            dic[str_menor].sort()
            dic[str_menor].reverse()

        elif 3 <= float(magnitud) < 4:
            str_menor = "3"
            if str_menor not in dic:
                dic[str_menor] = [[magnitud, fecha, hora, lugar]]
            elif str_menor in dic:
                dic[str_menor].append([magnitud, fecha, hora, lugar])
            dic[str_menor].sort()
            dic[str_menor].reverse()

        elif 4 <= float(magnitud) < 5:
            str_menor = "4"
            if str_menor not in dic:
                dic[str_menor] = [[magnitud, fecha, hora, lugar]]
            elif str_menor in dic:
                dic[str_menor].append([magnitud, fecha, hora, lugar])
            dic[str_menor].sort()
            dic[str_menor].reverse()

        elif 5 <= float(magnitud) < 6:
            str_menor = "5"
            if str_menor not in dic:
                dic[str_menor] = [[magnitud, fecha, hora, lugar]]
            elif str_menor in dic:
                dic[str_menor].append([magnitud, fecha, hora, lugar])
            dic[str_menor].sort()
            dic[str_menor].reverse()

        elif 6 <= float(magnitud) < 7:
            str_menor = "6"
            if str_menor not in dic:
                dic[str_menor] = [[magnitud, fecha, hora, lugar]]
            elif str_menor in dic:
                dic[str_menor].append([magnitud, fecha, hora, lugar])
            dic[str_menor].sort()
            dic[str_menor].reverse()

    for x in dic:
        c += 1
        mag = str(x)
        archi = open("mag" + mag + ".txt", "w")
        q = dic[x]
        s = "Fecha: {0}; Hora: {1}; Lugar: {2}; Magnitud: {3}.\n"
        archi.write(s.format(q[0][1], q[0][2], q[0][3], q[0][0]))
        archi.write(s.format(q[1][1], q[1][2], q[1][3], q[1][0]))
        archi.write(s.format(q[2][1], q[2][2], q[2][3], q[2][0]))

        archi.close()
    arch.close()


    return c
