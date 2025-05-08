package lastupload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Get a channel upload playlist from a channel ID.
func GetUploadsPlaylistID(channelID, YOUTUBE_KEY string) (string, error) {
	response, err := http.Get("https://www.googleapis.com/youtube/v3/channels?part=contentDetails&id=" + channelID + "&key=" + YOUTUBE_KEY)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("issue fetching data, status: %q\n", response.Status)
		return "", errors.New(errMsg)
	}

	var data *uploadsPlaylistSearch
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", err
	}

	return data.Items[0].ContentDetails.RelatedPlaylists.Uploads, nil
}
