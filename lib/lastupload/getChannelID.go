package lastupload

import (
	"encoding/json"
	"errors"
	"fmt"
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
	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("issue fetching data, status: %q\n", response.Status)
		return "", errors.New(errMsg)
	}

	var data *ChannelSearch
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", err
	}

	return data.Items[0].ID.ChannelID, nil
}
