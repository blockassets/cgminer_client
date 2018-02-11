package cgminer_client

type Map map[string]interface{}
type MapInt64 map[string]int64

/**
  "STATUS": [
    {
      "STATUS": "S",
      "When": 1517894216,
      "Code": 9,
      "Msg": "4 ASC(s)",
      "Description": "cgminer 4.10.0"
    }
  ],
*/
type Status struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
	Status      string `json:"STATUS"`
	When        int64  `json:"When"`
}

type DevCommon struct {
	Name string `json:"Name,omit"` // omit is special for prometheus_helper library to exclude it
	ASC  int    `json:"ASC"`
	ID   int    `json:"ID"`
}

type Dev struct {
	DevCommon
	Enabled               string  `json:"Enabled"`
	Status                string  `json:"Status"`
	Temperature           float64 `json:"Temperature"`
	MHSav                 float64 `json:"MHS av"`
	MHS5s                 float64 `json:"MHS 5s"`
	MHS1m                 float64 `json:"MHS 1m"`
	MHS5m                 float64 `json:"MHS 5m"`
	MHS15m                float64 `json:"MHS 15m"`
	Accepted              int64   `json:"Accepted"`
	Rejected              int64   `json:"Rejected"`
	HardwareErrors        int64   `json:"Hardware Errors"`
	Utility               float64 `json:"Utility"`
	TotalMH               int64   `json:"TotalMH"`
	Diff1Work             int64   `json:"Diff1 Work"`
	DifficultyAccepted    float64 `json:"Difficulty Accepted"`
	DifficultyRejected    float64 `json:"Difficulty Rejected"`
	LastShareDifficulty   float64 `json:"Last Share Difficulty"`
	LastValidWork         int64   `json:"Last Valid Work"`
	DeviceHardwarePercent float64 `json:"Device Hardware%"`
	DeviceRejectedPercent float64 `json:"Device Rejected%"`
	DeviceElapsed         int64   `json:"Device Elapsed"`
}

type DevsResponse struct {
	Status []Status `json:"STATUS"`
	Devs   []Dev    `json:"DEVS"`
	Id     int64    `json:"id"`
}

/**
  "DEVS": [
    {
      "ASC": 0,
      "Name": "ttyS1",
      "ID": 1,
      "1_accept": 154,
      "2_accept": 132,
      "3_accept": 138,
      "4_accept": 129,
*/
type ChipStat struct {
	DevCommon
	Accept MapInt64
}

type ChipStatResponse struct {
	Status    []Status `json:"STATUS"`
	Data      []Map    `json:"DEVS"` // data goes into the map and then gets transformed into ChipStat
	Id        int64    `json:"id"`
	ChipStats []ChipStat
}

//type Pool struct {
//	Pool				int		`json:"POOL"`
//	URL					string
//	Status				string
//	Priority               int64
//	Quota                  int64
//	Accepted               int64
//	Rejected               int64
//	User                   string
//	LastShareTime          int64   `json:"Last Share Time"`
//
//	//"POOL": 0,
//	//"URL": "stratum+tcp://us2.litecoinpool.org:3333",
//	//"Status": "Alive",
//	//"Priority": 16842752,
//	//"Quota": 1,
//	//"Long Poll": "Y",
//	//"Getworks": 0,
//	//"Accepted": 1273,
//	//"Rejected": 76,
//	//"Works": 0,
//	//"Discarded": 0,
//	//"Stale": 0,
//	//"Get Failures": 0,
//	//"Remote Failures": 0,
//	//"User": "hodl4now.10011",
//	//"Passwd": "123123",
//	//"Last Share Time": 1516693404,
//	//"Diff1 Shares": 2048,
//	//"Proxy Type": "",
//	//"Proxy": "",
//	//"Difficulty Accepted": 0.00000000,
//	//"Difficulty Rejected": 0.00000000,
//	//"Difficulty Stale": 0.00000000,
//	//"Last Share Difficulty": 0.00000000,
//	//"Has Stratum": true,
//	//"Stratum Active": true,
//	//"Stratum URL": "stratum+tcp://us2.litecoinpool.org:3333",
//	//"Has GBT": false,
//	//"Best Share": 0,
//	//"Pool Rejected%": 0.0000,
//	//"Pool Stale%": 0.0000,
//	//"Bad Work": 0
//
//
//
//	//Accepted               int64
//	//BestShare              int64   `json:"Best Share"`
//	//Diff1Shares            int64   `json:"Diff1 Shares"`
//	//DifficultyAccepted     float64 `json:"Difficulty Accepted"`
//	//DifficultyRejected     float64 `json:"Difficulty Rejected"`
//	//DifficultyStale        float64 `json:"Difficulty Stale"`
//	//Discarded              int64
//	//GetFailures            int64 `json:"Get Failures"`
//	//Getworks               int64
//	//HasGBT                 bool    `json:"Has GBT"`
//	//HasStratum             bool    `json:"Has Stratum"`
//	//LastShareDifficulty    float64 `json:"Last Share Difficulty"`
//	//LastShareTime          int64   `json:"Last Share Time"`
//	//LongPoll               string  `json:"Long Poll"`
//	//Pool                   int64   `json:"POOL"`
//	//PoolRejectedPercentage float64 `json:"Pool Rejected%"`
//	//PoolStalePercentage    float64 `json:"Pool Stale%"`
//	//Priority               int64
//	//ProxyType              string `json:"Proxy Type"`
//	//Proxy                  string
//	//Quota                  int64
//	//Rejected               int64
//	//RemoteFailures         int64 `json:"Remote Failures"`
//	//Stale                  int64
//	//Status                 string
//	//StratumActive          bool   `json:"Stratum Active"`
//	//StratumURL             string `json:"Stratum URL"`
//	//URL                    string
//	//User                   string
//	//Works                  int64
//}

