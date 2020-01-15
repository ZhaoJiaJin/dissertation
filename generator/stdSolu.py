#!/usr/bin/env python

import numpy as np
import time
from Solu import Solution

class StdSolu(Solution):
    def __init__(self,afile,tfile,m,n,lvl):
        super().__init__(afile,tfile,m,n,lvl)

    def findSolution(self):
        D = self.D
        print(D)
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
        for i in range(1,2):
            s = StdSolu("a.csv","t.csv",9,4,i)
            begin = time.time()
            print("level:",i)
            print(s.findSolution().shape)
            end = time.time()
            print("cost:",end - begin)
            f.write("{0} {1}\n".format(i,end - begin))
            f.flush()
