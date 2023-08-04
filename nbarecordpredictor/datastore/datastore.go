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

type RecordDataStore interface {
	GetAll() ([]Record, error)
	Get(years []string) ([]Record, error)
}

type CSVStore struct {
	files   map[string]string
	records map[string][]Record
}

func (c *CSVStore) GetAll() ([]Record, error) {
	if c == nil || len(c.records) == 0 {
		return []Record{}, fmt.Errorf("datastore.GetAll: CSV datastore not initialized properly")
	}

	records := []Record{}
	for _, rows := range c.records {
		records = append(records, rows...)
	}

	return records, nil
}

func (c *CSVStore) Get(years []string) ([]Record, error) {
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

func NewCSVStore(directoryPath string) (*CSVStore, error) {
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

	return &CSVStore{f, records}, nil
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

func parseRecord(row string) (Record, error) {
	r := strings.Split(row, ",")
	if len(r) <= 2 || r[0] == "GP" {
		return Record{}, fmt.Errorf("parseRecord: invalid row")
	}

	gamesPlayed, err := strconv.Atoi(r[0])
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read games played '%s': %w", r[0], err)
	}

	wins, err := strconv.Atoi(r[1])
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read wins '%s': %w", r[1], err)
	}

	losses, err := strconv.Atoi(r[2])
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read losses '%s': %w", r[2], err)
	}

	winPercentage, err := strconv.ParseFloat(r[3], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read win percentage '%s': %w", r[3], err)
	}

	minutes, err := strconv.ParseFloat(r[4], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read minutes '%s': %w", r[4], err)
	}

	points, err := strconv.ParseFloat(r[5], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read points '%s': %w", r[5], err)
	}

	fieldGoalsMade, err := strconv.ParseFloat(r[6], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read field goals made '%s': %w", r[6], err)
	}

	fieldGoalsAttempted, err := strconv.ParseFloat(r[7], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read field goals attempted '%s': %w", r[7], err)
	}

	fieldGoalPercentage, err := strconv.ParseFloat(r[8], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read field goal percentage '%s': %w", r[8], err)
	}

	threesMade, err := strconv.ParseFloat(r[9], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read threes made '%s': %w", r[9], err)
	}

	threesAttempted, err := strconv.ParseFloat(r[10], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read win threes attempted '%s': %w", r[10], err)
	}

	threePercentage, err := strconv.ParseFloat(r[11], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read three percentage '%s': %w", r[11], err)
	}

	freeThrowsMade, err := strconv.ParseFloat(r[12], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read free throws made '%s': %w", r[12], err)
	}

	freeThrowsAttempted, err := strconv.ParseFloat(r[13], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read free throws attempted '%s': %w", r[13], err)
	}

	freeThrowPercentage, err := strconv.ParseFloat(r[14], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read free throws percentage '%s': %w", r[14], err)
	}

	offensiveRebounds, err := strconv.ParseFloat(r[15], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read offensive rebounds '%s': %w", r[15], err)
	}

	defensiveRebounds, err := strconv.ParseFloat(r[16], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read defensive rebounds '%s': %w", r[16], err)
	}

	rebounds, err := strconv.ParseFloat(r[17], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read rebounds '%s': %w", r[17], err)
	}

	assists, err := strconv.ParseFloat(r[18], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read assists '%s': %w", r[18], err)
	}

	turnovers, err := strconv.ParseFloat(r[19], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read turnovers '%s': %w", r[19], err)
	}

	steals, err := strconv.ParseFloat(r[20], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read steals '%s': %w", r[20], err)
	}

	blocks, err := strconv.ParseFloat(r[21], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read blocks '%s': %w", r[21], err)
	}

	blocksAgainst, err := strconv.ParseFloat(r[22], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read blocks against '%s': %w", r[22], err)
	}

	personalFouls, err := strconv.ParseFloat(r[23], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read personal fouls '%s': %w", r[23], err)
	}

	personalFoulsAgainst, err := strconv.ParseFloat(r[24], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read personal fouls against '%s': %w", r[24], err)
	}

	plusMinus, err := strconv.ParseFloat(r[25], 64)
	if err != nil {
		return Record{}, fmt.Errorf("parseRecord - cannot read plus/minus '%s': %w", r[25], err)
	}

	return Record{
		GamesPlayed:          gamesPlayed,
		Wins:                 wins,
		Losses:               losses,
		WinPercentage:        winPercentage,
		Minutes:              minutes,
		Points:               points,
		FieldGoalsMade:       fieldGoalsMade,
		FieldGoalsAttempted:  fieldGoalsAttempted,
		FieldGoalPercentage:  fieldGoalPercentage,
		ThreesMade:           threesMade,
		ThreesAttempted:      threesAttempted,
		ThreePercentage:      threePercentage,
		FreeThrowsMade:       freeThrowsMade,
		FreeThrowsAttempted:  freeThrowsAttempted,
		FreeThrowPercentage:  freeThrowPercentage,
		OffensiveRebounds:    offensiveRebounds,
		DefensiveRebounds:    defensiveRebounds,
		Rebounds:             rebounds,
		Assists:              assists,
		Turnovers:            turnovers,
		Steals:               steals,
		Blocks:               blocks,
		BlocksAgainst:        blocksAgainst,
		PersonalFouls:        personalFouls,
		PersonalFoulsAgainst: personalFoulsAgainst,
		PlusMinus:            plusMinus,
	}, nil
}
