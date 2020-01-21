#ifndef _MATRIX_H
#define _MATRIX_H

#include <iostream>

class Matrix
{
public:
  Matrix(int x, int y);
  ~Matrix();
  void print();
  float* getdata();
  /*Matrix crop(int x, int y);
  Matrix mul(Matrix v);
  Matrix kronReshapeMul()
  // v shape has to be row*1
  float dot(Matrix v);*/

private:
  float* data;
  int row, col;
};

#endif