type SummaryResponse struct {
	Status  []Status  `json:"STATUS"`
	Summary []Summary `json:"SUMMARY"`
	Id      int64     `json:"id"`
}

//type poolsResponse struct {
//	Status []status `json:"STATUS"`
//	Pools  []Pool   `json:"POOLS"`
//	Id     int64    `json:"id"`
//}
//
//type addPoolResponse struct {
//	Status []status `json:"STATUS"`
//	Id     int64    `json:"id"`
//}

/*
   "Elapsed": 48417,
   "MHS av": 335.86,
   "MHS 5s": 357.15,
   "MHS 1m": 323.17,
   "MHS 5m": 357.15,
   "MHS 15m": 339.55,
   "temper1": 31,
   "temper2": 31,
   "temper3": 31,
   "temper4": 31,
   "Found Blocks": 0,
   "Getworks": 3236,
   "Accepted": 8215,
   "Rejected": 75,
   "Hardware Errors": 1654,
   "Utility": 10.18,
   "Discarded": 48716,
   "Stale": 0,
   "Get Failures": 0,
   "Local Work": 77113,
   "Remote Failures": 0,
   "Network Blocks": 366,
   "Total MH": 16261155,
   "Work Utility": 1.39,
   "Difficulty Accepted": 4538.375,
   "Difficulty Rejected": 42.875,
   "Difficulty Stale": 0,
   "Best Share": 8322,
   "Device Hardware%": 59.6036,
   "Device Rejected%": 3.8247,
   "Pool Rejected%": 0.9359,
   "Pool Stale%": 0,
   "Last getwork": 1517965137
*/
type Summary struct {
	Elapsed               int64   `json:"Elapsed"`
	MHSav                 float64 `json:"MHS av"`
	MHS5s                 float64 `json:"MHS 5s"`
	MHS1m                 float64 `json:"MHS 1m"`
	MHS5m                 float64 `json:"MHS 5m"`
	MHS15m                float64 `json:"MHS 15m"`
	Temper1               float64 `json:"temper1"`
	Temper2               float64 `json:"temper2"`
	Temper3               float64 `json:"temper3"`
	Temper4               float64 `json:"temper4"`
	FoundBlocks           int64   `json:"Found Blocks"`
	Getworks              int64   `json:"Getworks"`
	Accepted              int64   `json:"Accepted"`
	Rejected              int64   `json:"Rejected"`
	HardwareErrors        int64   `json:"Hardware Errors"`
	Utility               float64 `json:"Utility"`
	Discarded             int64   `json:"Discarded"`
	Stale                 int64   `json:"Stale"`
	GetFailures           int64   `json:"Get Failures"`
	LocalWork             int64   `json:"Local Work"`
	RemoteFailures        int64   `json:"Remote Failures"`
	NetworkBlocks         int64   `json:"Network Blocks"`
	TotalMH               float64 `json:"Total MH"`
	WorkUtility           float64 `json:"Work Utility"`
	DifficultyAccepted    float64 `json:"Difficulty Accepted"`
	DifficultyRejected    float64 `json:"Difficulty Rejected"`
	DifficultyStale       float64 `json:"Difficulty Stale"`
	BestShare             int64   `json:"Best Share"`
	DeviceHardwarePercent float64 `json:"Device Hardware%"`
	DeviceRejectedPercent float64 `json:"Device Rejected%"`
	PoolRejectedPercent   float64 `json:"Pool Rejected%"`
	PoolStalePercent      float64 `json:"Pool Stale%"`
	LastGetwork           int64   `json:"Last getwork"`
}
