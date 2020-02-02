#include "matrix.h"
#include "matrixcal.h"
#include "helper.h"
#include "solu.h"
#include <string>

#include <stdio.h> 
#include <stdlib.h>
#include <time.h> 

int main(int argc, char *argv[]){
    if (argc != 3){
        std::cout << "please provide method and level" << std::endl;
        return 1;
    }
    srand (time(NULL));

    std::string inputfa = "a.csv";
    std::string inputta = "t.csv";
    const int originm = 9;
    const int  originn = 4;
    int m = 3;
    int n = 3;

    char* method = argv[1];
    char* level_c = argv[2];
    std::cout << method << level_c << std::endl;
    
    int level = atoi(level_c);
    int bign = calN(level);
    std::cout << "level is " << level << " N is " << bign << std::endl;
    Matrix origina (originm,originn);
    Matrix origint (originm,originm);

    load_from_file(inputfa, origina);
    load_diagonal(inputta, origint);

    Matrix a(m,n);
    Matrix t(m,m);
    origina.crop(a);
    origint.crop(t);

    Matrix y(m*bign,1);
    randomMatrix(y);

    Solu sl(&a,&t,&y,m,n,bign);

    sl.solve();
    /*double residual = sl.verifyans();
    std::cout << "Real Residual:" << residual << std::endl;
    */
    return 0;
}
