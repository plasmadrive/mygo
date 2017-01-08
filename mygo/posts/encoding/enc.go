package encoding


import (
"mygo/posts"
)
func decode(b []byte) (posts.Post, error) {
	post := posts.Post{}
	err := json.Unmarshal(b,&post)
	return post,err
}

