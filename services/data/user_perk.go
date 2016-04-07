package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// UserPerkRedemp type
type UserPerkRedemp map[uint64]map[uint64]uint64

// PerkTermScore type
type PerkTermScore map[uint64]map[string]float64

// UserTermScore type
type UserTermScore map[uint64]map[string]float64

// CalcUserTermScore returns a UserTermScore from UserPerkRedemp and PerkTermScore
func CalcUserTermScore(upr UserPerkRedemp, pts PerkTermScore) UserTermScore {
	uts := make(UserTermScore)
	for uid, perkMap := range upr {
		// perkMap is perk_id: count
		for pid, rCount := range perkMap {
			termMap := pts[pid]
			// termMap is term: score
			for term, score := range termMap {
				userMap := uts[uid]
				if userMap == nil {
					userMap = make(map[string]float64)
				}
				// userMap is term: count_score
				userMap[term] += float64(rCount) * score
				uts[uid] = userMap
			}
		}
	}
	return uts
}

// Test CalcUserTermScore
func Test() {
	upr := make(UserPerkRedemp)
	upr[1] = make(map[uint64]uint64)
	upr[1][1] = 1
	upr[1][2] = 2

	upr[2] = make(map[uint64]uint64)
	upr[2][1] = 3
	upr[2][2] = 4

	pts := make(PerkTermScore)
	pts[1] = make(map[string]float64)
	pts[1]["1"] = 1
	pts[1]["2"] = 2

	pts[2] = make(map[string]float64)
	pts[2]["1"] = 3
	pts[2]["2"] = 4

	uts := CalcUserTermScore(upr, pts)
	fmt.Println(uts)
}

func calcFile(file string, redempMap UserPerkRedemp, multi int) UserPerkRedemp {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	// redempMap := make(UserPerkRedemp)

	// for i := 1; i <= 100; i++ {
	for {
		r, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("ERROR: ", err)
			return nil
		}

		uidS, pidS, cS := r[0], r[1], r[2]
		uid, _ := strconv.ParseUint(uidS, 10, 64)
		pid, _ := strconv.ParseUint(pidS, 10, 64)
		c, _ := strconv.ParseUint(cS, 10, 64)

		um := redempMap[uid]
		if um == nil {
			um = make(map[uint64]uint64)
		}
		um[pid] = c
		redempMap[uid] = um
	}
	return redempMap
}

// RedempData func
func RedempData() UserPerkRedemp {
	redemps := "/data/gocode/src/github.com/yiziz/collab/fixtures/redemps.csv"
	pageEvents := "/data/gocode/src/github.com/yiziz/collab/fixtures/pageevents.csv"
	// f, err := os.Open("/data/gocode/src/github.com/yiziz/collab/fixtures/redemps.csv")
	// f, err := os.Open("/data/gocode/src/github.com/yiziz/collab/fixtures/pageevents.csv")
	// f, err := os.Open("/data/gocode/src/github.com/yiziz/collab/fixtures/unweighted.csv")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// defer f.Close()
	//
	// reader := csv.NewReader(f)
	//
	redempMap := make(UserPerkRedemp)

	calcFile(redemps, redempMap, 10)
	calcFile(pageEvents, redempMap, 1)

	// for i := 1; i <= 100; i++ {
	// for {
	// 	r, err := reader.Read()
	//
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("ERROR: ", err)
	// 		return nil
	// 	}
	//
	// 	uidS, pidS, cS := r[0], r[1], r[2]
	// 	uid, _ := strconv.ParseUint(uidS, 10, 64)
	// 	pid, _ := strconv.ParseUint(pidS, 10, 64)
	// 	c, _ := strconv.ParseUint(cS, 10, 64)
	//
	// 	um := redempMap[uid]
	// 	if um == nil {
	// 		um = make(map[uint64]uint64)
	// 	}
	// 	um[pid] = c
	// 	redempMap[uid] = um
	// }
	return redempMap
}
