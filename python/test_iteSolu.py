import unittest
import numpy as np
from iteSolu import matvec_reshape

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

if __name__ == '__main__':
    unittest.main()
