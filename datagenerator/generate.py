#!/usr/bin/env python

"""compute D and D^2

because Q is very large can be represented using Kronecker product of P and D^2, so we wouldn't compute Q here
"""

import numpy as np
import time

L = "left"
R = "right"
U = "up"
D = "down"

def findNeighbors(i,n):
    """find all the neighbours of i
    Parameters
    ----------
    i : int
        the index of the point that we need to find its neighbours
    n : int
        the shape of original matrix is n*n
    """

    # all the neighours of i
    res = {
            L:i-n,
            R:i+n,
            U:i-1,
            D:i+1
            }

    # remove all the invalid neighbours
    if i % n == 0:#the first row
        res.pop(U)
    elif (i % n) == (n - 1):#the last row
        res.pop(D)
    if i < n:# the first column
        res.pop(L)
    elif i > n*n-1-n:# the last column
        res.pop(R)
    return res.values()


def generateD(lvl):
    """ compute matrix D
    Parameters:
    ----------
    lvl : int
        Power Level
    """
    size = 4**lvl
    ori_size = 2**lvl
    res = np.zeros((size,size),dtype=int)

    for i in range(0, size):
        # fill all the neighbours of i
        for pos in findNeighbors(i,ori_size):
            res[i][pos] = 1
    return res

def neighDSquare(i,n):
    """ find all the points that have common neighours with i and the sum of common neighours
    Parameters
    ----------
    i : int
        the index of the point that we need to find its neighbours
    n : int
        the shape of original matrix is n*n
    """
    neighs = {
            i:findNeighbors(i,n),

            i-n-1:[i-1,i-n],
            i-n+1:[i+1,i-n],
            i+n-1:[i-1,i+n],
            i+n+1:[i+1,i+n],

            i-2:[i-1],
            i+2:[i+1],
            i-2*n:[i-n],
            i+2*n:[i+n]
    }
    # remove unvalid points
    if int(i / n) - int((i-n-1) / n) != 1:
        neighs.pop(i-n-1)
    if int(i / n) - int((i-n+1) / n) != 1:
        neighs.pop(i-n+1)
    if int((i+n-1) / n) - int(i / n) != 1:
        neighs.pop(i+n-1)
    if int((i+n+1) / n) - int(i / n) != 1:
        neighs.pop(i+n+1)

    if int(i / n) - int((i-2) / n) != 0:
        neighs.pop(i-2)
    if int(i / n) - int((i+2) / n) != 0:
        neighs.pop(i+2)
    res = []

    for idx in list(neighs):
        if idx < 0 or idx > n*n-1:
            continue
        cnt = 0
        for neighidx in neighs[idx]:
            if neighidx >= 0 and neighidx <= n*n-1:
                cnt += 1
        res.append((idx,cnt))
    return res

def generateDSquare(lvl):
    """ compute D Square
    let D^D = M
    then M(i,j) means how many common neighours do i and j have
    Parameters:
    ----------
    lvl : int
        Power Level
    """
    size = 4**lvl
    ori_size = 2**lvl
    res = np.zeros((size,size),dtype=int)
    for i in range(0,size):
        for (neighid,neighvalue) in neighDSquare(i, ori_size):
            res[i][neighid] = neighvalue
    return res



if __name__ == "__main__":
    # compute D and D Square for power level from 1 to 8
    for lvl in range(1,8):
        print("lvl",lvl)
        begin = time.time()
        res = generateD(lvl)
        end = time.time()
        #print("D:",res)
        print("generate D, time cost(s):", end-begin)

        begin = time.time()
        new_q = generateDSquare(lvl)
        end = time.time()
        print("generate D Square time cost using new method(s):", end-begin)

        #begin = time.time()
        #np_q = res.dot(res)
        #end = time.time()
        print("generate D Square cost using numpy.dot(s):", end-begin)

        #print("compare result of D Squares from different method:",(np_q==new_q).all())
