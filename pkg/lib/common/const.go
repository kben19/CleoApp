package common

// Constant Messages
const (
	// Params
	IDParam        = "id"
	LimitParam     = "limit"
	OffsetParam    = "offset"
	OrderByParam   = "order_by"
	OrderTypeParam = "order_type"

	// Invalid Messages
	InvalidParamaterMsg          = "invalid parameter"
	InvalidIDParameterMsg        = "invalid id parameter"
	InvalidOffsetParameterMsg    = "invalid offset parameter"
	InvalidLowOffsetMsg          = "invalid offset smaller than 0"
	InvalidHighOffsetMsg         = "invalid offset higher than total"
	InvalidLimitParameterMsg     = "invalid limit parameter"
	InvalidOrderByParameterMsg   = "invalid order by parameter"
	InvalidOrderTypeParameterMsg = "invalid order type parameter"

	// Golang Messages
	FailedToMarshal       = "failed to marshal"
	FailedToWriteResponse = "failed to write response"
)

var (
	WhitelistOrderType = []string{"asc", "desc"}
)
