uf = int(input("Valor propiedad en UF (1500 - 13000) : "))
if 1500 <= uf <= 13000 :
    pie = int(input("Ingrese % de PIE (20% - 45%) : "))
    if 20 <= pie <= 45 :
        plazo = int(input("Ingrese plazo (20 - 25 - 30) : "))
        if plazo == 20 or plazo == 25 or plazo == 30:
            tipo = int(input("Tipo vivienda Casa(1) o Departamento(2) : "))
            if tipo == 1 or tipo == 2:
                estado = int(input("Estado Vivienda Nueva(1) o Usada(2) : "))
                if estado == 1 or estado == 2:

                    vp = uf-(uf*(pie/100))

                    if (tipo == 1 and estado == 1) :
                        if plazo == 20:
                            i = vp*0.25
                            pi = vp+i
                            s = (0.5+0.8)*12*plazo
                            pf = pi+s
                            dm = pf/(12*plazo)

                        elif plazo == 25:
                            i = vp*0.3
                            pi = vp+i
                            s = (0.5+0.8)*12*plazo
                            pf = pi+s
                            dm = pf/(12*plazo)

                        elif plazo == 30:
                            i = vp * 0.35
                            pi = vp + i
                            s = (0.5 + 0.8) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                    elif (tipo == 2 and estado == 1 ):
                        if plazo == 20:
                            i = vp * 0.28
                            pi = vp + i
                            s = (0.5 + 0.8 + 0.3) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                        elif plazo == 25:
                            i = vp * 0.33
                            pi = vp + i
                            s = (0.5 + 0.8 + 0.3) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                        elif plazo == 30:
                            i = vp * 0.41
                            pi = vp + i
                            s = (0.5 + 0.8 + 0.3) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                    elif (tipo == 1 and estado == 2 ):
                        if plazo == 20:
                            i = vp * 0.22
                            pi = vp + i
                            s = (0.5) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                        elif plazo == 25:
                            i = vp * 0.27
                            pi = vp + i
                            s = (0.5) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                        elif plazo == 30:
                            i = vp * 0.31
                            pi = vp + i
                            s = (0.5) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                    elif (tipo == 2 and estado == 2 ):
                        if plazo == 20:
                            i = vp * 0.26
                            pi = vp + i
                            s = (0.5 + 0.3) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)

                        elif plazo == 25:
                            i = vp * 0.32
                            pi = vp + i
                            s = (0.5 + 0.3) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)
                        elif plazo == 30:
                            i = vp * 0.37
                            pi = vp + i
                            s = (0.5 + 0.3) * 12 * plazo
                            pf = pi + s
                            dm = pf / (12 * plazo)
                    print("Total del credito a pagar", pf, "UF")
                    print("Dividendo mensual de", round(dm, 2), "UF")
                else:
                    print("Estado de vivienda invalido")
            else:
                print("Tipo de vivienda invalido")
        else:
            print("Plazo invalido")
    else:
        print("Error: Pie fuera de rango")
else:
    print("Error: Valor fuera de rango")

