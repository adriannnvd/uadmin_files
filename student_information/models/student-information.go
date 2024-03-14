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

// StudentInfo Model !
type StudentInfo struct {
	uadmin.Model
	Name       string `uadmin:"required;search;help:Firstname Middlename Lastname"`
	StudentNo  string `uadmin:"read_only"`
	Address    string `uadmin:"required"`
	Contact    string `uadmin:"required;default_value: ;display_name:Contact Number;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	Email      string
	Birthdate  string   `uadmin:"required;help:YYYY/MM/DD"`
	Age        int      `uadmin:"required;default_value: "`
	Gender     Gender   `uadmin:"required;filter"`

	Guardian   Guardian `uadmin:"required;help:Firstname Middlename Lastname"`
	GuardianID uint

	School     School `uadmin:"required;filter"`
	SchoolID   uint
	
	Course     string `uadmin:"required;filter;help:ex: BS Computer Engineering ..."`
	YearLevel  string `uadmin:"required;filter;help:ex: First Year, Second Year ..., Fifth Year;pattern:^(First|Second|Third|Fourth|Fifth) Year$;pattern_msg:Please input correct School Year Level. Follow the format."`
}

func (s *StudentInfo) Save() {
	studentNum := GenerateStudentNo(s.YearLevel)
	s.StudentNo = studentNum

	uadmin.Save(s)
}

func GenerateStudentNo(yearLevel string) string {

	yearLevelPrefix := map[string]string{
		"Fifth Year":  "19",
		"Fourth Year": "20",
		"Third Year":  "21",
		"Second Year": "22",
		"First Year":  "23",
	}

	prefix, ok := yearLevelPrefix[yearLevel]
	if !ok {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100000)

	return fmt.Sprintf("%s-%05d", prefix, randomNumber)

}

func (s StudentInfo) Validate() (errMsg map[string]string) {

	errMsg = map[string]string{}

	student_information := StudentInfo{}
	if uadmin.Count(&student_information, "name = ? AND id <> ?", s.Name, s.ID) != 0 {
		errMsg["Name"] = "This student is already in the system"
	}
	return
}
