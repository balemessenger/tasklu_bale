package mock

const (
	Activities = `{
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

	Session =`
{
  "ok": true,
  "status": "CREATED",
  "data": {
    "type": "User",
    "id": "5b029f6056ad663c26003ed7",
    "username": "person820",
    "name": "Bryce",
    "images": {
      "medium": "https://test.taskulu.com/static/users-0/medium.png",
      "large": "https://test.taskulu.com/static/users-0/large.png",
      "big": "https://test.taskulu.com/static/users-0/big.png"
    },
    "email": "person820@example.com",
    "locale": {
      "calendar": "gregorian",
      "language": "en",
      "time_zone": "UTC"
    },
    "api_key": "294b5885ca8406fad39dd03a82cfb042c87c4f7",
    "app_key": "c4d84C6g4CYRrgUsDJ_wDHrgGOW-YUWK39dV2fIYeT3teW4wwhcXaYFdrmpz_yN7WgGj1QbNC03BuPxlI7KiIwH4nPCXxGiEz5Fz-xsdC2OAu9OY7SQMzkj8rWEhIvRk",
    "session_id": "fa8fd25662d1e8bf301d501e31434a78"
  }
}
`
)