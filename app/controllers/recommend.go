package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/reiver/go-porterstemmer"
	"github.com/yiziz/collab/app/parsers"
	"github.com/yiziz/collab/services/data"
	"github.com/yiziz/collab/services/recommend"
)

// RecommendByUser action
func RecommendByUser(c *gin.Context) {
	// parse params
	params, err := parsers.RecommendByUser(c)
	if err != nil {
		log.Println(err)
		badRequest(c)
		return
	}
	pids := recommend.PerkByUser(params.UserID)
	jsonOK(c, pids)
}

// RecommendByPerk action
func RecommendByPerk(c *gin.Context) {
	// parse params
	params, err := parsers.RecommendByPerk(c)
	if err != nil {
		log.Println(err)
		badRequest(c)
		return
	}
	pids := recommend.PerkByPerk(params.PerkID)
	jsonOK(c, pids)
}

// RecommendByTerms action
func RecommendByTerms(c *gin.Context) {
	// parse params
	params, err := parsers.RecommendByTerms(c)
	if err != nil {
		log.Println(err)
		badRequest(c)
		return
	}

	termMap := make(map[string]float64)
	for _, term := range params.Terms {
		stemmedTerm := porterstemmer.StemString(term)
		termMap[stemmedTerm] = data.TermScores[stemmedTerm]
	}

	fmt.Println("TERMMAP")
	fmt.Println(termMap)

	pids := recommend.PerkByTermsThruPerk(termMap)
	pid_user := recommend.PerkByTerms(termMap)

	if len(pids) > 5 {
		pids = pids[:5]
	}

	if len(pid_user) > 5 {
		pid_user = pid_user[:5]
	}

	pids = append(pids, pid_user...)

	jsonOK(c, pids)
}
