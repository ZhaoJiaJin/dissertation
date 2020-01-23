#ifndef _HELPER_H
#define _HELPER_H

#include<fstream>
#include <cstdlib>
#include <cmath>
#include "matrix.h"

template <class T>
int loadFromFile(std::string fname, Matrix<T> *m);
template <class T>
Matrix<T> randomMatrix(int row, int col);
int calN(int level);


template <class T>
Matrix<T> randomMatrix(int row, int col){
    Matrix<T> res(row,col);
    std::vector<T> data = res.getdata();

    for (int i = 0; i < row*col; i ++){
        data[i] = rand() % 1000;
    }

    return res;
}

int calN(int level){
    return pow(4,level) * 12;
}


template <class T>
int loadFromFile(std::string fname, Matrix<T> *m){
    std::ifstream infile(fname);
    int i = 0;
    T t = 0.0000000000000000000;
    while(infile >> t){
        m->push_back(t);
        i++;
    }
    return 0;
}

#endif
