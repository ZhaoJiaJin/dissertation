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


int loadFromFile(std::string fname, Matrix *m){
    std::ifstream infile(fname);
    int i = 0;
    float t = 0.0000000000000000000;
    while(infile >> t){
        m->set_by_idx(i,t);
        i++;
    }
    return 0;
}

Matrix mul(Matrix a, Matrix b){
    int rowa, rowb,cola, colb;
    rowa = a.getrow();
    rowb = b.getrow();
    cola = a.getcol();
    colb = b.getcol();

    if (cola != rowb){
        throw "matrix do not match in matrix multiplication";
    }

    Matrix res(rowa,colb);

    return res;
}
