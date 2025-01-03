package service

import (
	"fmt"
	"os"

	"github.com/ams-api/internal/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

type ISendMail interface {
	SendAttendanceAlert()
	SendAttendanceAlertAccordingToUserType(userId uint, req *models.SendAlertRequest)
}

func (s Service) SendAttendanceAlert() {
	// Example usage
	apiKey := os.Getenv("SENDGRID_API_KEY")
	class, err := s.repo.FindAllClass(&models.ListClassRequest{
		InstructorId: 0})
	if err != nil {
		logrus.Errorf("error getting class: %v", err)
		return
	}
	for _, c := range *class {
		threshold := 75.0
		lowAttendanceStudents, err := s.repo.GetStudentsWithLowAttendance(c.ID, threshold)
		if err != nil {
			logrus.Errorf("error getting low attendance students: %v", err)
			return
		}
		for _, student := range lowAttendanceStudents {
			err := SendAttendanceAlert(apiKey, student, threshold)
			if err != nil {
				logrus.Printf("Error sending alert to %s %s: %v",
					student.FirstName, student.LastName, err)
				continue
			}
			logrus.Printf("Successfully sent attendance alert to %s %s",
				student.FirstName, student.LastName)
		}
	}
}

func (s Service) SendAttendanceAlertAccordingToUserType(userId uint, req *models.SendAlertRequest) {
	// Example usage
	apiKey := os.Getenv("SENDGRID_API_KEY")
	user, err := s.repo.FindUserById(userId)
	if err != nil {
		logrus.Errorf("error getting user: %v", err)
		return
	}
	institutionId := user.InstitutionId
	var class *[]models.Class
	if req.UserType == "teacher" {
		class, err = s.repo.FindClassByInstructor(userId)
		if err != nil {
			logrus.Errorf("error getting class: %v", err)
			return
		}
	}
	if req.UserType == "institution_admin" {
		class, err = s.repo.FindClassByInstitutionId(institutionId)
		if err != nil {
			logrus.Errorf("error getting class: %v", err)
			return
		}
	}
	if class != nil || len(*class) != 0 {
		for _, c := range *class {
			threshold := req.Threshold
			lowAttendanceStudents, err := s.repo.GetStudentsWithLowAttendance(c.ID, threshold)
			if err != nil {
				logrus.Errorf("error getting low attendance students: %v", err)
				return
			}
			for _, student := range lowAttendanceStudents {
				err := SendAttendanceAlert(apiKey, student, threshold)
				if err != nil {
					logrus.Printf("Error sending alert to %s %s: %v",
						student.FirstName, student.LastName, err)
					continue
				}
				logrus.Printf("Successfully sent attendance alert to %s %s",
					student.FirstName, student.LastName)
			}
		}
	}
}

// FormatAttendanceEmailHTML generates a well-formatted HTML email for attendance alerts
func FormatAttendanceEmailHTML(student models.StudentAttendanceAlert, threshold float64) string {
	// Create the attendance status bar color based on percentage
	var statusColor string
	switch {
	case student.AttendancePercentage < 50:
		statusColor = "#dc3545" // red
	case student.AttendancePercentage < 75:
		statusColor = "#ffc107" // yellow
	default:
		statusColor = "#28a745" // green
	}

	// Format recent attendance with colored status
	var recentAttendanceHTML string
	for _, att := range student.RecentAttendance {
		status := "Present"
		if att[len(att)-6:] == "Absent" {
			status = "Absent"
		}
		// statusClass := status == "Present" ? "color: #28a745;" : "color: #dc3545;"
		if status == "Present" {
			statusClass := "color: #28a745;"
			recentAttendanceHTML += fmt.Sprintf(`<tr>
				<td style="padding: 8px; border-bottom: 1px solid #ddd;">%s</td>
				<td style="padding: 8px; border-bottom: 1px solid #ddd; %s">%s</td>
			</tr>`, att[:8], statusClass, status)
		} else {
			statusClass := "color: #dc3545;"
			recentAttendanceHTML += fmt.Sprintf(`<tr>
				<td style="padding: 8px; border-bottom: 1px solid #ddd;">%s</td>
				<td style="padding: 8px; border-bottom: 1px solid #ddd; %s">%s</td>
			</tr>`, att[:8], statusClass, status)
		}
		// recentAttendanceHTML += fmt.Sprintf(`<tr>
		//     <td style="padding: 8px; border-bottom: 1px solid #ddd;">%s</td>
		//     <td style="padding: 8px; border-bottom: 1px solid #ddd; %s">%s</td>
		// </tr>`, att[:8], statusClass, status)
	}

	htmlContent := fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <title>Attendance Alert</title>
    </head>
    <body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
        <div style="background-color: #f8f9fa; padding: 20px; border-radius: 5px; margin-bottom: 20px;">
            <h2 style="color: #333; margin: 0;">Attendance Alert</h2>
            <p style="color: #666;">Course: %s (%s)</p>
        </div>

        <p>Dear %s %s,</p>

        <p>This is regarding your attendance in the above course.</p>

        <div style="background-color: #fff; border: 1px solid #ddd; border-radius: 5px; padding: 20px; margin: 20px 0;">
            <h3 style="margin-top: 0;">Attendance Summary</h3>
            
            <!-- Attendance Progress Bar -->
            <div style="background-color: #f0f0f0; border-radius: 10px; height: 20px; margin: 15px 0;">
                <div style="background-color: %s; width: %.1f%%; height: 100%%; border-radius: 10px; max-width: 100%%;">
                    <span style="color: white; font-size: 12px; line-height: 20px; padding: 0 10px;">%.1f%%</span>
                </div>
            </div>

            <table style="width: 100%%; border-collapse: collapse; margin: 15px 0;">
                <tr>
                    <td style="padding: 8px 0;">Current Attendance:</td>
                    <td style="padding: 8px 0;"><strong>%.1f%%</strong></td>
                </tr>
                <tr>
                    <td style="padding: 8px 0;">Classes Attended:</td>
                    <td style="padding: 8px 0;"><strong>%d out of %d</strong></td>
                </tr>
                <tr>
                    <td style="padding: 8px 0;">Required Attendance:</td>
                    <td style="padding: 8px 0;"><strong>%.0f%%</strong></td>
                </tr>
                <tr>
                    <td style="padding: 8px 0;">Classes Needed:</td>
                    <td style="padding: 8px 0;"><strong>%d more classes</strong></td>
                </tr>
            </table>
        </div>

        <div style="background-color: #fff; border: 1px solid #ddd; border-radius: 5px; padding: 20px; margin: 20px 0;">
            <h3 style="margin-top: 0;">Recent Attendance History</h3>
            <table style="width: 100%%; border-collapse: collapse;">
                <tr>
                    <th style="padding: 8px; text-align: left; border-bottom: 2px solid #ddd;">Date</th>
                    <th style="padding: 8px; text-align: left; border-bottom: 2px solid #ddd;">Status</th>
                </tr>
                %s
            </table>
        </div>

        <div style="margin: 20px 0;">
            <p><strong>Action Required:</strong></p>
            <p>To maintain the required attendance of %.0f%%, please ensure regular attendance in upcoming classes. If you have any concerns, please contact your instructor.</p>
        </div>

        <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #ddd; color: #666;">
            <p>Regards,<br>%s<br>Course Instructor</p>
        </div>
    </body>
    </html>
    `,
		student.CourseName,
		student.CourseCode,
		student.FirstName,
		student.LastName,
		statusColor,
		student.AttendancePercentage,
		student.AttendancePercentage,
		student.AttendancePercentage,
		student.ClassesAttended,
		student.TotalClasses,
		threshold,
		student.ClassesNeededForThreshold,
		recentAttendanceHTML,
		threshold,
		student.InstructorName)

	return htmlContent
}

// SendAttendanceAlert sends attendance alert email to a student
func SendAttendanceAlert(apiKey string, student models.StudentAttendanceAlert, threshold float64) error {
	config := EmailConfig{
		FromEmail: "bpokhrel140@gmail.com",
		FromName:  "AMS - Attendance Management System",
		ToEmail:   student.Email,
		ToName:    fmt.Sprintf("%s %s", student.FirstName, student.LastName),
		Subject:   fmt.Sprintf("Low Attendance Alert - %s (%s)", student.CourseName, student.CourseCode),
		PlainContent: fmt.Sprintf("Your attendance in %s is %.1f%%. Please check your email for detailed information.",
			student.CourseName, student.AttendancePercentage),
		HTMLContent: FormatAttendanceEmailHTML(student, threshold),
	}

	return SendEmail(apiKey, config)
}

type EmailConfig struct {
	FromEmail    string
	FromName     string
	ToEmail      string
	ToName       string
	Subject      string
	PlainContent string
	HTMLContent  string
}

// SendEmail sends an email using SendGrid
func SendEmail(apiKey string, config EmailConfig) error {
	from := mail.NewEmail(config.FromName, config.FromEmail)
	to := mail.NewEmail(config.ToName, config.ToEmail)

	message := mail.NewSingleEmail(from, config.Subject, to, config.PlainContent, config.HTMLContent)
	client := sendgrid.NewSendClient(apiKey)

	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("API error: status code %d, body: %s", response.StatusCode, response.Body)
	}

	return nil
}
