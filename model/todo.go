package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Text   string `gorm:"type:varchar(50);unique_index"`
	IsDone bool
}

func (t Todo) String() string {
	return fmt.Sprintf("%s", t.Text)
}
