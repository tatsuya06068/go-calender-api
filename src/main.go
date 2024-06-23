package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// サービスアカウントの認証情報を読み込む
	b, err := os.ReadFile("credential.json")
	if err != nil {
		log.Fatalf("Unable to read service account file: %v", err)
	}

	// Google Calendar APIクライアントを作成
	config, err := google.JWTConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse service account file to config: %v", err)
	}

	client := config.Client(ctx)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Calendar client: %v", err)
	}

	calendarId := "" // ユーザーのカレンダーID

	// 新しいイベントを作成
	event := &calendar.Event{
		Summary:     "New Event Title",
		Location:    "800 Howard St., San Francisco, CA 94103",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2024-07-01T09:00:00-07:00",
		},
		End: &calendar.EventDateTime{
			DateTime: "2024-07-01T17:00:00-07:00",
		},
	}

	// INSERT
	newEvent, err := srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
	}

	fmt.Printf("Event created: %s\n", newEvent.HtmlLink)

	// UPDATE
	// eventId := "your-event-id"        // 変更するイベントID
	// // イベントを取得
	// eve, err := srv.Events.Get(calendarId, eventId).Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve event: %v", err)
	// }

	// // イベントの内容を更新
	// event.Summary = "Updated Event Title"
	// event.Description = "Updated Event Description"

	// // イベントを更新
	// updatedEvent, err := srv.Events.Update(calendarId, eventId, event).Do()
	// if err != nil {
	// 	log.Fatalf("Unable to update event: %v", err)
	// }

	// READ
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calendarId).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}

}
