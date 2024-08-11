package serializer

type Error struct {
	Error   string `json:"error,omitempty"`
	Details string `json:"details,omitempty"`
}
