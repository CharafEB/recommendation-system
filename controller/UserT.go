package controller

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

func (trk *TrakerCron) UsersCSV(ctx context.Context) {
	// Schedule the cron job every 30 seconds
	trk.Cron.AddFunc("@every 10s", func() {
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

		if err := trk.Application.Storge.Traker.CSVTabls(ctx, writer, "client"); err != nil {
			trk.respondWithError("Error while writing to Movies.csv", err)
		}
		fmt.Println("UsersRatingsCSV done...")
	})
}
