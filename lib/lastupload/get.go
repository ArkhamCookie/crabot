package lastupload

func GetTimeOfLastUpload(username, YOUTUBE_KEY string) (*LastUploadResults, error) {
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
