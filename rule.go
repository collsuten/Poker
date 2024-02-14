package rule

import (
	"math/rand"
)

var LandlordRules = _rules{reserved: true}
var TeamRules = _rules{}

// RunFastRules 跑得快規則
var RunFastRules = _rules{reserved: true, isRunFast: true}

// 牌值的常量定义
const (
	SpecialS = 14
	SpecialX = 15
	SpecialD = 16
)

// 新规则结构体
type _rules struct {
	reserved  bool
	isRunFast bool
}

// 新增规则函数
func (r _rules) Value(key int) int {
	if key == 1 {
		return 12
	} else if key == 2 || key == SpecialS || key == SpecialX || key == SpecialD {
		return key
	} else if key > 16 {
		return key
	}
	return key - 2
}

// 新增规则函数
func (r _rules) IsStraight(faces []int, count int) bool {
	if len(faces) < 3 {
		return false
	}
	if faces[len(faces)-1]-faces[0] != len(faces)-1 {
		return false
	}
	if faces[len(faces)-1] > 12 {
		return false
	}
	return count == 1 || (count == 2 && len(faces) >= 3)
}

// 新增规则函数
func (r _rules) StraightBoundary() (int, int) {
	return 1, 16
}

// 新增规则函数
func (r _rules) Reserved() bool {
	return r.reserved
}

// 新增规则：玩家出牌后下一位玩家必须出只大1个数的牌，接龙、炸弹、2或对2
func (r _rules) FollowsPrevious(cards []int, previous []int) bool {
	if len(cards) == 1 && len(previous) == 1 {
		if cards[0] == SpecialD && previous[0] == 2 {
			return true
		}
		return cards[0]-previous[0] == 1
	}
	if len(cards) == 2 && len(previous) == 2 && cards[0] == cards[1] && previous[0] == previous[1] {
		if cards[0] == SpecialD && previous[0] == 2 {
			return true
		}
		return cards[0]-previous[0] == 1
	}
	if len(cards) == 3 && len(previous) == 3 && cards[0] == cards[1] && cards[1] == cards[2] && previous[0] == previous[1] && previous[1] == previous[2] {
		return cards[0]-previous[0] == 1
	}
	return false
}

// 新增规则：当玩家出牌别人都要不起跳过后，出牌的人要随机从剩余牌组里再摸一张
func (r _rules) DrawCardOnPass(playersPassed int, totalPlayers int) bool {
	return playersPassed == totalPlayers-1
}

// 新增规则：发牌比例规则
func DealRatio(playerCount int) (totalCards int, cardsPerPlayer int, extraCards int) {
	totalCards = 55
	switch playerCount {
	case 4:
		cardsPerPlayer = 5
		extraCards = 1
	case 8:
		cardsPerPlayer = 5
		extraCards = 0
	default:
		// 默认情况，可以根据需要修改
		cardsPerPlayer = 5
		extraCards = 1
	}
	return totalCards, cardsPerPlayer, extraCards
}

// 新增规则：随机选择下一把的地主
func ChooseNextLandlord(players []int64) int64 {
	return players[rand.Intn(len(players))]
}
