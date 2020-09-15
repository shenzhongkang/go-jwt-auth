package crud

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gome/api/models"
	"gome/api/utils/channels"
	"time"
)

// RepositoryPostsCRUD is the struct for the Post CRUD
type RepositoryPostsCRUD struct {
	db *gorm.DB
}

// NewRepositoryPostsCRUD returns a new repository with DB connection
func NewRepositoryPostsCRUD(db *gorm.DB) *RepositoryPostsCRUD  {
	return &RepositoryPostsCRUD{db}
}

// Save returns a new post created or an error
func (r *RepositoryPostsCRUD) Save(post models.Post) (models.Post, error)  {
	var err error
	done := make(chan bool)
	go func(ch chan <- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Create(&post).Related(&post.Author, "author_id").Error
		if err != nil {
			ch <- false
			return
		}
		ch <-  true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}

// FindAll returns all the posts from the DB
func (r *RepositoryPostsCRUD) FindAll() ([]models.Post, error) {
	var err error
	posts := []models.Post{}
	done := make(chan bool)
	go func(ch chan <- bool) {
		defer close(ch)
		err = r.db.Debug().Preload("Author").Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}

// FindByID returns a post from the DB
func (r *RepositoryPostsCRUD) FindByID(uid uint32) (models.Post, error)  {
	var err error
	post := models.Post{}
	done := make(chan bool)
	go func(ch chan <- bool) {
		close(ch)
		err = r.db.Debug().Preload("Author").Find(&post, uid).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.Post{}, errors.New("post not found")
	}
	return models.Post{}, err
}

// Update updates a post from the DB
func (r *RepositoryPostsCRUD) Update(uid uint32, post models.Post) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan <- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", uid).Take(&models.Post{}).UpdateColumn(
			map[string]interface{}{
				"title": post.Title,
				"content": post.Content,
				"updated_at": time.Now(),
			},
			)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// Delete removes a post from the DB
func (r *RepositoryPostsCRUD) Delete(uid uint32, authorid uint64) (int64, error)  {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan <- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ? and author_id = ?", uid, authorid).Take(&models.Post{}).Delete(&models.Post{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}