package student

import (
	"testing"
)

func TestGetDeatils(t *testing.T) {
	t.Run("Get details of student with id 1", func(t *testing.T) {
		got, _ := GetStudentDetail("1")
		want := 200
		if got.status != want {
			t.Errorf("Got %d ,want %d", got.status, want)
		}
	})
	t.Run("Get the name of student with id 1", func(t *testing.T) {
		got, _ := GetStudentDetail("1")
		want := "Leanne Graham"

		if got.student.Name != want {
			t.Errorf("Got %s ,want %s", got.student.Name, want)
		}
	})
}
