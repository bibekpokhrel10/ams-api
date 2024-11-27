package repository

import (
	"time"

	"github.com/ams-api/internal/models"
)

type IDashboard interface {
	GetSuperAdminDashboard() (*models.SuperAdminDashboard, error)
	GetInstitutionAdminDashboard(institutionID uint) (*models.InstitutionAdminDashboard, error)
	GetTeacherDashboard(teacherID uint) (*models.TeacherDashboard, error)
	GetStudentDashboard(studentID uint) (*models.StudentDashboard, error)
	GetProgramAdminDashboard(programID uint) (*models.ProgramAdminDashboard, error)
}

// GetSuperAdminDashboard retrieves dashboard data for super admin
func (r *Repository) GetSuperAdminDashboard() (*models.SuperAdminDashboard, error) {
	var dashboard models.SuperAdminDashboard

	// Get total counts
	var stats struct {
		TotalInstitutions int64
		TotalStudents     int64
		TotalTeachers     int64
		TotalCourses      int64
	}

	// Count institutions
	if err := r.db.Model(&models.Institution{}).Count(&stats.TotalInstitutions).Error; err != nil {
		return nil, err
	}

	// Count students and teachers
	if err := r.db.Model(&models.User{}).
		Select("COUNT(CASE WHEN user_type = 'student' THEN 1 END) as total_students, COUNT(CASE WHEN user_type = 'teacher' THEN 1 END) as total_teachers").
		Scan(&stats).Error; err != nil {
		return nil, err
	}

	// Get institutions with attendance
	institutions := []models.InstitutionAttendance{}

	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	if err := r.db.Raw(`
		SELECT 
			i.name,
			COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0) as attendance
		FROM institutions i
		LEFT JOIN users u ON u.institution_id = i.id
		LEFT JOIN enrollments e ON e.student_id = u.id
		LEFT JOIN attendances a ON a.class_id = e.class_id AND a.student_id = u.id
		WHERE a.date >= ?
		GROUP BY i.id, i.name`, thirtyDaysAgo).
		Scan(&institutions).Error; err != nil {
		return nil, err
	}

	dashboard.Stats.TotalInstitutions = int(stats.TotalInstitutions)
	dashboard.Stats.TotalStudents = int(stats.TotalStudents)
	dashboard.Stats.TotalTeachers = int(stats.TotalTeachers)
	dashboard.Stats.TotalCourses = int(stats.TotalCourses)
	dashboard.InstitutionsList = institutions

	return &dashboard, nil
}

// GetInstitutionAdminDashboard retrieves dashboard data for institution admin
func (r *Repository) GetInstitutionAdminDashboard(institutionID uint) (*models.InstitutionAdminDashboard, error) {
	var dashboard models.InstitutionAdminDashboard

	// Get institution stats
	var stats struct {
		TotalStudents     int64
		TotalTeachers     int64
		TotalCourses      int64
		AverageAttendance float64
	}

	// Count students and teachers in institution
	if err := r.db.Model(&models.User{}).
		Where("institution_id = ?", institutionID).
		Select(`
			COUNT(CASE WHEN user_type = 'student' THEN 1 END) as total_students,
			COUNT(CASE WHEN user_type = 'teacher' THEN 1 END) as total_teachers
		`).
		Scan(&stats).Error; err != nil {
		return nil, err
	}

	// Get monthly attendance trends
	trends := []models.MonthlyAttendance{}

	if err := r.db.Raw(`
		SELECT 
			DATE_FORMAT(a.date, '%Y-%m') as month,
			AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END) as attendance
		FROM attendances a
		JOIN classes c ON a.class_id = c.id
		JOIN users u ON c.instructor_id = u.id
		WHERE u.institution_id = ?
		GROUP BY DATE_FORMAT(a.date, '%Y-%m')
		ORDER BY month DESC
		LIMIT 6
	`, institutionID).Scan(&trends).Error; err != nil {
		return nil, err
	}

	dashboard.Stats.TotalStudents = int(stats.TotalStudents)
	dashboard.Stats.TotalTeachers = int(stats.TotalTeachers)
	dashboard.Stats.TotalCourses = int(stats.TotalCourses)
	dashboard.AttendanceTrends = trends

	return &dashboard, nil
}

// GetTeacherDashboard retrieves dashboard data for teachers
func (r *Repository) GetTeacherDashboard(teacherID uint) (*models.TeacherDashboard, error) {
	var dashboard models.TeacherDashboard

	// Get teacher's classes with attendance
	classes := []models.ClassInfo{}

	if err := r.db.Raw(`
		SELECT 
			CONCAT(c.year, ' - ', co.name) as name,
			COUNT(DISTINCT e.student_id) as students,
			COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0) as attendance
		FROM classes c
		JOIN courses co ON c.course_id = co.id
		LEFT JOIN enrollments e ON e.class_id = c.id
		LEFT JOIN attendances a ON a.class_id = c.id
		WHERE c.instructor_id = ?
		GROUP BY c.id, c.year, co.name
	`, teacherID).Scan(&classes).Error; err != nil {
		return nil, err
	}

	dashboard.Stats.TotalClasses = len(classes)
	var totalStudents int64
	var totalAttendance float64

	for _, class := range classes {
		totalStudents += int64(class.Students)
		totalAttendance += class.Attendance
	}

	dashboard.Stats.TotalStudents = int(totalStudents)
	if len(classes) > 0 {
		dashboard.Stats.AverageAttendance = int(totalAttendance / float64(len(classes)))
	}
	dashboard.Classes = classes

	return &dashboard, nil
}

