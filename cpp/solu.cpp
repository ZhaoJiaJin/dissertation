#include "solu.h"


Solu::Solu(Matrix &a_, Matrix &t_,Matrix y_, int m_, int n_, int bign_){
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
    a.t(A_t);
    Matrix res;
    
    Matrix AtT;
    mul(A_t, t, AtT);
    mul(AtT, a, AtTA);
}


float Solu::solve(Matrix& res){
    solveb();
} 


void Solu::solveb(){
    
}
