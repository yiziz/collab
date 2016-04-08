package recommend

import (
	"fmt"

	"github.com/muesli/regommend"
	"github.com/yiziz/collab/services/data"
)

// RecMap by regommend Add as the second param
type RecMap map[interface{}]float64

var Rm data.UserPerkRedemp
var Pm data.PerkTermScore
var Um data.UserTermScore
var Terms *regommend.RegommendTable
var PerkTerms *regommend.RegommendTable

// Bar blah
func Bar() {
	Rm = data.RedempData()
	Pm = data.PerkTermScores()
	Um = data.CalcUserTermScore(Rm, Pm)

	Terms = regommend.Table("user_terms")

	for uid, termMap := range Um {
		rm := make(RecMap)
		for term, score := range termMap {
			rm[term] = float64(score)
		}
		Terms.Add(uid, rm)
	}

	PerkTerms = regommend.Table("perk_terms")

	for perkId, termMap := range Pm {
		rm := make(RecMap)
		for term, score := range termMap {
			rm[term] = float64(score)
		}
		PerkTerms.Add(perkId, rm)
	}

	fmt.Println(Terms.Count())
	fmt.Println(PerkTerms.Count())

	//fmt.Println("derp recoommendsations:", PerkByTermsThruPerk(map[string]float64{"ticket": 1081.4701543587407, "entertain": 0, "concert": 52.09396263715782}))
	// fmt.Println(terms)

	// Check equinox related perks
	// fmt.Println("crunch RECOMMENDATIONS:", PerkByPerk(2788))
	//
	// for uid := range Um {
	// 	recs, _ := Terms.Recommend(uid)
	// 	// if len(recs) < 100 {
	// 	if len(recs) != 0 {
	// 		//fmt.Println("recs: ", uid, " ", len(recs))
	// 		//fmt.Println(recs[:10])
	// 		// fmt.Println(recs)
	// 	}
	//
	// 	perks := PerkByUser(uid)
	// 	fmt.Println(perks)
	// }
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
