package models

import "time"

type WhiteEmail struct {
	ID          uint          `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	LifeTime    time.Duration `json:"lifeTime"`
	Email       string        `gorm:"not null; unique_index" json:"email"`
	CreatedByID uint          `gorm:"not null" json:"createdByID"`
	CreatedBy   User          `gorm:"ForeignKey:CreatedByID" json:"createdBy"`
}

const (
	WhiteEmailPermanentLifeTime time.Duration = -1
)

func NewWhiteEmail(email string, user *User) *WhiteEmail {
	e := &WhiteEmail{
		LifeTime:    WhiteEmailPermanentLifeTime,
		Email:       email,
		CreatedByID: user.ID,
	}
	db.Create(e)
	e.CreatedBy = *user
	return e
}

func GetWhiteEmails() []WhiteEmail {
	res := make([]WhiteEmail, 0, 0)
	db.Order("id ASC").Find(&res)
	return res
}

func GetWhiteEmail(email string) *WhiteEmail {
	res := &WhiteEmail{}
	nf := db.Table("white_emails").Where("email = ?", email).Scan(res).RecordNotFound()
	if nf {
		return nil
	}
	return res
}

func DeleteWhiteEmail(id uint) error {
	return db.Delete(WhiteEmail{}, "id = ?", id).Error
}

func (e *WhiteEmail) FetchCreatedBy(email bool) {
	db.Table("users").Where("id = ?", e.CreatedByID).Scan(&e.CreatedBy)
	if !email {
		e.CreatedBy.Email = ""
	}
}
