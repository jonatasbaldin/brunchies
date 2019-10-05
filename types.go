package main

type Track struct {
	Name    string
	Artists string
}

type Week struct {
	Year   int
	Week   int
	Tracks []Track
}

func (w Week) Equal(w2 Week) bool {
	if w.Year == w2.Year {
		if w.Week == w2.Week {
			return true
		}
		return false
	}
	return false
}

func getSpecificWeek(toSearch []Week, item Week) int {
	for i, w := range toSearch {
		if w.Equal(item) {
			return i
		}
	}

	return 0
}

type PlaylistMetadata struct {
	Author    string
	Followers uint
	URL       string
	ImageURL  string
}

type Data struct {
	PlaylistMetadata
	Weeks []Week
}
