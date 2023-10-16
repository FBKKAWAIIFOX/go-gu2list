package commands

import "github.com/diamondburned/arikawa/v3/api"

var (
	Commands = []api.CreateCommandData{
		{
			Name:        "vprice",
			Description: "查詢廢土村民錠價格",
		},
		{
			Name:        "test",
			Description: "測試指令",
		},
	}
)
