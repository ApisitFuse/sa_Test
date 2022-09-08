package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
	// 1 user เป็นเจ้าของได้หลาย video
	Videos []Video `gorm:"foreignKey:OwnerID"`
	// 1 user เป็นเจ้าของได้หลาย playlist
	Playlists []Playlist `gorm:"foreignKey:OwnerID"`
}
type Video struct {
	gorm.Model
	Name string
	Url  string `gorm:uniqueIndex`
	// OwnerID ทำหน้าที่เป็น FK
	OwnerID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Owner       User
	WatchVideos []WatchVideo `gorm:"foreignKey:VideoID"`
}

type Playlist struct {
	gorm.Model
	Title string
	// OwnerID ทำหน้าที่เป็น FK
	OwnerID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Owner       User
	WatchVideos []WatchVideo `gorm:"foreingnKey:PlaylistID"`
}

type Resolution struct {
	gorm.Model
	Value       string
	WatchVideos []WatchVideo `gorm:"foreignKey:ResolutionID"`
}
type WatchVideo struct {
	gorm.Model
	WatchTime time.Time

	// Resolution ทำหน้าที่เป็น FK
	ResolutionID *uint
	Resolution   Resolution

	// PlaylistID ทำหน้าที่เป็น FK
	PlaylistID *uint
	Playlist   Playlist

	// VideoID ทำหน้าที่เป็น FK
	VideoID *uint
	Video   Video
}

// type User struct {
// 	gorm.Model
// 	Name      string
// 	Email     string     `gorm:"uniqueIndex"`
// 	Ac_hiss   []Ac_his   `gorm:foreignKey:UserID"`
// 	Activitys []Activity `gorm:foreignKey:UserID"`

// }
// type Ac_his struct {
// 	gorm.Model
// 	Achour     int
// 	date_start time.Time
// 	date_end   time.Time
// 	time_start time.Time
// 	time_end   time.Time
// 	ActivityID *uint
// 	Activity   Activity

// 	UserID *uint
// 	User   User

// 	StudentID *uint
// 	Student   Student
// }

// type Student struct {
// 	gorm.Model
// 	Sname                    string
// 	Surname                  string
// 	dob                      time.Time
// 	Saddress                 string
// 	Sparent                  string
// 	Sidentity_card           int
// 	admission_date           time.Time
// 	Sgrade                   float64
// 	Sphone_num               string
// 	Seducation_qualification string
// 	Sgraduated               string
// 	Ac_hiss                  []Ac_his `gorm:foreignKey:StudentID"`
// }

// type Activity struct {
// 	gorm.Model
// 	Acname     string
// 	date_start time.Time
// 	date_end   time.Time
// 	time_start time.Time
// 	time_end   time.Time
// 	UserID     *uint
// 	User       User
// 	Ac_hiss    []Ac_his `gorm:foreignKey:ActivityID"`
// }
