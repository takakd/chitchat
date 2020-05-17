package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAd time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.Uuid, &conv.CreatedAd); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}
