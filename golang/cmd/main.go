package main

import(
    "flag"
    "time"
    "conjugate/config"
    "conjugate/utils"
    "conjugate/solu"
    "strings"
    //"math/rand"
    "gonum.org/v1/gonum/mat"
    "log"
)

var(
    afile string
    tfile string
    lvl int
    n int
    m int
    threadNum int
    method string
    yfile string
)

func main(){
    flag.StringVar(&afile, "afile", "config/a.csv", "matrix A config file")
    flag.StringVar(&tfile, "tfile", "config/t.csv", "matrix T config file")
    flag.IntVar(&lvl, "lvl", 1, "level number")
    flag.IntVar(&m,"m",9,"value of m")
    flag.IntVar(&n,"n",4,"value of n")
    flag.IntVar(&threadNum,"th",10,"the number of thread")
    flag.StringVar(&method,"method","syl","use which method,  you can choose ite,std,and syl, and iteWoCorr")
    flag.StringVar(&yfile,"y","config/vectory","the vector y")
    flag.Parse()

    bigN := utils.CalN(lvl)
    sizey := m * bigN
    log.Printf("-----------lvl:%v, N:%v, m:%v, n:%v------------\n",lvl,bigN,m,n)
    log.Println("begin to load vector y")
    fully,err := utils.LoadY(yfile)
    log.Println("load vector y complete")
    if err != nil{
        panic(err)
    }
    y := fully[:sizey]
    //y := make([]float64,sizey)
    /*for i := range y{
        //y[i] = float64(rand.Intn(50))
        y[i] = float64(i)
    }*/
    config.InitConfig(afile,tfile,y)
    //log.Println(config.Conf.Y)

    a := config.GetMatrixA(m,n)
    t := config.GetMatrixT(m)
    allres := make(map[string]mat.Vector)
    methods := strings.Split(method,",")

    //stdsl := solu.NewStdSolu(a,t,m,n,bigN,lvl,y)
    //stdsl.FindSolution()
    //stdsl.Validate(tmpres)

    //log.Printf("A:\n%1.3f\n\n", mat.Formatted(a))
    //log.Printf("T:\n%1.3f\n\n", mat.Formatted(t))
    basesl := solu.NewIteSolu(a,t,m,n,bigN,lvl,y,threadNum)

    for _,mtd := range methods{
        if mtd == "ite" {
            log.Println("begin iteration method")
            var res *mat.VecDense
            begin := time.Now().UnixNano()
            sl := solu.NewIteSolu(a,t,m,n,bigN,lvl,y,threadNum)
            res = sl.FindSolution(true)
            end := time.Now().UnixNano()
            log.Println("iterate method time cost:",end - begin)
            allres[mtd] = res
            //log.Printf("res:\n%1.3f\n\n", mat.Formatted(res.T()))
            //stdsl.Validate(res)
	    residual := basesl.CalR(res)
	    log.Println("residual for iteration solution:",residual)
        }
        if mtd == "iteWoCorr" {
            log.Println("begin iteration method")
            var res *mat.VecDense
            begin := time.Now().UnixNano()
            sl := solu.NewIteSolu(a,t,m,n,bigN,lvl,y,threadNum)
            res = sl.FindSolution(false)
            end := time.Now().UnixNano()
            log.Println("iterate method without residual correction time cost:",end - begin)
            allres[mtd] = res
            //log.Printf("res:\n%1.3f\n\n", mat.Formatted(res.T()))
            //stdsl.Validate(res)
	    residual := basesl.CalR(res)
	    log.Println("residual for iteration solution without residual correction:",residual)
        }

        if mtd == "std" {
	    log.Println("begin standard method")
            var res *mat.VecDense
            begin := time.Now().UnixNano()
            sl := solu.NewStdSolu(a,t,m,n,bigN,lvl,y)
            res = sl.FindSolution()
            end := time.Now().UnixNano()
            log.Println("standard method time cost:",end - begin)
            allres[mtd] = res
            //log.Printf("res:\n%1.3f\n\n", mat.Formatted(res.T()))
            //stdsl.Validate(res)
	    residual := basesl.CalR(res)
	    log.Println("residual for standard solution:",residual)
        }
        if mtd == "syl"{
	    log.Println("begin sylvester method")
            var res *mat.VecDense
            begin := time.Now().UnixNano()
            res = solu.NewSylSolu(a,t,m,n,bigN,lvl,y,threadNum,1e-9)
            //res = sl.FindSolution()
            end := time.Now().UnixNano()
            log.Println("syl method time cost:",end - begin)
            allres[mtd] = res
            //log.Printf("res:\n%1.3f\n\n", mat.Formatted(res.T()))
            //stdsl.Validate(res)
	    residual := basesl.CalR(res)
	    log.Println("residual for sylvester solution:",residual)
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
            log.Printf("Distance between %s and %s is %v\n",k1,k2,mat.Dot(distance,distance))
        }
    }
}
