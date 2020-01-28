#include "matrixcal.h"

void randomMatrix(Matrix &m){
    for (int i = 0; i < m.getrow(); i ++){
        for (int j = 0; j < m.getcol(); j ++){
            m.set(i,j,rand() % 1000);
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

void kron_mul(Matrix &left, Matrix &middle, Matrix &right, Matrix &res){
    if(middle.getcol() != 1){
        throw "the middle matrix should have one column when performing kronecker reshape product";
    }
    int leftx,lefty,middlex,rightx,righty;
    leftx = left.getrow();
    lefty = left.getcol();
    middlex = middle.getrow();
    rightx = left.getrow();
    righty = left.getcol();

    if (lefty * rightx != middlex){
        throw "the matrix sizes do not match when performing kronecker reshape product";
    }
    res.alloc(leftx, righty);
    
    
}
