#ifndef _MATRIXCAL_H
#define _MATRIXCAL_H

#include "matrix.h"
#include <vector>
#include<fstream>
#include <cstdlib>

void randomMatrix(Matrix &m);
int load_from_file(std::string fname, Matrix &m);
int load_diagonal(std::string fname, Matrix &m);
void mul(Matrix &a, Matrix &b, Matrix &res);
void mul_kernel(float *a, float *b, float *c, int m, int n, int p);
void kron_mul(Matrix &ma, Matrix &mb, Matrix &mx, Matrix &res);

void kron_prod(Matrix &a,Matrix &b, Matrix &res);

void adjacency_mul(Matrix& x, Matrix &res, int rowx, int colx, int srcx,int srcy);
void adjacency_mul_kernel(float *x, float *res, int rowx, int colx, int srcx,int srcy);
void find_neighbour(int i, int m, int n, std::vector<int> res);

void matrix_sub(Matrix& a,Matrix& b,Matrix& c,Matrix& res);
void matrix_sub_kernel(float* a,float* b,float* c,float* res, int size);

float dot(Matrix &a,Matrix &b);
float dot_kernel(float *a,float *b, int size);

#endif
