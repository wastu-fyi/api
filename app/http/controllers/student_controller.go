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
		return resp.InternalServerError(ctx, "Gagal menghitung total data mahasiswa. Mohon coba lagi beberapa saat.",
			resp.WithMessage("count query error: "+err.Error()),
		)
	}

	var rows []models.Student
	if err := tx.Limit(limit).Offset(offset).Get(&rows); err != nil {
		return resp.InternalServerError(ctx, "Gagal mengambil daftar mahasiswa dari basis data.",
			resp.WithMessage("select query error: "+err.Error()),
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
	return resp.OK(ctx, items, "Daftar mahasiswa berhasil dimuat.", resp.WithMeta(meta))
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
		return resp.InternalServerError(ctx, "Gagal mengambil rekap program studi.", resp.WithMessage("aggregate query error: "+err.Error()))
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
	return resp.OK(ctx, items, "Data program studi berhasil dimuat.")
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
		return resp.InternalServerError(ctx, "Gagal mengambil rekap tahun masuk mahasiswa.", resp.WithMessage("aggregate query error: "+err.Error()))
	}

	type item struct {
		Year     int   `json:"year"`
		Students int64 `json:"students"`
	}
	items := make([]item, 0, len(rows))
	for _, r := range rows {
		items = append(items, item{Year: r.AdmissionYear, Students: r.Total})
	}
	return resp.OK(ctx, items, "Data tahun masuk berhasil dimuat.")
}

func (c *StudentController) Show(ctx http.Context) http.Response {
	idParam := strings.TrimSpace(ctx.Request().Query("id", ""))
	nimParam := strings.TrimSpace(ctx.Request().Query("student_id", ""))
	if idParam != "" && nimParam != "" {
		return resp.BadRequest(ctx, "Harap pilih salah satu parameter: id atau student_id, tidak keduanya.", nil)
	}
	if idParam == "" && nimParam == "" {
		return resp.BadRequest(ctx, "Parameter id atau student_id wajib diisi.", nil)
	}

	db := facades.Orm().Query()
	var s models.Student
	if idParam != "" {
		if _, err := strconv.ParseUint(idParam, 10, 64); err != nil {
			return resp.BadRequest(ctx, "Format id tidak valid. Gunakan angka positif.", nil)
		}
		if err := db.Model(&models.Student{}).Where("id = ?", idParam).First(&s); err != nil {
			return resp.NotFound(ctx, "Data mahasiswa dengan id tersebut tidak ditemukan.")
		}
	} else {
		if err := db.Model(&models.Student{}).Where("nim = ?", nimParam).First(&s); err != nil {
			return resp.NotFound(ctx, "Data mahasiswa dengan student_id tersebut tidak ditemukan.")
		}
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

	type studentDetail struct {
		ID             uint    `json:"id"`
		StudentID      uint64  `json:"student_id"`
		Name           string  `json:"name"`
		Gender         string  `json:"gender"`
		AdmissionYear  int     `json:"admission_year"`
		EducationLevel string  `json:"education_level"`
		StudyCode      int     `json:"study_code"`
		StudyName      string  `json:"study_name"`
		Status         *string `json:"status"`
	}

	detail := studentDetail{
		ID:             s.ID,
		StudentID:      studentIDNum,
		Name:           s.Name,
		Gender:         gender,
		AdmissionYear:  s.AdmissionYear,
		EducationLevel: enums.StudyLevel(s.Study),
		StudyCode:      s.Study,
		StudyName:      enums.StudyName(s.Study),
		Status:         s.Status,
	}
	return resp.OK(ctx, detail, "Detail mahasiswa berhasil dimuat.")
}
