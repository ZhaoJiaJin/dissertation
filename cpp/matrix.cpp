#include "matrix.h"

Matrix::Matrix(int x, int y){
    alloc(x,y);
    identity = false;
}

Matrix::Matrix(){
    identity = false;
}

Matrix::Matrix(int x, bool is_iden){
    identity = is_iden;
    row = x;
    col = x;
    data = nullptr;
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


Matrix::Matrix(const Matrix& other){
	identify = other.is_identity();
	row = other.getrow();
	col = other.getcol();

	data = new float[row*col];
	otherdata = other.get_data();
	std::copy(std::begin(otherdata), std::end(otherdata), std::begin(data));
}


Matirx& operator=(const Matrix& other){
	if(&other != this){
		delete[] data;
		data = nullptr;
		data = new float[row*col];
		std::copy( std::begin(other.data), std::end(other.data), std::begin(data));
		row = other.getrow();
		col = other.getcol();
		is_iden
	}
	return *this;
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
    if (is_identity()){
        if (i == j){
            return 1;
        }else{
            return 0;
        }
    }
    return data[cal_idx(i,j,row,col)];
}

void Matrix::set(int i,int j,float v){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    if (is_identity()){
        throw "identity matrix do not support set method";
    }
    data[cal_idx(i,j,row,col)] = v;
}

void Matrix::set_by_idx(int i,float v){
    if (i >= row*col){
        throw "x or y exceed matrix size!";
    }
    if (is_identity()){
        throw "identity matrix do not support set method";
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
    if (is_identity() && (data == nullptr)){
        alloc(row,col);
        for(int i = 0; i < row; i ++){
            for (int j = 0; j < col; j ++){
                if (i == j){
                    data[cal_idx(i,j,row,col)] = 1;
                }else{
                    data[cal_idx(i,j,row,col)] = 0;
                }
            }
        }
    }
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

int Matrix::resize(int newr, int newc){
    if (newr * newc != row * col){
        return -1;
    }
    if (is_identity()){
        throw "identity matrix do not support resize method";
    }
    row = newr;
    col = newc;
    return 0;
}
