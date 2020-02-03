#include "helper.h"


int calN(int level){
    return pow(4,level) * 1;
}

void most_square(int v, int& orix, int& oriy){
    int begin = (int)sqrt(v);
    for (int i = begin; i >= 1; i --){
        if (v % i == 0){
            orix = i;
            oriy = v/i;
        }
    }
    return;
}

int cal_idx(int x, int y, int row, int col){
    return row * y + x;
}


