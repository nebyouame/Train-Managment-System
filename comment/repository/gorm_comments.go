package repository

import (
	"TrainProject/comment"
	"TrainProject/entity"
	"github.com/jinzhu/gorm"
)

type CommentGormRepo struct {
	conn *gorm.DB
}


func NewCommentGormRepo(db *gorm.DB) comment.CommentRepository {
	return &CommentGormRepo{conn: db}
}


func (cmntRepo *CommentGormRepo) Comments() ([]entity.Comment, []error) {
	cmnts := []entity.Comment{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}


func (cmntRepo *CommentGormRepo) Comment(id uint) (*entity.Comment, []error) {
	cmnt := entity.Comment{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}


func (cmntRepo *CommentGormRepo) UpdateComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}


func (cmntRepo *CommentGormRepo) DeleteComment(id uint) (*entity.Comment, []error) {
	cmnt, errs := cmntRepo.Comment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}


func (cmntRepo *CommentGormRepo) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

