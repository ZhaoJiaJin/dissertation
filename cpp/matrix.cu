#include "matrix.h"
#include <iostream.h>

Matrix::Matrix(int x, int y){
    row = x;
    col = y;
    cudaMallocManaged(&data, x*y*sizeof(float));
}

Matrix::~dev_array(){
    cubaFree(x);
}

float* Matrix::getdata(){
    return x
}

void Matrix::print(){
    for (int i=0; i < row; i++) {
        std::cout << "[";
        for (int j=0; j < col; j ++){
            std::cout << x[i*col+j] << " ";
        }
        std::cout << "]" << endl;
    }
}


