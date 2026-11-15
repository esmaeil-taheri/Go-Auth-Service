package httpmsg

import (
	"gameauth_v2/pkg/richerror"
	"net/http"
)

func Error(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)

		msg := re.Message()
		code := mapKinparamHTTPStatusCode(re.Kind())
		if code >= 500 {
			msg = "internal server error"
		}
		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}

}

func mapKinparamHTTPStatusCode(king richerror.Kind) int {
	switch king {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
