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
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

type tasks20260212000000 struct {
	Estimation *int64 `xorm:"BIGINT null" json:"estimation"`
}

func (tasks20260212000000) TableName() string {
	return "tasks"
}

type taskWorklogs20260212000000 struct {
	ID          int64  `xorm:"bigint autoincr not null unique pk" json:"id"`
	TaskID      int64  `xorm:"bigint INDEX not null" json:"task_id"`
	UserID      int64  `xorm:"bigint not null" json:"user_id"`
	Duration    int64  `xorm:"bigint not null" json:"duration"`
	Description string `xorm:"text null" json:"description"`
	LogDate     string `xorm:"date not null 'log_date'" json:"log_date"`
	Created     int64  `xorm:"created not null" json:"created"`
	Updated     int64  `xorm:"updated not null" json:"updated"`
}

func (taskWorklogs20260212000000) TableName() string {
	return "task_worklogs"
}

func init() {
	migrations = append(migrations, &xormigrate.Migration{
		ID:          "20260212000000",
		Description: "Add estimation, and create task_worklogs table",
		Migrate: func(tx *xorm.Engine) error {

			// Add estimation column to tasks (Sync2 will add it if it doesn't exist)
			err := tx.Sync2(tasks20260212000000{})
			if err != nil {
				return err
			}

			// Create task_worklogs table
			return tx.Sync2(taskWorklogs20260212000000{})
		},
		Rollback: func(tx *xorm.Engine) error {
			// Drop task_worklogs table
			err := tx.DropTables(taskWorklogs20260212000000{})
			if err != nil {
				return err
			}

			// Remove estimation column (optional, commented out to avoid data loss)
			// _, err = tx.Exec("ALTER TABLE tasks DROP COLUMN estimation")
			return nil
		},
	})
}