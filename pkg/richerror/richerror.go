package richerror

type Kind int
type Op string

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation Op
	wrapperError error
	message string
	kind Kind
	meta map[string]interface{}
}

func New(op Op) RichError {
	return RichError{operation: op}
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op 
	return r
}

func (r RichError) WithErr(err error) RichError {
	r.wrapperError = err
	return r
}

func (r RichError) WithMsg(msg string) RichError {
	r.message = msg
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta
	return r
}

func (r RichError) Error() string {
	return r.message
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	re, ok := r.wrapperError.(RichError)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrapperError.(RichError)
	if !ok {
		return r.wrapperError.Error()
	}

	return re.Message()

}