package storage

import (
	"fmt"
	"strings"
	"time"

	r "github.com/dancannon/gorethink"
	m "github.com/mitchellh/mapstructure"
)

type Activity struct {
	Message
	URL string
}

func (s *Storage) SaveActivity(activity *Activity) {

	fmt.Printf("\n\nLooks like you're saving a activity: %#v\n\n", activity)

	if activity.Timestamp == 0 {
		activity.Timestamp = int(time.Now().Unix())
	}
	_, err := r.DB(s.Name).Table("activities").Insert(activity).RunWrite(s.Session)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func (s *Storage) FindAllActivities() ([]Activity, error) {
	rows, err := r.DB(s.Name).Table("activities").Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Activity
	err = rows.All(&list)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return list, nil
}

func (s *Storage) FindOneActivity(id string) (*Activity, error) {

	rows, err := r.DB(s.Name).Table("activities").Filter(
		r.Row.Field("id").Eq(id)).Run(s.Session)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activity Activity
	err = rows.One(&activity)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return &activity, nil
}

func (s *Storage) DeleteActivity(id string) (*Activity, error) {

	rows, err := r.DB(s.Name).Table("activities").Get(id).Delete(r.DeleteOpts{ReturnChanges: true}).Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var value r.WriteResponse
	rows.One(&value)

	var oldValueMap, ok = value.Changes[0].OldValue.(map[string]interface{})
	if !ok {
		fmt.Println("Type assertion failed :(")
	}
	fmt.Println("OldvalueMap: ", oldValueMap)

	var oldValueActivity Activity
	err = m.Decode(oldValueMap, &oldValueActivity)
	if err != nil {
		fmt.Println("err decoding: ", err)
	}

	fmt.Println("Oldvalueactivity: ", oldValueActivity)
	return &oldValueActivity, nil
}

func (s *Storage) SearchActivities(searchStrings []string) ([]Activity, error) {

	fmt.Printf("Searchterms: %s\n", searchStrings)

	//Append the strings into one regex string, e.g. bob|said|bananas
	searchTerms := strings.Join(searchStrings, "|")
	fmt.Printf("Filtered searchterms: %s\n", searchTerms)

	rows, err := r.Table("activities").Filter(func(activity r.Term) r.Term {
		return activity.Field("text").Match(searchTerms)
	}).Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []Activity{}
	err2 := rows.All(&list)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("Search result record %#v\n", list)

	return list, nil
}
