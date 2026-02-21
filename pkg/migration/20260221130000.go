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

package migration

import (
	"time"

	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

type task20260221130000 struct {
	Estimation *int64 `xorm:"BIGINT null"`
}

func (task20260221130000) TableName() string {
	return "tasks"
}

type taskWorklog20260221130000 struct {
	ID          int64     `xorm:"bigint autoincr not null unique pk"`
	TaskID      int64     `xorm:"bigint not null index"`
	UserID      int64     `xorm:"bigint not null index"`
	Duration    int64     `xorm:"bigint not null"`
	Description string    `xorm:"text null"`
	LogDate     time.Time `xorm:"timestamptz not null 'log_date'"`
	Created     int64     `xorm:"bigint not null"`
	Updated     int64     `xorm:"bigint not null"`
}

func (taskWorklog20260221130000) TableName() string {
	return "task_worklogs"
}

func init() {
	migrations = append(migrations, &xormigrate.Migration{
		ID:          "20260221130000",
		Description: "Add estimation to tasks and create task_worklogs table",
		Migrate: func(tx *xorm.Engine) error {
			// 1. Add estimation column to tasks table
			if err := tx.Sync2(task20260221130000{}); err != nil {
				return err
			}

			// 2. Create task_worklogs table
			return tx.Sync2(taskWorklog20260221130000{})
		},
		Rollback: func(tx *xorm.Engine) error {
			// Drop task_worklogs table
			if err := tx.DropTables(taskWorklog20260221130000{}); err != nil {
				return err
			}

			// Drop estimation column from tasks
			return dropTableColum(tx, "tasks", "estimation")
		},
	})
}
