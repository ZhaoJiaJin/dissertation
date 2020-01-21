package main

import(
    "flag"
    "os"
    "time"
    "conjugate/config"
    "conjugate/utils"
    "conjugate/solu"
    "math/rand"
    "gonum.org/v1/gonum/mat"
    "log"
    "runtime/pprof"
    "runtime"
)

var(
    afile string
    tfile string
    lvl int
    n int
    m int
    threadNum int
    method string
    cpuprofile string
    memprofile string
)

func main(){
    flag.StringVar(&afile, "afile", "config/a.csv", "matrix A config file")
    flag.StringVar(&tfile, "tfile", "config/t.csv", "matrix T config file")
    flag.IntVar(&lvl, "lvl", 1, "level number")
    flag.IntVar(&m,"m",9,"value of m")
    flag.IntVar(&n,"n",4,"value of n")
    flag.IntVar(&threadNum,"th",10,"the number of thread")
    flag.StringVar(&method,"method","ite","use which method, ite|std|both")
    flag.StringVar(&cpuprofile,"cpu", "", "write cpu profile to `file`")
    flag.StringVar(&memprofile,"mem", "", "write mem profile to `file`")
    flag.Parse()

    rand.Seed(time.Now().Unix())
    bigN := utils.CalN(lvl)
    sizey := m * bigN
    log.Printf("-----------lvl:%v, N:%v------------\n",lvl,bigN)
    y := make([]float64,sizey)
    for i := range y{
        y[i] = float64(rand.Intn(60))
    }
    config.InitConfig(afile,tfile,y)
    //fmt.Println(config.Conf.Y)

    a := config.GetMatrixA(m,n)
    t := config.GetMatrixT(m)
    var resStd *mat.VecDense
    var resIte *mat.VecDense
    if cpuprofile != "" {
        f, err := os.Create(cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }
    if memprofile != "" {
        defer func(){
            f, err := os.Create(memprofile)
            if err != nil {
                log.Fatal("could not create memory profile: ", err)
            }
            defer f.Close()
            runtime.GC() // get up-to-date statistics
            //defer runtime.GC() // get up-to-date statistics
            if err := pprof.Lookup("heap").WriteTo(f,0); err != nil {
                log.Fatal("could not write memory profile: ", err)
            }
        }()
    }



    if method == "ite" || method == "both"{
        begin := time.Now().Unix()
        sl := solu.NewIteSolu(a,t,m,n,bigN,lvl,y,threadNum)
        resIte = sl.FindSolution()
        end := time.Now().Unix()
        log.Println("iterate method time cost:",end - begin)
        sl.CalRealR(resIte)
    }
    if method == "std" || method == "both"{
        begin := time.Now().Unix()
        sl := solu.NewStdSolu(a,t,m,n,bigN,lvl,y)
        resStd = sl.FindSolution()
        end := time.Now().Unix()
        log.Println("standard method time cost:",end - begin)
    }
    if method == "both"{
        distance := new(mat.VecDense)
        distance.SubVec(resIte,resStd)
        log.Printf("||x-x0|| = %v\n",mat.Dot(distance,distance))
    }
}
