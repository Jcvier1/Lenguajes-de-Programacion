{-# OPTIONS_GHC -Wno-incomplete-patterns #-}

module E2
    ( main
    ) where

subconjunto :: Eq a => [a] -> [a] -> Bool
subconjunto [] _ = True 
subconjunto (x:xs) ys = elem x ys && subconjunto xs ys

main :: IO ()
main  = do
    let lista = [1, 2, 3, 6]
    let lista2 = [1, 6]

    print(subconjunto lista2 lista)
    --True