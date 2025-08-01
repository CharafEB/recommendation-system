package controller

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

func (trk *TrakerCron) MoviesCSV(ctx context.Context) {
	//Create the temp Movies.csv file
	trk.Cron.AddFunc("@every 24h", func() {
		fmt.Print("MoviesCSV started...")
		file, err := os.Create("/temp/Movies.csv")
		if err != nil {
			trk.respondWithError("Movies.csv not created", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := trk.Application.Storge.Traker.MoviesCSV(ctx, writer); err != nil {
			trk.respondWithError("Error while writing on Movies.csv", err)
		}
	})

}
