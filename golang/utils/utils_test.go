package utils

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "math"
)

func TestCalN(t *testing.T){
    assert.Equal(t,  48,CalN(1), "CalN(1) is wrong.")
    assert.Equal(t, 16 *12,CalN(2),  "CalN(2) is wrong.")
    assert.Equal(t, 16 * 4 *12, CalN(3), "CalN(3) is wrong.")
}

func TestMostSqure(t *testing.T){
    for i:=1; i < 12; i ++{
        x,y := MostSqure(CalN(i))
        assert.Equal(t, int(math.Pow(2,float64(i))*3), x, "MostSqure Wrong")
        assert.Equal(t, int(math.Pow(2,float64(i))*4), y, "MostSqure Wrong")
    }
}
