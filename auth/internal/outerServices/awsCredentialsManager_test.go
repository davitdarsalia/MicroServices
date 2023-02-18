package outerServices

import (
	"fmt"
	"testing"
)

func TestGetAWSCredentials(t *testing.T) {
	a := GetAWSCredentials()
	fmt.Println(a)
}
