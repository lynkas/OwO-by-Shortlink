package main

var UnID = int64(-1)

type Record struct {
	From string `json:"from"`
	Identify string `json:"identify"`
	Format string `json:"format"`
	Content string `json:"content"`
	TextSerial string `json:"text_serial"`
	Prefix string `json:"prefix"`
	Serial int64 `json:"serial"`
}

type MessageGroup struct {
	Format string `json:"format"`
	Content string `json:"content"`
}

type NewRecord struct {
	Token string `json:"token"`
	Identify string `json:"identify"`
	Format string `json:"format"`
	Content string `json:"content"`
}

type IDGroup struct {
	TextSerial string `json:"text_serial"`
	Prefix string `json:"prefix"`
	Serial int64 `json:"serial"`
}

type User struct {
	From string `json:"from"`
	Identify string `json:"identify"`
}

func (r *Record)toMessageGroup() *MessageGroup {
	return &MessageGroup{
		Format:  r.Format,
		Content: r.Content,
	}
}

func  (r *Record)toUser() *User {
	return &User{
		From:     r.From,
		Identify: r.Identify,
	}
}

func (r *Record)toIDGroup() *IDGroup {
	return &IDGroup{
		TextSerial: r.TextSerial,
		Prefix:     r.Prefix,
		Serial:     r.Serial,
	}
}