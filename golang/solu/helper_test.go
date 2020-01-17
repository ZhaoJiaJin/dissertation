package solu

import(
    "testing"
    "gonum.org/v1/gonum/mat"
    "github.com/stretchr/testify/assert"
    "conjugate/utils"
)

func TestFindNeigh(t *testing.T){
    assert.Equal(t, map[string]int{D:1,R:3},
        findNeighbors(0, 3, 3), "findNeighbors is wrong.")
    assert.Equal(t, map[string]int{L:2,U:4,R:8},
        findNeighbors(5, 3, 3), "findNeighbors is wrong.")
}

func TestIdentifyMatrix(t *testing.T){
    v := GenerateI(3)
    utils.PrintMatrix(v)
    v = GenerateI(5)
    utils.PrintMatrix(v)
}


func Testvectorsub(t *testing.T){
    n := 5
    avl := make([]float64, n)
    for i := range avl{
        avl[i] = randv()
    }
    ma := mat.NewVecDense(n,avl)

    bvl := make([]float64, n)
    for i := range bvl{
        bvl[i] = randv()
    }
    mb := mat.NewVecDense(n,bvl)

    cvl := make([]float64, n)
    for i := range cvl{
        cvl[i] = randv()
    }
    mc := mat.NewVecDense(n,cvl)

    res := vectorsub(ma,mb,mc)
    resx,resy := res.Dims()
    if resx != n || resy != 1{
        t.Fatal("vectorsub failed")
    }
    for i:=0; i < n; i++{
        if res.At(i,0) != avl[i] - bvl[i] - cvl[i]{
            t.Fatal("vectorsub failed")
        }
    }

}
