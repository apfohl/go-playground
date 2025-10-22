package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func Test(t *testing.T) {
	fmt.Println(strings.SplitN("user_info->>'first_name'", "->>", 2))
	fmt.Println(strings.SplitN("name", "->>", 2))
}

func Test_ReplaceAll(t *testing.T) {
	s := strings.ReplaceAll("", "foo", "bar")
	println(s)
}

func Test_escape(t *testing.T) {
	fmt.Printf("%%\"user_id\": \"%s\"%%", "foo")
}

func alterArray(slice *[]string) {
	*slice = append(*slice, "foo")
}

func Test_pass_array(t *testing.T) {
	slice := []string{"bar"}
	alterArray(&slice)
	fmt.Println(slice)
}

func Test_regex_empty(t *testing.T) {
	println(regexp.MustCompile(`^[A-Za-z_0-9\-'>]+$`).MatchString("a; DROP TABLE items; --"))
}

func Test_map_not_found(t *testing.T) {
	m := map[string]*string{}
	fmt.Println(m["foo"])
}

func Test_validate_UUID(t *testing.T) {
	input := "courses.course.27cc183b-52b3-4887-8beb-5f744b6457"

	parts := strings.Split(input, ".")

	if len(parts) != 3 {
		print("ERROR")
	}

	print(validation.Validate(&parts[2], is.UUID))
}

func Test_parse_time(t *testing.T) {
	tim, _ := time.Parse(time.RFC3339, "0")
	print(tim.Format(time.RFC3339))
}
