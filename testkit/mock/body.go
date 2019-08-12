package mock

const (
	ChangeStatusActivities = `{
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
                        "value": "github"
                    },
                    {
                        "type": "text",
                        "value": "Doing"
                    },
                    {
                        "type": "text",
                        "value": "Done"
                    }
                ],
                "message": "شما وضعیت کار %[0] را از %[1] به %[2] تغییر دادید."
            },
            "created_at": 1565091225
        },
        {
            "by": "56151bcafa1bc7a1810027ca",
            "content": {
                "keys": [
                    {
                        "type": "text",
                        "ids": {
                            "project_id": "5d088afd56ad6678a4df44dc",
                            "task_id": "5d4d74e856ad667f22044dfc"
                        },
                        "value": "پیگیری رک"
                    },
                    {
                        "type": "text",
                        "value": "مولانا"
                    }
                ],
                "message": "شما کار %[0] را در لیست %[1] ایجاد کردید."
            },
            "created_at": 1565357288
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
Project = `
{
  "ok": true,
  "status": "OK",
  "data": {
    "id": "5b029ec356ad663c2600095d",
    "created_at": 1526898371,
    "created_by": "5b029ec356ad663c2600095c",
    "type": "Project",
    "title": "veniam ut ex",
    "description": "Illum voluptatem aut itaque accusantium.",
    "closed": false,
    "user_roles": [
      {
        "id": "5b029ec356ad663c26000966",
        "name": "admin"
      },
      {
        "id": "5b029ec356ad663c26000965",
        "name": "viewer"
      }
    ],
    "permissions": [
      "edit project",
      "manage sheets"
    ],
    "new_members": [],
    "color_labels": [
      "blue",
      "purple",
      "red",
      "orange",
      "green",
      "pink"
    ],
    "members_count": 1,
    "membership_condition": "approved",
    "attachment_size_limit": 524288000,
    "pinned": false,
    "users": [
      {
        "type": "User",
        "id": "5b029ec356ad663c2600095c",
        "username": "person121",
        "name": "Rylee",
        "images": {
          "medium": "https://test.taskulu.com/static/users-0/medium.png",
          "large": "https://test.taskulu.com/static/users-0/large.png",
          "big": "https://test.taskulu.com/static/users-0/big.png"
        }
      }
    ],
    "sheets": [
	{
		"id": "5a8d1fff56ad660b0dd0d345",
		"type": "Sheet",
		"title": "فروغ",
		"permission": "edit",
		"default": true,
		"task_lists": [
			{
				"id": "5d46aff456ad6671e8008b99",
				"type": "TaskList",
				"title": "فروغ",
				"task_order": [
					"5d46b04456ad667202008c23"
				],
				"archived": false,
				"permission": "edit",
				"sections": [
					{
						"type": "Section",
						"name": "Todo",
						"count": 0,
						"order": 0,
						"tasks": []
					},
					{
						"type": "Section",
						"name": "Doing",
						"count": 0,
						"order": 1,
						"tasks": []
					},
					{
						"type": "Section",
						"name": "Done",
						"count": 1,
						"order": 2,
						"tasks": [
							{
								"id": "5d46b04456ad667202008c23",
								"created_at": 1564913732,
								"title": "راه اندازی gitlab واسه CICD",
								"description": "",
								"type": "Task",
								"status": 2,
								"assigned_to": [
									"5a8d1fde56ad660aabd0d2ec"
								],
								"permission": "edit",
								"color_label": 0,
								"start_time": 0,
								"deadline": 1565094600,
								"tags": [],
								"has_check_list": false,
								"comment_count": 0,
								"attachment_count": 0,
								"timelogging_count": 0,
								"check_list_count": 0,
								"weight_value": 0
							}
						]
					}
				]
			}
		]
	},
	{
	"id": "5b029ec356ad663c26000962",
	"type": "Sheet",
	"title": "aperiam dolorem sint",
	"permission": "edit",
	"default": false,
	"task_lists": []
	}
    ]
  }
}
`

)