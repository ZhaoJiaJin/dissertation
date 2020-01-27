#include "matrix.h"

Matrix::Matrix(int x, int y){
    alloc(x,y);
}

Matrix::Matrix(){
}

void Matrix::alloc(int x, int y){
    row = x;
    col = y;
    //TODO:
    //cudaMallocManaged(&data, x*y*sizeof(float));
    data = new float[x*y];
}


Matrix::~Matrix(){
    delete[] data;   
}

void Matrix::print(){
    //std::cout << data[0] << std::endl;
    //return;
    
    std::cout << std::setiosflags(std::ios_base::showpoint);
    for (int i=0; i < row; i++) {
	    std::cout << "[";
        for (int j=0; j < col; j ++){
		    std::cout << data[i*col+j] << " ";
        }
	    std::cout << "]" << std::endl; 
    }
}

float Matrix::get(int i, int j){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    return data[i*col + j];
}

void Matrix::set(int i,int j,float v){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    data[i*col+j] = v;
}

void Matrix::set_by_idx(int i,float v){
    if (i >= row*col){
        throw "x or y exceed matrix size!";
    }
    data[i] = v;
}



void Matrix::crop(Matrix &newm){
    int x,y;
    x = newm.getrow();
    y = newm.getcol();
    if ((x > row) || (y > col)){
        throw "x or y exceed matrix size!";
    }
    for (int i=0; i < x; i ++){
        for (int j = 0; j < y; j ++){
            newm.set(i,j,get(i,j));
        }
    }
}

int Matrix::getrow(){
    return row;
}

int Matrix::getcol(){
    return col;
}

float* Matrix::get_data(){
    return data;
}


