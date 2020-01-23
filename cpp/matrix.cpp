#include "matrix.h"

template <class T>
Matrix<T>::Matrix(int x, int y){
    row = x;
    col = y;
    //TODO:
    //cudaMallocManaged(&data, x*y*sizeof(float));
    //data = new float[x*y];
}


template <class T>
std::vector<T> *Matrix<T>::getdata(){
    return &data;
}

template <class T>
void Matrix<T>::print(){
    //std::cout << data[0] << std::endl;
    //return;
    
    std::cout << std::setiosflags(std::ios_base::showpoint);
    for (int i=0; i < row; i++) {
	    std::cout << "[";
        for (int j=0; j < col; j ++){
		    std::cout << data.at(i*col+j) << " ";
        }
	    std::cout << "]" << std::endl; 
    }
}

template <class T>
T Matrix<T>::get(int i, int j){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    return data.at(i*col + j);
}

template <class T>
void Matrix<T>::set(int i,int j,T v){
    if ((i >= row) || (j >= col)){
        throw "x or y exceed matrix size!";
    }
    data[i*col+j] = v;
}

template <class T>
Matrix<T> Matrix<T>::crop(int x, int y){
    if ((x > row) || (y > col)){
        throw "x or y exceed matrix size!";
    }
    Matrix<T> newm(x,y);
    for (int i=0; i < x; i ++){
        for (int j = 0; j < y; j ++){
            newm.push_back(get(i,j));
        }
    }
    return newm;
}

template <class T>
void Matrix<T>::push_back(T v){
    data.push_back(v);
}

template class Matrix<float>;
