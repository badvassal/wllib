package action

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

const (
	lootItemLen         = 2
	lootCashLen         = 3
	lootCashPrefixFixed = 0xde
	lootCashPrefixRand  = 0x5e
	lootItemTerminator  = 0xff
)

type LootItem struct {
	Fixed  bool // B0b7
	ID     int  // B0b0,6
	Amount int  // B1
}

type LootCash struct {
	Fixed  bool // B0
	Amount int  // B1,2
}

// Loot represents a loot bag definiton in an MSQ block.
type Loot struct {
	ToClass    int // B0
	ToSelector int // B1
	Items      []LootItem
	Cash       *LootCash
}

// DecodeLootItem decodes a loot bag item from a sequence of bytes.
func DecodeLootItem(data []byte) (*LootItem, error) {
	if len(data) < lootItemLen {
		return nil, wlerr.Errorf(
			"data length too short: have=%d want>=%d: data=%+v",
			len(data), lootItemLen, data)
	}

	item := &LootItem{}

	if data[0] == lootItemTerminator {
		return nil, wlerr.Errorf(
			"loot item has invalid value: have=%d want!=%d",
			lootItemTerminator, lootItemTerminator)
	}

	if data[0]&0x80 != 0 {
		item.Fixed = true
	}
	item.ID = int(data[0] & 0x7f)

	item.Amount = int(data[1])

	return item, nil
}

// DecodeLootCash decodes a loot bag cash element from a sequence of bytes.
func DecodeLootCash(data []byte) (*LootCash, error) {
	if len(data) < lootCashLen {
		return nil, wlerr.Errorf(
			"data length too short: have=%d want>=%d: data=%+v",
			len(data), lootCashLen, data)
	}

	cash := &LootCash{}

	switch data[0] {
	case lootCashPrefixFixed:
		cash.Fixed = true

	case lootCashPrefixRand:
		cash.Fixed = false

	default:
		return nil, wlerr.Errorf(
			"loot cash starts with invalid byte: "+
				"have=0x%02x want=0x%02x||0x%02x: data=%+v",
			data[0], lootCashPrefixFixed, lootCashPrefixRand, data)
	}

	amount, err := gen.ReadUint16(data[1:3])
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to read loot cash amount")
	}
	cash.Amount = amount

	return cash, nil
}

// decodeLootElem decodes a single element in a loot list (either an item or
// cash) from a sequence of bytes.  Only one the returned values can be
// non-nil, indicating the type of element decoded.  If both returned pointers
// are nil, the loot list terminator was encountered.  The returned int
// indicates the number of bytes that were read from the input buffer.
func decodeLootElem(data []byte) (*LootItem, *LootCash, int, error) {
	switch data[0] {
	case lootItemTerminator:
		return nil, nil, 1, nil

	case lootCashPrefixFixed, lootCashPrefixRand:
		cash, err := DecodeLootCash(data)
		if err != nil {
			return nil, nil, 0, err
		}
		return nil, cash, lootCashLen, nil

	default:
		item, err := DecodeLootItem(data)
		if err != nil {
			return nil, nil, 0, err
		}
		return item, nil, lootItemLen, nil
	}
}

// DecodeLoot decodes a loot bag from a sequence of bytes.  It returns the
// decoded loot and its length (in bytes).
func DecodeLoot(data []byte) (*Loot, int, error) {
	wrapErr := wlerr.MakeWrapper("failed to decode loot")

	if len(data) < 2 {
		return nil, 0, wrapErr(nil,
			"data length too short: have=%d want>=2", len(data))
	}

	loot := &Loot{}
	off := 0

	if data[off] > 0x0f {
		return nil, 0, wrapErr(nil,
			"loot contains invalid action class: have=0x%02x want<=0x0f",
			data[off])
	}
	loot.ToClass = int(data[off])
	off++

	loot.ToSelector = int(data[off])
	off++

	for {
		if off >= len(data) {
			return nil, 0, wrapErr(nil, "loot list missing terminator byte")
		}

		item, cash, sz, err := decodeLootElem(data[off:])
		if err != nil {
			return nil, 0, wrapErr(err, "")
		}

		off += sz

		if item != nil {
			loot.Items = append(loot.Items, *item)
		} else if cash != nil {
			if loot.Cash != nil {
				return nil, 0, wrapErr(nil,
					"loot list contains more than one cash element")
			}
			loot.Cash = cash
		} else {
			// List terminated.
			break
		}
	}

	return loot, off, nil
}

// DecodeLootTable decotes a set of loot bags from a table of byte buffers.
func DecodeLootTable(table gen.Table) ([]*Loot, error) {
	var loots []*Loot

	for i, elem := range table.Elems {
		if len(elem) == 0 {
			loots = append(loots, nil)
		} else {
			loot, _, err := DecodeLoot(elem)
			if err != nil {
				return nil, wlerr.Wrapf(err, "lootidx=%d", i)
			}
			loots = append(loots, loot)
		}
	}

	return loots, nil
}

// EncodeLootItem encodes a single loot bag item to a byte sequence.
func EncodeLootItem(item LootItem) []byte {
	b0 := byte(item.ID)
	if item.Fixed {
		b0 |= 0x80
	}

	return []byte{b0, byte(item.Amount)}
}

// EncodeLootCash encodes a loot bag cash object to a byte sequence.
func EncodeLootCash(cash LootCash) []byte {
	b := make([]byte, lootCashLen)

	if cash.Fixed {
		b[0] = lootCashPrefixFixed
	} else {
		b[0] = lootCashPrefixRand
	}

	copy(b[1:], gen.WriteUint16(uint16(cash.Amount)))

	return b
}

// EncodeLoot encodes a loot bag to a byte sequence.
func EncodeLoot(loot Loot) []byte {
	var b []byte

	b = append(b, byte(loot.ToClass))
	b = append(b, byte(loot.ToSelector))

	for _, item := range loot.Items {
		b = append(b, EncodeLootItem(item)...)
	}

	if loot.Cash != nil {
		b = append(b, EncodeLootCash(*loot.Cash)...)
	}

	b = append(b, lootItemTerminator)

	return b
}
