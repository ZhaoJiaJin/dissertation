#ifndef _HELPER_H
#define _HELPER_H

#include<fstream>
#include <cstdlib>
#include <cmath>
#include "matrix.h"

int loadFromFile(std::string fname, Matrix *m);
Matrix randomMatrix(int row, int col);
int calN(int level);

Matrix mul(Matrix a, Matrix b);



#endif
