package models

type (
	// Spell describes a spell
	Spell struct {
		Name          string `json:"name,omitempty"`
		Level         string `json:"level,omitempty"`
		School        string `json:"school,omitempty"`
		Classes       string `json:"classes,omitempty"` // TODO make this an array one day, have to edit SQL
		AtHigherLevel string `json:"at_higher_level,omitempty"`
		CastingTime   string `json:"casting_time,omitempty"`
		Components    string `json:"components,omitempty"`
		Concentration string `json:"concentration,omitempty"`
		Description   string `json:"description,omitempty"`
		Duration      string `json:"duration,omitempty"`
		ID            string `json:"id,omitempty"`
		Page          string `json:"page,omitempty"`
		Range         string `json:"range,omitempty"`
		Ritual        string `json:"ritual,omitempty"`
		Slug          string `json:"slug,omitempty"`
		Source        string `json:"source,omitempty"`
	}

	// Spells is an array of Spell structs
	Spells []Spell
)
