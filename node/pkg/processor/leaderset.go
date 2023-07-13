package processor

import "encoding/binary"

func calculateLeaderSetSize(guardianSetSize int) int {
	return guardianSetSize/3 + 1
}

func isLeader(hash []byte, myGuardianSetIndex int, numGuardians int) bool {
	// determine if this guardian is responsible for this observation
	hashId := binary.BigEndian.Uint64(hash)
	targetIdx := int(hashId % uint64(numGuardians)) // TODO support variable size guardian set
	r := (targetIdx + myGuardianSetIndex) % numGuardians
	return r < calculateLeaderSetSize(numGuardians)
}
