package responses

import "EasyTutor/data/data"

type Offer struct {
	ID 			string	`json:"id"`
	data.Offer
}
