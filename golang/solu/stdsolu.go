package solu

import(
    "gonum.org/v1/gonum/mat"
    "fmt"
)

type StdSolu struct{
    Q mat.Matrix
    B mat.Matrix
    C mat.Matrix
    y mat.Vector
}

func NewStdSolu(a,t mat.Matrix,m,n,bign,lvl int,y []float64)*StdSolu{
    res := new(StdSolu)
    d := generateD(lvl)

    d2 := new(mat.Dense)
    d2.Product(d,d)
    q := new(mat.Dense)
    q.Kronecker(GenerateI(n), d2)
    res.Q = q

    IMat := GenerateI(bign)
    B := new(mat.Dense)
    B.Kronecker(a,IMat)
    res.B = B
    C := new(mat.Dense)
    C.Kronecker(t,IMat)
    res.C = C

    yvec := mat.NewVecDense(len(y),y)
    res.y = yvec
    return res
}

func (sl *StdSolu)FindSolution()(*mat.VecDense){
    BtC := new(mat.Dense)
    BtC.Product(sl.B.T(), sl.C)

    BtCB := new(mat.Dense)
    BtCB.Product(BtC, sl.B)

    A := new(mat.Dense)
    A.Add(sl.Q, BtCB)

    b := new(mat.VecDense)
    b.MulVec(BtC, sl.y)

    res := new(mat.VecDense)
    res.SolveVec(A,b)
    return res
}


func (sl *StdSolu) Validate(x mat.Vector){
    BtC := new(mat.Dense)
    BtC.Product(sl.B.T(), sl.C)

    BtCB := new(mat.Dense)
    BtCB.Product(BtC, sl.B)

    A := new(mat.Dense)
    A.Add(sl.Q, BtCB)

    b := new(mat.VecDense)
    b.MulVec(BtC, sl.y)

    res := new(mat.Dense)
    res.Product(A,x)
    fmt.Printf("expect :\n%1.2f\n\n", mat.Formatted(b.T()))
    fmt.Printf("got:\n%1.2f\n\n", mat.Formatted(res.T()))
}
