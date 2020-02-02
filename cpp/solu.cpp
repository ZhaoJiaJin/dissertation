#include "solu.h"


Solu::Solu(Matrix *a_, Matrix *t_,Matrix *y_, int m_, int n_, int bign_){
    a = a_;
    t = t_;
    y = y_;
    m = m_;
    n = n_;
    bign = bign_;
    init();
}


void Solu::init(){
    most_square(bign, srcx,srcy);   
    // get A transpose * T * A
    //get A transpose
    Matrix A_t;
    a->t(A_t);
    
    Matrix AtT;
    mul(A_t, *t, AtT);
    mul(AtT, *a, AtTA);
    //TODO: store all neighbour beforehead
}


float Solu::solve(){
    calb();

    int leng = b.getrow();
    answer.alloc(leng,1);
    Matrix v1;
    Matrix v2;
    calQx(answer, v1);
    calBtCBX(answer,v2);
    Matrix r;
    matrix_sub(b,v1,v2,r);
    Matrix p;
    p.copy(r);

    float r_k_norm = dot(r,r);
    std::cout << r_k_norm << std::endl;
    
    return 0;
} 


void Solu::calb(){
    y->resize(bign,m);
    Matrix yt;
    mul(*y,*t,yt);
    mul(yt,*a,b);
    b.resize(bign*n,1);
}

void Solu::calQx(Matrix &x, Matrix &res){
    Matrix tmpres;
    adjacency_mul(x, tmpres, bign,n,srcx,srcy);
    adjacency_mul(tmpres,res,bign,n,srcx,srcy);
}


void Solu::calBtCBX(Matrix &x, Matrix &res){
    x.resize(bign,n);
    mul(x,AtTA,res);
    res.resize(bign*n,1);
}
