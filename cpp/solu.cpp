#include "solu.h"


Solu::Solu(Matrix *a_, Matrix *t_,Matrix *y_, int m_, int n_, int bign_){
    a = a_;
    t = t_;
    y = y_;
    m = m_;
    n = n_;
    bign = bign_;
    init();
}


void Solu::init(){
    most_square(bign, srcx,srcy);   
    // get A transpose * T * A
    //get A transpose
    Matrix A_t;
    a->t(A_t);
    
    Matrix AtT;
    mul(A_t, *t, AtT);
    mul(AtT, *a, AtTA);
    //TODO: store all neighbour beforehead
}


double Solu::solve(){
    calb();
    b.printraw("b:");
    int leng = b.getrow();
    answer.alloc(leng,1);
    Matrix v1;
    Matrix v2;
    calQx(answer, v1);
    calBtCBX(answer,v2);
    v1.printraw("v1:");
    v2.printraw("v2:");
    Matrix r;
    matrix_sub(b,v1,v2,r);
    r.printraw("r:");
    Matrix p;
    p.copy(r);
    double r_k_norm = dot(r,r);
    //std::cout << "r_k_norm:" << r_k_norm << std::endl;
    printf("r_k_norm:%.17g\n", r_k_norm);
    //std::cout << r_k_norm << std::endl;
    r.printraw("r:");
    p.printraw("p:");
    answer.printraw("ans:");
    for(int i = 1; i < 2*leng; i ++){
    //for(int i = 1; i < 10; i ++){
    	//p.print("Matrix p");
        std::cout << "Ite:" << i << std::endl;
        calQx(p, v1);
        calBtCBX(p,v2);
        Matrix q;
        matrix_add(v1,v2,q);
	q.printraw("q:");
        double alpha = r_k_norm / dot(p,q);
	//std::cout << "alpha:" << alpha << std::endl;
    	printf("alpha:%.17g\n", alpha);
        matrix_add_scale(answer,p,alpha,answer);
	answer.printraw("ans:");
        //std::cout << alpha << std::endl;

	if(i%5 != 0){
        matrix_add_scale(r,q, alpha*-1,r);
	}else{

        calQx(answer, v1);
        calBtCBX(answer,v2);
        matrix_sub(b,v1,v2,r);
	}
	r.printraw("r:");

        double r_k1_norm = dot(r,r);
        //std::cout << "NewR:" << r_k1_norm << std::endl;
    	printf("NewR:%.17g\n", r_k1_norm);
        double beta = r_k1_norm/r_k_norm;
        r_k_norm = r_k1_norm;

        if (r_k1_norm < 1e-14){
            //std::cout << "Stop:" << i << " "<< r_k1_norm << std::endl;
    	    printf("Stop:%d %.17g\n", i,r_k1_norm);
            break;
        }
        matrix_add_scale(r,p,beta,p);
	p.printraw("p:");
    }
    
    return 0;
} 


void Solu::calb(){
    y->resize(bign,m);
    Matrix yt;
    y->printraw("y:");
    t->printraw("t:");
    mul(*y,*t,yt);
    yt.printraw("yt:");
    mul(yt,*a,b);
    b.printraw("yta:");
    b.resize(bign*n,1);
}

void Solu::calQx(Matrix &x, Matrix &res){
    Matrix tmpres;
    adjacency_mul(x, tmpres, bign,n,srcx,srcy);
    adjacency_mul(tmpres,res,bign,n,srcx,srcy);
}


void Solu::calBtCBX(Matrix &x, Matrix &res){
    x.resize(bign,n);
    mul(x,AtTA,res);
    x.resize(bign*n,1);
    res.resize(bign*n,1);
}

double Solu::verifyans(){
    Matrix v1;
    Matrix v2;
    calQx(answer, v1);
    calBtCBX(answer,v2);
    Matrix r;
    matrix_sub(b,v1,v2,r);
    return dot(r,r);
}
