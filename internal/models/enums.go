package models

// FormStatus enum
type FormStatus string

const (
	FormStatusDraft     FormStatus = "DRAFT"
	FormStatusPublished FormStatus = "PUBLISHED"
	FormStatusClosed    FormStatus = "CLOSED"
	FormStatusArchived  FormStatus = "ARCHIVED"
)

// FormPageType enum
type FormPageType string

const (
	FormPageTypeSingle  FormPageType = "SINGLE"
	FormPageTypeStepper FormPageType = "STEPPER"
)

// ResponseStatus enum
type ResponseStatus string

const (
	ResponseStatusStarted   ResponseStatus = "STARTED"
	ResponseStatusCompleted ResponseStatus = "COMPLETED"
	ResponseStatusAbandoned ResponseStatus = "ABANDONED"
	ResponseStatusPartial   ResponseStatus = "PARTIAL"
)

// TransactionType enum
type TransactionType string

const (
	TransactionTypeCredit   TransactionType = "CREDIT"
	TransactionTypeDebit    TransactionType = "DEBIT"
	TransactionTypeTransfer TransactionType = "TRANSFER"
	TransactionTypeResponse TransactionType = "RESPONSE"
	TransactionTypeRefund   TransactionType = "REFUND"
	TransactionTypeAiUsage  TransactionType = "AI_USAGE"
	TransactionTypeGift     TransactionType = "GIFT"
)

// AnalyticsEventType enum
type AnalyticsEventType string

const (
	AnalyticsEventTypeFormOpened   AnalyticsEventType = "FORM_OPENED"
	AnalyticsEventTypeFormStarted  AnalyticsEventType = "FORM_STARTED"
	AnalyticsEventTypeFormSubmitted AnalyticsEventType = "FORM_SUBMITTED"
)

// QuestionType enum
type QuestionType string

const (
	QuestionTypeTextShort         QuestionType = "TEXT_SHORT"
	QuestionTypeTextLong          QuestionType = "TEXT_LONG"
	QuestionTypeNumber            QuestionType = "NUMBER"
	QuestionTypeDate              QuestionType = "DATE"
	QuestionTypeFileAny           QuestionType = "FILE_ANY"
	QuestionTypeFileImageOrVideo  QuestionType = "FILE_IMAGE_OR_VIDEO"
	QuestionTypeChoiceSingle      QuestionType = "CHOICE_SINGLE"
	QuestionTypeChoiceMultiple    QuestionType = "CHOICE_MULTIPLE"
	QuestionTypeChoiceDropdown    QuestionType = "CHOICE_DROPDOWN"
	QuestionTypeChoicePicture     QuestionType = "CHOICE_PICTURE"
	QuestionTypeChoiceCheckbox    QuestionType = "CHOICE_CHECKBOX"
	QuestionTypeChoiceBool        QuestionType = "CHOICE_BOOL"
	QuestionTypeInfoEmail         QuestionType = "INFO_EMAIL"
	QuestionTypeInfoPhone         QuestionType = "INFO_PHONE"
	QuestionTypeInfoUrl           QuestionType = "INFO_URL"
	QuestionTypeUserDetail        QuestionType = "USER_DETAIL"
	QuestionTypeUserAddress       QuestionType = "USER_ADDRESS"
	QuestionTypeScreenWelcome     QuestionType = "SCREEN_WELCOME"
	QuestionTypeScreenEnd         QuestionType = "SCREEN_END"
	QuestionTypeScreenStatement   QuestionType = "SCREEN_STATEMENT"
	QuestionTypeRatingZeroToTen   QuestionType = "RATING_ZERO_TO_TEN"
	QuestionTypeRatingStar        QuestionType = "RATING_STAR"
	QuestionTypeRatingRank        QuestionType = "RATING_RANK"
	QuestionTypeLegal             QuestionType = "LEGAL"
	QuestionTypeRedirectToUrl     QuestionType = "REDIRECT_TO_URL"
)

// CouponDiscountType enum
type CouponDiscountType string

const (
	CouponDiscountTypeFlat    CouponDiscountType = "FLAT"
	CouponDiscountTypePercent CouponDiscountType = "PERCENT"
)

// CouponType enum
type CouponType string

const (
	CouponTypeGeneral    CouponType = "GENERAL"
	CouponTypePlanBased  CouponType = "PLAN_BASED"
)

// TeamSubscriptionStatus enum
type TeamSubscriptionStatus string

const (
	TeamSubscriptionStatusActive  TeamSubscriptionStatus = "ACTIVE"
	TeamSubscriptionStatusGrace   TeamSubscriptionStatus = "GRACE"
	TeamSubscriptionStatusBlocked TeamSubscriptionStatus = "BLOCKED"
)