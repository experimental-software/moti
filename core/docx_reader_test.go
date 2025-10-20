package core

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	result := ReadDocxFile("./t/man-know-thyself.docx")
	fmt.Print(result)
}
