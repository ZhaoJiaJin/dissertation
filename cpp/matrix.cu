#include "matrix.h"

Matrix::Matrix(int x, int y){
    row = x;
    col = y;
    cudaMallocManaged(&data, x*y*sizeof(float));
}

Matrix::~Matrix(){
    cudaFree(data);
}

float* Matrix::getdata(){
    return data;
}

void Matrix::print(){
    for (int i=0; i < row; i++) {
	std::cout << "[";
        for (int j=0; j < col; j ++){
		std::cout << data[i*col+j] << " ";
        }
	std::cout << "]" << std::endl; 
    }
}


