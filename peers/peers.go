package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP `bencode:"ip"`
	Port uint16 `bencode:"port"`
}

func Unmarshal(peers []byte) ([]*Peer, error) {
	peerSize := 6 //4 for ip and 2 for port
	if len(peers)%peerSize != 0 {
		err := fmt.Errorf("malformed peer size %v", len(peers))
		return nil, err
	}

	var peerList []*Peer
	for i := 0; i < len(peers); i += 6 {
		var peer Peer
		peer.IP = net.IP(peers[i : i+4])
		peer.Port = binary.BigEndian.Uint16(peers[i+4 : i+6])
		peerList = append(peerList, &peer)
	}
	return peerList, nil
}

func (p *Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
