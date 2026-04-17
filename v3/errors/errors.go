package errors

type SmsutilsError struct {
	message         string
	driverName      string
	driverCode      string
	driverMessage   string
	driverRequestID string
	driverResponse  any
}

func (e *SmsutilsError) Error() string {
	return e.message
}

func (e *SmsutilsError) WithDriverName(driverName string) *SmsutilsError {
	e.driverName = driverName
	return e
}

func (e *SmsutilsError) DriverName() string {
	return e.driverName
}

func (e *SmsutilsError) WithDriverCode(code string) *SmsutilsError {
	e.driverCode = code
	return e
}

func (e *SmsutilsError) DriverCode() string {
	return e.driverCode
}

func (e *SmsutilsError) WithDriverMessage(message string) *SmsutilsError {
	e.driverMessage = message
	return e
}

func (e *SmsutilsError) DriverMessage() string {
	return e.driverMessage
}

func (e *SmsutilsError) WithDriverRequestID(requestID string) *SmsutilsError {
	e.driverRequestID = requestID
	return e
}

func (e *SmsutilsError) DriverRequestID() string {
	return e.driverRequestID
}

func (e *SmsutilsError) WithDriverResponse(driverResponse any) *SmsutilsError {
	e.driverResponse = driverResponse
	return e
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

	return err
}

var ErrDriverNotRegistered = New("driver not registered")
var ErrDriverCredentialInvalid = New("driver credential invalid")
var ErrDriverSendFailed = New("driver send failed")
