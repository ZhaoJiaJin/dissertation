#!/usr/bin/env python

import numpy as np
import time
from Solu import Solution

class StdSolu(Solution):
    def __init__(self,afile,tfile,m,n,lvl,y):
        super().__init__(afile,tfile,m,n,lvl,y)

    def findSolution(self):
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
        return np.squeeze(self.u)




