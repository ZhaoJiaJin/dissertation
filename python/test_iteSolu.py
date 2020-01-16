import unittest
import numpy as np
from iteSolu import matvec_reshape,IteSolu
from helper import calNFromLvl

class TestHelperFunctions(unittest.TestCase):

    def test_matvec_reshape(self):
        shapea = (3,6)
        shapeb = (4,5)
        A = np.random.randint(0,50,shapea)
        #A = np.array([[1,2],[3,4]])
        B = np.random.randint(0,50,shapeb)
        #B = np.array([[1,0],[0,1]])
        x = np.random.randint(0,50,(shapea[1] * shapeb[1]))
        #x = np.array([1,2,3,4])
        res = matvec_reshape(A,B,x)
        expect = np.kron(A,B).dot(x)
        np.testing.assert_array_equal(res,expect,"wrong answer")

    def test_DmulX(self):
        m = 2
        n = 2
        sourcefilea = "a.csv"
        sourcefilet = "t.csv"
        lvl = 1
        y = 5 * np.random.randint(0,10,(m * calNFromLvl(lvl), 1))
        ite_s = IteSolu(sourcefilea,sourcefilet,m,n,lvl,y,0.001)
        ite_s.getDSize()

        X = np.random.randint(0,10,(ite_s.num_N,ite_s.num_n))
        print("X:",X)
        res = ite_s.DmulX(X)

        D = ite_s.getD()
        print("D:",D)
        expect = D.dot(X)
        np.testing.assert_array_equal(res,expect,"wrong answer")

    def test_QX(self):
        m = 2
        n = 2
        sourcefilea = "a.csv"
        sourcefilet = "t.csv"
        lvl = 1
        y = 5 * np.random.randint(0,10,(m * calNFromLvl(lvl), 1))
        ite_s = IteSolu(sourcefilea,sourcefilet,m,n,lvl,y,0.001)
        ite_s.getDSize()

        x = np.random.randint(0,10,(ite_s.num_N*ite_s.num_n))
        print("x:",x)
        res = ite_s.calQx(x)

        ite_s.getD()
        Q = ite_s.getQ()
        print("Q:",Q)
        expect = Q.dot(x)
        np.testing.assert_array_equal(res,expect,"wrong answer")





if __name__ == '__main__':
    unittest.main()
