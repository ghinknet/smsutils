package errors

import "go.gh.ink/toolbox/pointer"

type SmsutilsError struct {
	message         string
	driverName      string
	driverCode      string
	driverMessage   string
	driverRequestID string
	driverResponse  any

	raw error
}

func (e *SmsutilsError) Error() string {
	return e.message
}

func (e *SmsutilsError) Is(err error) bool {
	return e.raw == err
}

func (e *SmsutilsError) Unwrap() error {
	return e.raw
}

func (e *SmsutilsError) WithDriverName(driverName string) *SmsutilsError {
	ne := pointer.Copy(e)
	ne.driverName = driverName
	return ne
}

func (e *SmsutilsError) DriverName() string {
	return e.driverName
}

func (e *SmsutilsError) WithDriverCode(code string) *SmsutilsError {
	ne := pointer.Copy(e)
	ne.driverCode = code
	return ne
}

func (e *SmsutilsError) DriverCode() string {
	return e.driverCode
}

func (e *SmsutilsError) WithDriverMessage(message string) *SmsutilsError {
	ne := pointer.Copy(e)
	ne.driverMessage = message
	return ne
}

func (e *SmsutilsError) DriverMessage() string {
	return e.driverMessage
}

func (e *SmsutilsError) WithDriverRequestID(requestID string) *SmsutilsError {
	ne := pointer.Copy(e)
	ne.driverRequestID = requestID
	return ne
}

func (e *SmsutilsError) DriverRequestID() string {
	return e.driverRequestID
}

func (e *SmsutilsError) WithDriverResponse(driverResponse any) *SmsutilsError {
	ne := pointer.Copy(e)
	ne.driverResponse = driverResponse
	return ne
}

func (e *SmsutilsError) DriverResponse() any {
	return e.driverResponse
}

type Option func(*SmsutilsError)

func WithDriverName(driverName string) Option {
	return func(e *SmsutilsError) {
		e.driverName = driverName
	}
}

func WithDriverCode(driverCode string) Option {
	return func(e *SmsutilsError) {
		e.driverCode = driverCode
	}
}

func WithDriverMessage(driverMessage string) Option {
	return func(e *SmsutilsError) {
		e.driverMessage = driverMessage
	}
}

func WithDriverRequestID(driverRequestID string) Option {
	return func(e *SmsutilsError) {
		e.driverRequestID = driverRequestID
	}
}

func WithDriverResponse(driverResponse any) Option {
	return func(e *SmsutilsError) {
		e.driverResponse = driverResponse
	}
}

func New(c string, options ...Option) *SmsutilsError {
	err := &SmsutilsError{message: c}

	for _, option := range options {
		option(err)
	}

	err.raw = err

	return err
}

var ErrDriverNotRegistered = New("driver not registered")
var ErrDriverCredentialInvalid = New("driver credential invalid")
var ErrDriverSendFailed = New("driver send failed")
