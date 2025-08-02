package controller

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

func (trk *TrakerCron) MoviesCSV(ctx context.Context) {
	// Schedule the cron job every 30 seconds
	trk.Cron.AddFunc("@every 10s", func() {
		fmt.Println("MoviesCSV started...")

		// Ensure the directory exists
		dirPath := "tempf"
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			trk.respondWithError("Failed to create directory for Movies.csv", err)
			return
		}

		// Create the Movies.csv file
		filePath := dirPath + "/Movies.csv"
		file, err := os.Create(filePath)
		if err != nil {
			trk.respondWithError("Movies.csv not created", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := trk.Application.Storge.Traker.MoviesCSV(ctx, writer); err != nil {
			trk.respondWithError("Error while writing to Movies.csv", err)
		}
		fmt.Println("MoviesCSV done...")
	})
}

