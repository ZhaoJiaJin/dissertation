#!/usr/bin/env python

"""compute D and D^2

because Q is very large can be represented using Kronecker product of P and D^2, so we wouldn't compute Q here
"""

import math
import numpy as np
import time

L = "left"
R = "right"
U = "up"
D = "down"

def findNeighbors(i,m,n):
    """find all the neighbours of i
    Parameters
    ----------
    i : int
        the index of the point that we need to find its neighbours
    m : int
        the number of rows of original matrix is n*n
    n : int
        the number of columns of original matrix is n*n
    """

    # all the neighours of i
    res = {
            L:i-m,
            R:i+m,
            U:i-1,
            D:i+1
            }

    # remove all the invalid neighbours
    if i % m == 0:#the first row
        res.pop(U)
    elif (i % n) == (n - 1):#the last row
        res.pop(D)
    if i < m:# the first column
        res.pop(L)
    elif i > n*m-1-m:# the last column
        res.pop(R)
    return res.values()

def mostSqure(v):
    sqrt = int(math.sqrt(v))
    for i in range(sqrt, 0, -1):
        if v % i == 0:
            return (i, v/i)


def generateD(lvl):
    """ compute matrix D
    Parameters:
    ----------
    lvl : int
        Power Level
    """
    size = calNFromLvl(lvl)
    (m,n) = mostSqure(size)
    res = np.zeros((size,size),dtype=int)

    for i in range(0, size):
        # fill all the neighbours of i
        diag = 0
        for pos in findNeighbors(i,m,n):
            res[i][pos] = 1
            diag += 1
        res[i][i] = -1*diag
    return res


def calNFromLvl(lvl):
    return 4**lvl*12

if __name__ == "__main__":
    # compute D and D Square for power level from 1 to 8
    for lvl in range(1,8):
        print("lvl",lvl)
        begin = time.time()
        res = generateD(lvl)
        end = time.time()
        print("D:",res)
        print("generate D, time cost(s):", end-begin)

        begin = time.time()
        new_q = generateDSquare(lvl)
        end = time.time()
        print("generate D Square time cost using new method(s):", end-begin)

        #begin = time.time()
        #np_q = res.dot(res)
        #end = time.time()
        #print("generate D Square cost using numpy.dot(s):", end-begin)

        #print("compare result of D Squares from different method:",(np_q==new_q).all())
