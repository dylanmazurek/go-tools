package date

import (
	"encoding/json"
	"strconv"
	"time"
)

type Date struct {
	Date time.Time
}

func Parse(date string) (Date, error) {
	t, err := time.Parse(time.DateOnly, date)

	newDate := Date{
		Date: t,
	}

	return newDate, err
}

func ParseDate(date time.Time) (Date, error) {
	stripTime := date.Format(time.DateOnly)
	d, err := time.Parse(time.DateOnly, stripTime)

	newDate := Date{
		Date: d,
	}

	return newDate, err
}

func (t *Date) String() string {
	formattedDate := t.Date.Format(time.DateOnly)
	return formattedDate
}

func (t *Date) MarshalJSON() ([]byte, error) {
	type Alias Date
	marshaledJSON, err := json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	})

	return marshaledJSON, err
}

func (t *Date) UnmarshalJSON(data []byte) error {
	dateString := string(data)
	if dateString == "null" {
		return nil
	}

	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	t.Date, err = time.Parse(time.DateOnly, s)

	return err
}
