package lastupload

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetChannelID(username, YOUTUBE_KEY string) (string, error) {
	response, err := http.Get("https://www.googleapis.com/youtube/v3/search?part=snippet&q=" + username + "&type=channel&key=" + YOUTUBE_KEY)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data *ChannelSearch
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", err
	}

	return data.Items[0].ID.ChannelID, nil
}
