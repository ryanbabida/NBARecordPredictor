package datastore

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Record is a per-game statistic
type Record struct {
	GamesPlayed          int     `json:"gamesPlayed"`
	Wins                 int     `json:"wins"`
	Losses               int     `json:"losses"`
	WinPercentage        float64 `json:"winPercentage"`
	Minutes              float64 `json:"minutes"`
	Points               float64 `json:"points"`
	FieldGoalsMade       float64 `json:"fieldGoalsMade"`
	FieldGoalsAttempted  float64 `json:"fieldGoalAttempted"`
	FieldGoalPercentage  float64 `json:"fieldGoalPercentage"`
	ThreesMade           float64 `json:"threesMade"`
	ThreesAttempted      float64 `json:"threesAttempted"`
	ThreePercentage      float64 `json:"threePercentage"`
	FreeThrowsMade       float64 `json:"freeThrowsMade"`
	FreeThrowsAttempted  float64 `json:"freeThrowsAttempted"`
	FreeThrowPercentage  float64 `json:"freeThrowPercentage"`
	OffensiveRebounds    float64 `json:"offensiveRebounds"`
	DefensiveRebounds    float64 `json:"defensiveRebounds"`
	Rebounds             float64 `json:"rebounds"`
	Assists              float64 `json:"assists"`
	Turnovers            float64 `json:"turnovers"`
	Steals               float64 `json:"steals"`
	Blocks               float64 `json:"blocks"`
	BlocksAgainst        float64 `json:"blocksAgainst"`
	PersonalFouls        float64 `json:"personalFouls"`
	PersonalFoulsAgainst float64 `json:"personalFoulsAgainst"`
	PlusMinus            float64 `json:"plusMinus"`
}

type RecordData struct {
	Features [][]float64 `json:"features"`
	Results  []float64   `json:"results"`
}

type RecordDataStore interface {
	GetAll() ([]Record, error)
	Get(years []string) ([]Record, error)
	GetDataSet() (RecordData, error)
}

type csvStore struct {
	files   map[string]string
	records map[string][]Record
}

var _ RecordDataStore = (*csvStore)(nil)

func (c *csvStore) GetAll() ([]Record, error) {
	if c == nil || len(c.records) == 0 {
		return []Record{}, fmt.Errorf("datastore.GetAll: CSV datastore not initialized properly")
	}

	records := []Record{}
	for _, rows := range c.records {
		records = append(records, rows...)
	}

	return records, nil
}

func (c *csvStore) Get(years []string) ([]Record, error) {
	if c == nil || len(c.records) == 0 {
		return []Record{}, fmt.Errorf("datastore.Get: CSV datastore not initialized properly")
	}

	records := []Record{}
	for _, year := range years {
		if rows, ok := c.records[year]; ok {
			records = append(records, rows...)
		}
	}

	return records, nil
}

func (c *csvStore) GetDataSet() (RecordData, error) {
	featuresSet := [][]float64{}
	resultsSet := []float64{}
	for _, recordsByYear := range c.records {
		for _, r := range recordsByYear {
			resultsSet = append(resultsSet, r.WinPercentage)

			features := []float64{
				r.Minutes,
				r.Points,
				r.FieldGoalsMade,
				r.FieldGoalsAttempted,
				r.FieldGoalPercentage,
				r.ThreesMade,
				r.ThreesAttempted,
				r.ThreePercentage,
				r.FreeThrowsMade,
				r.FreeThrowsAttempted,
				r.FreeThrowPercentage,
				r.OffensiveRebounds,
				r.DefensiveRebounds,
				r.Rebounds,
				r.Assists,
				r.Turnovers,
				r.Steals,
				r.Blocks,
				r.BlocksAgainst,
				r.PersonalFouls,
				r.PersonalFoulsAgainst,
				r.PlusMinus,
			}

			featuresSet = append(featuresSet, features)
		}
	}
	return RecordData{Features: featuresSet, Results: resultsSet}, nil
}

