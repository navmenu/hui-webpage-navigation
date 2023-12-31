// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package webnavigation

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNKNOWN.String() && e.Code == 400
}

func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_UNKNOWN.String(), fmt.Sprintf(format, args...))
}

func IsMissingParam(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MISSING_PARAM.String() && e.Code == 400
}

func ErrorMissingParam(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_MISSING_PARAM.String(), fmt.Sprintf(format, args...))
}

func IsBadParam(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BAD_PARAM.String() && e.Code == 400
}

func ErrorBadParam(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_BAD_PARAM.String(), fmt.Sprintf(format, args...))
}

func IsAlreadyExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ALREADY_EXIST.String() && e.Code == 400
}

func ErrorAlreadyExist(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ALREADY_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_EXIST.String() && e.Code == 404
}

func ErrorNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsDbError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DB_ERROR.String() && e.Code == 500
}

func ErrorDbError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DB_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAdminNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADMIN_NOT_FOUND.String() && e.Code == 401
}

func ErrorAdminNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_ADMIN_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsAdminNotCreated(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADMIN_NOT_CREATED.String() && e.Code == 204
}

func ErrorAdminNotCreated(format string, args ...interface{}) *errors.Error {
	return errors.New(204, ErrorReason_ADMIN_NOT_CREATED.String(), fmt.Sprintf(format, args...))
}

func IsAdminNoPermission(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADMIN_NO_PERMISSION.String() && e.Code == 403
}

func ErrorAdminNoPermission(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_ADMIN_NO_PERMISSION.String(), fmt.Sprintf(format, args...))
}
