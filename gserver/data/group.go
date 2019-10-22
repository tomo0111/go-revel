package data

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var grInstance GroupRepository

type GroupRepository interface {
	// Get groups all data
	FindAll() ([]*entity.Group, *model.ErrorResBody)

	// Get group from groups.name
	FindByName(name string) (*entity.Group, *model.ErrorResBody)

	// Generate groups, user_groups, service_groups
	SaveWithUserGroupWithServiceGroup(group entity.Group) (*entity.Group, *model.ErrorResBody)
}

type GroupRepositoryImpl struct {
	Db *gorm.DB
}

func GetGroupRepositoryInstance(db *gorm.DB) GroupRepository {
	if grInstance == nil {
		grInstance = NewGroupRepository(db)
	}
	return grInstance
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	log.Logger.Info("New `GroupRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `GroupRepository`")
	return GroupRepositoryImpl{Db: db}
}


func (gr GroupRepositoryImpl) FindAll() ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	if err := gr.Db.Find(&groups).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindByName(name string) (*entity.Group, *model.ErrorResBody) {
	var groups *entity.Group
	if err := gr.Db.Where("name = ?", name).Find(&groups).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) SaveWithUserGroupWithServiceGroup(group entity.Group) (*entity.Group, *model.ErrorResBody) {
	tx := gr.Db.Begin()

	// Save groups
	if err := tx.Create(&group).Error; err != nil {
		log.Logger.Warn("Failed to save groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit user data.")
		}

		return nil, model.InternalServerError()
	}

	// Save service_groups
	serviceGroup := entity.ServiceGroup{
		GroupId: group.Id,
		ServiceId: ctx.GetServiceId().(int),
	}
	if err := tx.Create(&serviceGroup).Error; err != nil {
		log.Logger.Warn("Failed to save service_groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service data.")
		}

		return nil, model.InternalServerError()
	}

	// Save user_groups
	userGroup := entity.UserGroup{
		UserId: ctx.GetUserId().(int),
		GroupId: group.Id,
	}
	if err := tx.Create(&userGroup).Error; err != nil {
		log.Logger.Warn("Failed to save user_groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service data.")
		}

		return nil, model.InternalServerError()
	}

	tx.Commit()

	return &group, nil
}
