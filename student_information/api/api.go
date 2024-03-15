package api

import (
	"net/http"
	"strings"
)

// Handler !
func Handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/api/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/student_info_list") {
		StudentInfoListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/guardian_list") {
		GuardianListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/school_list") {
		SchoolListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/course_list") {
		CourseListAPIHandler(w, r)
		return
	}
}
