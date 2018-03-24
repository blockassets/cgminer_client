package cgminer_client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

var (
	// Makefile build
	version = ""
)

type Client struct {
	server  string
	timeout time.Duration
}

// New returns a Client pointer, which is used to communicate with a running
// cgminer instance. Note that New does not attempt to connect to the miner.
func New(hostname string, port int64, timeout time.Duration) *Client {
	miner := &Client{}
	miner.server = fmt.Sprintf("%s:%d", hostname, port)
	miner.timeout = time.Second * timeout
	return miner
}

func (miner *Client) runCommand(command, argument string) (string, error) {
	conn, err := net.DialTimeout("tcp", miner.server, miner.timeout)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	type commandRequest struct {
		Command   string `json:"command"`
		Parameter string `json:"parameter,omitempty"`
	}

	request := &commandRequest{
		Command: command,
	}

	if argument != "" {
		request.Parameter = argument
	}

	requestBody, err := jsoniter.Marshal(request)
	if err != nil {
		return "", err
	}

	fmt.Fprintf(conn, "%s", requestBody)
	result, err := bufio.NewReader(conn).ReadString('\x00')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(result, "\x00"), nil
}

// Devs returns basic information on the miner.
func (miner *Client) Devs() (*[]Dev, error) {
	response, err := miner.runCommand("devs", "")
	if err != nil {
		return nil, err
	}

	result, err := processDevs(response)
	return &result.Devs, err
}

/* Makes testing easier */
func processDevs(response string) (*DevsResponse, error) {
	devsResponse := &DevsResponse{}
	err := jsoniter.Unmarshal([]byte(response), devsResponse)
	if err != nil {
		return nil, err
	}

	return devsResponse, err
}

func (miner *Client) ChipStat() (*[]ChipStat, error) {
	response, err := miner.runCommand("chipstat", "")
	if err != nil {
		return nil, err
	}

	result, err := processChipStat(response)
	return &result.ChipStats, err
}

/* Makes testing easier */
func processChipStat(response string) (*ChipStatResponse, error) {
	chipStatResponse := &ChipStatResponse{}
	err := jsoniter.Unmarshal([]byte(response), chipStatResponse)
	if err != nil {
		return nil, err
	}

	chipStatResponse.ChipStats = make([]ChipStat, len(chipStatResponse.Data))
	// All the json contents now lives in the generic []Data map.
	for i, data := range chipStatResponse.Data {

		// There is no good way of mapping json -> struct when there is an
		// unknown set of properties on an object. In this, case it is all
		// the 1_accept properties. Map the known properties by hand, delete
		// them and then map all the accept properties.
		chipStat := &chipStatResponse.ChipStats[i]
		chipStat.ASC = int(data["ASC"].(float64))
		chipStat.Name = data["Name"].(string)
		chipStat.ID = int(data["ID"].(float64))

		delete(data, "ASC")
		delete(data, "Name")
		delete(data, "ID")

		chipStat.Accept = make(map[string]int64)
		for key, val := range data {
			chipStat.Accept[key] = int64(val.(float64))
		}
	}

	return chipStatResponse, err
}

// Summary returns basic information on the miner.
func (miner *Client) Summary() (*Summary, error) {
	response, err := miner.runCommand("summary", "")
	if err != nil {
		return nil, err
	}

	result, err := processSummary(response)
	return &result.Summary[0], err
}

/* Makes testing easier */
func processSummary(response string) (*SummaryResponse, error) {
	summaryResponse := &SummaryResponse{}
	err := jsoniter.Unmarshal([]byte(response), summaryResponse)
	if err != nil {
		return nil, err
	}

	if len(summaryResponse.Summary) != 1 {
		return nil, errors.New("Received multiple Summary objects")
	}

	return summaryResponse, err
}

func (miner *Client) Restart() error {
	_, err := miner.runCommand("restart", "")
	return err
}

//
func (miner *Client) Quit() error {
	_, err := miner.runCommand("quit", "")
	return err
}

//// Pools returns a slice of Pool structs, one per pool.
//func (miner *Client) Pools() ([]Pool, error) {
//	result, err := miner.runCommand("pools", "")
//	if err != nil {
//		return nil, err
//	}
//
//	var poolsResponse poolsResponse
//	err = jsoniter.Unmarshal([]byte(result), &poolsResponse)
//	if err != nil {
//		return nil, err
//	}
//
//	var pools = poolsResponse.Pools
//	return pools, nil
//}

//// AddPool adds the given URL/username/password combination to the miner's
//// pool list.
//func (miner *Client) AddPool(url, username, password string) error {
//	// TODO: Don't allow adding a pool that's already in the pool list
//	// TODO: Escape commas in the URL, username, and password
//	parameter := fmt.Sprintf("%s,%s,%s", url, username, password)
//	result, err := miner.runCommand("addpool", parameter)
//	if err != nil {
//		return err
//	}
//
//	var addPoolResponse addPoolResponse
//	err = jsoniter.Unmarshal([]byte(result), &addPoolResponse)
//	if err != nil {
//		// If there an error here, it's possible that the pool was actually added
//		return err
//	}
//
//	status := addPoolResponse.Status[0]
//
//	if status.Status != "S" {
//		return errors.New(fmt.Sprintf("%d: %s", status.Code, status.Description))
//	}
//
//	return nil
//}
//
//func (miner *Client) Enable(pool *Pool) error {
//	parameter := fmt.Sprintf("%d", pool.Pool)
//	_, err := miner.runCommand("enablepool", parameter)
//	return err
//}
//
//func (miner *Client) Disable(pool *Pool) error {
//	parameter := fmt.Sprintf("%d", pool.Pool)
//	_, err := miner.runCommand("disablepool", parameter)
//	return err
//}
//
//func (miner *Client) Delete(pool *Pool) error {
//	parameter := fmt.Sprintf("%d", pool.Pool)
//	_, err := miner.runCommand("removepool", parameter)
//	return err
//}
//
//func (miner *Client) SwitchPool(pool *Pool) error {
//	parameter := fmt.Sprintf("%d", pool.Pool)
//	_, err := miner.runCommand("switchpool", parameter)
//	return err
//}
//
