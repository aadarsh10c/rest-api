package student

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type StudentResponse struct {
	student Student
	status  int
}

func GetStudentDetail(id string) (*StudentResponse, error) {
	//create a proper url
	var url string = fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	//get the response from url
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return &StudentResponse{}, err
	}
	defer res.Body.Close()
	//extract the response body from the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return &StudentResponse{}, err
	}

	//unmarshal the json result into struct
	student := Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		log.Fatal(err)
		return &StudentResponse{}, err
	}

	//return the status code
	return &StudentResponse{student, res.StatusCode}, nil
}
