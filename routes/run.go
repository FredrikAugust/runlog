package routes

import (
	"log/slog"
	"net/http"

	"github.com/fredrikaugust/runlog/storage"

	"github.com/tkrajina/gpxgo/gpx"
)

func Upload(db *storage.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("gpx")
		if err != nil {
			http.Error(w, "invalid gpx file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		slog.Info("received file", "filename", fileHeader.Filename)

		gpxFile, err := gpx.Parse(file)
		if err != nil {
			slog.Error("failed to read content of gpx file", "error", err.Error())
			http.Error(w, "could not read file", http.StatusInternalServerError)
			return
		}

		slog.Info("parsed gpx file", "filename", fileHeader.Filename)

		for _, track := range gpxFile.Tracks {
			for _, segment := range track.Segments {
				for _, point := range segment.Points {
					slog.Info("point", "lat", point.Latitude, "lon", point.Longitude, "segment", segment.GetTrackPointsNo())
				}
			}
		}
	}
}
