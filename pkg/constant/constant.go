package constant

const (
	XFF             = "X-Forwarded-For"
	UA              = "User-Agent"
	Ticket          = "ticket"
	ValidateService = "validate_service"
	SubsystemAlias  = "ha"
	Renew           = "renew"
	SSOAccessToken  = "access-token"
	SSOSubsystem    = "subsystem"
)

const (
	TimeTTLDefault    = 600
	MinHttpDnsTimeTTL = 30

	TimeSecondsWithOneMinute = 60
	TimeSecondsWithOneHour   = 60 * TimeSecondsWithOneMinute
	TimeSecondsWithOneDay    = 24 * TimeSecondsWithOneHour
	TimeSecondsWithOneWeek   = 7 * TimeSecondsWithOneDay
	TimeSecondsWithOneMonth  = 30 * TimeSecondsWithOneDay
	TimeSecondsWithOneYear   = 365 * TimeSecondsWithOneDay
)

const (
	DependenceVerify          = "DEPENDENCE_VERIFY"
	DoNotEnableAuthentication = "DO_NOT_ENABLE_AUTHENTICATION"
)

const (
	UserContextName = "user"
)
