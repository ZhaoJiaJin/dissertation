#!/usr/bin/env python

import numpy as np
from helper import generateD,calNFromLvl

class Solution:
    def __init__(self,afile,tfile,m,n,lvl,y):
        self.af = afile
        self.tf = tfile
        self.num_m = m
        self.num_n = n
        self.num_lvl = lvl
        self.num_N = calNFromLvl(lvl)
        self.A = None
        self.B = None

        self.D = None
        self.Q = None
        self.u = None

        self.T = None
        self.N = None
        self.C = None

        self.y = y
        self.u = None
        self.N = np.eye(self.num_N)
        self.P = np.eye(self.num_n)
        # init all variable
        self.loadFromFile()
        self.getD()
        self.DSqure = None

    def getDSqure(self):
        DSqure = self.D.dot(self.D)
        self.DSqure = DSqure
        return self.DSqure


    def loadFromFile(self):
        f = open(self.af,"r")
        A = np.loadtxt(f,delimiter=",")
        f.close()
        self.A = A[:self.num_m, :self.num_n]
        f = open(self.tf,"r")
        T = np.loadtxt(f,delimiter=",")
        f.close()
        self.T = np.zeros((self.num_m,self.num_m))
        for i in range(0,self.num_m):
            self.T[i,i] = T[i]

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
        self.C = np.kron(self.T,self.N)
        return self.C

