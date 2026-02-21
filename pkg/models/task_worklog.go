// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"time"
	"code.vikunja.io/api/pkg/log"
	"code.vikunja.io/api/pkg/user"
	"code.vikunja.io/api/pkg/web"
	"xorm.io/xorm"
)

// TaskWorklog represents a work log entry for a task
type TaskWorklog struct {
		// The unique, numeric id of this worklog.
		ID int64 `xorm:"bigint autoincr not null unique pk" json:"id" param:"id"`
		// The task this worklog belongs to.
		TaskID int64 `xorm:"bigint INDEX not null" json:"-" param:"task"`
		// The user who created this worklog.
		UserID int64 `xorm:"bigint not null" json:"-"`
		// The user object who created this worklog.
		User *user.User `xorm:"-" json:"user"`
		// Duration of work in seconds.
		Duration int64 `xorm:"bigint not null" json:"duration"`
		// Description of the work done.
		Description string `xorm:"text null" json:"description"`
		// The date when this work was done.
		LogDate time.Time `xorm:"timestamptz not null 'log_date'" json:"log_date"`
		// A timestamp when this worklog was created (unix timestamp).
		Created int64 `xorm:"bigint created not null" json:"created"`
		// A timestamp when this worklog was last updated (unix timestamp).
		Updated int64 `xorm:"bigint updated not null" json:"updated"`
	
		web.CRUDable    `xorm:"-" json:"-"`
		web.Permissions `xorm:"-" json:"-"`
}

// TableName returns the table name for worklogs
func (*TaskWorklog) TableName() string {
	return "task_worklogs"
}
// Validate validates the worklog data
func (wl *TaskWorklog) Validate() error {
	// Validate TaskID
	if wl.TaskID <= 0 {
		return &ErrTaskDoesNotExist{ID: wl.TaskID}
	}

	// Validate Duration
	if wl.Duration <= 0 {
		return &ErrInvalidWorklogDuration{Duration: wl.Duration}
	}

	// Validate LogDate (optional, set to today if not provided)
	if wl.LogDate.IsZero() {
		wl.LogDate = time.Now()
	}

	return nil
}

func (wl *TaskWorklog) Create(s *xorm.Session, a web.Auth) error {
	wl.UserID = a.GetID()

	if wl.Duration <= 0 {
		return &ErrInvalidWorklogDuration{Duration: wl.Duration}
	}

	if wl.LogDate.IsZero() {
		wl.LogDate = time.Now()
	}

	_, err := s.Insert(wl)
	return err
}
// ReadAll gets all worklogs for a task
func (wl *TaskWorklog) ReadAll(s *xorm.Session, a web.Auth, search string, page int, perPage int) (result interface{}, resultCount int, totalItems int64, err error) {
	// Get all worklogs for this task
	worklogs := []*TaskWorklog{}
	err = s.Where("task_id = ?", wl.TaskID).
		OrderBy("log_date DESC, created DESC").
		Find(&worklogs)
	if err != nil {
		return nil, 0, 0, err
	}

	// Get all user IDs
	userIDs := []int64{}
	for _, w := range worklogs {
		userIDs = append(userIDs, w.UserID)
	}

	// Get all users
	users := make(map[int64]*user.User)
	if len(userIDs) > 0 {
		err = s.In("id", userIDs).Find(&users)
		if err != nil {
			return nil, 0, 0, err
		}
	}

	// Attach user objects
	for _, w := range worklogs {
		if u, exists := users[w.UserID]; exists {
			w.User = u
		}
	}

	return worklogs, len(worklogs), int64(len(worklogs)), nil
}

// Update updates a worklog
func (wl *TaskWorklog) Update(s *xorm.Session, a web.Auth) error {
	_, err := s.ID(wl.ID).
		Cols("duration", "description", "log_date").
		Update(wl)
	return err
}

// ReadOne gets a single worklog by id
func (wl *TaskWorklog) ReadOne(s *xorm.Session, a web.Auth) (err error) {
	exists, err := s.
		Where("id = ?", wl.ID).
		And("task_id = ?", wl.TaskID).
		Get(wl)

	if err != nil {
		return err
	}

	if !exists {
		return &ErrWorklogDoesNotExist{ID: wl.ID}
	}

	return nil
}

// Delete deletes a worklog
func (wl *TaskWorklog) Delete(s *xorm.Session, a web.Auth) error {
	log.Infof("DELETE worklog: id=%d task_id=%d", wl.ID, wl.TaskID)

	affected, err := s.
	Where("id = ?", wl.ID).
	And("task_id = ?", wl.TaskID).
	Delete(&TaskWorklog{})

	log.Infof("DELETE affected rows: %d", affected)

	return err
}
// Delete deletes a worklog
func (tw *TaskWorklog) CanDelete(s *xorm.Session, a web.Auth) (bool, error) {
	if tw == nil {
		return false, ErrGenericForbidden{}
	}

	if tw.TaskID == 0 {
		return false, ErrTaskDoesNotExist{}
	}

	task, err := GetTaskByIDSimple(s, tw.TaskID)
	if err != nil {
		return false, err
	}

	return task.CanWrite(s, a)
}
// CanRead checks if the user can read worklogs for this task
func (wl *TaskWorklog) CanRead(s *xorm.Session, a web.Auth) (bool, int, error) {
	return true, 0, nil
}
func (wl *TaskWorklog) CanCreate(s *xorm.Session, a web.Auth) (bool, error) {
	return true, nil
}

func (wl *TaskWorklog) GetTaskID() int64 {
	return wl.TaskID
}

func (wl *TaskWorklog) SetTaskID(id int64) {
	wl.TaskID = id
}