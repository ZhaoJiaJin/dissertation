#!/usr/bin/env python

import numpy as np
from helper import generateD
import time

class StdSolu:
    def __init__(self,afile,tfile,m,n,lvl):
        self.af = afile
        self.tf = tfile
        self.num_m = m
        self.num_n = n
        self.num_lvl = lvl
        self.num_N = 4**lvl
        self.A = None
        self.B = None

        self.D = None
        self.Q = None
        self.u = None
        self.y = None

        self.T = None
        self.N = None
        self.C = None

        self.y = None
        self.u = None

    def loadFromFile(self):
        A = np.loadtxt(open(self.af,"r"),delimiter=",")
        self.A = A[:self.num_m, :self.num_n]
        T = np.loadtxt(open(self.tf,"r"),delimiter=",")
        self.T = np.zeros((self.num_m,self.num_m))
        for i in range(0,self.num_m):
            self.T[i,i] = T[i]
        self.y = 5 * np.random.random_sample((self.num_m * self.num_N, 1))

    def getD(self):
        self.D = generateD(self.num_lvl)
        return self.D

    def getQ(self):
        #if self.D == None:
        #    raise Exception("D is None")
        qsize = self.num_n * self.num_N
        Q = np.zeros((qsize,qsize), dtype = int)
        DSqure = self.D.dot(self.D)
        for s in range(0, qsize, self.num_N):
            Q[s:s+self.num_N, s:s+self.num_N] = DSqure
        self.Q = Q
        return self.Q

    def getB(self):
        self.B = np.kron(self.A, np.eye(self.num_N))
        return self.B

    def getC(self):
        self.C = np.kron(self.T,np.eye(self.num_N))
        return self.C

    def findSolution(self):
        # init all variable
        self.loadFromFile()
        D = self.getD()
        Q = self.getQ()

        B = self.getB()
        BTrans = np.transpose(self.B)
        C = self.getC()
        y = self.y

        BTransC = BTrans.dot(C)

        BTransCB = BTransC.dot(B)
        leftmost = Q + BTransCB

        rightmost = BTransC.dot(y)
        leftmostInv = np.linalg.inv(leftmost)
        self.u = leftmostInv.dot(rightmost)
        return self.u



if __name__ == "__main__":
    with open("record","w") as f:
        for i in range(3,5):
            s = StdSolu("a.csv","t.csv",9,4,i)
            begin = time.time()
            print(i)
            print(s.findSolution().shape)
            end = time.time()
            print("cost:",end - begin)
            f.write("{0} {1}\n".format(i,end - begin))
            f.flush()
