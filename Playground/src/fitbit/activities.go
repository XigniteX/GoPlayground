package fitbit

type ActivityContainer struct {
	Activities []struct {
		ActiveDuration float64 `json:"activeDuration"`
		ActivityLevel  []struct {
			Minutes float64 `json:"minutes"`
			Name    string  `json:"name"`
		} `json:"activityLevel"`
		ActivityName          string  `json:"activityName"`
		ActivityTypeId        float64 `json:"activityTypeId"`
		Calories              float64 `json:"calories"`
		Distance              float64 `json:"distance"`
		DistanceUnit          string  `json:"distanceUnit"`
		Duration              float64 `json:"duration"`
		LastModified          string  `json:"lastModified"`
		LogId                 float64 `json:"logId"`
		LogType               string  `json:"logType"`
		ManualValuesSpecified struct {
			Calories bool `json:"calories"`
			Distance bool `json:"distance"`
			Steps    bool `json:"steps"`
		} `json:"manualValuesSpecified"`
		OriginalDuration  float64 `json:"originalDuration"`
		OriginalStartTime string  `json:"originalStartTime"`
		Speed             float64 `json:"speed"`
		StartTime         string  `json:"startTime"`
		Steps             float64 `json:"steps"`
	} `json:"activities"`
	Pagination struct {
		BeforeDate string  `json:"beforeDate"`
		Limit      float64 `json:"limit"`
		Next       string  `json:"next"`
		Offset     float64 `json:"offset"`
		Previous   string  `json:"previous"`
		Sort       string  `json:"sort"`
	} `json:"pagination"`
}

