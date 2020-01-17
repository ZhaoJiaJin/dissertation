package solu

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "conjugate/utils"
)

func TestFindNeigh(t *testing.T){
    assert.Equal(t, map[string]int{D:1,R:3},
        findNeighbors(0, 3, 3), "findNeighbors is wrong.")
    assert.Equal(t, map[string]int{L:2,U:4,R:8},
        findNeighbors(5, 3, 3), "findNeighbors is wrong.")
}

func TestIdentifyMatrix(t *testing.T){
    v := GenerateI(3)
    utils.PrintMatrix(v)
    v = GenerateI(5)
    utils.PrintMatrix(v)
}
