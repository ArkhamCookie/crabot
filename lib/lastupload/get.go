package lastupload

import (
	"regexp"
	"time"
)

// Get the data related to latest upload based on YouTube username.
func GetData(username, YOUTUBE_KEY string) ([]LastUploadResults, error) {
	channelID, err := GetChannelID(username, YOUTUBE_KEY)
	if err != nil {
		return nil, err
	}

	uploadPlaylistID, err := GetUploadsPlaylistID(channelID, YOUTUBE_KEY)
	if err != nil {
		return nil, err
	}

	lastUploadData, err := GetLastUploadData(uploadPlaylistID, YOUTUBE_KEY)
	if err != nil {
		return nil, err
	}

	return lastUploadData, nil
}

// Get the date of the lastest upload based on LastUploadResults.
//
// LastUploadResults can be gotten by running the `lastupload.Get` functiion.
func GetDate(lastUploadResults []LastUploadResults) (time.Time, error) {
	// Get last upload date from lastUploadResults
	lastUpload := lastUploadResults[0].Snippet.PublishedAt.String()

	regex, err := regexp.Compile("T")
	if err != nil {
		return time.Now(), err
	}

	// Replace the **SECOND** 'T' in lastUpload
	count := 0
	reformatted := regex.ReplaceAllStringFunc(lastUpload, func(s string) string {
		count++
		if count == 2 {
			return " "
		}
		return s
	})

	regex, err = regexp.Compile("Z")
	if err != nil {
		return time.Now(), err
	}

	// Replace all 'Z's in reformatted
	reformatted = regex.ReplaceAllString(reformatted, " ")

	regex, err = regexp.Compile(" .0000 UTC")
	if err != nil {
		return time.Now(), err
	}

	reformatted = regex.ReplaceAllString(reformatted, "")

	layout := "2006-01-02 15:04:05"
	uploadTime, err := time.Parse(layout, reformatted)
	if err != nil {
		return time.Now(), err
	}

	return uploadTime, nil
}
