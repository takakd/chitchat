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


// @note: why is this method defined here.
// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a rrow and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAd)
	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
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

// Get a thread by UUID
func TheadByUUID(uuid string)(conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAd )
	return
}
