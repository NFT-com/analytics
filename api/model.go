package api

// Filter contains all of the supported query parameters for all endpoints.
type Filter struct {
	Chain       string `query:"chain"`
	Collection  string `query:"collection"`
	Marketplace string `query:"marketplace"`
	TokenID     string `query:"token_id"`
	Start       string `query:"start"` // FIXME: Change these two types
	End         string `query:"end"`
}

// FIXME: Validate format if parameters are set.
// FIXME: Think of event-specific filters.
// FIXME: Add filters for start and end height.
