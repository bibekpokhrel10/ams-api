package repository

import (
	"math"
	"time"

	"github.com/ams-api/internal/models"
)

type IDashboard interface {
	GetSuperAdminDashboard() (*models.SuperAdminDashboard, error)
	GetInstitutionAdminDashboard(institutionID uint) (*models.InstitutionAdminDashboard, error)
	GetTeacherDashboard(teacherID uint) (*models.TeacherDashboard, error)
	GetStudentDashboard(studentID uint) (*models.StudentDashboard, error)
	GetProgramAdminDashboard(programID uint) (*models.ProgramAdminDashboard, error)
	GetStudentClassAttendance(studentID uint, classID uint) (*models.StudentClassAttendance, error)
	GetTeacherClassAttendance(classID uint) (*models.TeacherClassAttendance, error)
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

	// Count courses
	if err := r.db.Model(&models.Course{}).Where("deleted_at IS NULL").Count(&stats.TotalCourses).Error; err != nil {
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

	if err := r.db.Raw(`
	SELECT COUNT(DISTINCT c.id) as total_courses
	FROM courses c
	JOIN semesters s ON c.semester_id = s.id
	JOIN programs p ON s.program_id = p.id
	WHERE p.institution_id = ? AND c.deleted_at IS NULL
`, institutionID).Scan(&stats).Error; err != nil {
		return nil, err
	}

	// Get monthly attendance trends
	trends := []models.MonthlyAttendance{}

	if err := r.db.Raw(`
    SELECT 
        TO_CHAR(a.date, 'YYYY-MM') as month,
 	       AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END) as attendance
    FROM attendances a
    JOIN classes c ON a.class_id = c.id
    JOIN users u ON c.instructor_id = u.id
    WHERE u.institution_id = ?
    GROUP BY TO_CHAR(a.date, 'YYYY-MM')
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
            c.id as class_id,
            CONCAT(c.year, ' - ', co.name) as name,
            COUNT(DISTINCT e.student_id) as students,
            ROUND(COALESCE(AVG(CASE WHEN a.is_present THEN 100.0 ELSE 0.0 END), 0),2) as attendance
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
		avgAttendance := math.Round(totalAttendance / float64(len(classes)))
		dashboard.Stats.AverageAttendance = int(avgAttendance)
	}

	dashboard.Classes = classes
	return &dashboard, nil
}

// GetStudentDashboard retrieves dashboard data for students
func (r *Repository) GetStudentDashboard(studentID uint) (*models.StudentDashboard, error) {
	var dashboard models.StudentDashboard

	// Updated query to leverage model relationships
	classes := []models.ClassDetails{}

	if err := r.db.Raw(`
		SELECT 
			c.id as class_id,
			co.code as course_code,
			co.name as course_name,
			co.credits,
			p.name as program_name,
			s.name as semester_name,
			s.year as semester_year,
			s.time_period as semester_period,
			c.year,
			c.schedule,
			u.first_name as instructor_name,
			ROUND(COALESCE(AVG(CASE WHEN a.is_present THEN 100.0 ELSE 0.0 END), 0), 2) as class_attendance,
			CASE WHEN e.student_id IS NOT NULL THEN true ELSE false END as is_enrolled
		FROM enrollments e
		JOIN classes c ON e.class_id = c.id
		JOIN courses co ON c.course_id = co.id
		JOIN semesters s ON co.semester_id = s.id
		JOIN programs p ON s.program_id = p.id
		JOIN users u ON c.instructor_id = u.id
		LEFT JOIN attendances a ON a.class_id = c.id AND a.student_id = e.student_id
		WHERE e.student_id = ?
		GROUP BY 
			c.id, co.code, co.name, co.credits, 
			p.name, s.name, s.year, s.time_period, 
			c.year, c.schedule, u.first_name, e.student_id
	`, studentID).Scan(&classes).Error; err != nil {
		return nil, err
	}

	// Monthly attendance data for the student across all enrolled classes
	monthlyAttendance := []models.MonthlyAttendance{}

	if err := r.db.Raw(`
		SELECT 
			TO_CHAR(a.date, 'YYYY-MM') as month,
			ROUND(COALESCE(AVG(CASE WHEN a.is_present THEN 100.0 ELSE 0.0 END), 0), 2) as attendance
		FROM enrollments e
		JOIN classes c ON e.class_id = c.id
		JOIN attendances a ON a.class_id = c.id AND a.student_id = e.student_id
		WHERE e.student_id = ?
		GROUP BY TO_CHAR(a.date, 'YYYY-MM')
		ORDER BY month DESC
		LIMIT 6
	`, studentID).Scan(&monthlyAttendance).Error; err != nil {
		return nil, err
	}

	// Calculate stats
	dashboard.Stats.TotalClasses = len(classes)
	var totalAttendance float64
	for _, class := range classes {
		totalAttendance += class.ClassAttendance
	}

	if len(classes) > 0 {
		avgAttendance := math.Round(totalAttendance / float64(len(classes)))
		dashboard.Stats.AverageAttendance = int(avgAttendance)
	}

	dashboard.Classes = classes
	dashboard.MonthlyAttendance = monthlyAttendance

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
			ROUND(COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0), 2) as average_attendance
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
			ROUND(COALESCE(AVG(CASE WHEN a.is_present THEN 100 ELSE 0 END), 0), 2) as attendance
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
	dashboard.Stats.AverageAttendance = int(math.Round(stats.AverageAttendance))

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

// GetStudentClassAttendance retrieves detailed attendance data for a specific student in a class
func (r *Repository) GetStudentClassAttendance(studentID uint, classID uint) (*models.StudentClassAttendance, error) {
	var attendance models.StudentClassAttendance

	// Get class and student basic info
	if err := r.db.Raw(`
        SELECT 
            c.id as class_id,
            co.code as course_code,
            co.name as course_name,
            CONCAT(u.first_name, ' ', u.last_name) as instructor_name,
            c.schedule,
            CONCAT(s.first_name, ' ', s.last_name) as student_name,
            s.id as roll_number,
            COALESCE(
                (SELECT COUNT(*) 
                FROM attendances 
                WHERE class_id = c.id 
                AND student_id = e.student_id 
                AND is_present = true), 0
            ) as present_classes,
            COALESCE(
                (SELECT COUNT(*) 
                FROM attendances 
                WHERE class_id = c.id 
                AND student_id = e.student_id), 0
            ) as total_classes,
            COALESCE(
                (SELECT COUNT(*) 
                FROM attendances 
                WHERE class_id = c.id 
                AND student_id = e.student_id 
                AND is_present = true 
                AND date >= CURRENT_DATE - INTERVAL '7 days'), 0
            ) as current_streak
        FROM enrollments e
        JOIN classes c ON e.class_id = c.id
        JOIN courses co ON c.course_id = co.id
        JOIN users u ON c.instructor_id = u.id
        JOIN users s ON e.student_id = s.id
        WHERE e.student_id = ? AND c.id = ?
    `, studentID, classID).Scan(&attendance.BasicInfo).Error; err != nil {
		return nil, err
	}

	// Get monthly attendance data
	if err := r.db.Raw(`
        SELECT 
            TO_CHAR(date, 'Month') as month,
            COUNT(CASE WHEN is_present = true THEN 1 END) as present,
            COUNT(*) as total,
            COALESCE(
                CAST(COUNT(CASE WHEN is_present = true THEN 1 END) AS FLOAT) * 100.0 / 
                NULLIF(COUNT(*), 0), 
                0
            ) as attendance
        FROM attendances
        WHERE student_id = ? AND class_id = ?
        GROUP BY TO_CHAR(date, 'Month'), TO_CHAR(date, 'MM')
        ORDER BY TO_CHAR(date, 'MM') DESC
    `, studentID, classID).Scan(&attendance.MonthlyAttendance).Error; err != nil {
		return nil, err
	}

	// Get recent attendance records
	if err := r.db.Raw(`
        SELECT 
            date,
            CONCAT('Week ', EXTRACT(WEEK FROM date)) as week,
            CASE WHEN is_present THEN 'Present' ELSE 'Absent' END as status
        FROM attendances
        WHERE student_id = ? AND class_id = ?
        ORDER BY date DESC
        LIMIT 10
    `, studentID, classID).Scan(&attendance.RecentAttendance).Error; err != nil {
		return nil, err
	}

	return &attendance, nil
}

// GetTeacherClassAttendance retrieves class attendance data for a teacher's view
func (r *Repository) GetTeacherClassAttendance(classID uint) (*models.TeacherClassAttendance, error) {
	var attendance models.TeacherClassAttendance

	// Get class basic info and overall stats
	if err := r.db.Raw(`
        SELECT 
            c.id as class_id,
            co.code as course_code,
            co.name as course_name,
            CONCAT(u.first_name, ' ', u.last_name) as instructor_name,
            c.schedule,
            COUNT(DISTINCT e.student_id) as total_students,
            COALESCE(
                CAST(COUNT(CASE WHEN a.is_present = true THEN 1 END) AS FLOAT) * 100.0 / 
                NULLIF(COUNT(*), 0), 
                0
            ) as overall_attendance,
            COUNT(DISTINCT a.date) as total_classes
        FROM classes c
        JOIN courses co ON c.course_id = co.id
        JOIN users u ON c.instructor_id = u.id
        LEFT JOIN enrollments e ON c.id = e.class_id
        LEFT JOIN attendances a ON c.id = a.class_id
        WHERE c.id = ?
        GROUP BY c.id, co.code, co.name, u.first_name, u.last_name, c.schedule
    `, classID).Scan(&attendance.BasicInfo).Error; err != nil {
		return nil, err
	}

	// Get weekly attendance data
	if err := r.db.Raw(`
        WITH weekly_data AS (
            SELECT 
                date_trunc('week', date) as week_start,
                COUNT(CASE WHEN is_present = true THEN 1 END) as present,
                COUNT(*) as total
            FROM attendances
            WHERE class_id = ?
            GROUP BY date_trunc('week', date)
        )
        SELECT 
            CONCAT('Week ', ROW_NUMBER() OVER (ORDER BY week_start)) as week,
            present,
            total,
            COALESCE(
                CAST(present AS FLOAT) * 100.0 / NULLIF(total, 0), 
                0
            ) as attendance
        FROM weekly_data
        ORDER BY week_start DESC
        LIMIT 4
    `, classID).Scan(&attendance.WeeklyData).Error; err != nil {
		return nil, err
	}

	// Get monthly attendance data
	if err := r.db.Raw(`
        SELECT 
            TO_CHAR(date, 'Month') as month,
            COUNT(CASE WHEN is_present = true THEN 1 END) as present,
            COUNT(*) as total,
            COALESCE(
                CAST(COUNT(CASE WHEN is_present = true THEN 1 END) AS FLOAT) * 100.0 / 
                NULLIF(COUNT(*), 0), 
                0
            ) as attendance
        FROM attendances
        WHERE class_id = ?
        GROUP BY TO_CHAR(date, 'Month'), TO_CHAR(date, 'MM')
        ORDER BY TO_CHAR(date, 'MM') DESC
        LIMIT 3
    `, classID).Scan(&attendance.MonthlyData).Error; err != nil {
		return nil, err
	}

	// Get student-wise attendance data
	if err := r.db.Raw(`
        SELECT 
            u.id,
            CONCAT(u.first_name, ' ', u.last_name) as name,
            COALESCE(
                CAST(COUNT(CASE WHEN a.is_present = true THEN 1 END) AS FLOAT) * 100.0 / 
                NULLIF(COUNT(*), 0), 
                0
            ) as attendance
        FROM enrollments e
        JOIN users u ON e.student_id = u.id
        LEFT JOIN attendances a ON a.student_id = e.student_id AND a.class_id = e.class_id
        WHERE e.class_id = ?
        GROUP BY u.id, u.first_name, u.last_name
        ORDER BY attendance DESC
    `, classID).Scan(&attendance.StudentList).Error; err != nil {
		return nil, err
	}

	return &attendance, nil
}
