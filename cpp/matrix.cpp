#include "matrix.h"

Matrix::Matrix(int x, int y){
    alloc(x,y);
}

Matrix::Matrix(){
}

Matrix::Matrix(int x, bool is_iden){
    identity = is_iden;
    row = x;
    col = x;
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
	std::cout << std::endl;
    std::cout << std::setiosflags(std::ios_base::showpoint);
    for (int i=0; i < row; i++) {
	    std::cout << "[";
        for (int j=0; j < col; j ++){
		    std::cout << get(i,j) << " ";
        }
	    std::cout << "]" << std::endl; 
    }
}

float Matrix::get(int i, int j){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    return data[cal_idx(i,j,row,col)];
}

void Matrix::set(int i,int j,float v){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    data[cal_idx(i,j,row,col)] = v;
}

void Matrix::set_by_idx(int i,float v){
    if (i >= row*col){
        throw "x or y exceed matrix size!";
    }
    data[i] = v;
}

void Matrix::set_diagonal(int i, float v){
    set(i,i,v);
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


void Matrix::t(Matrix& t){
    int ox,oy;
    ox = getrow();
    oy = getcol();
    t.alloc(oy,ox);
    for(int i = 0; i < ox; i ++){
        for (int j = 0; j < oy; j ++){
            t.set(j,i, get(i,j));
        }
    }
    return;
}

bool Matrix::is_identity(){
    return identity;
}
