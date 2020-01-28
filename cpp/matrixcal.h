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
void kron_mul(Matrix &ma, Matrix &mb, Matrix &mx, Matrix &res);

void kron_prod(Matrix &a,Matrix &b, Matrix &res);


#endif
