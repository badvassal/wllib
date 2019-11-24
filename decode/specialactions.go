package decode

// SpecialActinos represents an MSQ block's set of special actions (action class 6).
type SpecialActions struct {
	Actions []byte
}

// DecodeSpecialActions parses a set of special actions from a sequence of
// bytes.  baseOff is the offset of the start of the special actions relative
// to the start of the MSQ block's encoded section.
func DecodeSpecialActions(data []byte, baseOff int) (*SpecialActions, error) {
	return &SpecialActions{
		Actions: data,
	}, nil
}

// EncodeSpecialActions encodes the special actions sequence to a byte
// sequence.
func EncodeSpecialActions(sa SpecialActions) []byte {
	return sa.Actions
}
