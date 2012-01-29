package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getPosts(path string, params map[string]string) (output []Post, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(postsCollection))
	collection := parsed_response.(*postsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// AllPosts returns all posts in site 
func (session Session) AllPosts(params map[string]string) (output []Post, error os.Error) {
	return session.getPosts("posts", params)
}

// Posts returns the posts with the given ids
func (session Session) Posts(ids []int, params map[string]string) (output []Post, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";")}, "/")
	return session.getPosts(request_path, params)
}