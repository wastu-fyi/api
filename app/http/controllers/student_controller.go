package controllers

import (
	"strconv"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/app/models"
	"wastu/pkg/enums"
	"wastu/pkg/resp"
)

type StudentController struct{}

func NewStudentController() *StudentController {
	return &StudentController{}
}

func (c *StudentController) Index(ctx http.Context) http.Response {
	q := strings.TrimSpace(ctx.Request().Query("q", ""))
	limitStr := ctx.Request().Query("limit", "10")
	pageStr := ctx.Request().Query("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 25 {
		limit = 25
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit
	db := facades.Orm().Query()
	tx := db.Model(&models.Student{}).Order("CAST(nim AS UNSIGNED) DESC")

	if q != "" {
		tx = tx.Where("name LIKE ?", "%"+q+"%")
	}

	total, err := tx.Count()
	if err != nil {
		return resp.InternalServerError(ctx, "Failed to count students",
			resp.WithMessage(err.Error()),
		)
	}

	var rows []models.Student
	if err := tx.Limit(limit).Offset(offset).Get(&rows); err != nil {
		return resp.InternalServerError(ctx, "Failed to fetch students",
			resp.WithMessage(err.Error()),
		)
	}

	type studentItem struct {
		ID        uint   `json:"id"`
		StudentID uint64 `json:"student_id"`
		Name      string `json:"name"`
		Study     int    `json:"study"`
		StudyName string `json:"study_name"`
	}

	items := make([]studentItem, 0, len(rows))
	for _, s := range rows {
		var nimNum uint64
		if n, err := strconv.ParseUint(s.Nim, 10, 64); err == nil {
			nimNum = n
		}
		items = append(items, studentItem{
			ID:        s.ID,
			StudentID: nimNum,
			Name:      s.Name,
			Study:     s.Study,
			StudyName: enums.StudyName(s.Study),
		})
	}

	meta := resp.BuildPaginationMeta(total, limit, page)
	return resp.OK(ctx, items, "Students fetched", resp.WithMeta(meta))
}

func (c *StudentController) Studies(ctx http.Context) http.Response {
	db := facades.Orm().Query()
	type row struct {
		Study int   `gorm:"column:study"`
		Total int64 `gorm:"column:total"`
	}
	var rows []row
	if err := db.Model(&models.Student{}).
		Select("study, COUNT(*) as total").
		Group("study").
		Order("study ASC").
		Get(&rows); err != nil {
		return resp.InternalServerError(ctx, "Failed to fetch studies", resp.WithMessage(err.Error()))
	}

	type item struct {
		Code           int    `json:"code"`
		Name           string `json:"name"`
		EducationLevel string `json:"education_level"`
		Students       int64  `json:"students"`
	}
	items := make([]item, 0, len(rows))
	for _, r := range rows {
		items = append(items, item{
			Code:           r.Study,
			Name:           enums.StudyName(r.Study),
			EducationLevel: enums.StudyLevel(r.Study),
			Students:       r.Total,
		})
	}
	return resp.OK(ctx, items, "Studies fetched")
}

func (c *StudentController) Years(ctx http.Context) http.Response {
	db := facades.Orm().Query()
	type row struct {
		AdmissionYear int   `gorm:"column:admission_year"`
		Total         int64 `gorm:"column:total"`
	}
	var rows []row
	if err := db.Model(&models.Student{}).
		Select("admission_year, COUNT(*) as total").
		Group("admission_year").
		Order("admission_year DESC").
		Get(&rows); err != nil {
		return resp.InternalServerError(ctx, "Failed to fetch years", resp.WithMessage(err.Error()))
	}

	type item struct {
		Year     int   `json:"year"`
		Students int64 `json:"students"`
	}
	items := make([]item, 0, len(rows))
	for _, r := range rows {
		items = append(items, item{Year: r.AdmissionYear, Students: r.Total})
	}
	return resp.OK(ctx, items, "Years fetched")
}

func (c *StudentController) Show(ctx http.Context) http.Response {
	studentID := strings.TrimSpace(ctx.Request().Route("student_id"))
	if studentID == "" {
		return resp.BadRequest(ctx, "student_id is required", map[string]any{"field": "student_id"})
	}

	db := facades.Orm().Query()
	var s models.Student
	if err := db.Model(&models.Student{}).Where("nim = ?", studentID).First(&s); err != nil {
		return resp.NotFound(ctx, "Student not found")
	}

	gender := ""
	switch s.Gender {
	case 1:
		gender = "M"
	case 2:
		gender = "F"
	}

	var studentIDNum uint64
	if n, err := strconv.ParseUint(s.Nim, 10, 64); err == nil {
		studentIDNum = n
	}

	data := map[string]any{
		"student_id":      studentIDNum,
		"name":            s.Name,
		"gender":          gender,
		"admission_year":  s.AdmissionYear,
		"education_level": enums.StudyLevel(s.Study),
		"study_code":      s.Study,
		"study_name":      enums.StudyName(s.Study),
		"status":          s.Status,
	}
	return resp.OK(ctx, data, "Student fetched")
}
