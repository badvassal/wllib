package action

import (
	"github.com/badvassal/wllib/gen"
)

// These constants are table indices.  The action table with a given index
// contains the specified type of data.
const (
	IDShop       = 6
	IDTransition = 10
)

// Tables is the set of action tables in a single MSQ block.
type Tables struct {
	T0          gen.Table
	T1          gen.Table
	T2          gen.Table
	T3          gen.Table
	T4          gen.Table
	T5          gen.Table
	T6          gen.Table
	T7          gen.Table
	T8          gen.Table
	T9          gen.Table
	Transitions []*Transition // Table 10.
	T11         gen.Table
	T12         gen.Table
	T13         gen.Table
	T14         gen.Table
	T15         gen.Table
}
