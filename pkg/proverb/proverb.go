package proverb

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/labstack/gommon/log"
)

/*
{
            "quote": "Success is leaving a good path, or even better, leaving no path at all. Maxime Lagacé",
            "author": null,
            "likes": 0,
            "tags": [
                "success"
            ],
            "pk": 465548,
            "image": null,
            "language": "en"
        }
*/

const data = `[
       {
           "quote": "Success is leaving a good path, or even better, leaving no path at all. Maxime Lagacé",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465548,
           "image": null,
           "language": "en"
       },
       {
           "quote": "My success is not who I am. Judith Guest",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465549,
           "image": null,
           "language": "en"
       },
       {
           "quote": "Success is how high you bounce when you hit bottom. George S. Patton",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465550,
           "image": null,
           "language": "en"
       },
       {
           "quote": "You have reached the pinnacle of success as soon as you become uninterested in money, compliments, or publicity. Thomas Wolfe",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465551,
           "image": null,
           "language": "en"
       },
       {
           "quote": "To be able to look back upon one’s life in satisfaction, is to live twice. Kahlil Gibran",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465552,
           "image": null,
           "language": "en"
       },
       {
           "quote": "I don’t dwell on success. Maybe that’s one reason I’m successful. Calvin Klein",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465553,
           "image": null,
           "language": "en"
       },
       {
           "quote": "I do not pray for success, I ask for faithfulness. Mother Teresa",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465554,
           "image": null,
           "language": "en"
       },
       {
           "quote": "A minute’s success pays the failure of years. Robert Browning",
           "author": null,
           "likes": 0,
           "tags": [
               "success"
           ],
           "pk": 465555,
           "image": null,
           "language": "en"
       }
]`

type Proverb struct {
	Quote  string `json:"quote,omitempty"`
	Author string `json:"author,omitempty"`
	//Likes  int    `json:"likes,omitempty"`
	//Tags     []string
	//Pk       int
	//Language string
}

func New() *Proverb {
	return &Proverb{}
}

//// Put in entity
//type ProverbData struct {
//}

// Get proverbs from db
func (p *Proverb) GetProverbs() ([]Proverb, error) {

	proverbs := make([]Proverb, 0, 8)

	err := json.Unmarshal([]byte(data), &proverbs)
	if err != nil {
		log.Errorf("Unable to parse data: %v", err.Error())
		return nil, err
	}

	for i, v := range proverbs {
		fmt.Printf("(%d)Proverb: %s\t, Author: %s\n", i, v.Quote, v.Author)
	}
	return proverbs, nil // TODO []Proverb is already a pointer - check pointer return
}

func (p *Proverb) LoadProverb() error {
	// Get data from api, then store in db

	proverbs := make([]Proverb, 0, 8)

	err := json.Unmarshal([]byte(data), &proverbs)
	if err != nil {
		log.Errorf("Unable to parse data: %v", err.Error())
		return err
	}

	// map to storage type
	data := make([]Data, 0, len(proverbs))
	for _, pvb := range proverbs {
		d := Data{Proverb: pvb.Quote, Author: &pvb.Author}
		data = append(data, d)
	}

	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/silvade")
	panicOnError(err)
	defer db.Close()

	// insert into store
	store := NewStorage(db)

	for _, d := range data {
		err := store.Create(context.TODO(), &d)
		if err != nil {
			log.Errorf("Unable to parse data: %v", err.Error())
		}
	}

	// return count
	return nil
}
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
