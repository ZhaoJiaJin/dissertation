#include "matrixcal.h"

void randomMatrix(Matrix &m){
    for (int i = 0; i < m.getrow(); i ++){
        for (int j = 0; j < m.getcol(); j ++){
            m.set(i,j,rand() % 10);
        }
    }
}

int load_from_file(std::string fname, Matrix &m){
    std::ifstream infile(fname);
    int i = 0;
    double t = 0.0000000000000000000;
    while(infile >> t){
        m.set_by_idx(i,t);
        i++;
    }
    return 0;
}

int load_diagonal(std::string fname, Matrix &m){
    std::ifstream infile(fname);
    int i = 0;
    double t = 0.0000000000000000000;
    while(infile >> t){
        m.set_diagonal(i,t);
        i++;
    }
    return 0;

}

void mul(Matrix &a, Matrix &b, Matrix &res){
    int rowa, rowb,cola, colb;
    rowa = a.getrow();
    rowb = b.getrow();
    cola = a.getcol();
    colb = b.getcol();

    if (cola != rowb){
        throw "matrix do not match in matrix multiplication";
    }

    res.alloc(rowa,colb);

    mul_kernel(a.get_data(),b.get_data(), res.get_data(), rowa,cola,colb);
}


void mul_kernel(double *a, double *b, double *c, int m, int n, int p){
    double tmp;
    for(int x = 0; x < m; x ++){
        for(int y = 0; y < p; y ++){
            tmp = 0;
            for(int t = 0; t < n; t++){
                tmp += (a[cal_idx(x,t,m,n)] * b[cal_idx(t,y,n,p)]);
            }
            c[cal_idx(x,y,m,p)] = tmp;
        }
    }
}

/*
 * kron_mul perform reshape and kronecker product, the caller should perform transpose of the left matrix by itself.
 */
void kron_mul(Matrix &ma, Matrix &mb, Matrix &mx, Matrix &res){
    if(mx.getcol() != 1){
        throw "the middle matrix should have one column when performing kronecker reshape product";
    }
    int bcol, xrow, arow; 
    arow = ma.getrow();
    xrow = mx.getrow();
    bcol = mb.getcol();

    if (arow * bcol != xrow){
        throw "the matrix sizes do not match when performing kronecker reshape product";
    }
    //res.alloc(leftx, righty);
    
    if (mx.resize(bcol,arow) != 0){
        throw "matrix resize failed";
    }
    Matrix tmpres;
    if (mb.is_identity()){
        tmpres.copy(mx);
    }else{
        mul(mb,mx,tmpres);
    }

    if (ma.is_identity()){
        res.copy(tmpres);
    }else{
        mul(tmpres, ma, res);
    }

    res.resize(xrow,1);
    return;
}

void kron_prod(Matrix &a, Matrix &b, Matrix &res){
    int arow,acol, brow,bcol;
    arow = a.getrow();
    acol = a.getcol();
    brow = b.getrow();
    bcol = b.getcol();

    res.alloc(arow*brow, acol*bcol);

    for(int outi = 0; outi < arow; outi ++){
        for(int outj = 0; outj < acol; outj ++){
            for(int inneri = 0; inneri < brow; inneri ++){
                for(int innerj = 0; innerj < bcol; innerj ++){
                    res.set(outi * brow+inneri, outj*bcol + innerj, a.get(outi,outj) * b.get(inneri,innerj));
                }
            }
        }
    }
}


void adjacency_mul(Matrix& x, Matrix& res, int rowx, int colx, int srcx,int srcy){

    res.alloc(rowx*colx,1);
    //std::cout << "full size" << rowx*colx << std::endl;
    double* rawx = x.get_data();
    double* rawres = res.get_data();
    //res.alloc(rowx,colx);
    adjacency_mul_kernel(rawx,rawres,rowx,colx,srcx,srcy);
}


