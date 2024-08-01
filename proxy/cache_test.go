package proxy

import "testing"

func TestFindGoodLimit100(t *testing.T) {
    goodCache := NewGoodCache()
    res := goodCache.FindGoodLimit100()
    t.Log(res)
}
