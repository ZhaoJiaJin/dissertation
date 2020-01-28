#ifndef _MATRIX_H
#define _MATRIX_H

#include <iomanip>
#include <vector>
#include <iostream>

class Matrix
{
public:
  Matrix(int x, int y);
  Matrix();
  ~Matrix();
  void alloc(int x, int y);
  void print();
  void crop(Matrix &newm);
  float get(int i, int j);
  void set(int i, int j,float v);
  void set_by_idx(int i, float v);
  //void push_back(float v);
  int getrow();
  int getcol();
  float* get_data();
  void t(Matrix& t);
  /*Matrix kronReshapeMul()
  // v shape has to be row*1
  float dot(Matrix v);*/

private:
  float* data;
  int row, col;

};

#endif
