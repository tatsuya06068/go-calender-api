// 新しいイベントを作成
	newEvent := &calendar.Event{
		Summary:     "新しいイベントのタイトル",
		Location:    "イベントの場所",
		Description: "イベントの詳細",
		Start: &calendar.EventDateTime{
			DateTime: "2024-08-10T10:00:00+09:00", // 開始日時
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: "2024-08-10T12:00:00+09:00", // 終了日時
			TimeZone: "Asia/Tokyo",
		},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: "sample-12345", // 一意のリクエストID (任意のユニークな文字列)
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet", // Google Meetを指定
				},
			},
		},
	}

	// イベントの作成
	event, err := svc.Events.Insert("primary", newEvent).ConferenceDataVersion(1).Do()
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}

	// Meet URLを取得して表示
	if event.ConferenceData != nil && len(event.ConferenceData.EntryPoints) > 0 {
		for _, entryPoint := range event.ConferenceData.EntryPoints {
			if entryPoint.EntryPointType == "video" {
				fmt.Printf("Google Meet URL: %s\n", entryPoint.Uri)
			}
		}
	} else {
		fmt.Println("Meet URL is not available.")
	}
