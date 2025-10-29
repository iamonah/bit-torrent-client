## 🧩 BitTorrent Client in Go

### Overview

This project is a **minimal BitTorrent client** implemented in **Go**.
It demonstrates how peer-to-peer (P2P) communication works in the BitTorrent protocol — including connecting to trackers, parsing torrent metadata, establishing peer handshakes, and downloading file pieces from multiple peers concurrently.

The goal is to provide a clear, well-structured implementation of BitTorrent fundamentals rather than a fully featured production client.

---

### ✨ Features

* 🧱 **Torrent file parsing** (`.torrent` files)
* 🌐 **Tracker communication** via HTTP/UDP
* 🔗 **Peer discovery and handshake**
* 📦 **Piece downloading and verification (SHA-1)**
* ⚡ **Concurrent downloading from multiple peers**
* 🧠 **Piece selection strategy** (e.g., rarest-first)
* 🧾 **Resume and integrity check support (optional)**
* 🧰 Clean, idiomatic Go structure (service, transport, etc.)


### 🚀 Getting Started

#### 1. Prerequisites

* Go **1.23+**
* Internet connection
* A `.torrent` file (for testing)

#### 2. Clone the repository

```bash
git clone https://github.com/yourusername/bittorrent-client.git
cd bittorrent-client
```

#### 3. Build

```bash
go build -o bittorrent cmd/main.go
```

#### 4. Run

```bash
./bittorrent -torrent testdata/ubuntu.torrent -output ./downloads
```

---

### ⚙️ Configuration Flags

| Flag         | Description                        | Example                          |
| ------------ | ---------------------------------- | -------------------------------- |
| `-torrent`   | Path to the `.torrent` file        | `-torrent testdata/file.torrent` |
| `-output`    | Directory to save downloaded files | `-output ./downloads`            |
| `-max-peers` | Limit number of peers              | `-max-peers 50`                  |
| `-debug`     | Enable verbose logging             | `-debug`                         |

---

### 📡 How It Works

1. **Parse `.torrent` file**

   * Extracts `announce`, `info hash`, `piece length`, and `pieces`.

2. **Connect to Tracker**

   * Announces the client’s presence and retrieves a list of peers.

3. **Peer Handshake**

   * Performs the BitTorrent handshake to verify shared torrent info.

4. **Message Exchange**

   * Exchanges standard messages (`interested`, `request`, `piece`, etc.) between peers.

5. **Piece Download & Verification**

   * Downloads file pieces concurrently.
   * Verifies SHA-1 hash for each piece.

6. **Assemble Final File**

   * Merges pieces into the target file once all pieces are verified.

---

### 🧠 Key Concepts

* **Info Hash** → Unique identifier for a torrent file.
* **Tracker** → Coordinates peers sharing the same torrent.
* **Peer** → Another client in the swarm that uploads/downloads pieces.
* **Piece** → A fixed-size chunk of the file(s) being shared.
* **Swarm** → All peers connected via the same torrent.

---

### 🧩 Example Output

```text
[INFO] Connecting to tracker: http://tracker.example.com:8080/announce
[INFO] Found 12 peers
[INFO] Downloading piece 3/248 (1.2%)
[INFO] Verified piece 3 ✅
[INFO] Download complete! File saved to ./downloads/
```

---

### 🔒 Notes

* This implementation follows **BitTorrent v1** specification (BEP 3).
* For simplicity, **DHT**, **magnet links**, and **PEX** are not yet implemented.
* For educational use — not intended for copyrighted materials.

---

### 🧰 Future Enhancements

* Support for **DHT** and **magnet URIs**
* Implement **uploading (seeding)** mode
* Add **rate limiting** and **choke/unchoke** strategy
* Support for **multi-file torrents**
* Improve **resumable downloads**

