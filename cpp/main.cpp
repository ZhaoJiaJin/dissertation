#include "matrix.h"
#include "matrixcal.h"
#include "helper.h"
#include <string>

int main(int argc, char *argv[]){
    if (argc != 3){
        std::cout << "please provide method and level" << std::endl;
        return 1;
    }
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

    origina.print();
    origint.print();
    /*Matrix a(m,n);
    Matrix t(m,m);
    origina.crop(a);
    origint.crop(t);

    Matrix y(m*bign,1);
    randomMatrix(y);*/

    Matrix t(3,true);
    Matrix a(m,n);
    //Matrix t(m,m);
    randomMatrix(a);
    //randomMatrix(t);
    a.print();
    t.print();
    Matrix kres;
    kron_prod(a,t,kres);
    //std::cout << "result" << std::endl;


    Matrix y(9,1);
    randomMatrix(y);
    Matrix res;
    mul(kres,y,res);
    res.print();

    Matrix at;
    a.t(at);
    Matrix newres;
    kron_mul(at, t, y,newres);
    newres.print();
    return 0;
}
