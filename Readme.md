# Go-NihonVoice

Go-NihonVoice is a Go-based CLI tool for converting Japanese text into natural-sounding speech using **Google Text-to-Speech**. It supports multiple voices, adjustable speech rate, pitch, and can output audio in **MP3** or **WAV** format.

This tool is ideal for developers, language learners, or anyone needing programmatic Japanese TTS in Go projects.

---

## Features

- Convert Japanese text to speech.
- Supports **8 voice options**:
  - Standard Voices: `stand-a`, `stand-b`, `stand-c`, `stand-d`
  - Wavenet Voices: `wave-a`, `wave-b`, `wave-c`, `wave-d`
- Adjustable **speech rate** (`0.25 ~ 4.0`) and **pitch** (`-20.0 ~ 20.0`).
- Output audio formats: **MP3** and **WAV**.
- Simple CLI interface with clear error messages for invalid input.

---

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/go-nihonvoice.git
cd go-nihonvoice
```

## Install dependencies:

```bash
Copy code
go mod tidy
Add your Google service account key for TTS API access:
```

**Place it in the project root as sample-service-account-key.json.**

Usage
Run directly with go run
```bash
Copy code
go run main.go -text "こんにちは" -voice stand-a -rate 1.0 -pitch 0.0 -o output.mp3
Or build a binary first
bash
Copy code
go build -o nihonvoice
./nihonvoice -text "おはようございます" -voice wave-b -rate 1.2 -pitch 2.0 -o morning.mp3
CLI Flags
Flag	Description	Default	Notes
-text	Text to convert to speech	(required)	Cannot be empty
-voice	Speaker's voice	stand-a	Options: stand-a, stand-b, stand-c, stand-d, wave-a, wave-b, wave-c, wave-d
-rate	Speech rate	1.0	Valid range: 0.25 ~ 4.0
-pitch	Speaking pitch	0.0	Valid range: -20.0 ~ 20.0
-o	Output audio file	(required)	Must end with .mp3 or .wav
```
**Validation Rules**
Text: Cannot be empty.

Voice: Must match one of the 8 supported voices.

Rate: 0.25 ≤ rate ≤ 4.0.

Pitch: -20.0 ≤ pitch ≤ 20.0.

Output File: Must be .mp3 or .wav.

---


## Example Commands
Standard Voice Example
```bash
Copy code
go run main.go -text "こんばんは" -voice stand-c -rate 1.0 -pitch 0 -o evening.mp3
Wavenet Voice with custom rate/pitch
```
```bash
Copy code
go run main.go -text "おはようございます" -voice wave-b -rate 1.5 -pitch 3.0 -o morning.mp3
WAV Output Example
```
```bash
Copy code
go run main.go -text "テスト音声" -voice stand-d -rate 1.0 -pitch -2.0 -o test.wav
```
Output:

yaml
Copy code
## ßmp3 created successfully at: morning.mp3
Project Structure
pgsql
Copy code
Go-NihonVoice/
├── cmd/                         # CLI entrypoint
│   └── cmd.go                   # CLI logic and argument parsing
├── speech/                      # Google Text-to-Speech wrapper
│   └── speaker.go                # Handles TTS API interaction
├── main.go                      # Main entrypoint
├── go.mod                        # Go module file
├── go.sum                        # Dependency versions
└── sample-service-account-key.json # Google service account key
Contributing
Fork the repository.

Create a feature branch: git checkout -b feature-name.

Commit your changes: git commit -m "Add feature".

Push to the branch: git push origin feature-name.

Open a Pull Request.

License
This project is licensed under the MIT License.

**Notes**
Requires Go 1.20+.

Ensure your Google Cloud TTS API is enabled and the service account key has proper permissions.

Designed for Japanese language (ja-JP), though TTS API supports multiple languages.

yaml
Copy code
