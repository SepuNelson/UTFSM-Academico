C1 = float(input("Nevarro Nummies consumidas (en unidades): "))
C2 = float(input("Space Soup consumida (en [ml]): "))
C3 = float(input("Carne de rana consumida (en [g]): "))

g = round(C1*(1.90) + C2*(10.0/285) + C3*(0.30/100),2)
c = round(C1*(6.00) + C2*(12.0/285),2)
p = round(C1*(0.80) + C2*(11.0/285) + C3*(16.0/100),2)

ct = int(round(g*9,0) + round(c*4,0) + round(p*4,0))

print("Grogu a consumido: ")
print(g,"[g] de grasas")
print(c,"[g] de carbohidratos")
print(p,"[g] de proteínas")
print("dando un total de")
print(ct,"[calorías]")

