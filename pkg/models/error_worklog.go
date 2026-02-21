// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.

package models

import (
	"fmt"
	"net/http"

	"code.vikunja.io/api/pkg/web"
)

// ErrWorklogDoesNotExist represents an error where a worklog does not exist
type ErrWorklogDoesNotExist struct {
	ID int64
}

// IsErrWorklogDoesNotExist checks if an error is ErrWorklogDoesNotExist.
func IsErrWorklogDoesNotExist(err error) bool {
	_, ok := err.(*ErrWorklogDoesNotExist)
	return ok
}

func (err *ErrWorklogDoesNotExist) Error() string {
	return fmt.Sprintf("Worklog does not exist [ID: %d]", err.ID)
}

// ErrCodeWorklogDoesNotExist holds the unique world-error code of this error
const ErrCodeWorklogDoesNotExist = 10001

// HTTPError holds the http error description
func (err *ErrWorklogDoesNotExist) HTTPError() web.HTTPError {
	return web.HTTPError{
		HTTPCode: http.StatusNotFound,
		Code:     ErrCodeWorklogDoesNotExist,
		Message:  "This worklog does not exist.",
	}
}

// ErrInvalidWorklogDuration represents an error where a worklog duration is invalid
type ErrInvalidWorklogDuration struct {
	Duration int64
}

func (err *ErrInvalidWorklogDuration) Error() string {
	return fmt.Sprintf("Invalid worklog duration [Duration: %d]", err.Duration)
}

const ErrCodeInvalidWorklogDuration = 10002

func (err *ErrInvalidWorklogDuration) HTTPError() web.HTTPError {
	return web.HTTPError{
		HTTPCode: http.StatusBadRequest,
		Code:     ErrCodeInvalidWorklogDuration,
		Message:  "Worklog duration must be greater than 0.",
	}
}

// ErrWorklogCannotBeUpdated represents an error where a worklog cannot be updated
type ErrWorklogCannotBeUpdated struct {
	ID int64
}

func (err *ErrWorklogCannotBeUpdated) Error() string {
	return fmt.Sprintf("Worklog cannot be updated [ID: %d]", err.ID)
}

const ErrCodeWorklogCannotBeUpdated = 10003

func (err *ErrWorklogCannotBeUpdated) HTTPError() web.HTTPError {
	return web.HTTPError{
		HTTPCode: http.StatusForbidden,
		Code:     ErrCodeWorklogCannotBeUpdated,
		Message:  "You cannot update this worklog.",
	}
}

// ErrWorklogCannotBeDeleted represents an error where a worklog cannot be deleted
type ErrWorklogCannotBeDeleted struct {
	ID int64
}

func (err *ErrWorklogCannotBeDeleted) Error() string {
	return fmt.Sprintf("Worklog cannot be deleted [ID: %d]", err.ID)
}

const ErrCodeWorklogCannotBeDeleted = 10004

func (err *ErrWorklogCannotBeDeleted) HTTPError() web.HTTPError {
	return web.HTTPError{
		HTTPCode: http.StatusForbidden,
		Code:     ErrCodeWorklogCannotBeDeleted,
		Message:  "You cannot delete this worklog.",
	}
}