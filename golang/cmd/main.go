package main

import(
    "flag"
    "time"
    "conjugate/config"
    "conjugate/utils"
    "conjugate/solu"
    "strings"
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
    flag.IntVar(&m,"m",3,"value of m")
    flag.IntVar(&n,"n",2,"value of n")
    flag.IntVar(&threadNum,"th",10,"the number of thread")
    flag.StringVar(&method,"method","syl","use which method,  you can choose ite,std,and syl")
    flag.Parse()

    bigN := utils.CalN(lvl)
    sizey := m * bigN
    fmt.Printf("-----------lvl:%v, N:%v------------\n",lvl,bigN)
    y := make([]float64,sizey)
    for i := range y{
        y[i] = float64(rand.Intn(60))
    }
    config.InitConfig(afile,tfile,y)
    //fmt.Println(config.Conf.Y)

    a := config.GetMatrixA(m,n)
    t := config.GetMatrixT(m)
    allres := make(map[string]mat.Vector)
    methods := strings.Split(method,",")
    for _,mtd := range methods{
        if method == "ite" {
            var res mat.Vector
            begin := time.Now().Unix()
            sl := solu.NewIteSolu(a,t,m,n,bigN,lvl,y,threadNum)
            res = sl.FindSolution()
            end := time.Now().Unix()
            fmt.Println("iterate method time cost:",end - begin)
            allres[mtd] = res
        }
        if method == "std" {
            var res mat.Vector
            begin := time.Now().Unix()
            sl := solu.NewStdSolu(a,t,m,n,bigN,lvl,y)
            res = sl.FindSolution()
            end := time.Now().Unix()
            fmt.Println("standard method time cost:",end - begin)
            allres[mtd] = res
        }
        if method == "syl"{
            var res mat.Vector
            begin := time.Now().Unix()
            res = solu.NewSylSolu(a,t,m,n,bigN,lvl,y,threadNum,1e-5)
            //res = sl.FindSolution()
            end := time.Now().Unix()
            fmt.Println("syl method time cost:",end - begin)
            allres[mtd] = res

        }
    }
    for idx1 := 0; idx1 < len(methods); idx1 ++{
        for idx2 := idx1+1; idx2 < len(methods); idx2 ++{
            k1 := methods[idx1]
            k2 := methods[idx2]
            tmpres1 := allres[k1]
            tmpres2 := allres[k2]
            distance := new(mat.VecDense)
            distance.SubVec(tmpres1,tmpres2)
            fmt.Printf("Distance between %s and %s is %v\n",k1,k2,mat.Dot(distance,distance))
        }
    }
}
