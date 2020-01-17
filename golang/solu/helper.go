package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
)

const(
    L = "left"
    R = "right"
    U = "up"
    D = "down"
)

// GenerateI generate identify matrix with size n
func GenerateI(n int)mat.Matrix{
    v := make([]float64, n)
    for i := range v{
        v[i] = 1
    }

    return mat.NewDiagDense(n,v)
}

func findNeighbors(i,m,n int)map[string]int{
    res := map[string]int{
            L:i-m,
            R:i+m,
            U:i-1,
            D:i+1,
        }

    // remove all the invalid neighbours
    if i % m == 0{//the first row
        delete(res,U)
    }
    if (i % n) == (n - 1){//the last row
        delete(res, D)
    }
    if i < m{ //the first column
        delete(res, L)
    }
    if i > n*m-1-m{// the last column
        delete(res, R)
    }
    return res
}


func generateD(lvl int)*mat.Dense{
    size := utils.CalN(lvl)
    m,n := utils.MostSqure(size)
    res := make([]float64,size * size)

    for i:=0; i < size; i ++{
        allneig := findNeighbors(i,m,n)
        for _,pos := range allneig{
            res[i*size + pos] = 1
        }
        res[i*size + i] = float64(-1*len(allneig))
    }

    return mat.NewDense(size,size, res)
}
