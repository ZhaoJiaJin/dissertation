package utils

import(
    "math"
    "gonum.org/v1/gonum/mat"
    "fmt"
)

// CalN calculate N via lvl
func CalN(lvl int)int{
    return int(math.Pow(4,float64(lvl)) * 12)
}

//MostSqure find the most squired width and height for v
func MostSqure(v int)(int,int){
    begin := int(math.Sqrt(float64(v)))
    for i := begin; i >= 1; i --{
        if v % i == 0{
            return i, v/i
        }
    }
    panic("should not go here")
}


//PrintMatrix print matrix
func PrintMatrix(a mat.Matrix){
    fc := mat.Formatted(a, mat.Prefix("         "), mat.Squeeze())
    fmt.Printf("matrix = %v\n", fc)
}
