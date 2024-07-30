package strategy

import "testing"

func TestIStrategy(t *testing.T) {
    operator := Operator{}

    operator.setStrategy(NewScoreLeaderboard())
    result := operator.calculate()
    t.Log(result)

    operator.setStrategy(NewCoinLeaderboard())
    result = operator.calculate()
    t.Log(result)
}
