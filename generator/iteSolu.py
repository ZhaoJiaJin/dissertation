#!/usr/bin/env python

import numpy as np
from Solu import Solution

def matvec_reshape(A,B,x,isASymm = False):
    (rowa,cola) = A.shape
    (rowb,colb) = B.shape
    X = np.reshape(x, (colb,cola),order='F')
    if isASymm:
        Y = B.dot(X.dot(A))
    else:
        Y = B.dot(X.dot(np.transpose(A)))
    return np.reshape(Y,(rowa*rowb),order='F')


class IteSolu(Solution):
    def __init__(self,afile,tfile,m,n,lvl,toler):
        super().__init__(afile,tfile,m,n,lvl)
        self.toler = toler
        self.b = None
        self.ATrans = np.transpose(self.A)


    def getb(self):
        Y = np.reshape(self.y,(self.num_N, self.num_m), order='F')
        # b = reshape(NYTA)
        TA = self.T.dot(self.A)
        YTA = Y.dot(TA)
        B = self.N.dot(YTA)

        self.b = np.reshape(B,(self.num_n * self.num_N), order='F')
        return self.b

    def calQx(self, x):
        #TODO: can be optmised using Algo. 1
        return matvec_reshape(self.P, self.DSqure, x, True)

    def calBTCBx(self,x):
        X = np.reshape(x, (self.num_N, self.num_n), order='F')
        X = X.dot(self.ATrans)
        X = X.dot(self.T)
        X = self.N.dot(X)
        X = X.dot(self.A)
        return np.reshape(X,(self.num_N * self.num_n),order='F')

    def findSolution(self):
        # init b(Gx = b)
        b = self.getb()
        self.getDSqure()
        n = len(b)
        x = np.ones(n)

        r = b - self.calQx(x) - self.calBTCBx(x)
        p = r

        r_k_norm = np.dot(r,r)
        origin_r_norm = r_k_norm
        for i in range(2*n*2):
            #rold = r
            q = self.calQx(p) + self.calBTCBx(p)
            alpha = r_k_norm / np.dot(p,q)
            x += alpha * p
            r -= alpha * q

            r_k1_norm = np.dot(r,r)
            beta = r_k1_norm/r_k_norm
            r_k_norm = r_k1_norm
            if r_k1_norm < 1e-10 * origin_r_norm:
                print('Itr:', i)
                break
            p = r + beta * p
        print(r_k_norm)
        return x

if __name__ == "__main__":
    s = IteSolu("a.csv","t.csv",9,4,1,0.1)
    print(s.findSolution())




