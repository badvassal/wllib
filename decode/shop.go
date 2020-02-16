package decode

// ShopData represents an MSQ block's set of shops.
type ShopData struct {
	Shops []byte
}

// DecodeShopData parses a set of shops from a sequence of bytes.  baseOff is
// the offset of the start of the shop data relative to the start of the MSQ
// block's secure section.
func DecodeShopData(data []byte, baseOff int) (*ShopData, error) {
	return &ShopData{
		Shops: data,
	}, nil
}
