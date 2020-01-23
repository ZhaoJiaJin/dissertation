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
    Matrix <float> a (originm,originn);
    Matrix <float> t (originm,1);

    loadFromFile<float>(inputfa, &a);
    loadFromFile<float>(inputta, &t);

    a.print();
    t.print();
    a = a.crop(m,n);
    t = t.crop(m,1);
    //Matrix y = randomMatrix(m*bign,1);
    a.print();
    t.print();
    return 0;
}
