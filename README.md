## 🧩 BitTorrent Client (Go)

### Overview

This project is a **lightweight BitTorrent client** implemented in **Go**.
It can download files from peers using a `.torrent` file — demonstrating how the BitTorrent protocol works under the hood: parsing torrent metadata, connecting to peers, and downloading verified pieces.

The goal is educational clarity and simplicity rather than a full-featured production client.

---

### ✨ Features

* Parses `.torrent` files
* Connects to peers discovered via tracker
* Downloads file pieces concurrently
* Verifies each piece using SHA-1 hashes
* Assembles the complete file from verified pieces
* Simple, minimal entry point with no CLI flags

### 🚀 Getting Started

#### 1. Prerequisites

* Go **1.23+**
* Internet connection
* A valid `.torrent` file

#### 2. Clone the Repository

```bash
git clone https://github.com/onahvictor/torrent-client.git
cd torrent-client
```

#### 3. Build the Binary

```bash
go build -o torrent-client ./cmd
```

#### 4. Run the Client

The client takes **two arguments**:

1. Path to the `.torrent` file
2. Output directory for the downloaded file

Example:

```bash
./torrent-client ./testdata/example.torrent ./downloads
```

---

### ⚙️ How It Works

1. **Open the torrent file**

   * Reads and parses the `.torrent` metadata (announce URL, info hash, piece hashes, etc.).
2. **Announce to tracker**

   * Contacts the tracker to get a list of available peers.
3. **Handshake with peers**

   * Performs the BitTorrent handshake to confirm they share the same torrent.
4. **Download pieces**

   * Requests and downloads file pieces concurrently from multiple peers.
5. **Verify and assemble**

   * Verifies each piece using SHA-1 hash and assembles the final file once all pieces are complete.

---

### 📦 Example Output

```text
Opening torrent file: example.torrent
Connecting to tracker...
Found 8 peers
Downloading piece 4/256
Piece 4 verified ✅
Download complete: ./downloads/example.iso
```

---

### 🧠 Key Concepts

| Concept          | Description                                                       |
| ---------------- | ----------------------------------------------------------------- |
| **Torrent file** | A metadata file describing the files and trackers for the torrent |
| **Tracker**      | A server that helps peers discover each other                     |
| **Peer**         | Another client sharing the same torrent                           |
| **Piece**        | A fixed-size chunk of the full file                               |
| **Info hash**    | Unique identifier for the torrent content                         |

---

### 🧩 Future Improvements

* Add support for **flags** (e.g., `-torrent`, `-out`)
* Add **DHT** and **magnet link** support
* Implement **upload (seeding)**
* Handle **multi-file torrents**
* Improve download progress display

