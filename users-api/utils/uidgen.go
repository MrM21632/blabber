package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
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

type UniqueID uint64

type UniqueIDGenerator struct {
	mutex sync.Mutex
	epoch time.Time

	serverId uint64
	currTime uint64
	currSeq  uint64
}

func InitializeNode() (*UniqueIDGenerator, error) {
	if SequenceLength+ServerIdLength+TimestampLength != 64 {
		err := errors.New("initialization failed: combined length of sections is invalid")
		log.Error(err.Error())
		return nil, err
	}

	serverId := GetUidgenServerID()
	epoch := GetUidgenEpoch()
	epochTime := time.Unix(int64(epoch)/1000, (int64(epoch)%1000)*1000000)

	result := UniqueIDGenerator{}
	result.serverId = serverId
	if result.serverId > uint64(MaxServerId) {
		err := errors.New(
			"initialization failed: server ID must between 0 and " + strconv.FormatInt(int64(MaxServerId), 10),
		)
		log.Error(err.Error())
		panic("")
	}

	now := time.Now()
	result.epoch = now.Add(epochTime.Sub(now))
	return &result, nil
}

func (node *UniqueIDGenerator) GeanerateId() UniqueID {
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
	result := UniqueID(
		(uint64(now) << uint64(timestampShift)) | (node.serverId << uint64(serverIdShift)) | node.currSeq,
	)
	return result
}
