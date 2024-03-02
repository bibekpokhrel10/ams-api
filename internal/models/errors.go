package models

import "errors"

var (
	ErrEmailRequired            = errors.New("required email")
	ErrFirstNameRequired        = errors.New("required firstname")
	ErrNameRequired             = errors.New("required name")
	ErrUsernameRequired         = errors.New("required username")
	ErrInvalidEmail             = errors.New("invalid email")
	ErrInvalidUsername          = errors.New("username must be alpha-numeric character only")
	ErrPasswordRequired         = errors.New("required password")
	ErrParentIDRequired         = errors.New("required parent id")
	ErrRequiredOldPassword      = errors.New("required old passowd")
	ErrSamePasswordAsOld        = errors.New("required different than old passowd")
	ErrRequiredNewPassword      = errors.New("required new passowd")
	ErrRequiredConfrimPassword  = errors.New("required Conformation passowd")
	ErrPasswordMinLength        = errors.New("password should be atleast 8 characters")
	ErrPasswordUpperCaseLetter  = errors.New("password must contain at least one uppercase letter")
	ErrPasswordLowerCaseLetter  = errors.New("password must contain at least one lowercase letter")
	ErrPasswordSpecialCharacter = errors.New("password must contain at least one special character")
	ErrPasswordOneDigit         = errors.New("password must contain at least one digit")
	ErrPasswordMismatch         = errors.New("password provided do not match")
	ErrPasswordHashGenerate     = errors.New("unable to generate hash password")
	ErrInvalidEntity            = errors.New("invalid model")

	//Transaction Pin Error
	ErrInvalidTxnPinRequest = errors.New("invalid request for transaction pin")
	ErrRequredOldTxnPin     = errors.New("required old transaction pin")
	ErrRequredNewTxnPin     = errors.New("required new transaction pin")
	ErrRequredConfirmTxnPin = errors.New("required confirm transaction pin")
	ErrRequred4Digits       = errors.New("required 4 digits pin")
	ErrInvalidNewTxnPin     = errors.New("invalid new_txn_pin")
	ErrInvalidConfirmTxnPin = errors.New("invalid confirm_txn_pin")
	ErrTxnPinMustNotSame    = errors.New("new and confirm transaction pin must be same")
	ErrInvalidTxnPin        = errors.New("invalid transaction pin format")
	ErrOldAndNewTxnPinSame  = errors.New("old and new txn pin must be different")

	//Provider Error
	ErrRequiredProviderName  = errors.New("required provider name")
	ErrInvalidProviderName   = errors.New("invalid provider name")
	ErrRequiredAPIKey        = errors.New("required provider api key")
	ErrRequiredAPISecretKey  = errors.New("required provider api secret key")
	ErrRequiredAPILicenseKey = errors.New("required provider api license key")
	ErrRequiredAPIUrl        = errors.New("required provider api url")

	//Theme Error
	ErrRequiredThemeName = errors.New("required theme name")

	//Notice Error
	ErrRequiredNoticeTitle   = errors.New("required notice title")
	ErrRequiredNoticeContent = errors.New("required notice content")
	ErrRequiredNoticeType    = errors.New("required notice type")

	//Transfer
	ErrNilTransferRequest   = errors.New("transaction request is nil")
	ErrRequiredAmount       = errors.New("required transfer amount value")
	ErrRequiredAction       = errors.New("required transfer action")
	ErrRequiredTxnInitiator = errors.New("required transfer initiator")
	ErrRequiredTxnPin       = errors.New("required transaction pin")
)
