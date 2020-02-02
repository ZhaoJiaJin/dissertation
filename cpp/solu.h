#ifndef _SOLU_H
#define _SOLU_H

#include "matrix.h"
#include "helper.h"
#include "matrixcal.h"

class Solu
{
public:
    Solu(Matrix *a_, Matrix *t_,Matrix *y_, int m_, int n_, int bign_);
    double solve();
    double verifyans();

private:
    Matrix *a;
    Matrix *t;
    Matrix *y;
    int m;
    int n;
    int bign;
    int srcx;
    int srcy;
    Matrix AtTA;
    Matrix b;
    Matrix answer;

    void init();
    void calb();
    void calQx(Matrix &x, Matrix &res);
    void calBtCBX(Matrix &x, Matrix &res);
};


#endif
