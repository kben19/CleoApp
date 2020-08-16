package types

import (
	"github.com/kben19/CleoApp/pkg/lib/common"
	"github.com/kben19/CleoApp/pkg/lib/utils"
)

func (p StandardListParam) ValidateProductParam() error {
	whileListColumns := []string{"id", "article", "color", "price"}
	if err := p.ValidatePagination(); err != nil {
		return err
	}
	if !utils.FindMatchString(p.OrderBy, whileListColumns) {
		return common.ErrInvalidOrderByParameter
	}
	if !utils.FindMatchString(p.OrderType, common.WhitelistOrderType) {
		return common.ErrInvalidOrderTypeParameter
	}
	return nil
}

func (p StandardListParam) ValidatePagination() error {
	if p.Offset < 0 {
		return common.ErrInvalidLowOffsetParameter
	}
	if p.Limit <= 0 {
		return common.ErrInvalidLimitParameter
	}
	return nil
}
