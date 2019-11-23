#!/usr/bin/env python

import numpy as np

L = "left"
R = "right"
U = "up"
D = "down"

def findneighbors(i,n):
    res = {
            L:i-n,
            R:i+n,
            U:i-1,
            D:i+1
            }
    if i % n == 0:#the first row
        res.pop(U)
    elif (i % n) == (n - 1):#the last row
        res.pop(D)
    if i < n:# the first column
        res.pop(L)
    elif i > n*n-1-n:
        res.pop(R)
    return res.values()




def generateD(lvl):
    size = 4**lvl
    ori_size = 2**lvl
    res = np.zeros((size,size),dtype=int)

    for i in range(0, size):
        for pos in findneighbors(i,ori_size):
            res[i][pos] = 1
    return res


if __name__ == "__main__":
    print(generateD(10).shape)
