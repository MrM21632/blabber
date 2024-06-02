package uidgen

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
	"users/envvars"
)

const (
	TimestampLength uint8 = 42 // Timestamp section
	ServerIdLength  uint8 = 11 // Server ID section
	SequenceLength  uint8 = 11 // UID Sequence section

	MaxSequence  uint16 = 1<<SequenceLength - 1  // Max possible UID sequence
	MaxServerId  uint16 = 1<<ServerIdLength - 1  // Max possible server ID
	MaxTimestamp uint64 = 1<<TimestampLength - 1 // Max possible timestamp

	serverIdShift  = SequenceLength
	timestampShift = ServerIdLength + SequenceLength
)

type UniqueId uint64

type UniqueIdGenerator struct {
	mutex sync.Mutex
	epoch time.Time

	serverId uint64
	currTime uint64
	currSeq  uint64
}

// Epoch is configurable using the UIDGEN_EPOCH_MS environment variable. This allows for quick, reconfigurable
// deployemnts of the UID generation service, especially when scale is required.
//
// If the epoch is not found, we default to the epoch used by Twitter's original Snowflake algorithm.
func GetEpoch() uint64 {
	var epoch uint64
	var err error
	if epoch, err = envvars.GetenvInteger("UIDGEN_EPOCH_MS"); err != nil {
		fmt.Println("get epoch failed: envvar UIDGEN_EPOCH_MS not found, using default")
		return 1288834974657
	}

	return epoch
}

// Server ID is configurable using the UIDGEN_NODE_ID environment variable. This allows for quick, reconfigurable
// deployments of the UID generation service, especially when scale is required.
//
// If the server ID is not found, we default to 0.
func GetServerId() uint64 {
	var serverId uint64
	var err error
	if serverId, err = envvars.GetenvInteger("UIDGEN_NODE_ID"); err != nil {
		fmt.Println("get server id failed: envvar UIDGEN_NODE_ID not found, using default")
		return 0
	}

	return serverId
}

func InitializeNode() (*UniqueIdGenerator, error) {
	if SequenceLength+ServerIdLength+TimestampLength != 64 {
		return nil, errors.New("initialization failed: combined length of sections is invalid")
	}

	serverId := GetServerId()
	epoch := GetEpoch()
	epochTime := time.Unix(int64(epoch)/1000, (int64(epoch)%1000)*1000000)

	result := UniqueIdGenerator{}
	result.serverId = serverId
	if result.serverId > uint64(MaxServerId) {
		return nil, errors.New(
			"initialization failed: server ID must between 0 and " + strconv.FormatInt(int64(MaxServerId), 10),
		)
	}

	now := time.Now()
	result.epoch = now.Add(epochTime.Sub(now))
	return &result, nil
}

func (node *UniqueIdGenerator) GeanerateId() UniqueId {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	now := time.Since(node.epoch).Milliseconds()
	if now == int64(node.currTime) {
		node.currSeq = (node.currSeq + 1) & uint64(MaxSequence)
		if node.currSeq == 0 {
			for now <= int64(node.currTime) {
				now = int64(time.Since(node.epoch).Milliseconds())
			}
		}
	} else {
		node.currSeq = 0
	}

	node.currTime = uint64(now)
	result := UniqueId(
		(uint64(now) << uint64(timestampShift)) | (node.serverId << uint64(serverIdShift)) | node.currSeq,
	)
	return result
}
