#ifndef _MATRIX_H
#define _MATRIX_H

#include <iomanip>
#include <iostream>
#include <algorithm>
#include "helper.h"

/*
 * due to how kronecker and reshape works, 
 * we use column-first layout in our data array to represent our matrix
 */
class Matrix
{
public:
  Matrix(int x, int y);
  Matrix();
  Matrix(int c, bool is_iden);
  ~Matrix();

  Matrix(const Matrix& other);
  Matirx& operator=(const Matrix& other);
  void alloc(int x, int y);
  void print();
  void crop(Matrix &newm);
  float get(int i, int j);
  void set(int i, int j,float v);
  void set_by_idx(int i, float v);
  void set_diagonal(int i, float v);
  //void push_back(float v);
  int getrow();
  int getcol();
  float* get_data();
  void t(Matrix& t);
  bool is_identity();
  int resize(int newr,int newc);
  /*Matrix kronReshapeMul()
  // v shape has to be row*1
  float dot(Matrix v);*/

private:
  float* data;
  int row, col;
  bool identity;

};

#endif
