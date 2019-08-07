package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"taskulu/internal"
	"testing"
)

func TestTaskuluResponse(t *testing.T) {
	projectId := "123456"
	input := `{
    "ok": true,
    "status": "OK",
    "data": [
        {
            "by": "56151bcafa1bc7a1810027ca",
            "content": {
                "keys": [
                    {
                        "type": "text",
                        "ids": {
                            "project_id": "5a8d1fff56ad660b0dd0d343",
                            "task_id": "5d46b04456ad667202008c23"
                        },
                        "value": "راه اندازی gitlab واسه CICD"
                    },
                    {
                        "type": "text",
                        "value": "Done"
                    },
                    {
                        "type": "text",
                        "value": "Doing"
                    }
                ],
                "message": "شما وضعیت کار %[0] را از %[1] به %[2] تغییر دادید."
            },
            "created_at": 1565091225
        }
	]
}`
	mockTaskulu.AddHandler(func(writer http.ResponseWriter, bytes []byte) {
		_, err := writer.Write([]byte(input))
		if err != nil {
			t.Error(err)
		}
	})

	taskulu := internal.NewTaskulu("http://127.0.0.1:12346")

	err, b := taskulu.GetActivities("sdfsdf", "ssdfsdfsd", projectId)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, b.Data[0].By, "56151bcafa1bc7a1810027ca")
	assert.Equal(t, b.Data[0].CreatedAt, 1565091225)

}

func TestRealResponse(t *testing.T) {
	taskulu := internal.NewTaskulu("https://taskulu.com")
	err, body := taskulu.GetActivities("oRjJhqNBKGsSeZz5DOfbTAxCV_qAqalolbQMqLqisW7OmVyKvf5cxQYDpiSwGAePf5WBx74jH9IP09_QKxa4xhTDVTWurosLWpOK6VFGzIsRVLighsGEL_KOyXJe9on7", "77a426f5d99770633459fcb99dbb2975", "5a8d1fff56ad660b0dd0d343")

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, body.Ok, true)
	assert.Equal(t, body.Status, "OK")

}

func TestBaleToken(t *testing.T) {
	mockBaleHook.AddHandler(func(w http.ResponseWriter, bytes []byte) {
		assert.Equal(t, string(bytes), "{\"text\":\"salam\"}")
	})
	bale := internal.NewBale("http://127.0.0.1:12345", "1234")
	err := bale.Send("salam")
	if err != nil {
		t.Error(err)
	}
}
