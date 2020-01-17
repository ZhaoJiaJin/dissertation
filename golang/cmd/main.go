package main

import(
    "flag"
    "conjugate/config"
    "conjugate/utils"
    "conjugate/solu"
    "math/rand"
)

var(
    afile string
    tfile string
    lvl int
    n int
    m int
    threadNum int
    std bool
)

func main(){
    flag.StringVar(&afile, "afile", "config/a.csv", "matrix A config file")
    flag.StringVar(&tfile, "tfile", "config/t.csv", "matrix T config file")
    flag.IntVar(&lvl, "lvl", 1, "level number")
    flag.IntVar(&m,"m",9,"value of m")
    flag.IntVar(&n,"n",4,"value of n")
    flag.IntVar(&threadNum,"th",10,"the number of thread")
    flag.BoolVar(&std,"std",false,"use standard method")
    flag.Parse()

    sizey := m * utils.CalN(lvl)
    y := make([]float64,sizey)
    for i := range y{
        y[i] = float64(rand.Intn(60))
    }
    config.InitConfig(afile,tfile,y)
    //fmt.Println(config.Conf.Y)

    a := config.GetMatrixA(m,n)
    t := config.GetMatrixT(m)
    sl := solu.NewIteSolu(a,t,m,n,utils.CalN(lvl),y,threadNum)
    sl.FindSolution()
}
