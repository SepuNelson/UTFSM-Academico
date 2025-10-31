import math

def distancia_euclidiana(p1, p2):
    return math.sqrt((p1[0] - p2[0])**2 + (p1[1] - p2[1])**2)

def distancia_minima_dividir_y_conquistar(puntos):
    if len(puntos) <= 3:
        return distancia_minima_bruta(puntos)
    
    mid = len(puntos) // 2
    Q = puntos[:mid]
    R = puntos[mid:]
    
    distancia_minima_izquierda = distancia_minima_dividir_y_conquistar(Q)
    distancia_minima_derecha = distancia_minima_dividir_y_conquistar(R)
    
    distancia_minima = min(distancia_minima_izquierda, distancia_minima_derecha)
    
    puntos_cercanos_medio = []
    for punto in puntos:
        if abs(punto[0] - puntos[mid][0]) < distancia_minima:
            puntos_cercanos_medio.append(punto)
    
    puntos_cercanos_medio.sort(key=lambda punto: punto[1])
    
    return min(distancia_minima, distancia_minima_entre_medio(puntos_cercanos_medio, distancia_minima))

def distancia_minima_entre_medio(puntos, distancia_minima):
    for i in range(len(puntos)):
        for j in range(i + 1, len(puntos)):
            if (puntos[j][1] - puntos[i][1]) >= distancia_minima:
                break
            distancia_minima = min(distancia_minima, distancia_euclidiana(puntos[i], puntos[j]))
    return distancia_minima

def distancia_minima_bruta(puntos):
    distancia_minima = float('inf')
    for i in range(len(puntos)):
        for j in range(i + 1, len(puntos)):
            distancia_minima = min(distancia_minima, distancia_euclidiana(puntos[i], puntos[j]))
    return distancia_minima

def leer_entrada():
    with open('input-1.dat', 'r') as file:
        input_data = file.read().strip().split()
        i = 0
        while i < len(input_data):
            n = int(input_data[i])
            i += 1
            puntos = []
            for _ in range(n):
                x = int(input_data[i])
                y = int(input_data[i + 1])
                puntos.append((x, y))
                i += 2
            puntos.sort()
            distancia = distancia_minima_dividir_y_conquistar(puntos)
            truncado = math.floor(distancia * 100) / 100
            print(f"{truncado:.2f}")

if __name__ == "__main__":
    leer_entrada()
