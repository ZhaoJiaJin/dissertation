package utils

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalN(t *testing.T){
    assert.Equal(t,  48.0,CalN(1), "CalN(1) is wrong.")
    assert.Equal(t, 16 *12.0,CalN(2),  "CalN(2) is wrong.")
    assert.Equal(t, 16 * 4 *12.0, CalN(3), "CalN(3) is wrong.")
}

func TestMostSqure(t *testing.T){
    x,y := MostSqure(4)
    assert.Equal(t, 2, x, "MostSqure Wrong")
    assert.Equal(t, 2, y, "MostSqure Wrong")
    x,y = MostSqure(6)
    assert.Equal(t, 2, x, "MostSqure Wrong")
    assert.Equal(t, 3, y, "MostSqure Wrong")
}
