package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/uadmin/uadmin"
)

// Gender Field
type Gender int

func (Gender) Male() Gender {
	return 1
}
func (Gender) Female() Gender {
	return 2
}

type YearLevel int

func (YearLevel) FirstYear() YearLevel {
	return 1
}
func (YearLevel) SecondYear() YearLevel {
	return 2
}
func (YearLevel) ThirdYear() YearLevel {
	return 3
}
func (YearLevel) FourthYear() YearLevel {
	return 4
}
func (YearLevel) FifthYear() YearLevel {
	return 5
}

// StudentInfo Model !
type StudentInfo struct {
	uadmin.Model
	Name      string `uadmin:"search;help:Firstname Middlename Lastname"`
	StudentNo string `uadmin:"read_only"`
	Address   string
	Contact   string `uadmin:"default_value: ;display_name:Contact Number;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	Email     string
	Birthdate string `uadmin:"list_excludehelp:YYYY/MM/DD"`
	Age       int    `uadmin:"list_exclude;default_value: "`
	Gender    Gender `uadmin:"filter"`

	Guardian   Guardian `uadmin:"help:Firstname Middlename Lastname"`
	GuardianID uint

	School   School `uadmin:"filter"`
	SchoolID uint

	Course    string    `uadmin:"filter;help:ex: BS Computer Engineering ..."`
	YearLevel YearLevel `uadmin:"filter"`
}

func (s *StudentInfo) Save() {
	level := s.YearLevel
	studentID := s.StudentNo
	student_information := StudentInfo{}

	if studentID == "" {
		if uadmin.Count(&student_information, "name = ? AND id <> ?", s.Name, s.ID) == 0 {

			prefix := ""
			if level == 1 {
				prefix = "23"
			} else if level == 2 {
				prefix = "22"
			} else if level == 3 {
				prefix = "21"
			} else if level == 4 {
				prefix = "20"
			} else if level == 5 {
				prefix = "19"
			} else {
				prefix = ""
			}

			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(100000)

			studentNum := fmt.Sprintf("%s-%05d", prefix, randomNumber)
			s.StudentNo = studentNum
		}
	} else {
		s.StudentNo = studentID
	}

	uadmin.Save(s)
}

func (s StudentInfo) Validate() (errMsg map[string]string) {

	errMsg = map[string]string{}

	student_information := StudentInfo{}
	if uadmin.Count(&student_information, "name = ? AND id <> ?", s.Name, s.ID) != 0 {
		errMsg["Name"] = "This student is already in the system"
	}
	return
}
