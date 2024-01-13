package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "strconv"

    "github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
    for event := range analyticsChannel {
        fmt.Println("Command Events")
        fmt.Println(event.Timestamp)
        fmt.Println(event.Command)
        fmt.Println(event.Parameters)
        fmt.Println(event.Event)
        fmt.Println()
    }
}

func main() {
    os.Setenv("SLACK_BOT_TOKEN", "xoxb-5665290907731-5665464699283-TWTkT1ESy9MOpdxyXi9MBBsa")
    os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05KGFASXMK-5662527141989-85f9b7a799b7ed06ff54fa2154d9a32c579eae2ee0fdf318b7d5c80e4fca2f0d")

    // creating the bot
    age_bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

    // when pass command to slack bot, the way it handles the event, it is going to print it out
    go printCommandEvents(age_bot.CommandEvents())

    age_bot.Command("my yob is <year>", &slacker.CommandDefinition{
        Description: "yob calculator",
        //Example:     "my yob is 2000",
        Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
            year := request.Param("year")
            // to convert to numerical form
            yob, err := strconv.Atoi(year)
            if err != nil {
                println("error")
            }
            age := 2023 - yob
            current_age := fmt.Sprintf("age is %d", age)
            response.Reply(current_age)
        },
        // event is handled using a function with takes the bot context
    })

    //
    ctx, cancel := context.WithCancel(context.Background())
    // makes sure function is called toward the end
    defer cancel()

    err := age_bot.Listen(ctx)
    if err != nil {
        // logs out of error
        log.Fatal(err)
    }

}
