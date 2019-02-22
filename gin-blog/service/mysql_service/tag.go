package mysql_service

import (
	"gin-blog/models"
)

func GetTags(pageNum int, pageSize int, maps interface{}) []models.Tag {
	tags := models.GetTags(pageNum, pageSize, maps)
	return tags
}

func GetTagTotal(maps interface{}) int {
	count := models.GetTagTotal(maps)
	return count

}

func AddTag(t *models.Tag) (int, error) {
	id, err := models.AddTag(t)
	return id, err
}
