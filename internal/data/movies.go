package data

import "time"

type Movie struct {
	ID        int64     // Unique integer ID for the movie
	CreatedAt time.Time // Timestamp for when the movie is added to our database
	Title     string    // Movie title
	Year      int32     // Movie release year
	Runtime   int32     // Movie runtime (in minutes)
	Genres    []string  // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     // The version number starts at 1 and will be incremented each // time the movie information is updated
}

// type Movie struct {
// 	ID        int64     `json:"id"`
// 	CreatedAt time.Time `json:"created_at"`
// 	Title     string    `json:"title"`
// 	Year      int32     `json:"year"`
// 	Runtime   int32     `json:"runtime"`
// 	Genres    []string  `json:"genres"`
// 	Version   int32     `json:"version"`
// }
