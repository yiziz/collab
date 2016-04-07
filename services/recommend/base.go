package recommend

import (
	"fmt"
	"sort"

	"github.com/muesli/regommend"
	"github.com/yiziz/collab/services/data"
)

// RecMap by regommend Add as the second param
type RecMap map[interface{}]float64

var Rm data.UserPerkRedemp
var Pm data.PerkTermScore
var Um data.UserTermScore
var Terms *regommend.RegommendTable

// Bar blah
func Bar() {
	Rm = data.RedempData()
	Pm = data.PerkTermScores()
	Um = data.CalcUserTermScore(Rm, Pm)

	Terms = regommend.Table("terms")

	for uid, termMap := range Um {
		rm := make(RecMap)
		for term, score := range termMap {
			rm[term] = float64(score)
		}
		Terms.Add(uid, rm)
	}

	fmt.Println(Terms.Count())
	// fmt.Println(terms)

	for uid := range um {
		recs, _ := Terms.Recommend(uid)
		// if len(recs) < 100 {
		if len(recs) != 0 {
			//fmt.Println("recs: ", uid, " ", len(recs))
			//fmt.Println(recs[:10])
			// fmt.Println(recs)
		}
		neighbors, _ := Terms.Neighbors(uid)
		if len(neighbors) != 0 {
			// if len(neighbors) < 100 {
			// 	fmt.Println(uid)
			//fmt.Println("neighbors: ", uid, " ", len(neighbors))
			fmt.Println(neighbors[:10])
			// 	fmt.Println(neighbors)

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
			}
			fmt.Println("Id:", uid, "data:", Rm[uid])
		}
	}
	fmt.Println("end")
}

// Foo blah
func Foo() {
	dataMap := data.RedempData()

	// Accessing a new regommend table for the first time will create it.
	perks := regommend.Table("perks")

	for uid, perkMap := range dataMap {
		rm := make(RecMap)
		for pid, count := range perkMap {
			rm[pid] = float64(count)
		}
		perks.Add(uid, rm)
	}

	fmt.Println(perks.Count())
	// recs, _ := perks.Recommend(39)
	// fmt.Println(recs)
	// for _, rec := range recs {
	// 	fmt.Println("Recommending", rec.Key, "with score", rec.Distance)
	// }

	// neighbors, _ := perks.Neighbors(39)
	// fmt.Println(neighbors)
	for uid := range dataMap {
		recs, _ := perks.Recommend(uid)
		if len(recs) < 100 {
			fmt.Println("recs: ", uid, " ", len(recs))
			// fmt.Println(recs)
		}
		neighbors, _ := perks.Neighbors(uid)
		if len(neighbors) < 100 {
			// 	fmt.Println(uid)
			fmt.Println("neighbors: ", uid, " ", len(neighbors))
			// 	fmt.Println(neighbors)
		}
	}
	fmt.Println("end")
}
