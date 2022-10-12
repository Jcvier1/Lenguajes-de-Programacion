%Eliminar un elemento de una lista
borrar(_, [], []).
borrar(A, [A|X], R):- 
    borrar(A, X, R), !.
borrar(X, [A|X], [A|R]):- 
    borrar(X, X, R).