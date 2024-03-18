package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	Name      string `uadmin:"required;search;help:Firstname Middlename Lastname"`
	StudentNo string `uadmin:"read_only"`
	Address   string
	Contact   string `uadmin:"default_value: ;display_name:Contact Number;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	Email     string
	Birthdate string `uadmin:"list_exclude;help:YYYY/MM/DD"`
	Age       int    `uadmin:"list_exclude;default_value: "`
	Gender    Gender `uadmin:"filter"`

	Guardian   Guardian `uadmin:"help:Firstname Middlename Lastname"`
	GuardianID uint

	School   School `uadmin:"filter"`
	SchoolID uint
	Code     string `uadmin:"read_only;list_exclude;hidden"`

	Course   Course `uadmin:"filter;help:ex. BS Computer Engineering ..."`
	CourseID uint

	YearLevel YearLevel `uadmin:"required;filter;help:Select your year level"`
}

func (s *StudentInfo) Save() {
	student_information := StudentInfo{}
	uadmin.Preload(s)

	level := s.YearLevel
	studentID := s.StudentNo

	currentYear := time.Now().Year() % 100

	name := s.Name

	schoolCode := s.School.SchoolCode
	s.Code = schoolCode

	if studentID == "" {

		baseCount := uadmin.Count(&student_information, "year_level = ? AND code = ?", s.YearLevel, s.Code)
		recentCount := baseCount + 1

		prefix := ""
		if level == 1 {
			prefix = strconv.Itoa(currentYear - 1)
		} else if level == 2 {
			prefix = strconv.Itoa(currentYear - 2)
		} else if level == 3 {
			prefix = strconv.Itoa(currentYear - 3)
		} else if level == 4 {
			prefix = strconv.Itoa(currentYear - 4)
		} else if level == 5 {
			prefix = strconv.Itoa(currentYear - 5)
		}
		rand.Seed(time.Now().UnixNano())

		// Name Initials
		words := strings.Fields(name)
		initials := ""
		for _, word := range words {
			initials += string(word[0])
		}
		initials = strings.ToUpper(initials)

		//Random Numbers
		randomNumber := rand.Intn(999)

		studentNum := fmt.Sprintf("%v%s-%v-%d-%05d", schoolCode, prefix, initials, randomNumber, recentCount)
		s.StudentNo = studentNum
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
