package models

import "time"

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
	Name string `json:"name"`
	// Attendance float64 `json:"attendance"`
}

// Institution Admin Dashboard
type InstitutionAdminDashboard struct {
	Stats struct {
		TotalStudents int `json:"total_students"`
		TotalTeachers int `json:"total_teachers"`
		TotalCourses  int `json:"total_courses"`
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
	ClassId    uint    `json:"class_id"`
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

type StudentClassAttendance struct {
	BasicInfo struct {
		ClassID        uint   `json:"class_id"`
		CourseCode     string `json:"course_code"`
		CourseName     string `json:"course_name"`
		InstructorName string `json:"instructor_name"`
		Schedule       string `json:"schedule"`
		StudentName    string `json:"student_name"`
		RollNumber     uint   `json:"roll_number"`
		PresentClasses int    `json:"present_classes"`
		TotalClasses   int    `json:"total_classes"`
		CurrentStreak  int    `json:"current_streak"`
	} `json:"basic_info"`
	MonthlyAttendance []struct {
		Month      string  `json:"month"`
		Present    int     `json:"present"`
		Total      int     `json:"total"`
		Attendance float64 `json:"attendance"`
	} `json:"monthly_attendance"`
	RecentAttendance []struct {
		Date   time.Time `json:"date"`
		Week   string    `json:"week"`
		Status string    `json:"status"`
	} `json:"recent_attendance"`
}

type TeacherClassAttendance struct {
	BasicInfo struct {
		ClassID           uint    `json:"class_id"`
		CourseCode        string  `json:"course_code"`
		CourseName        string  `json:"course_name"`
		InstructorName    string  `json:"instructor_name"`
		Schedule          string  `json:"schedule"`
		TotalStudents     int     `json:"total_students"`
		OverallAttendance float64 `json:"overall_attendance"`
		TotalClasses      int     `json:"total_classes"`
	} `json:"basic_info"`
	WeeklyData []struct {
		Week       string  `json:"week"`
		Present    int     `json:"present"`
		Total      int     `json:"total"`
		Attendance float64 `json:"attendance"`
	} `json:"weekly_data"`
	MonthlyData []struct {
		Month      string  `json:"month"`
		Present    int     `json:"present"`
		Total      int     `json:"total"`
		Attendance float64 `json:"attendance"`
	} `json:"monthly_data"`
	StudentList []struct {
		ID         uint    `json:"id"`
		Name       string  `json:"name"`
		Attendance float64 `json:"attendance"`
	} `json:"student_list"`
}
