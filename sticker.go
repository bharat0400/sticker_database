package sticker_database

import (
	"database/sql"
	"log"
	"time"
)

func getTrendingStickers() []map[string]string {
	db, err := sql.Open("mysql", "root:0400@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()
	var stickers []map[string]string
	results, err := db.Query("select * from stickers where end_hour > ? and start_hour < ? order by priority desc", time.Now().Hour(), time.Now().Hour())
	if err != nil {
		log.Fatal("Error when fetching stickers table rows:", err)
	}
	defer results.Close()
	err = results.Scan(&stickers)
	if err != nil {
		log.Fatal("Unable to parse row:", err)
	}
	// for results.Next() {
	// 	var sticker sticker_model.Sticker
	// 	err = results.Scan(&sticker)
	// 	if err != nil {
	// 		log.Fatal("Unable to parse row:", err)
	// 	}
	// 	stickers = append(stickers, sticker)
	// }
	return stickers
}
