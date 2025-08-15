package controller

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

func (trk *TrakerCron) UsersRatingsCSV(ctx context.Context) {
	// Schedule the cron job every 22 hours
	trk.Cron.AddFunc("@every 22h", func() {
		fmt.Println("UsersRatingsCSV started...")

		// Ensure the directory exists
		dirPath := "tempf"
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			trk.respondWithError("Failed to create directory for UsersRating.csv", err)
			return
		}
		
		// Create the Movies.csv file
		filePath := dirPath + "/UsersRating.csv"
		file, err := os.Create(filePath)
		if err != nil {
			trk.respondWithError("UsersRating.csv not created", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := trk.Application.Storge.Traker.CSVTabls(ctx, writer, "ratings"); err != nil {
			trk.respondWithError("Error while writing to UsersRatingsCSV.csv", err)
		}
	})
}
