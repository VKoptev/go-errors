package errors

import "errors"

// Is is alias for builtin errors.Is. It implements such interface as OneOf and EachOf.
// IMPORTANT! Is uses only first error from variadic argument.
func Is(err error, target ...error) bool {
	if len(target) == 0 {
		return err == nil
	}

	return errors.Is(err, target[0])
}

// As is alias for builtin errors.As. It's needed to do not import default library.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// OneOf reports whether any error in err's chain matches at least one of target errors.
func OneOf(err error, target ...error) bool {
	for _, e := range target {
		if errors.Is(err, e) {
			return true
		}
	}

	return false
}

// EachOf reports whether any error in err's chain matches each of target errors.
func EachOf(err error, target ...error) bool {
	for _, e := range target {
		if !errors.Is(err, e) {
			return false
		}
	}

	return true
}
