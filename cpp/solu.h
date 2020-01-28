#ifndef _SOLU_H
#define _SOLU_H

#include "matrix.h"
#include "helper.h"

class Solu
{
public:
    Solu(Matrix &a_, Matrix &t_,Matrix y_, int m_, int n_, int bign_);
    float solve(Matrix &res);

private:
    Matrix a;
    Matrix t;
    Matrix y;
    int m;
    int n;
    int bign;
    int srcx;
    int srcy;
    Matrix AtTA;
    Matrix b;

    void init();
    void solveb();
};


#endif
