
#Pido variables

s1 = input("Ingrese nombre de la sucursal : ")
cx1 = int(input("Coordenada x : "))
cy1 = int(input("Coordenada y : "))

s2 = input("Ingrese nombre de la sucursal : ")
cx2 = int(input("Coordenada x : "))
cy2 = int(input("Coordenada y : "))

s3 = input("Ingrese nombre de la sucursal : ")
cx3 = int(input("Coordenada x : "))
cy3 = int(input("Coordenada y : "))

cs1 = 0
cs2 = 0
cs3 = 0

a = "si"
sumaf = 0

while a == "si":
    suma = 0
    print(" ")
    n = int(input("Ingrese número del plato :"))
    if n == 1:
        suma = 4000
    elif n == 2:
        suma = 3000
    elif n == 3:
        suma = 3500

    if n == 1 or n == 2 or n == 3 or n == -1:
        while n != -1 and (n == 1 or n == 2 or n == 3):
            n = int(input("Ingrese número del plato : "))
            if n == 1:
                suma += 4000
            elif n == 2:
                suma += 3000
            elif n == 3:
                suma += 3500
        if n != 1 and n != 2 and n != 3 and n != -1:
            print("Error en el valor ingresado")
    else:
        print("Error en el valor ingresado")

    sumaf += suma

    print("Total del pedido $", suma)

    x = int(input("Ingrese coordenada x cliente : "))
    y = int(input("Ingrese coordenada y cliente : "))

    from math import sqrt

    d1 = sqrt( (x - cx1 )**2 + (y - cy1 )**2 )
    d2 = sqrt( (x - cx2 )**2 + (y - cy2 )**2 )
    d3 = sqrt( (x - cx3 )**2 + (y - cy3 )**2 )

    if d1 < d2 and d1 < d3:
        cs1 += 1
        print("Pedido será entregado por ",s1)

    elif d2 < d1 and d2 < d3:
        cs2 += 1
        print("Pedido será entregado por ",s2)

    elif d3 < d1 and d3 < d2:
        cs3 += 1
        print("Pedido será entregado por ",s3)

    a = input("¿Desea registrar otro pedido? : ")
    a = a.lower()

    print(" ")

print("##### Estadísticas Finales ####")
print("Monto total recaudado $", sumaf)
print(s1, "entregó", cs1 , "pedidos")
print(s2, "entregó", cs2 , "pedidos")
print(s3, "entregó", cs3 , "pedidos")