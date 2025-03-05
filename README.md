# CryptoSong Cracker

## Intro

[CryptoSong](https://github.com/Samuel-BlankAmber/CryptoSong) is an app for encrypting messages using music. It works by SHA-256 hashing the `song title || artist` and using this as the key to an AES-GCM cipher.

There is only a finite number of songs and artists, so in theory this should be trivial to bruteforce. However, Shazam's song database is not public, so this script does not exhaustively crack all ciphertexts.

This script uses the [Million Song Dataset](http://millionsongdataset.com), but it can easily be repurposed to use other wordlists.

## How to use

1. Download the dataset. This only needs the titles and artists, so the summary is sufficient:

`wget http://millionsongdataset.com/sites/default/files/AdditionalFiles/msd_summary_file.h5`

2. Parse this file and construct `keys.txt`:

`./parse.py`

3. Profit.

`go run crack.go -encrypted=pVvRULGUvgfsP6eieusttSI7/yrbPugZ0B6dcjuToOIdspU0ZEh3`

Output:
```
Decrypted: 'hello world' with key: 'BeautifulEminem'
```
