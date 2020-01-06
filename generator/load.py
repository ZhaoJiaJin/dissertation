#!/usr/bin/env python

import numpy as np

def loadA(fname):
    return np.loadtxt(open(fname,"r"),delimiter=",")
def loadt(fname):
    return np.loadtxt(open(fname,"r"),delimiter=",")

if __name__ == "__main__":
    print(loadA("./a.csv").shape)
    print(loadA("./t.csv").shape)
