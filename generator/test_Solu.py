import unittest
from iteSolu import IteSolu
from stdSolu import StdSolu
import time
import numpy as np
from helper import calNFromLvl
import sys

class TestSolution(unittest.TestCase):

    def test_IterateSolution(self):
        m = 2
        n = 2
        sourcefilea = "a.csv"
        sourcefilet = "t.csv"

        for i in range(1,10):
            print("-------------lvl:{0}, N:{1}------------".format(i, calNFromLvl(i)))
            y = 5 * np.random.randint(0,10,(m * calNFromLvl(i), 1))
            #iterate method
            start = time.time()
            ite_s = IteSolu(sourcefilea,sourcefilet,m,n,i,y,0.001)
            ite_x = ite_s.findSolution()
            end = time.time()
            print("iterate method took:{0}s".format(end - start))
            sys.stdout.flush()

            #std solution
            start = time.time()
            std_s = StdSolu(sourcefilea,sourcefilet,m,n,i,y)
            std_x = std_s.findSolution()
            end = time.time()
            print("standard method took:{0}s".format(end - start))
            sys.stdout.flush()

            distance = std_x - ite_x
            print("distance square:", np.dot(distance,distance))
            sys.stdout.flush()
            #for i in range(0, len(ite_x)):
            #    print(ite_x[i], std_x[i])
            #np.testing.assert_array_equal(ite_x,std_x)

if __name__ == '__main__':
    unittest.main()
