% Código de la Cerradura
cerradura(1, 4, 5, 1, 0).

% Formula de la cercanía o lejanía del código y la clave ingresada
% Obtiene el valor del código, de cada "celda"
% Calcula la diferencia absoluta entre dos números
% Calcula la diferencia absoluta entre dos números
% Calcula la diferencia absoluta entre dos números
% Calcula la diferencia absoluta entre dos números
% Calcula la diferencia absoluta entre dos números
% Calcula el número que determina si está "Cerca", "Lejos", o encontró la clave con éxito
% Da la respuesta en relación al número generado anteriormente
verificar(X1, X2, X3, X4, X5, R) :-
    cerradura(C1, C2, C3, C4, C5),
    calcular_distancia(X1, C1, D1),
    calcular_distancia(X2, C2, D2),
    calcular_distancia(X3, C3, D3),
    calcular_distancia(X4, C4, D4),
    calcular_distancia(X5, C5, D5),
    Distancia is (D1 + D2 + D3 + D4 + D5) / 5,
    evaluar_distancia(Distancia, R).


% Calcula el número que determina si está "Cerca", "Lejos", o encontró la clave con éxito
calcular_distancia(X, C, D) :-
    D is abs(X - C).


% Da la respuesta en relación al número generado anteriormente
evaluar_distancia(Distancia, "Contraseña desubierta") :- Distancia =:= 0, !.
evaluar_distancia(Distancia, "Lejos") :- Distancia > 1, !.
evaluar_distancia(_, "Cerca").