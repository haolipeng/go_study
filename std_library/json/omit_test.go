package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func TestOmitJson(t *testing.T) {
	user := User{
		Email: "1078285863@.qq.com",
		//Password: "123456",
	}

	buf, err := json.Marshal(user)
	if err != nil {
		t.Log("json.Marshal failed!")
	}

	fmt.Println(string(buf))
}
