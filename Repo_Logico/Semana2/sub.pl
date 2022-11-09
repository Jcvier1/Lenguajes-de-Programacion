% SUB CONJUNTO
Mem(N, [N | _]).
mem(N, [_ | T]) :- mem(N, T).
sub([], []) :- !.
sub([], [_|_]) :- !.
sub([_|_], []) :- false.
sub([H | T], L) :- mem(H, L), sub(T, L), !.
sub([H | _], L) :- not(mem(H, L)), false.