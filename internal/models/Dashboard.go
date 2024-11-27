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
		TotalInstitutions int `json:"totalInstitutions"`
		TotalStudents     int `json:"totalStudents"`
		TotalTeachers     int `json:"totalTeachers"`
		TotalCourses      int `json:"totalCourses"`
	} `json:"stats"`
	InstitutionsList []InstitutionAttendance `json:"institutionsList"`
}

type InstitutionAttendance struct {
	Name       string  `json:"name"`
	Attendance float64 `json:"attendance"`
}

// Institution Admin Dashboard
type InstitutionAdminDashboard struct {
	Stats struct {
		TotalStudents     int `json:"totalStudents"`
		TotalTeachers     int `json:"totalTeachers"`
		TotalCourses      int `json:"totalCourses"`
		AverageAttendance int `json:"averageAttendance"`
	} `json:"stats"`
	AttendanceTrends []MonthlyAttendance `json:"attendanceTrends"`
}

type MonthlyAttendance struct {
	Month      string  `json:"month"`
	Attendance float64 `json:"attendance"`
}

// Program Admin Dashboard
type ProgramAdminDashboard struct {
	Stats struct {
		TotalStudents     int `json:"totalStudents"`
		TotalCourses      int `json:"totalCourses"`
		AverageAttendance int `json:"averageAttendance"`
	} `json:"stats"`
	SemesterData []SemesterInfo `json:"semesterData"`
}

type SemesterInfo struct {
	Name       string  `json:"name"`
	Attendance float64 `json:"attendance"`
}

// Teacher Dashboard
type TeacherDashboard struct {
	Stats struct {
		TotalClasses      int `json:"totalClasses"`
		TotalStudents     int `json:"totalStudents"`
		AverageAttendance int `json:"averageAttendance"`
	} `json:"stats"`
	Classes []ClassInfo `json:"classes"`
}

type ClassInfo struct {
	Name       string  `json:"name"`
	Students   int     `json:"students"`
	Attendance float64 `json:"attendance"`
}

// Student Dashboard
type StudentDashboard struct {
	Stats struct {
		TotalCourses      int `json:"totalCourses"`
		AverageAttendance int `json:"averageAttendance"`
		CurrentSemester   int `json:"currentSemester"`
	} `json:"stats"`
	Courses []CourseInfo `json:"courses"`
}

type CourseInfo struct {
	Name       string  `json:"name"`
	Grade      string  `json:"grade"`
	Attendance float64 `json:"attendance"`
}
