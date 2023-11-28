package errorcode

import (
	"slices"

	cs3rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

// FromCS3Status converts a CS3 status code into a corresponding local Error representation.
//
// It evaluates the provided CS3 status code and returns an equivalent graph Error.
// If the CS3 status code does not have a direct equivalent within the app,
// or is ignored, a general purpose Error is returned.
//
// This function is particularly useful when dealing with CS3 responses,
// and a unified error handling within the application is necessary.
func FromCS3Status(status *cs3rpc.Status, ignore ...cs3rpc.Code) *Error {
	err := &Error{errorCode: GeneralException, msg: "unspecified error has occurred"}

	if status != nil {
		err.msg = status.GetMessage()
	}

	code := status.GetCode()
	switch {
	case slices.Contains(ignore, status.GetCode()):
		fallthrough
	case code == cs3rpc.Code_CODE_OK:
		err = nil
	case code == cs3rpc.Code_CODE_NOT_FOUND:
		err.errorCode = ItemNotFound
	case code == cs3rpc.Code_CODE_PERMISSION_DENIED:
		err.errorCode = AccessDenied
	case code == cs3rpc.Code_CODE_UNAUTHENTICATED:
		err.errorCode = Unauthenticated
	case code == cs3rpc.Code_CODE_INVALID_ARGUMENT:
		err.errorCode = InvalidRequest
	case code == cs3rpc.Code_CODE_ALREADY_EXISTS:
		err.errorCode = NameAlreadyExists
	case code == cs3rpc.Code_CODE_FAILED_PRECONDITION:
		err.errorCode = PreconditionFailed
	case code == cs3rpc.Code_CODE_UNIMPLEMENTED:
		err.errorCode = NotSupported
	}

	return err
}

// FromStat transforms a *provider.StatResponse object and an error into an *Error.
//
// It takes a stat of type *provider.StatResponse, an error, and a variadic parameter of type cs3rpc.Code.
// If the error is not nil, it creates an Error object with the error message and a GeneralException code.
// If the error is nil, it invokes the FromCS3Status function with the StatResponse Status and the ignore codes.
func FromStat(stat *provider.StatResponse, err error, ignore ...cs3rpc.Code) *Error {
	switch {
	case err != nil:
		return &Error{msg: err.Error(), errorCode: GeneralException}
	default:
		return FromCS3Status(stat.GetStatus(), ignore...)
	}
}
