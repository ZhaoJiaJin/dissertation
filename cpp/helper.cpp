#include "helper.h"

Matrix randomMatrix(int row, int col){
    Matrix res(row,col);

    for (int i = 0; i < row; i ++){
        for (int j = 0; j < col; j ++){
            res.set(i,j,rand() % 1000);
        }
    }

    return res;
}

int calN(int level){
    return pow(4,level) * 12;
}


int loadFromFile(std::string fname, Matrix &m){
    std::ifstream infile(fname);
    int i = 0;
    float t = 0.0000000000000000000;
    while(infile >> t){
        m.set_by_idx(i,t);
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

    matrix_mul(a.get_data(),b.get_data(), res.get_data(), rowa,cola,colb);
}



void matrix_mul(float *a, float *b, float *c, int m, int n, int p){
    float tmp;
    for(int x = 0; x < m; x ++){
        for(int y = 0; y < p; y ++){
            tmp = 0;
            for(int t = 0; t < n; t++){
                tmp += (a[x*n + t] * b[t*p+y]);
            }
            c[p*x + y] = tmp;
        }
    }
}

void most_square(int v, int& orix, int& oriy){
    int begin = (int)sqrt(v);
    for (int i = begin; i >= 1; i --){
        if (v % i == 0){
            orix = i;
            oriy = v/i;
        }
    }
    return;
}
