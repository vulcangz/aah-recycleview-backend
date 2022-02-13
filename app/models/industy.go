// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

// Based on go-aah examples:
package models

import (
	"errors"
	"sync"
	"time"
)

var (
	industryStore  = make(map[int64]*Industry)
	lastIndustryID int64
	idMx           = &sync.Mutex{}
	industryMx     = &sync.RWMutex{}
)

type Industry struct {
	IndustryID   int64     `json:"IndustryID,omitempty"`
	IndustryName string    `json:"IndustryName,omitempty"`
	Nature       string    `json:"Nature,omitempty"`
	Count        int64     `json:"Count,omitempty"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// IndustryAlias is used for time format to ISO 8061
type IndustryAlias Industry

// AllIndustries returns all the industries.
func AllIndustries() []interface{} {
	industryMx.RLock()
	defer industryMx.RUnlock()

	industries := make([]interface{}, 0)
	for _, industry := range industryStore {
		industries = append(industries, ToIndustryAlias(industry))
	}
	return industries
}

// CreateIndustry method creates the new industry entry.
func CreateIndustry(industry *Industry) (*Industry, int64) {
	industryMx.Lock()
	defer industryMx.Unlock()
	id := createIndustryID()
	industry.IndustryID = id
	industry.Count = 0
	industry.CreatedAt = time.Now()
	industry.UpdatedAt = industry.CreatedAt
	industryStore[id] = industry
	return industryStore[id], id
}

// GetIndustry method return the industry for given ID.
func GetIndustry(id int64) (*Industry, error) {
	industryMx.RLock()
	defer industryMx.RUnlock()
	if _, found := industryStore[id]; !found {
		return nil, errors.New("industry not found")
	}

	industryStore[id].Count++

	return industryStore[id], nil
}

// UpdateIndustry method updates the given info with industry store.
func UpdateIndustry(industry *Industry) (*Industry, error) {
	industryMx.Lock()
	defer industryMx.Unlock()
	if _, found := industryStore[industry.IndustryID]; !found {
		return nil, errors.New("industry not found")
	}

	industryStore[industry.IndustryID].IndustryName = industry.IndustryName
	industryStore[industry.IndustryID].Nature = industry.Nature
	industryStore[industry.IndustryID].Count = industry.Count
	industryStore[industry.IndustryID].UpdatedAt = time.Now()
	return industryStore[industry.IndustryID], nil
}

// DeleteIndustry method deletes the industry for given ID.
func DeleteIndustry(id int64) error {
	industryMx.Lock()
	defer industryMx.Unlock()
	if _, found := industryStore[id]; !found {
		return errors.New("industry not found")
	}

	delete(industryStore, id)
	return nil
}

// ToIndustryAlias method formats the time to RFC3339 using alias.
func ToIndustryAlias(industry *Industry) interface{} {
	return &struct {
		*IndustryAlias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		IndustryAlias: (*IndustryAlias)(industry),
		CreatedAt:     industry.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     industry.UpdatedAt.Format(time.RFC3339),
	}
}

func createIndustryID() int64 {
	idMx.Lock()
	defer idMx.Unlock()
	lastIndustryID++
	return lastIndustryID
}

func init() {
	industryStore[1] = &Industry{IndustryID: 1, IndustryName: "Electronics", Nature: "1", Count: 4}
	industryStore[2] = &Industry{IndustryID: 2, IndustryName: "Mobile & Tablets", Nature: "1", Count: 3}
	industryStore[3] = &Industry{IndustryID: 3, IndustryName: "Computers & Acc.", Nature: "2", Count: 3}
	industryStore[4] = &Industry{IndustryID: 4, IndustryName: "Electronic Material", Nature: "1", Count: 3}
	industryStore[5] = &Industry{IndustryID: 5, IndustryName: "Food", Nature: "1", Count: 3}
	industryStore[6] = &Industry{IndustryID: 6, IndustryName: "Dairy", Nature: "1", Count: 3}
	industryStore[7] = &Industry{IndustryID: 7, IndustryName: "Fashion", Nature: "3", Count: 3}
	industryStore[8] = &Industry{IndustryID: 8, IndustryName: "Stationery", Nature: "4", Count: 3}
	industryStore[9] = &Industry{IndustryID: 9, IndustryName: "Sports & Fitness", Nature: "1", Count: 3}
	industryStore[10] = &Industry{IndustryID: 10, IndustryName: "Specs & Opticals", Nature: "1", Count: 3}
	industryStore[11] = &Industry{IndustryID: 11, IndustryName: "Sanitary", Nature: "2", Count: 3}
	industryStore[12] = &Industry{IndustryID: 12, IndustryName: "Plastic Ware", Nature: "1", Count: 3}
	industryStore[13] = &Industry{IndustryID: 13, IndustryName: "Pharma", Nature: "2", Count: 3}
	industryStore[14] = &Industry{IndustryID: 14, IndustryName: "Paint", Nature: "1", Count: 3}
	industryStore[15] = &Industry{IndustryID: 15, IndustryName: "Audio & Video Store & Gaming", Nature: "1", Count: 3}
	industryStore[16] = &Industry{IndustryID: 16, IndustryName: "Automobiles", Nature: "1", Count: 3}
	industryStore[17] = &Industry{IndustryID: 17, IndustryName: "Beauty & Grooming", Nature: "1", Count: 3}
	industryStore[18] = &Industry{IndustryID: 18, IndustryName: "Books", Nature: "1", Count: 3}
	industryStore[19] = &Industry{IndustryID: 19, IndustryName: "Building Material", Nature: "1", Count: 3}
	industryStore[20] = &Industry{IndustryID: 20, IndustryName: "Cultery Crockery", Nature: "1", Count: 3}
	lastIndustryID = 20
}
