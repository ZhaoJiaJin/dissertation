#ifndef _MATRIX_H
#define _MATRIX_H

#include <iomanip>
#include <vector>
#include <iostream>

template <class T>
class Matrix
{
public:
  Matrix(int x, int y);
  void print();
  std::vector<T> *getdata();
  Matrix crop(int x, int y);
  T get(int i, int j);
  void set(int i, int j,T v);
  void push_back(T v);
  /*Matrix mul(Matrix v);
  Matrix kronReshapeMul()
  // v shape has to be row*1
  float dot(Matrix v);*/

private:
  std::vector<T> data;
  int row, col;
};

#endif
