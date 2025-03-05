#!/usr/bin/env python3
import h5py
from tqdm import tqdm

FILE_NAME = "msd_summary_file.h5"

if __name__ == "__main__":
    with h5py.File(FILE_NAME, 'r') as f:
        artist_names = f["metadata"]["songs"]["artist_name"]
        song_titles = f["metadata"]["songs"]["title"]
        with open("keys.txt", "w", encoding="utf-8") as keys_file:
            for artist_name, song_title in tqdm(zip(artist_names, song_titles)):
                key = song_title.decode('utf-8') + artist_name.decode('utf-8')
                keys_file.write(key + "\n")