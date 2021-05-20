package getui

import (
	"fmt"
	"testing"
)

var account *Account

func init() {
	account = NewGetUiAccount(
		"9QOz0GyEBv7xGa8MeULlt1",
		"D2uZ6uxnNm96wgJDvXHvN5",
		"vV771kd1S97DcfoY1agq91",
		"1BtSNqcitK7g2aSjuqInE9")
	res := account.AuthSign()
	fmt.Println(account.Auth, res)
}

func TestAccount(t *testing.T) {

	var tt = []PushSingleParams{
		PushSingleParams{
			Title:               "你好",
			Text:                "今天是2019.03.23",
			TransmissionContent: "好日子",
			Cid:                 []string{"0615a6190fba0484db5e83079ca4c17b"},
			AppKey:              account.AppKey,
		},
	}
	for _, t := range tt {
		fmt.Println(account.SaveListBodyAndPushList(t))

	}
	//RASL_0511_2fa1c9c874a64e379c09737eac152a88
}

func TestPushResultQueryByTasks(tests *testing.T) {
	fmt.Println(account.PushResultQueryByTask([]string{"0615a6190fba0484db5e83079ca4c17b"}))
}