func NewCSVStore(directoryPath string) (*csvStore, error) {
	f := map[string]string{
		"1997": "1996-1997.csv",
		"1998": "1997-1998.csv",
		"1999": "1998-1999.csv",
		"2000": "1999-2000.csv",
		"2001": "2000-2001.csv",
		"2002": "2001-2002.csv",
		"2003": "2002-2003.csv",
		"2004": "2003-2004.csv",
		"2005": "2004-2005.csv",
		"2006": "2005-2006.csv",
		"2007": "2006-2007.csv",
		"2008": "2007-2008.csv",
		"2009": "2008-2009.csv",
		"2010": "2009-2010.csv",
		"2011": "2010-2011.csv",
		"2012": "2011-2012.csv",
		"2013": "2012-2013.csv",
		"2014": "2013-2014.csv",
		"2015": "2014-2015.csv",
		"2016": "2015-2016.csv",
		// "2017": "2016-2017.csv",
	}

	records := map[string][]Record{}

	for year, filename := range f {
		rows, err := readCSV(directoryPath + filename)
		if err != nil {
			return nil, fmt.Errorf("NewCSVStore: %w", err)
		}

		records[year] = rows
	}

	return &csvStore{f, records}, nil
}

func readCSV(filepath string) ([]Record, error) {
	records := []Record{}

	readFile, err := os.Open(filepath)

	if err != nil {
		return nil, fmt.Errorf("ReadCSV - failed to read filepath '%s': %w", filepath, err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		record, err := parseRecord(fileScanner.Text())
		if err != nil {
			log.Println(err)
			continue
		}

		records = append(records, record)
	}

	readFile.Close()

	return records, nil
}

type columnIdxMap[T int | float64] struct {
	column      string
	idx         int
	valToAssign *T
}

func parseRecord(row string) (Record, error) {
	r := strings.Split(row, ",")
	if len(r) <= 2 || r[0] == "GP" {
		return Record{}, fmt.Errorf("parseRecord: invalid row")
	}

	rec := Record{}
	intColumnIdxMap := []columnIdxMap[int]{
		{"Games Played", 0, &rec.GamesPlayed},
		{"Wins", 1, &rec.Wins},
		{"Losses", 2, &rec.Losses},
	}

	for _, idxMap := range intColumnIdxMap {
		v, err := strconv.Atoi(r[idxMap.idx])
		if err != nil {
			return Record{}, fmt.Errorf("parseRecord - cannot read '%s' ('%d'): %w", idxMap.column, v, err)
		}

		*idxMap.valToAssign = v
	}

	floatColumnIdxMap := []columnIdxMap[float64]{
		{"WinPercentage", 3, &rec.WinPercentage},
		{"Minutes", 4, &rec.Minutes},
		{"Points", 5, &rec.Points},
		{"FieldGoalsMade", 6, &rec.FieldGoalsMade},
		{"FieldGoalsAttempted", 7, &rec.FieldGoalsAttempted},
		{"FieldGoalPercentage", 8, &rec.FieldGoalPercentage},
		{"ThreesMade", 9, &rec.ThreesMade},
		{"Threes ttempted", 10, &rec.ThreesAttempted},
		{"ThreePercentage", 11, &rec.ThreePercentage},
		{"FreeThrows Made", 12, &rec.FreeThrowsMade},
		{"FreeThrowsAttempted", 13, &rec.FreeThrowsAttempted},
		{"FreeThrowPercentage", 14, &rec.FreeThrowPercentage},
		{"OffensiveRebounds", 15, &rec.OffensiveRebounds},
		{"DefensiveRebounds", 16, &rec.DefensiveRebounds},
		{"Rebounds", 17, &rec.Rebounds},
		{"Assists", 18, &rec.Assists},
		{"Turnovers", 19, &rec.Turnovers},
		{"Steals", 20, &rec.Steals},
		{"Blocks", 21, &rec.Blocks},
		{"BlocksAgainst", 22, &rec.BlocksAgainst},
		{"Personal Fouls", 23, &rec.PersonalFouls},
		{"Personal Fouls Against", 24, &rec.PersonalFoulsAgainst},
		{"PlusMinus", 25, &rec.PlusMinus},
	}

	for _, idxMap := range floatColumnIdxMap {
		v, err := strconv.ParseFloat(r[idxMap.idx], 64)
		if err != nil {
			return Record{}, fmt.Errorf("parseRecord - cannot read '%s' ('%f'): %w", idxMap.column, v, err)
		}

		*idxMap.valToAssign = v
	}

	return rec, nil
}
