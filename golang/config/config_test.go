package config

import(
    "testing"
    "conjugate/utils"
)


func TestInitConfig(t *testing.T){
    y := []float64{1,2,3,4}
    InitConfig("a.csv","t.csv",y)
    if len(Conf.A) != 4*9 || Conf.A[0] != 1 || Conf.A[len(Conf.A)-1] != 0.01524107{
        t.Fatal("matrix A verify failed",Conf.A)
    }
    if len(Conf.T) != 9 || Conf.T[0] != 629881.6 || Conf.T[len(Conf.T)-1] != 3.08641975e+07{
        t.Fatal("matrix T verify failed",Conf.T)
    }
}

func TestSlice(t *testing.T){
    y := []float64{1,2,3,4}
    InitConfig("a.csv","t.csv",y)
    a := GetMatrixA(4,3)
    utils.PrintMatrix(a)
    b := GetMatrixT(3)
    utils.PrintMatrix(b)
}
