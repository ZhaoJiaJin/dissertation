package solu

import(
    "sync"
    "gonum.org/v1/gonum/mat"
    "log"
)

var pool *MPool
var debuglog = false

func init(){
    pool = &MPool{
        Data:make(map[int]chan []float64),
    }
}

type MPool struct{
    sync.Mutex
    Data map[int]chan []float64
}

func printlog(v ...interface{}){
    if debuglog{
        log.Println(v...)
    }
}

func printflog(format string, v ...interface{}){
    if debuglog{
        log.Printf(format,v...)
    }
}

func (p *MPool)PutRawData(v []float64){
    p.Lock()
    defer p.Unlock()

    l := len(v)
    if _,ok := p.Data[l]; !ok{
        p.Data[l] = make(chan []float64,10)
    }
    select{
    case p.Data[l] <- v:
        printlog("pool: put array of length:",l)
    default:
        printlog("pool: fail to put array of length:",l)
    }
}

func (p *MPool)PutVector(v *mat.VecDense){
    printlog("pool:put vector back to pool")
    p.PutRawData(v.RawVector().Data)
}


func (p *MPool)PutMatrix(v *mat.Dense){
    printlog("pool:put matrix back to pool")
    p.PutRawData(v.RawMatrix().Data)
}

func (p *MPool)GetMatrix(row,col int)*mat.Dense{
    printlog("pool:get matrix from pool")
    ret := mat.NewDense(row,col,p.GetRawData(row*col))
    ret.Reset()
    ret.ReuseAs(row,col)
    return ret
}

func (p *MPool)GetVector(l int)*mat.VecDense{
    printlog("pool:get vector from pool")
    ret := mat.NewVecDense(l,p.GetRawData(l))
    ret.Reset()
    ret.ReuseAsVec(l)
    return ret
}

func (p *MPool)GetRawData(l int)[]float64{
    p.Lock()
    defer p.Unlock()
    if _,ok := p.Data[l]; !ok{
        printflog("pool:get generate new raw array of size:%v\n",l)
        return make([]float64, l)
    }

    select{
    case rawArr := <-p.Data[l]:
        printflog("pool:get raw array of size: %v\n", l)
        return rawArr
    default:
        printflog("pool:get generate new raw array of size:%v\n", l)
        return make([]float64, l)
    }

}
