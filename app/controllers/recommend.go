package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yiziz/collab/app/parsers"
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
	mod := 10.0
	for _, term := range params.Terms {
		termMap[term] = mod
		mod /= 1.5
	}

	fmt.Println("TERMMAP")
	fmt.Println(termMap)

	pids := recommend.PerkByTerms(termMap)
	jsonOK(c, pids)
}
