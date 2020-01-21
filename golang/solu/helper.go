package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
    "sync"
    "log"
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

func findAllNeigh(m,n,N,threadNum int)[][]uint32{
    log.Println("find all neighbours")
    res := make([][]uint32,N)
    var wg sync.WaitGroup
    for i:=0; i < threadNum; i ++{
        wg.Add(1)
        go func(tid int){
            defer wg.Done()
            for idx := tid; idx < N; idx += threadNum{
                res[idx] = findNeighbors(idx,m,n)
            }
        }(i)
    }
    wg.Wait()
    log.Println("find all neighbours done")
    return res
}

func findNeighbors(i,m,n int)[]uint32{
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
    r := make([]uint32,len(res))
    idx := 0
    for _,v := range res{
        r[idx] = uint32(v)
        idx ++
    }
    return r
}


func generateD(lvl int)*mat.Dense{
    size := utils.CalN(lvl)
    m,n := utils.MostSqure(size)
    res := make([]float64,size * size)

    for i:=0; i < size; i ++{
        allneig := findNeighbors(i,m,n)
        for _,pos := range allneig{
            res[i*size + int(pos)] = 1
        }
        res[i*size + i] = float64(-1*len(allneig))
    }

    return mat.NewDense(size,size, res)
}

func vectorsub(a,b,c mat.Vector)*mat.VecDense{
    //res := new(mat.VecDense)
    res := pool.GetVector(a.Len())
    res.SubVec(a,b)
    res.SubVec(res,c)
    return res
}
