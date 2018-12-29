package actress

//ActressSearch の結果を代入
type ActressSearch struct {
	Request struct {
		Parameters struct {
			AffiliateID string `json:"affiliate_id"`
			APIID       string `json:"api_id"`
			Keyword     string `json:"keyword"`
			Output      string `json:"output"`
		} `json:"parameters"`
	} `json:"request"`
	Result struct {
		Status        string `json:"status"`
		ResultCount   int    `json:"result_count"`
		TotalCount    string `json:"total_count"`
		FirstPosition int    `json:"first_position"`
		Actress       []struct {
			ID          string      `json:"id"`
			Name        string      `json:"name"`
			Ruby        string      `json:"ruby"`
			Bust        interface{} `json:"bust"`
			Waist       interface{} `json:"waist"`
			Hip         interface{} `json:"hip"`
			Height      interface{} `json:"height"`
			Birthday    interface{} `json:"birthday"`
			BloodType   interface{} `json:"blood_type"`
			Hobby       interface{} `json:"hobby"`
			Prefectures interface{} `json:"prefectures"`
			ListURL     struct {
				Digital string `json:"digital"`
				Monthly string `json:"monthly"`
				Ppm     string `json:"ppm"`
				Mono    string `json:"mono"`
				Rental  string `json:"rental"`
			} `json:"listURL"`
			ImageURL struct {
				Small string `json:"small"`
				Large string `json:"large"`
			} `json:"imageURL,omitempty"`
			Cup string `json:"cup,omitempty"`
		} `json:"actress"`
	} `json:"result"`
}