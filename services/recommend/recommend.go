package recommend

import (
	"fmt"
	"sort"

	"github.com/yiziz/collab/services/data"
)

// PerkByUser returns perk id recommendations using uid
func PerkByUser(uid uint64) []uint64 {
	var perks []uint64
	neighbors, _ := Terms.Neighbors(uid)

	fmt.Println("NEIGHBORS LENGTH: ", len(neighbors))

	if len(neighbors) != 0 {
		unseenPerks := make(map[uint64]float64)

		for _, neighbor := range neighbors[:5] {
			nuid := neighbor.Key.(uint64)
			score := neighbor.Distance

			if score > 0.85 {
				for perkId, perkNeighborScore := range Rm[nuid] {
					youScore, ok := Rm[uid][perkId]
					unseenPerks[perkId] += float64(int64(perkNeighborScore))

					if ok {
						unseenPerks[perkId] -= float64(int64(youScore))
					}
				}
			}
		}

		for perkId, _ := range Rm[uid] {
			delete(unseenPerks, perkId)
		}

		var scoreArray data.WordScoreArray
		for perkId, score := range unseenPerks {
			scoreObj := new(data.WordScore)
			scoreObj.PerkID = perkId
			scoreObj.Score = score

			scoreArray = append(scoreArray, scoreObj)
		}

		sort.Stable(sort.Reverse(scoreArray))

		for _, score := range scoreArray {
			fmt.Print("{id:", score.PerkID, ", score:", score.Score, "}")
			perks = append(perks, score.PerkID)
		}
		fmt.Println("Id:", uid, "data:", Rm[uid])
	}

	if len(perks) > 10 {
		perks = perks[:10]
	}
	return perks
}

// PerkByPerk returns perk id recommendations using pid
func PerkByPerk(pid uint64) []uint64 {
	var perks []uint64
	neighbors, err := PerkTerms.Neighbors(pid)

	if err != nil {
		fmt.Println(err)
	}

	for _, neighbor := range neighbors[:5] {
		perks = append(perks, neighbor.Key.(uint64))
	}

	if len(perks) > 10 {
		perks = perks[:10]
	}
	return perks
}

// PerkByTerms returns perk id recommendations using terms (sl)
func PerkByTerms(sl map[string]float64) []uint64 {
	newUserId := uint64(99999)

	recMap := make(RecMap)
	for term, score := range sl {
		recMap[term] = score
	}

	Terms.Add(newUserId, recMap)
	perks := PerkByUser(newUserId)
	fmt.Println("PPPPPPPPP")
	fmt.Println(recMap)
	fmt.Println(perks)
	Terms.Delete(newUserId)

	if len(perks) > 10 {
		perks = perks[:10]
	}
	return perks
}