//TODO:change to cuda
void adjacency_mul_kernel(double *x, double *res, int rowx, int colx, int srcx, int srcy){
    std::vector<int> neighs;
    for(int i=0; i < rowx; i ++){
        find_neighbour(i, srcx,srcy, neighs);
        int neigh_size = neighs.size();
        for (int j = 0; j < colx; j++){
            double val = 0.0f;
            for (int vstart = 0; vstart < neigh_size; vstart++){
                val += (x[j*rowx+neighs[vstart]]);
            }
            val -= (neigh_size * x[j*rowx + i]);
            //std::cout << j*rowx+i << std::endl;
            res[j*rowx+i] = val;
        }
    }
}


//TODO: change to device_vector
void find_neighbour(int i, int m, int n, std::vector<int> res){
    res.clear();
    int idxes[4] = {i-m,i+m,i-1,i+1};
    bool exist[4] = {true,true,true,true};
    if(i % m == 0){
        exist[2] = false;
    }
    if((i % n) == (n - 1)){
        exist[3] = false;
    }
    if(i < m){
        exist[0] = false;
    }
    if(i > (n*m-1-m)){
        exist[1] = false;
    }
    for(int s = 0; s < 4; s ++){
        if(exist[s]){
            res.push_back(idxes[s]);
        }
    }
}

void matrix_sub(Matrix& a,Matrix& b,Matrix& c,Matrix& res){
    int arow = a.getrow();
    int acol = a.getcol();
    int brow = b.getrow();
    int bcol = b.getcol();
    int crow = c.getrow();
    int ccol = c.getcol();
    if (arow != brow || arow != crow){
        throw "matrix sub failed.";
    }
    if (acol != bcol || acol != ccol){
        throw "matrix sub failed.";
    }
    res.alloc(arow,acol);
    matrix_sub_kernel(a.get_data(),b.get_data(),c.get_data(),res.get_data(),arow*acol);
}


//TODO: change to gpu
void matrix_sub_kernel(double* a,double* b,double* c,double* res, int size){
    for(int i=0; i < size; i ++){
        res[i] = a[i] - b[i] - c[i];
    }
}

double dot(Matrix &a,Matrix &b){
    if (a.getcol() != 1||b.getcol() != 1 || a.getrow() != b.getrow()) {
        throw "dot product failed.";
    }
    return dot_kernel(a.get_data(),b.get_data(),a.getrow());
}



//TODO: change to gpu
double dot_kernel(double *a,double *b, int size){
    double res = 0.0f;
    for(int i = 0; i < size; i ++){
        res += (a[i]*b[i]);
    }
    return res;
}

void matrix_add(Matrix& a,Matrix& b,Matrix& res){
    int arow = a.getrow();
    int acol = a.getcol();
    int brow = b.getrow();
    int bcol = b.getcol();
    if (arow != brow){
        throw "matrix sub failed.";
    }
    if (acol != bcol){
        throw "matrix sub failed.";
    }
    res.alloc(arow,acol);
    matrix_add_kernel(a.get_data(),b.get_data(),res.get_data(),arow*acol);
}


//TODO: change to gpu
void matrix_add_kernel(double* a,double* b,double* res, int size){
    for(int i=0; i < size; i ++){
        res[i] = a[i] + b[i] ;
    }
}

void matrix_add_scale(Matrix& a,Matrix& b,double scale,Matrix &res){
    int arow = a.getrow();
    int acol = a.getcol();
    int brow = b.getrow();
    int bcol = b.getcol();
    if (arow != brow){
        throw "matrix sub failed.";
    }
    if (acol != bcol){
        throw "matrix sub failed.";
    }
    matrix_add_scale_kernel(a.get_data(),b.get_data(),scale,res.get_data(),arow*acol);
}


//TODO: change to gpu
void matrix_add_scale_kernel(double* a,double* b,double scale,double* res, int size){
    for(int i=0; i < size; i ++){
        res[i] = a[i] + scale*b[i] ;
    }
}


