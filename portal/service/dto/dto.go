// author: JGZ
// time:   2021-02-04 10:00
package dto

// 用户信息dto
type UserInfoDto struct {
	Id        int64
	AccessKey string
	Username  string
	Email     string
	Nickname  string
}

// 上传视频dto
type UploadVideoDto struct {
	VideoSourcePath   string
	VideoCoverPath    string
	VideoSize         int64
	VideoName         string
	VideoIntroduction string
}
