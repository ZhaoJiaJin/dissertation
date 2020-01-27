#include "matrix.h"
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
    Matrix origint (originm,1);

    loadFromFile(inputfa, origina);
    loadFromFile(inputta, origint);

    origina.print();
    origint.print();
    Matrix a(m,n);
    Matrix t(m,1);
    origina.crop(a);
    origint.crop(t);
    //Matrix y = randomMatrix(m*bign,1);
    a.print();
    t.print();
    Matrix res;
    mul(a,t,res);
    std::cout << "result" << std::endl;
    res.print();
    return 0;
}
