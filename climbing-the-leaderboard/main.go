package main

import (
	"fmt"
	"sort"
)

/*
 * [Medium]
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	scoreMap := make(map[int32]int32)
	var index int32
	for _, r := range ranked {
		if _, ok := scoreMap[r]; !ok {
			index++
			scoreMap[r] = index
		}
	}
	var scores []int
	for k := range scoreMap {
		scores = append(scores, int(k))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	fmt.Printf("%15s: %v\n", "scores", scores)
	fmt.Printf("%15s: %v\n", "player", player)

	var playerRanks []int32

	for _, score := range player {
		for len(scores) > 0 && score >= int32(scores[len(scores)-1]) {
			scores = scores[:len(scores)-1]
		}
		playerRanks = append(playerRanks, int32(len(scores)+1))
	}

	return playerRanks
}

func main() {
	ranked := []int32{100, 100, 50, 40, 40, 20, 10}
	player := []int32{5, 25, 50, 120}

	fmt.Println(climbingLeaderboard(ranked, player))
}
