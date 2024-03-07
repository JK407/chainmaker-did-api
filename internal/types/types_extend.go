package types

type CommonResponse[T any] interface {
	BuildErrorResponse(code int, err error) T
	BuildSuccessResponse(code int, msg string, data T) T
	SetData(data T) T
}

func (n NewAccountResponse) BuildErrorResponse(code int, err error) NewAccountResponse {
	return NewAccountResponse{
		Code: code,
		Msg:  err.Error(),
	}
}

func (n NewAccountResponse) BuildSuccessResponse(code int, msg string, data NewAccountData) NewAccountResponse {
	return NewAccountResponse{
		Code: code,
		Msg:  msg,
		Data: &data,
	}
}

func (n NewAccountResponse) SetData(data NewAccountData) NewAccountResponse {
	return NewAccountResponse{
		Code: n.Code,
		Msg:  n.Msg,
		Data: &data,
	}
}
