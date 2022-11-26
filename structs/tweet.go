package structs

import "time"

type Tweet struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userid"`
	Content   string    `json:"content"`
	Img       Image     `json:"img"`
	CreatedAt time.Time `json:"createdAt"`
}

type Image struct {
	Src    string `json:"src"`
	Width  int16  `json:"width"`
	Height int16  `json:"height"`
}
