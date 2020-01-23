#ifndef _OPERATE_H
#define _OPERATE_H

#include "matrix.h"

int matrixMul(Matrix a, Matrix b, Matrix c);
float vectorDot(Matrix a, Matrix b, Matrix d);
Matrix Qx(Matrix x);
Matrix Dx(Matrix x);

#endif
