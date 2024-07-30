package strategy

import "fmt"

// 策略模式
const (
    scoreLeaderboardKey = "ScoreLeaderboardKey"
    coinLeaderboardKey  = "CoinLeaderboardKey"
)

// IStrategy 定义一个策略类
type IStrategy interface {
    GetRankDataFormZSet() string
}

// ScoreLeaderboard 游戏分数排行榜
type ScoreLeaderboard struct {
    Key string
}

func NewScoreLeaderboard() IStrategy {
    return &ScoreLeaderboard{Key: scoreLeaderboardKey}
}

// GetRankDataFormZSet 从 redis 的有序集合里拿数据，不同排行榜除了 key 不同，value 都是用户的 uid, score 为其所获取的游戏积分
func (s *ScoreLeaderboard) GetRankDataFormZSet() string {
    /*
       从 redis 拿数据
    */
    return fmt.Sprintf("从 key 为：%v 的 redis 有序集合中拿到数据了", s.Key)
}

// CoinLeaderboard 游戏币排行榜
type CoinLeaderboard struct {
    Key string
}

func NewCoinLeaderboard() IStrategy {
    return &CoinLeaderboard{Key: coinLeaderboardKey}
}

// GetRankDataFormZSet 从 redis 的有序集合里拿数据，不同排行榜除了 key 不同，value 都是用户的 uid, score 为其所获取的游戏积分
func (s *CoinLeaderboard) GetRankDataFormZSet() string {
    /*
       从 redis 拿数据
    */
    return fmt.Sprintf("从 key 为：%v 的 redis 有序集合中拿到数据了", s.Key)
}

// --------------------- 以下可理解为在项目中的 logic 模块(写业务逻辑时初始化策略执行者，设置策略后执行策略的方法) ---------------------

// Operator 具体策略的执行者
type Operator struct {
    strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
    operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate() string {
    return operator.strategy.GetRankDataFormZSet()
}
