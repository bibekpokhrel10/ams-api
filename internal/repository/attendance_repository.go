package repository

import (
	"fmt"
	"math"
	"time"

	"github.com/ams-api/internal/models"
)

type IAttendance interface {
	CreateAttendance(data *models.Attendance) (*models.Attendance, error)
	FindAllAttendance(req *models.ListAttendanceRequest) (*[]models.Attendance, int, error)
	FindAllAttendancesByClassIdAndDate(classId uint, date string) (*[]models.Attendance, error)
	FindAttendancesByClassIdAndDateAndStudentId(classId uint, date time.Time, studentId uint) (*models.Attendance, error)
	FindAttendanceById(id uint) (*models.Attendance, error)
	UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.Attendance, error)
	DeleteAttendance(id uint) error
	GetStudentsWithLowAttendance(classID uint, thresholdPercentage float64) ([]models.StudentAttendanceAlert, error)
}

func (r *Repository) CreateAttendance(data *models.Attendance) (*models.Attendance, error) {
	err := r.db.Model(&models.Attendance{}).Save(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAttendancesByClassIdAndDateAndStudentId(classId uint, date time.Time, studentId uint) (*models.Attendance, error) {
	data := &models.Attendance{}
	err := r.db.Model(models.Attendance{}).Where("class_id = ? AND date = ? AND student_id = ?", classId, date, studentId).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err

}

func (r *Repository) FindAllAttendance(req *models.ListAttendanceRequest) (*[]models.Attendance, int, error) {
	datas := &[]models.Attendance{}
	count := 0
	f := r.db.Model(models.Attendance{})
	if req.ClassId != 0 {
		f = f.Where("class_id = ?", req.ClassId)
	}
	if req.StudentId != 0 {
		f = f.Where("student_id = ?", req.StudentId)
	}
	if req.Date != "" {
		f = f.Where("date = ?", req.Date)
	}
	f = f.Count(&count)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) FindAllAttendancesByClassIdAndDate(classId uint, date string) (*[]models.Attendance, error) {
	datas := &[]models.Attendance{}
	err := r.db.Model(models.Attendance{}).Where("class_id = ? AND date = ?", classId, date).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindAttendanceById(id uint) (*models.Attendance, error) {
	data := &models.Attendance{}
	err := r.db.Model(models.Attendance{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.Attendance, error) {
	data := &models.Attendance{}
	err := r.db.Model(&models.Attendance{}).Where("id = ?", id).Updates(map[string]interface{}{
		"student_id": req.StudentId,
		"class_id":   req.ClassId,
		"date":       req.Date,
		"is_present": req.IsPresent,
		"status":     req.Status,
		"remarks":    req.Remarks,
	}).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteAttendance(id uint) error {
	return r.db.Model(&models.Attendance{}).Where("id = ?", id).Delete(&models.Attendance{}).Error
}

// GetStudentsWithLowAttendance retrieves students whose attendance is below the specified threshold
func (r *Repository) GetStudentsWithLowAttendance(classID uint, thresholdPercentage float64) ([]models.StudentAttendanceAlert, error) {
	var lowAttendanceStudents []models.StudentAttendanceAlert

	// Query to get attendance details for all students in the class
	if err := r.db.Raw(`
       WITH student_attendance AS (
  SELECT
    u.id as student_id,
    u.first_name,
    u.last_name,
    u.email,
    c.id as class_id,
    co.name as course_name,
    co.code as course_code,
    CONCAT(i.first_name, ' ', i.last_name) as instructor_name,
    (SELECT COUNT(*) FROM attendances where student_id = u.id AND class_id = c.id AND is_present = true) as classes_attended,
    (SELECT COUNT(*) FROM attendances where student_id = u.id AND class_id = c.id) as total_classes,
    COALESCE(
      CAST((SELECT COUNT(*) FROM attendances where student_id = u.id AND class_id = c.id AND is_present = true) AS FLOAT) * 100.0 /
      NULLIF((SELECT COUNT(*) FROM attendances where student_id = u.id AND class_id = c.id), 0),
      0
    ) as attendance_percentage
  FROM users u
  JOIN enrollments e ON e.student_id = u.id
  JOIN classes c ON e.class_id = c.id
  JOIN courses co ON c.course_id = co.id
  JOIN users i ON c.instructor_id = i.id
  WHERE c.id = ?
  GROUP BY
    u.id, u.first_name, u.last_name, u.email,
    c.id, co.name, co.code,
    i.first_name, i.last_name
  HAVING (SELECT COUNT(*) FROM attendances where student_id = u.id AND class_id = c.id) > 0
)
        SELECT 
            student_id,
            first_name,
            last_name,
            email,
            course_name,
            course_code,
            instructor_name,
            classes_attended,
            total_classes,
            ROUND(attendance_percentage::numeric, 2) as attendance_percentage
        FROM student_attendance
        WHERE attendance_percentage < ?
        ORDER BY attendance_percentage ASC
    `, classID, thresholdPercentage).Scan(&lowAttendanceStudents).Error; err != nil {
		return nil, err
	}

	// Add additional attendance details for each student
	for i := range lowAttendanceStudents {
		// Get recent attendance status
		var recentAttendance []struct {
			Date   time.Time
			Status string
		}
		if err := r.db.Raw(`
            SELECT 
                date,
                CASE WHEN is_present THEN 'Present' ELSE 'Absent' END as status
            FROM attendances
            WHERE student_id = ? AND class_id = ?
            ORDER BY date DESC
            LIMIT 5
        `, lowAttendanceStudents[i].StudentID, classID).Scan(&recentAttendance).Error; err != nil {
			return nil, err
		}

		// Format recent attendance for notification
		var recentStatus []string
		for _, att := range recentAttendance {
			recentStatus = append(recentStatus, fmt.Sprintf("%s: %s",
				att.Date.Format("Jan 02"), att.Status))
		}
		lowAttendanceStudents[i].RecentAttendance = recentStatus

		// Calculate classes needed to reach threshold
		currentAttendance := float64(lowAttendanceStudents[i].ClassesAttended)
		totalClasses := float64(lowAttendanceStudents[i].TotalClasses)
		targetAttendance := (thresholdPercentage * totalClasses) / 100

		classesNeeded := int(math.Ceil(targetAttendance - currentAttendance))
		if classesNeeded < 0 {
			classesNeeded = 0
		}
		lowAttendanceStudents[i].ClassesNeededForThreshold = classesNeeded
	}

	return lowAttendanceStudents, nil
}
