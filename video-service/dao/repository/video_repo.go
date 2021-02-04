// Code generated by sqlxgen. DO NOT EDIT.
package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type videoRepo struct {
	db *sqlx.DB
}

// execute fail return -1 and err
// execute success return count rows and nil
func (repo *videoRepo) Count() (int64, error) {
	var count int64
	err := repo.db.Get(&count, "select count(*) from video")
	if err != nil {
		return -1, err
	}
	return count, nil
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *videoRepo) Insert(video *Video) (int64, error) {
	result, err := repo.db.Exec("insert into video(id, access_key, create_time, update_time, delete_status, customer_id, video_name, video_introduction, video_cover_path, video_source_path, video_size, video_thumbs_count, video_play_count, video_comment_count) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", video.Id, video.AccessKey, video.CreateTime, video.UpdateTime, video.DeleteStatus, video.CustomerId, video.VideoName, video.VideoIntroduction, video.VideoCoverPath, video.VideoSourcePath, video.VideoSize, video.VideoThumbsCount, video.VideoPlayCount, video.VideoCommentCount)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *videoRepo) UpdateByPrimaryKey(primaryKey int64, video *Video) (int64, error) {
	result, err := repo.db.Exec("update video set id=?, access_key=?, create_time=?, update_time=?, delete_status=?, customer_id=?, video_name=?, video_introduction=?, video_cover_path=?, video_source_path=?, video_size=?, video_thumbs_count=?, video_play_count=?, video_comment_count=? where id = ?", video.Id, video.AccessKey, video.CreateTime, video.UpdateTime, video.DeleteStatus, video.CustomerId, video.VideoName, video.VideoIntroduction, video.VideoCoverPath, video.VideoSourcePath, video.VideoSize, video.VideoThumbsCount, video.VideoPlayCount, video.VideoCommentCount, primaryKey)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *videoRepo) DeleteByPrimaryKey(primaryKey int64) (int64, error) {
	result, err := repo.db.Exec("delete from video where id = ?", primaryKey)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return nil and err
// execute success but no record return nil and nil
// execute success has record return data and nil
func (repo *videoRepo) SelectByPrimaryKey(primaryKey int64) (*Video, error) {
	video := new(Video)
	err := repo.db.Get(video, "select id, access_key, create_time, update_time, delete_status, customer_id, video_name, video_introduction, video_cover_path, video_source_path, video_size, video_thumbs_count, video_play_count, video_comment_count from video where id = ?", primaryKey)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return video, nil
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *videoRepo) UpdateByExample(example *VideoExample, video *Video) (int64, error) {
	criteria := example.criteria
	var params []interface{}

	params = append(params, video.Id)
	params = append(params, video.AccessKey)
	params = append(params, video.CreateTime)
	params = append(params, video.UpdateTime)
	params = append(params, video.DeleteStatus)
	params = append(params, video.CustomerId)
	params = append(params, video.VideoName)
	params = append(params, video.VideoIntroduction)
	params = append(params, video.VideoCoverPath)
	params = append(params, video.VideoSourcePath)
	params = append(params, video.VideoSize)
	params = append(params, video.VideoThumbsCount)
	params = append(params, video.VideoPlayCount)
	params = append(params, video.VideoCommentCount)

	var condition = ""
	if len(criteria) > 0 {
		var fragments []string
		for _, e := range criteria {
			fragments = append(fragments, e.fragment)
			if !e.noValue {
				params = append(params, e.param1)
			}
			if e.betweenValue {
				params = append(params, e.param2)
			}
		}
		condition += "where " + strings.TrimLeft(strings.Join(fragments, " "), "and")
	}
	query, args, err := sqlx.In("update video set id=?, access_key=?, create_time=?, update_time=?, delete_status=?, customer_id=?, video_name=?, video_introduction=?, video_cover_path=?, video_source_path=?, video_size=?, video_thumbs_count=?, video_play_count=?, video_comment_count=? "+condition, params...)
	if err != nil {
		return -1, err
	}
	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *videoRepo) DeleteByExample(example *VideoExample) (int64, error) {
	criteria := example.criteria
	var params []interface{}
	var condition = ""
	if len(criteria) > 0 {
		var fragments []string
		for _, e := range criteria {
			fragments = append(fragments, e.fragment)
			if !e.noValue {
				params = append(params, e.param1)
			}
			if e.betweenValue {
				params = append(params, e.param2)
			}
		}
		condition += "where " + strings.TrimLeft(strings.Join(fragments, " "), "and")
	}
	query, args, err := sqlx.In("delete from video "+condition, params...)
	if err != nil {
		return -1, err
	}
	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return empty slice and err
// execute success but no record return empty slice and nil
// execute success has record return data and nil
func (repo *videoRepo) SelectByExample(example *VideoExample) ([]*Video, error) {
	var list []*Video
	criteria := example.criteria
	var params []interface{}
	var condition = ""
	if len(criteria) > 0 {
		var fragments []string
		for _, e := range criteria {
			fragments = append(fragments, e.fragment)
			if !e.noValue {
				params = append(params, e.param1)
			}
			if e.betweenValue {
				params = append(params, e.param2)
			}
		}
		condition += "where " + strings.TrimLeft(strings.Join(fragments, " "), "and")
	}
	query, args, err := sqlx.In("select id, access_key, create_time, update_time, delete_status, customer_id, video_name, video_introduction, video_cover_path, video_source_path, video_size, video_thumbs_count, video_play_count, video_comment_count from video "+condition, params...)
	if err != nil {
		return list, err
	}
	return list, repo.db.Select(&list, query, args...)
}
