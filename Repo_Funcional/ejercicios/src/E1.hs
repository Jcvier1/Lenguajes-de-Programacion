{-# OPTIONS_GHC -Wno-incomplete-patterns #-}
module E1
    ( main
    ) where


isInclude :: Eq a => [a] -> [a] -> Bool
isInclude [] _ = True
isInclude _ [] = False
isInclude xs ys | xs == take (length xs) ys = True
                | otherwise = isInclude xs (tail ys)

contieneR :: [a] -> [[a]] -> [[a]]
contieneR _ [] = []
contieneR s (x:xs)  = isInclude s xs


main :: IO ()
main = do
    let s = "hola"
    let c = "ho"

    let d = ["hola","a"]
    let f = "a"

    print(isInclude c s)
    print(contieneR f d)