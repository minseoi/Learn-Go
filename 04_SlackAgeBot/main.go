package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// 채널. 채널은 고루틴 간의 통신을 위한 파이프라인이다.
// chan<- 은 송신전용 채널
// <-chan 은 수신전용 채널
func printCommandEvnets(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("------------------")
		fmt.Println("Command Evnets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("------------------")
	}
}

func main() {
	//봇 만들기
	os.Setenv("SLACK_BOT_TOKEN", "") // Need to fill
	os.Setenv("SLACK_APP_TOKEN", "") // Need to fill
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//디버깅용
	go printCommandEvnets(bot.CommandEvents())

	//봇에게 명령어를 등록한다.
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{
			"my yob is 2000",
		},
		//실제 명령어가 실행될 때 호출되는 함수
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	//context는 고루틴에서 공유할 수 있는 값을 저장하는 작업 명세서같은 것
	//context.Background()는 빈 context를 반환한다.
	//context.WithCancel()은 context와 cancel 함수를 반환한다.
	ctx, cancel := context.WithCancel(context.Background())

	//main 함수가 종료되면 cancel 함수를 호출한다.
	defer cancel()

	//봇을 실행한다.
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
