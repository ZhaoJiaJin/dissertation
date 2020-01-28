#ifndef _MATRIXCAL_H
#define _MATRIXCAL_H

#include "matrix.h"
#include<fstream>
#include <cstdlib>

void randomMatrix(Matrix &m);
int load_from_file(std::string fname, Matrix &m);
int load_diagonal(std::string fname, Matrix &m);
void mul(Matrix &a, Matrix &b, Matrix &res);
void mul_kernel(float *a, float *b, float *c, int m, int n, int p);
void kron_mul(Matrix &left, Matrix &middle, Matrix &right, Matrix &res);
//void kron_mul_left_kernel(float *a, float );
//void kron_mul_right_kernel(float *a, float );



#endif
