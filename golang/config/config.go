package config

import(
    "gonum.org/v1/gonum/mat"
    "io/ioutil"
    "strings"
    "strconv"
)

const(
    //ROWA row size of matrix A
    ROWA = 9
    //COLA column size of matrix A
    COLA = 4
    //LENT size of square matrix T
    LENT = 9
)

// Conf contains matrix config
var Conf CfgMatrix

// CfgMatrix is the struct for config matrixs
type CfgMatrix struct{
    A []float64
    //ma *mat.Dense
    T []float64
    //mt *mat.DiagDense
    Y []float64
}

// GetMatrixA return part of matrix A
func GetMatrixA(m,n int)(mat.Matrix){
    ma := mat.NewDense(ROWA,COLA,Conf.A)
    return ma.Slice(0,m,0,n)
}
// GetMatrixT return part of matrix TG
func GetMatrixT(m int)(mat.Matrix){
    return mat.NewDiagDense(m, Conf.T[:m])
}

// InitConfig init all config matrixs,which is A,T and y
func InitConfig(afile,tfile string,y []float64){
    Conf.Y = y
    adata, err := ioutil.ReadFile(afile)
    if err != nil{
        panic(err)
    }
    Conf.A = make([]float64,ROWA*COLA)
    i := 0
    for _,l := range(strings.Split(string(adata),"\n")){
        for _,v := range(strings.Split(l,",")){
            v = strings.TrimSpace(v)
            if len(v) == 0{
                continue
            }
            vflt,err := strconv.ParseFloat(v,64)
            if err != nil{
                panic(err)
            }
            Conf.A[i] = vflt
            i++
        }
    }

    tdata, err := ioutil.ReadFile(tfile)
    if err != nil{
        panic(err)
    }

    Conf.T = make([]float64, LENT)
    i = 0
    for _,l := range(strings.Split(string(tdata),"\n")){
        for _,v := range(strings.Split(l,",")){
            v = strings.TrimSpace(v)
            if len(v) == 0{
                continue
            }
            vflt,err := strconv.ParseFloat(v,64)
            if err != nil{
                panic(err)
            }
            Conf.T[i] = vflt
            i++
        }
    }

}


