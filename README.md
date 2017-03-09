# LANShare
Sharing folders over LAN.

## Features
- Explore
- Download Files and Folders (as zip)
- Upload Files
- Simple WEB client
- Full API (List, Download, Upload)
- IPC

## How It Works
The first instance run http server and tcp server (for IPC).

All others instances will just add the folder to the main instance.

The added folder is the first parameter or the current directory.
