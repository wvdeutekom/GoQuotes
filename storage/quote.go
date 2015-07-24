package storage

import (
	"fmt"
	"strings"
	"time"

	r "github.com/dancannon/gorethink"
	m "github.com/mitchellh/mapstructure"
)

// Quote struct for slash command
type Quote Message

type Storage struct {
	Name    string
	URL     string
	Session *r.Session
}

func (s *Storage) SaveQuote(quote *Quote) {

	fmt.Printf("\n\nLooks like you're saving a quote: %#v\n\n", quote)

	if quote.Timestamp == 0 {
		quote.Timestamp = int(time.Now().Unix())
	}
	_, err := r.DB(s.Name).Table("quotes").Insert(quote).RunWrite(s.Session)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func (s *Storage) FindAllQuotes() ([]Quote, error) {
	rows, err := r.DB(s.Name).Table("quotes").Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Quote
	err = rows.All(&list)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return list, nil
}

func (s *Storage) FindOneQuote(id string) (*Quote, error) {

	rows, err := r.DB(s.Name).Table("quotes").Filter(
		r.Row.Field("id").Eq(id)).Run(s.Session)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quote Quote
	err = rows.One(&quote)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return &quote, nil
}

func (s *Storage) GetLatestQuote() Quote {

	rows, err := r.Table("quotes").OrderBy(r.Desc("timestamp")).Run(s.Session)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var quote Quote
	err2 := rows.One(&quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("Fetch one record %#v\n", quote)

	return quote
}

func (s *Storage) DeleteQuote(id string) (*Quote, error) {

	rows, err := r.DB(s.Name).Table("quotes").Get(id).Delete(r.DeleteOpts{ReturnChanges: true}).Run(s.Session)
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

	var oldValueQuote Quote
	err = m.Decode(oldValueMap, &oldValueQuote)
	if err != nil {
		fmt.Println("err decoding: ", err)
	}

	fmt.Println("OldvalueQuote: ", oldValueQuote)
	return &oldValueQuote, nil
}

func (s *Storage) SearchQuotes(searchStrings []string) ([]Quote, error) {

	fmt.Printf("Searchterms: %s\n", searchStrings)

	//Append the strings into one regex string, e.g. bob|said|bananas
	searchTerms := strings.Join(searchStrings, "|")
	fmt.Printf("Filtered searchterms: %s\n", searchTerms)

	rows, err := r.Table("quotes").Filter(func(quote r.Term) r.Term {
		return quote.Field("text").Match(searchTerms)
	}).Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []Quote{}
	err2 := rows.All(&list)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("Search result record %#v\n", list)

	return list, nil
}
