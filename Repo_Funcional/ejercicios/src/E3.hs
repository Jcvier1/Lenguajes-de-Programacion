module E3
    ( main
    ) where
-- Funcion para aplanar listas
aplanar :: [[a]] -> [a]
aplanar [] = []
aplanar (xs:xss) =  mappend xs (aplanar xss) 

main :: IO ()
main = do
    let lista = [[1,2], [3, 4] ,[5], [6,7]]
    print(aplanar lista)
    -- [1,2,3,4,5,6,7]


