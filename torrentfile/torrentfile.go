package torrentfile

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

type bencodeTorrent struct {
	Announce string       `bencode:"announce"`
	Info     bencodedInfo `bencode:"info"`
}

type bencodedFileInfo struct { // multi file data
	Length int      `bencode:"length"`
	Path   []string `bencode:"path"`
}

type bencodedInfo struct {
	Pieces      string             `bencode:"pieces"` //a slices of pieces
	PieceLength int                `bencode:"piece lenght"`
	Name        string             `bencode:"name"`
	Length      int                `bencode:"length,omitempty"` //single file
	Files       []bencodedFileInfo `bencode:"files,omitempty"`  //multi file
}

func (i *bencodedInfo) hash() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *i)
	if err != nil {
		return [20]byte{}, err
	}

	h := sha1.Sum(buf.Bytes())
	return h, nil
}

func (i *bencodedInfo) splitPieceHashes() ([][20]byte, error) {
	hashLen := 20 //len of SHA-1 hashes
	buf := []byte(i.Pieces)

	if len(buf)%hashLen != 0 {
		err := fmt.Errorf("recieved malformed piece of length %d", len(buf))
		return nil, err
	}

	var hashPieces [][20]byte

	for i := 0; i < len(buf); i += 20 {
		chunk := buf[i : i+20]
		hashPieces = append(hashPieces, [20]byte(chunk))
	}
	return hashPieces, nil
}

// torrent file is bencoded
func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return TorrentFile{}, err
	}
	defer file.Close()

	bto := bencodeTorrent{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		return TorrentFile{}, err
	}
	return bto.toTorrentFile()
}

func (bto *bencodeTorrent) toTorrentFile() (TorrentFile, error) {
	infoHash, err := bto.Info.hash()
	if err != nil {
		return TorrentFile{}, err
	}
	hash, err := bto.Info.splitPieceHashes()
	if err != nil {
		return TorrentFile{}, err
	}
	t := TorrentFile{
		Announce:     bto.Announce,
		InfoHash:     infoHash,
		PieceLength:  bto.Info.Length,
		Name:         bto.Info.Name,
		Length:       bto.Info.Length,
		PiecesHashes: hash,
	}
	return t, nil
}
