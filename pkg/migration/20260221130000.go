package migration

import (
	"time"

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
	UserID      int64     `xorm:"bigint not null"`
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
	registerMigration("20260221130000_add_estimation_and_task_worklogs", migrate20260221130000)
}

func migrate20260221130000(tx *xorm.Engine) error {
	// 1. Add estimation column to tasks table
	if err := tx.Sync(new(task20260221130000)); err != nil {
		return err
	}

	// 2. Create task_worklogs table
	if err := tx.Sync(new(taskWorklog20260221130000)); err != nil {
		return err
	}

	// 3. Add foreign key constraint (PostgreSQL only)
	if tx.Dialect().URI().DBType == "postgres" {
		_, err := tx.Exec(`
			ALTER TABLE task_worklogs
			ADD CONSTRAINT IF NOT EXISTS task_worklogs_task_id_fkey
			FOREIGN KEY (task_id)
			REFERENCES tasks(id)
			ON DELETE CASCADE
		`)
		if err != nil {
			return err
		}
	}

	return nil
}
