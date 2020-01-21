#include "matrix.h"

int main(void){
    Matrix a = Matrix(2,3);
    float* data = a.getdata();
    for (int i=1;i <= 2*3; i ++){
        data[i-1] = i;
    }
    a.print();
    return 0;
}
