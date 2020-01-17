package main

import(
    "flag"
    "time"
    "conjugate/config"
    "conjugate/utils"
    "conjugate/solu"
    "math/rand"
    "gonum.org/v1/gonum/mat"
    "fmt"
)

var(
    afile string
    tfile string
    lvl int
    n int
    m int
    threadNum int
    method string
)

func main(){
    flag.StringVar(&afile, "afile", "config/a.csv", "matrix A config file")
    flag.StringVar(&tfile, "tfile", "config/t.csv", "matrix T config file")
    flag.IntVar(&lvl, "lvl", 1, "level number")
    flag.IntVar(&m,"m",9,"value of m")
    flag.IntVar(&n,"n",4,"value of n")
    flag.IntVar(&threadNum,"th",10,"the number of thread")
    flag.StringVar(&method,"method","ite","use which method, ite|std|both")
    flag.Parse()

    bigN := utils.CalN(lvl)
    sizey := m * bigN
    fmt.Printf("-----------lvl:%v, N:%v------------",lvl,sizey)
    y := make([]float64,sizey)
    for i := range y{
        y[i] = float64(rand.Intn(60))
    }
    config.InitConfig(afile,tfile,y)
    //fmt.Println(config.Conf.Y)

    a := config.GetMatrixA(m,n)
    t := config.GetMatrixT(m)
    var resStd mat.Vector
    var resIte mat.Vector
    if method == "std" || method == "both"{
        begin := time.Now().Unix()
        sl := solu.NewStdSolu(a,t,m,n,bigN,lvl,y)
        resStd = sl.FindSolution()
        end := time.Now().Unix()
        fmt.Println("standard method time cost:",end - begin)
    }
    if method == "ite" || method == "both"{
        begin := time.Now().Unix()
        sl := solu.NewIteSolu(a,t,m,n,bigN,lvl,y,threadNum)
        resIte = sl.FindSolution()
        end := time.Now().Unix()
        fmt.Println("iterate method time cost:",end - begin)
    }
    if method == "both"{
        distance := new(mat.VecDense)
        distance.SubVec(resIte,resStd)
        fmt.Printf("||x-x0|| = %v\n",mat.Dot(distance,distance))
    }
}
