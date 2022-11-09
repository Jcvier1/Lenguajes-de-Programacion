compare([],_).
compare([SB|C1],[P|C2]) :-
    SB = P,
    compare(C1,C2).
SubS(_,[],[]).
SubS(SB,[L|C1],[L|F]) :- 

     cod(SB,S1), 
     cod(L,S2),
     compare(S1,S2), 
     SubS(SB,C1,F), !. 

SubS(SB,[_|C1],F) :- 
     SubS(SB,C1,F), !.