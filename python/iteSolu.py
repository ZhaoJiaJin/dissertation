#!/usr/bin/env python

import numpy as np
from Solu import Solution
from helper import findNeighbors

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
    def __init__(self,afile,tfile,m,n,lvl,y,toler):
        super().__init__(afile,tfile,m,n,lvl,y)
        self.toler = toler
        self.b = None
        self.ATrans = np.transpose(self.A)
        self.A_TTA = self.ATrans.dot(self.T.dot(self.A))

    def DmulX(self, x):
        # x can be a N*z matrix,z can be bigger than 1
        (row,col) = x.shape
        res = np.zeros((self.num_N,col))
        for i in range(0,self.num_N):
            neighs = findNeighbors(i,self.Dsize[0],self.Dsize[1])
            neighs_len = len(neighs)
            for j in range(0,col):
                for idx in neighs:
                    res[i][j] += x[idx][j]
                res[i][j] -= (neighs_len* x[i][j])
        return res

    def getb(self):
        Y = np.reshape(self.y,(self.num_N, self.num_m), order='F')
        # b = reshape(NYTA)
        TA = self.T.dot(self.A)
        YTA = Y.dot(TA)
        #B = self.N.dot(YTA)
        B = YTA

        self.b = np.reshape(B,(self.num_n * self.num_N), order='F')
        return self.b

    def calQx(self, x):
        #return matvec_reshape(self.P, self.DSqure, x, True)
        X = np.reshape(x, (self.num_N,self.num_n),order='F')
        X = self.DmulX(X)
        X = self.DmulX(X)
        return np.reshape(X,(self.num_N*self.num_n),order='F')


    def calBTCBx(self,x):
        X = np.reshape(x, (self.num_N, self.num_n), order='F')
        #X = self.N.dot(X.dot(self.A_TTA))
        # N is identity matrix
        X = X.dot(self.A_TTA)
        return np.reshape(X,(self.num_N * self.num_n),order='F')

    def findSolution(self):
        # init b(Gx = b)
        b = self.getb()
        self.getDSize()
        #self.getDSqure()
        n = len(b)
        x = np.ones(n)

        r = b - self.calQx(x) - self.calBTCBx(x)
        p = r

        r_k_norm = np.dot(r,r)
        origin_r_norm = r_k_norm
        for i in range(2*n):
            #rold = r
            q = self.calQx(p) + self.calBTCBx(p)
            alpha = r_k_norm / np.dot(p,q)
            x += alpha * p
            if i % 50 == 0:
                r = b - self.calQx(x) - self.calBTCBx(x)
            else:
                r -= alpha * q

            r_k1_norm = np.dot(r,r)
            beta = r_k1_norm/r_k_norm
            r_k_norm = r_k1_norm
            #if r_k1_norm < 1e-10 * origin_r_norm:
            if r_k1_norm < 1e-5:
                #newr = b - self.calQx(x) - self.calBTCBx(x)
                #newr_norm = np.dot(newr,newr)
                #print('Itr:', i, r_k1_norm,newr_norm)
                print('Itr:', i, r_k1_norm)
                break
            p = r + beta * p
        return x





