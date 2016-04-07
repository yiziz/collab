package recommend

import (
	"fmt"

	"github.com/muesli/regommend"
	"github.com/yiziz/collab/services/data"
)

// RecMap by regommend Add as the second param
type RecMap map[interface{}]float64

// Bar blah
func Bar() {
	rm := data.RedempData()
	pm := data.PerkTermScores()
	um := data.CalcUserTermScore(rm, pm)

	terms := regommend.Table("terms")
	for uid, termMap := range um {
		rm := make(RecMap)
		for term, score := range termMap {
			rm[term] = float64(score)
		}
		terms.Add(uid, rm)
		// fmt.Println(uid, ": ", rm)
		if uid == 301171 || uid == 480566 {
			fmt.Println(uid, ": ", rm)
		}
	}

	fmt.Println(terms.Count())
	// fmt.Println(terms)

	for uid := range um {
		recs, _ := terms.Recommend(uid)
		// if len(recs) < 100 {
		if len(recs) != 0 {
			fmt.Println("recs: ", uid, " ", len(recs))
			// fmt.Println(recs)
		}
		neighbors, _ := terms.Neighbors(uid)
		if len(neighbors) != 0 {
			// if len(neighbors) < 100 {
			// 	fmt.Println(uid)
			fmt.Println("neighbors: ", uid, " ", len(neighbors))
			fmt.Println(neighbors[:10])
			// 	fmt.Println(neighbors)
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
