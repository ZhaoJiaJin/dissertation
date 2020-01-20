package solu

import(
    "testing"
)

func TestPool(t *testing.T){
    a := []float64{1,2}
    pool.PutRawData(a)
    res := pool.GetRawData(2)
    if res[0] != 1 || res[1] != 2{
        t.Fatal("get fail")
    }
    pool.PutRawData(res)
    res = pool.GetRawData(2)
    res = pool.GetRawData(2)
    if res[0] != 0 || res[1] != 0{
        t.Fatal("get fail")
    }
    pool.PutRawData(res)
}
