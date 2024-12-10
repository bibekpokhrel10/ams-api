package models

type DashboardResponse struct {
	SuperAdmin       *SuperAdminDashboard       `json:"super_admin,omitempty"`
	InstitutionAdmin *InstitutionAdminDashboard `json:"institution_admin,omitempty"`
	ProgramAdmin     *ProgramAdminDashboard     `json:"program_admin,omitempty"`
	Teacher          *TeacherDashboard          `json:"teacher,omitempty"`
	Student          *StudentDashboard          `json:"student,omitempty"`
}

// SuperAdmin Dashboard
type SuperAdminDashboard struct {
	Stats struct {
		TotalInstitutions int `json:"total_institutions"`
		TotalStudents     int `json:"total_students"`
		TotalTeachers     int `json:"total_teachers"`
		TotalCourses      int `json:"total_courses"`
	} `json:"stats"`
	InstitutionsList []InstitutionAttendance `json:"institutions_list"`
}

type InstitutionAttendance struct {
	Name       string  `json:"name"`
	Attendance float64 `json:"attendance"`
}

// Institution Admin Dashboard
type InstitutionAdminDashboard struct {
	Stats struct {
		TotalStudents     int `json:"total_students"`
		TotalTeachers     int `json:"total_teachers"`
		TotalCourses      int `json:"total_courses"`
		AverageAttendance int `json:"average_attendance"`
	} `json:"stats"`
	AttendanceTrends []MonthlyAttendance `json:"attendance_trends"`
}

type MonthlyAttendance struct {
	Month      string  `json:"month"`
	Attendance float64 `json:"attendance"`
}

// Program Admin Dashboard
type ProgramAdminDashboard struct {
	Stats struct {
		TotalStudents     int `json:"total_students"`
		TotalCourses      int `json:"total_courses"`
		AverageAttendance int `json:"average_attendance"`
	} `json:"stats"`
	SemesterData []SemesterInfo `json:"semester_data"`
}

type SemesterInfo struct {
	Name       string  `json:"name"`
	Attendance float64 `json:"attendance"`
}

// Teacher Dashboard
type TeacherDashboard struct {
	Stats struct {
		TotalClasses      int `json:"total_classes"`
		TotalStudents     int `json:"total_students"`
		AverageAttendance int `json:"average_attendance"`
	} `json:"stats"`
	Classes []ClassInfo `json:"classes"`
}

type ClassInfo struct {
	Name       string  `json:"name"`
	Students   int     `json:"students"`
	Attendance float64 `json:"attendance"`
}

type StudentDashboard struct {
	Stats struct {
		TotalClasses      int `json:"total_classes"`
		AverageAttendance int `json:"average_attendance"`
	} `json:"stats"`
	Classes           []ClassDetails      `json:"classes"`
	MonthlyAttendance []MonthlyAttendance `json:"monthly_attendance"`
}

type ClassDetails struct {
	ClassID         uint    `json:"class_id"`
	CourseCode      string  `json:"course_code"`
	CourseName      string  `json:"course_name"`
	CourseCredits   int     `json:"course_credits"`
	ProgramName     string  `json:"program_name"`
	SemesterName    string  `json:"semester_name"`
	SemesterYear    string  `json:"semester_year"`
	SemesterPeriod  string  `json:"semester_period"`
	Year            int     `json:"year"`
	Schedule        string  `json:"schedule"`
	InstructorName  string  `json:"instructor_name"`
	ClassAttendance float64 `json:"class_attendance"`
	IsEnrolled      bool    `json:"is_enrolled"`
}
