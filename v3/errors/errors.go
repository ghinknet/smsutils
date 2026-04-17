package errors

type SmsutilsError struct {
	message         string
	driverName      string
	driverCode      int
	driverMessage   string
	driverRequestID string
	driverResponse  any
}

func (e *SmsutilsError) Error() string {
	return e.message
}

func (e *SmsutilsError) WithDriverName(driverName string) {
	e.driverName = driverName
}

func (e *SmsutilsError) DriverName() string {
	return e.driverName
}

func (e *SmsutilsError) WithDriverCode(code int) {
	e.driverCode = code
}

func (e *SmsutilsError) DriverCode() int {
	return e.driverCode
}

func (e *SmsutilsError) WithDriverMessage(message string) {
	e.driverMessage = message
}

func (e *SmsutilsError) DriverMessage() string {
	return e.driverMessage
}

func (e *SmsutilsError) WithDriverRequestID(requestID string) {
	e.driverRequestID = requestID
}

func (e *SmsutilsError) DriverRequestID() string {
	return e.driverRequestID
}

func (e *SmsutilsError) WithDriverResponse(driverResponse any) {
	e.driverResponse = driverResponse
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

func WithDriverCode(driverCode int) Option {
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

func New(c string, options ...Option) error {
	err := &SmsutilsError{message: c}

	for _, option := range options {
		option(err)
	}

	return err
}

var ErrDriverNotRegistered = New("driver not registered")
var ErrDriverSendFailed = New("driver send failed")
