package models

// Perk model
type Perk struct {
	Model
	PerkID      uint64
	TagLine     string
	Description string
}
