package torrentfile

import (
	"context"
	"crypto/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/jackpal/bencode-go"
	"github.com/onahvictor/torrent-client/p2p"
	"github.com/onahvictor/torrent-client/peers"
)

const Port uint16 = 6881

type trackerResponse struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

// application specific
type TorrentFile struct {
	Announce     string
	InfoHash     [20]byte
	PiecesHashes [][20]byte
	PieceLength  int
	Length       int
	Name         string
}

func (t *TorrentFile) DownlaodFile(path string) error {
	var peerID [20]byte
	_, err := rand.Read(peerID[:])
	if err != nil {
		return err
	}

	peers, err := t.requestPeersFromTrackers(peerID, Port)
	if err != nil {
		return err
	}

	torrent := p2p.Torrent{
		Peers:       peers,
		PeerID:      peerID,
		InfoHash:    t.InfoHash,
		PieceHashes: t.PiecesHashes,
		PieceLength: t.PieceLength,
		Length:      t.Length,
		Name:        t.Name,
	}
	buf, err := torrent.Download()
	if err != nil {
		return err
	}

	outFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = outFile.Write(buf)
	if err != nil {
		return err
	}
	return nil

}
func (t *TorrentFile) buildTrackerUrl(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}

	params := make(url.Values)
	params.Set("info_hash", string(t.InfoHash[:]))
	params.Set("peer_id", string(peerID[:]))
	params.Set("port", strconv.Itoa(int(port)))
	params.Set("uploaded", "0")
	params.Set("downloaded", "0")
	params.Set("compact", "1")
	params.Set("left", strconv.Itoa(t.Length))

	base.RawQuery = params.Encode()
	return base.String(), nil
}

func (t *TorrentFile) requestPeersFromTrackers(peerID [20]byte, port uint16) ([]peers.Peer, error) {
	url, err := t.buildTrackerUrl(peerID, port)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var trackerResp trackerResponse
	err = bencode.Unmarshal(resp.Body, &trackerResp)
	if err != nil {
		return nil, err
	}

	peers, err := peers.Unmarshal([]byte(trackerResp.Peers))
	if err != nil {
		return nil, err
	}
	return peers, nil
}
