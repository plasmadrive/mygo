package encoding




func TestPostJsonDecode(t *testing.T) {
	testData := `{ "id": 1,"content": "Hello,World!", "author" : {"id": 1, "name": "Gareth"}, "comments": [{"id": 1,"content": "First comment", "author": "Sandra"}, {"id": 2,"content": "Second comment", "author": "Paul"}]}`
	post,err := decode([]byte(testData))
	if err != nil {
		t.Error(err)
	}

	if post.Id != 1 {
		t.Error("Wrong post id. Expecting 1 but got",post.Id)
	}

	if post.Content != "Hello,World!" {
		t.Error("Wrong post content. Expecting `Hello,World!`, but got",post.Content)
	}


	
	if post.Author.Id != 1 {
		t.Error("Wrong author id. Expecting 1 but got",post.Author.Id)
	}

	if post.Author.Name != "Gareth" {
		t.Error("Wrong author name. Expecting  `Gareth` but got",post.Author.Name)
	}

	if len( post.Comments) != 2  {
		t.Fatal("Expected  `comments` array to have two elements but got",len(post.Comments))
	}

	if post.Comments[0].Id != 1 {
		t.Error("Expected comment id to be 1 but got", post.Comments[0].Id)
	}

	if post.Comments[0].Content != "First comment" {
		t.Error("Expected comment content to be `First post` but got",post.Comments[0].Content)
	}

	if post.Comments[0].Author != "Sandra" {
		t.Error("Expected author to be `Sandra` but got",post.Comments[0].Author)
	}
}


func TestPostJsonEncode(t *testing.T) {
	post := Post {
		Id : 1,
		Content: "Hello,World!", 
		Author : Author{Id: 1, Name: "Gareth"},
		Comments: []Comment {
			Comment{Id: 1,Content: "First comment", Author: "Sandra"}, 
			Comment{Id: 2,Content: "Second comment", Author: "Paul"},
		},
	}
	
	json ,err := encode(post)
	if err == nil {
		t.Fatal("Error marshalling post to bytes",err)
	}

	compareData := `{ "id": 1,"content": "Hello,World!", "author" : {"id": 1, "name": "Gareth"}, "comments": [{"id": 1,"content": "First comment", "author": "Sandra"}, {"id": 2,"content": "Second comment", "author": "Paul"}]}`
1

	if json  != compareData {
		t.Errorf("Expected  %s but got %s",`{ "id": 1,"content": "Hello,World!", "author" : {"id": 1, "name": "Gareth"}, "comments": [{"id": 1,"content": "First comment", "author": "Sandra"}, {"id": 2,"content": "Second comment", "author": "Paul"}]}`,json)
	}

	

}
