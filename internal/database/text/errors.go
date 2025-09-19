package text

import (
	"errors"
)

var (
	// ErrNotConnected is the not connected error.
	ErrNotConnected = errors.New(`not connected`)
	// ErrNoSuchFileOrDirectory is the no such file or directory error.
	ErrNoSuchFileOrDirectory = errors.New(`no such file or directory`)
	// ErrCannotIncludeDirectories is the cannot include directories error.
	ErrCannotIncludeDirectories = errors.New(`cannot include directories`)
	// ErrMissingDSN is the missing dsn error.
	ErrMissingDSN = errors.New(`missing dsn`)
	// ErrNoPreviousTransactionExists is the no previous transaction exists error.
	ErrNoPreviousTransactionExists = errors.New(`no previous transaction exists`)
	// ErrPreviousTransactionExists is the previous transaction exists error.
	ErrPreviousTransactionExists = errors.New(`previous transaction exists`)
	// ErrPasswordAttemptsExhausted is the exhausted password attempts error.
	ErrPasswordAttemptsExhausted = errors.New(`password attempts exhausted`)
	// ErrSingleTransactionCannotBeUsedWithInteractiveMode is the single transaction cannot be used with interactive mode error.
	ErrSingleTransactionCannotBeUsedWithInteractiveMode = errors.New(`--single-transaction cannot be used with interactive mode`)
	// ErrNoEditorDefined is the no editor defined error.
	ErrNoEditorDefined = errors.New(`no editor defined`)
	// ErrUnknownCommand is the unknown command error.
	ErrUnknownCommand = errors.New(`unknown command`)
	// ErrMissingRequiredArgument is the missing required argument error.
	ErrMissingRequiredArgument = errors.New(`missing required argument`)
	// ErrDriverNotAvailable is the driver not available error.
	ErrDriverNotAvailable = errors.New(`driver not available`)
	// ErrPasswordNotSupportedByDriver is the password not supported by driver error.
	ErrPasswordNotSupportedByDriver = errors.New(`\password not supported by driver`)
	// ErrUnterminatedQuotedString is the unterminated quoted string error.
	ErrUnterminatedQuotedString = errors.New(`unterminated quoted string`)
	// ErrNoShellAvailable is the no SHELL available error.
	ErrNoShellAvailable = errors.New(`no SHELL available`)
	// ErrNotInteractive is the not interactive error.
	ErrNotInteractive = errors.New(`not interactive`)
	// ErrInvalidType is the invalid type error.
	ErrInvalidType = errors.New(`invalid -TYPE: TYPE must be password, string, int, uint, float, or bool`)
	// ErrInvalidIdentifier is the invalid identifier error.
	ErrInvalidIdentifier = errors.New(`invalid identifier`)
	// ErrInvalidValue is the invalid value error.
	ErrInvalidValue = errors.New(`invalid value`)
	// ErrTooManyRows is the too many rows error.
	ErrTooManyRows = errors.New(`too many rows`)
	// ErrInvalidFormatType is the invalid format type error.
	ErrInvalidFormatType = errors.New(`\pset: allowed formats are unaligned, aligned, wrapped, html, asciidoc, latex, latex-longtable, troff-ms, json, csv`)
	// ErrInvalidFormatPagerType is the invalid format pager error.
	ErrInvalidFormatPagerType = errors.New(`\pset: allowed pager values are on, off, always`)
	// ErrInvalidFormatExpandedType is the invalid format expanded error.
	ErrInvalidFormatExpandedType = errors.New(`\pset: allowed expanded values are on, off, auto`)
	// ErrInvalidFormatLineStyle is the invalid format line style error.
	ErrInvalidFormatLineStyle = errors.New(`\pset: allowed line styles are ascii, old-ascii, unicode`)
	// ErrInvalidFormatBorderLineStyle is the invalid format border line style error.
	ErrInvalidFormatBorderLineStyle = errors.New(`\pset: allowed Unicode border line styles are single, double`)
	// ErrInvalidTimezoneLocation is the invalid timezone location error.
	ErrInvalidTimezoneLocation = errors.New(`\pset: invalid timezone location`)
	// ErrGraphicsNotSupported is the graphics not supported error.
	ErrGraphicsNotSupported = errors.New(`\chart: graphics not supported in terminal`)
	// ErrNoNumericColumns is the no numeric columns error.
	ErrNoNumericColumns = errors.New(`\chart: no numeric columns found`)
	// ErrInvalidQuotedString is the invalid quoted string error.
	ErrInvalidQuotedString = errors.New(`invalid quoted string`)
	// ErrInvalidFormatOption is the invalid format option error.
	ErrInvalidFormatOption = errors.New(`invalid format option`)
	// ErrInvalidWatchDuration is the invalid watch duration error.
	ErrInvalidWatchDuration = errors.New(`invalid watch duration`)
	// ErrUnableToNormalizeURL is the unable to normalize URL error.
	ErrUnableToNormalizeURL = errors.New(`unable to normalize URL`)
	// ErrInvalidIsolationLevel is the invalid isolation level error.
	ErrInvalidIsolationLevel = errors.New(`invalid isolation level`)
	// ErrNotSupported is the not supported error.
	ErrNotSupported = errors.New(`not supported`)
	// ErrWrongNumberOfArguments is the wrong number of arguments error.
	ErrWrongNumberOfArguments = errors.New(`wrong number of arguments`)
	// ErrUnknownFileType is the unknown file type error.
	ErrUnknownFileType = errors.New(`unknown file type`)
	// ErrNamedConnectionIsNotAURL is the named connection is not a url error.
	ErrNamedConnectionIsNotAURL = errors.New(`named connection is not a url`)
	// ErrInvalidConfig is the invalid config error.
	ErrInvalidConfig = errors.New(`invalid config`)
	// ErrIfEscaped is the if escaped error.
	ErrIfEscaped = errors.New(`\if escaped`)
	// ErrEndIfNoMatchingIf is the endif no matching if error.
	ErrEndIfNoMatchingIf = errors.New(`\endif: no matching \if`)
)
