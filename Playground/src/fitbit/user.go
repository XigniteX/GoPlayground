package fitbit

type UserContainer struct {
	User struct {
		Leeftijd               float64 `json:"age"`
		Avatar                 string  `json:"avatar"`
		Avatar150              string  `json:"avatar150"`
		AverageDailySteps      float64 `json:"averageDailySteps"`
		ClockTimeDisplayFormat string  `json:"clockTimeDisplayFormat"`
		Corporate              bool    `json:"corporate"`
		CorporateAdmin         bool    `json:"corporateAdmin"`
		DateOfBirth            string  `json:"dateOfBirth"`
		DisplayName            string  `json:"displayName"`
		DisplayNameSetting     string  `json:"displayNameSetting"`
		DistanceUnit           string  `json:"distanceUnit"`
		EncodedId              string  `json:"encodedId"`
		Features               struct {
			ExerciseGoal bool `json:"exerciseGoal"`
		} `json:"features"`
		FirstName               string  `json:"firstName"`
		FullName                string  `json:"fullName"`
		Gender                  string  `json:"gender"`
		GlucoseUnit             string  `json:"glucoseUnit"`
		Height                  float64 `json:"height"`
		HeightUnit              string  `json:"heightUnit"`
		LastName                string  `json:"lastName"`
		Locale                  string  `json:"locale"`
		MemberSince             string  `json:"memberSince"`
		MfaEnabled              bool    `json:"mfaEnabled"`
		OffsetFromUTCMillis     float64 `json:"offsetFromUTCMillis"`
		StartDayOfWeek          string  `json:"startDayOfWeek"`
		StrideLengthRunning     float64 `json:"strideLengthRunning"`
		StrideLengthRunningType string  `json:"strideLengthRunningType"`
		StrideLengthWalking     float64 `json:"strideLengthWalking"`
		StrideLengthWalkingType string  `json:"strideLengthWalkingType"`
		SwimUnit                string  `json:"swimUnit"`
		Timezone                string  `json:"timezone"`
		TopBadges               []struct {
			BadgeGradientEndColor   string        `json:"badgeGradientEndColor"`
			BadgeGradientStartColor string        `json:"badgeGradientStartColor"`
			BadgeType               string        `json:"badgeType"`
			Category                string        `json:"category"`
			Cheers                  []interface{} `json:"cheers"`
			DateTime                string        `json:"dateTime"`
			Description             string        `json:"description"`
			EarnedMessage           string        `json:"earnedMessage"`
			EncodedId               string        `json:"encodedId"`
			Image100px              string        `json:"image100px"`
			Image125px              string        `json:"image125px"`
			Image300px              string        `json:"image300px"`
			Image50px               string        `json:"image50px"`
			Image75px               string        `json:"image75px"`
			MarketingDescription    string        `json:"marketingDescription"`
			MobileDescription       string        `json:"mobileDescription"`
			Name                    string        `json:"name"`
			ShareImage640px         string        `json:"shareImage640px"`
			ShareText               string        `json:"shareText"`
			ShortDescription        string        `json:"shortDescription"`
			ShortName               string        `json:"shortName"`
			TimesAchieved           float64       `json:"timesAchieved"`
			Value                   float64       `json:"value"`
		} `json:"topBadges"`
		Weight     float64 `json:"weight"`
		WeightUnit string  `json:"weightUnit"`
	} `json:"user"`
}
