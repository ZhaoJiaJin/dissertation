package solu

import(
    "sync"
)

type DXOpe struct{
    X []float64
    Res []float64
    ThreadNum int
    Rowx int
    Colx int
    Srcx int
    Srcy int
    AllNeigh [][]uint32
    wg sync.WaitGroup
}

func (ope *DXOpe)Calculate(){
    for i:=0; i < ope.ThreadNum; i ++{
        ope.wg.Add(1)
        go ope.cal(i)
    }
    ope.wg.Wait()
}

func (ope *DXOpe)cal(tid int){
    defer ope.wg.Done()
    for i:=tid; i < ope.Rowx; i+=ope.ThreadNum{
        //neighs := findNeighbors(i,ope.Srcx,ope.Srcy)
        neighs := ope.AllNeigh[i]
        len_neighs := float64(len(neighs))
        for j := 0; j < ope.Colx; j ++{
            val := 0.0
            colAtj := ope.getXCol(j)
            for _,idx := range neighs{
                val += colAtj[idx]
            }
            val -= (len_neighs * colAtj[i])
            ope.setR(i,j,val)
        }
    }
}



func (ope *DXOpe)getXCol(j int)[]float64{
    // x is N*n matrix
    return ope.X[ope.Rowx*j:ope.Rowx*(j+1)]
}

func (ope *DXOpe)getX(i,j int)float64{
    // x is N*n matrix
    return ope.X[ope.Rowx*j + i]
}

func (ope *DXOpe)setR(i,j int,v float64){
    ope.Res[ope.Rowx*j + i] = v
}