// GetStudentDashboard retrieves dashboard data for students
func (r *Repository) GetStudentDashboard(studentID uint) (*models.StudentDashboard, error) {
	var dashboard models.StudentDashboard

	courses := []models.CourseInfo{}

	if err := r.db.Raw(`
		SELECT 
			co.name,
			'N/A' as grade,
			COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0) as attendance
		FROM enrollments e
		JOIN classes c ON e.class_id = c.id
		JOIN courses co ON c.course_id = co.id
		LEFT JOIN attendances a ON a.class_id = c.id AND a.student_id = e.student_id
		WHERE e.student_id = ?
		GROUP BY co.id, co.name
	`, studentID).Scan(&courses).Error; err != nil {
		return nil, err
	}

	dashboard.Stats.TotalCourses = len(courses)
	var totalAttendance float64
	for _, course := range courses {
		totalAttendance += course.Attendance
	}

	if len(courses) > 0 {
		dashboard.Stats.AverageAttendance = int(totalAttendance / float64(len(courses)))
	}
	dashboard.Courses = courses

	return &dashboard, nil
}

func (r *Repository) GetProgramAdminDashboard(programID uint) (*models.ProgramAdminDashboard, error) {
	var dashboard models.ProgramAdminDashboard

	// Get program stats
	var stats struct {
		TotalStudents     int64
		TotalCourses      int64
		AverageAttendance float64
	}

	// Count total students enrolled in the program
	if err := r.db.Model(&models.ProgramEnrollment{}).
		Where("program_id = ?", programID).
		Count(&stats.TotalStudents).Error; err != nil {
		return nil, err
	}

	// Count total courses in the program (through semesters)
	if err := r.db.Table("courses").
		Joins("JOIN classes ON classes.course_id = courses.id").
		Joins("JOIN semesters ON classes.semester_id = semesters.id").
		Where("semesters.program_id = ?", programID).
		Count(&stats.TotalCourses).Error; err != nil {
		return nil, err
	}

	// Get overall program attendance
	if err := r.db.Raw(`
		SELECT 
			COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0) as average_attendance
		FROM program_enrollments pe
		JOIN enrollments e ON e.student_id = pe.student_id
		JOIN classes c ON e.class_id = c.id
		JOIN semesters s ON c.semester_id = s.id
		JOIN attendances a ON a.class_id = c.id AND a.student_id = pe.student_id
		WHERE pe.program_id = ?
		AND a.date >= DATE_SUB(NOW(), INTERVAL 30 DAY)
	`, programID).Scan(&stats.AverageAttendance).Error; err != nil {
		return nil, err
	}

	// Get semester-wise data
	var semesterData []struct {
		Name       string
		Attendance float64
	}

	if err := r.db.Raw(`
		SELECT 
			CONCAT(s.name, ' - ', s.year) as name,
			COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0) as attendance
		FROM semesters s
		JOIN classes c ON c.semester_id = s.id
		LEFT JOIN enrollments e ON e.class_id = c.id
		LEFT JOIN attendances a ON a.class_id = c.id AND a.student_id = e.student_id
		WHERE s.program_id = ?
		GROUP BY s.id, s.name, s.year
		ORDER BY s.year DESC, s.time_period DESC
	`, programID).Scan(&semesterData).Error; err != nil {
		return nil, err
	}

	// Get class performance by semester (optional additional data)
	var classPerformance []struct {
		SemesterName string
		CourseName   string
		Students     int64
		Attendance   float64
	}

	if err := r.db.Raw(`
		SELECT 
			CONCAT(s.name, ' - ', s.year) as semester_name,
			co.name as course_name,
			COUNT(DISTINCT e.student_id) as students,
			COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0) as attendance
		FROM semesters s
		JOIN classes c ON c.semester_id = s.id
		JOIN courses co ON c.course_id = co.id
		LEFT JOIN enrollments e ON e.class_id = c.id
		LEFT JOIN attendances a ON a.class_id = c.id AND a.student_id = e.student_id
		WHERE s.program_id = ?
		GROUP BY s.id, s.name, s.year, co.id, co.name
		ORDER BY s.year DESC, s.time_period DESC, co.name
	`, programID).Scan(&classPerformance).Error; err != nil {
		return nil, err
	}

	// Populate dashboard
	dashboard.Stats.TotalStudents = int(stats.TotalStudents)
	dashboard.Stats.TotalCourses = int(stats.TotalCourses)
	dashboard.Stats.AverageAttendance = int(stats.AverageAttendance)

	// Convert semester data to the required format
	dashboard.SemesterData = make([]models.SemesterInfo, len(semesterData))
	for i, sem := range semesterData {
		dashboard.SemesterData[i] = models.SemesterInfo{
			Name:       sem.Name,
			Attendance: sem.Attendance,
		}
	}

	return &dashboard, nil
}
