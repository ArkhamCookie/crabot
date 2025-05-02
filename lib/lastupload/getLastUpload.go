package lastupload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetLastUpload(upload_playlist_id, YOUTUBE_KEY string) (*LastUploadResults, error) {
	response, err := http.Get("https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&playlistId=" + upload_playlist_id + "&maxResults=1&key=" + YOUTUBE_KEY)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("issue fetching data, status: %q\n", response.Status)
		return nil, errors.New(errMsg)
	}

	var data *LastUploadSearch
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}

	return data.Items, nil
}
