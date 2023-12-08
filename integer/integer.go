package integer

func New(value int) *int {
	return &value
}

func NewUint(value uint) *uint {
	return &value
}

func NonNil(b *int) int {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func IsNil(b *int) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

func IsUnitNil(b *uint) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

func IsUnit32Nil(b *uint32) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}
