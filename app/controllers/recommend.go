package controllers

import (
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
	pids := recommend.PerkByTerms(params.Terms)
	// pids = []uint64{98}
	jsonOK(c, pids)
}
