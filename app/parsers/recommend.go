package parsers

import "github.com/gin-gonic/gin"

// RecommendByUserParams holds params for RecommendByUser action
type RecommendByUserParams struct {
	UserID uint64 `json:"userId" binding:"required"`
}

// RecommendByPerkParams holds params for RecommendByPerk action
type RecommendByPerkParams struct {
	PerkID uint64 `json:"perkId" binding:"required"`
}

// RecommendByTermsParams holds params for RecommendByTerms action
type RecommendByTermsParams struct {
	Terms []string `json:"terms" binding:"required"`
}

// RecommendByUser parses request params into RecommendByUserParams struct
func RecommendByUser(c *gin.Context) (*RecommendByUserParams, error) {
	p := new(RecommendByUserParams)
	if err := parse(c, p); err != nil {
		return nil, err
	}
	return p, nil
}

// RecommendByPerk parses request params into RecommendByPerkParams struct
func RecommendByPerk(c *gin.Context) (*RecommendByPerkParams, error) {
	p := new(RecommendByPerkParams)
	if err := parse(c, p); err != nil {
		return nil, err
	}
	return p, nil
}

// RecommendByTerms parses request params into RecommendByTermsParams struct
func RecommendByTerms(c *gin.Context) (*RecommendByTermsParams, error) {
	p := new(RecommendByTermsParams)
	if err := parse(c, p); err != nil {
		return nil, err
	}
	return p, nil
}
