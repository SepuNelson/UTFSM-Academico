% Cifrados de bits con ADN

cifrado([0,0], a).
cifrado([0,1], g).
cifrado([1,0], c).
cifrado([1,1], t).

% Orden del código
% Si el mensaje está vacío, la respuesta también lo será
%
% Relacionar los 2 bits con una base
% Llamada recursiva para el resto del mensaje

descifrar([], []).
descifrar([X, Y | Resto_del_Mensaje], [Base | Resto_de_las_Bases]) :-
    cifrado([X, Y], Base),
    descifrar(Resto_del_Mensaje, Resto_de_las_Bases).