#include "matrixcal.h"

void randomMatrix(Matrix &m){
    for (int i = 0; i < m.getrow(); i ++){
        for (int j = 0; j < m.getcol(); j ++){
            m.set(i,j,rand() % 10);
        }
    }
}

int load_from_file(std::string fname, Matrix &m){
    std::ifstream infile(fname);
    int i = 0;
    float t = 0.0000000000000000000;
    while(infile >> t){
        m.set_by_idx(i,t);
        i++;
    }
    return 0;
}

int load_diagonal(std::string fname, Matrix &m){
    std::ifstream infile(fname);
    int i = 0;
    float t = 0.0000000000000000000;
    while(infile >> t){
        m.set_diagonal(i,t);
        i++;
    }
    return 0;

}

void mul(Matrix &a, Matrix &b, Matrix &res){
    int rowa, rowb,cola, colb;
    rowa = a.getrow();
    rowb = b.getrow();
    cola = a.getcol();
    colb = b.getcol();

    if (cola != rowb){
        throw "matrix do not match in matrix multiplication";
    }

    res.alloc(rowa,colb);

    mul_kernel(a.get_data(),b.get_data(), res.get_data(), rowa,cola,colb);
}


void mul_kernel(float *a, float *b, float *c, int m, int n, int p){
    float tmp;
    for(int x = 0; x < m; x ++){
        for(int y = 0; y < p; y ++){
            tmp = 0;
            for(int t = 0; t < n; t++){
                tmp += (a[cal_idx(x,t,m,n)] * b[cal_idx(t,y,n,p)]);
            }
            c[cal_idx(x,y,m,p)] = tmp;
        }
    }
}

/*
 * kron_mul perform reshape and kronecker product, the caller should perform transpose of the left matrix by itself.
 */
void kron_mul(Matrix &ma, Matrix &mb, Matrix &mx, Matrix &res){
    if(mx.getcol() != 1){
        throw "the middle matrix should have one column when performing kronecker reshape product";
    }
    int bcol, xrow, arow; 
    arow = ma.getrow();
    xrow = mx.getrow();
    bcol = mb.getcol();

    if (arow * bcol != xrow){
        throw "the matrix sizes do not match when performing kronecker reshape product";
    }
    //res.alloc(leftx, righty);
    
    if (mx.resize(bcol,arow) != 0){
        throw "matrix resize failed";
    }
    Matrix tmpres;
    if (mb.is_identity()){
        tmpres = mx;
    }else{
      mul(mb,mx,tmpres);
    }

    if (ma.is_identity()){
        res = tmpres;
    }else{
      mul(tmpres, ma, res);
    }

    res.resize(xrow,1);
    return;
}

void kron_prod(Matrix &a, Matrix &b, Matrix &res){
    int arow,acol, brow,bcol;
    arow = a.getrow();
    acol = a.getcol();
    brow = b.getrow();
    bcol = b.getcol();

    res.alloc(arow*brow, acol*bcol);

    for(int outi = 0; outi < arow; outi ++){
        for(int outj = 0; outj < acol; outj ++){
            for(int inneri = 0; inneri < brow; inneri ++){
                for(int innerj = 0; innerj < bcol; innerj ++){
                    res.set(outi * brow+inneri, outj*bcol + innerj, a.get(outi,outj) * b.get(inneri,innerj));
                }
            }
        }
    }
}
