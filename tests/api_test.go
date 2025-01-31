package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
)

const ADDR string = "http://localhost:8000"

func Test_GetTest_Should200(t *testing.T) {
    client := resty.New()

    response, _ := client.R().Get(ADDR + "/test/hello")

    if response.StatusCode() != 200 {
        t.Fatalf("Response code: Got: %d, want: %d\n", response.StatusCode(), 200)
    }

    want := []byte("Hello string")

    if !reflect.DeepEqual(response.Body(), want) {
        t.Fatalf("Body:\n\tGot: '%s'\n\twant: '%s'\n", string(response.Body()), string(want))
    }
}

func Test_IsUserExist_Should400_InvalidData(t *testing.T) {
    client := resty.New()

    response, _ := client.R().SetBody(map[string]string{"emil": "not-exist@user.domain"}).Post(ADDR + "/user/is_exist")

    if response.StatusCode() != 400 {
        t.Fatalf("Response code: Got: %d, want: %d\n", response.StatusCode(), 400)
    }

    want := map[string]string{"status": "not exist", "message": "Invalid data"}
    var result map[string]string
    err := json.Unmarshal(response.Body(), &result)

    if err != nil {
        t.Fatalf("Got error on Unmarshal = '%#v', want = nil\n", err)
    }

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("Body:\n\tGot: '%#v'\n\twant: '%#v'\n", result, want)
    }
}

func Test_IsUserExist_Should400_NotExist(t *testing.T) {
    client := resty.New()

    response, _ := client.R().SetBody(map[string]string{"email": "not-exist@user.domain"}).Post(ADDR + "/user/is_exist")

    if response.StatusCode() != 400 {
        t.Fatalf("Response code: Got: %d, want: %d\n", response.StatusCode(), 400)
    }

    want := map[string]string{"status": "not exist", "message": "Can't find user in DB"}
    var result map[string]string
    err := json.Unmarshal(response.Body(), &result)

    if err != nil {
        t.Fatalf("Got error on Unmarshal = '%#v', want = nil\n", err)
    }

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("Body:\n\tGot: '%#v'\n\twant: '%#v'\n", result, want)
    }
}

func Test_RegisterUser_Should400_InvalidData(t *testing.T) {
    client := resty.New()

    response, _ := client.R().SetBody(map[string]string{"wrongField": "user_1@someEmail.domain", "password": "12345"}).Post(ADDR + "/user/")

    if response.StatusCode() != 400 {
        t.Fatalf("Response code:\n\tGot: %d\n\twant: %d\n", response.StatusCode(), 400)
    }

    want := map[string]string{"status": "error", "message": "Invalid data"}
    var result map[string]string
    err := json.Unmarshal(response.Body(), &result)

    if err != nil {
        t.Fatalf("Got error on Unmarshal = '%#v', want = nil\n", err)
    }

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("Body:\n\tGot: '%#v'\n\twant: '%#v'\n", result, want)
    }
}

