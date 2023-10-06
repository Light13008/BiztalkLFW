# BiztalkLFW
#### Biztalk Network and File Paths Watcher Service
Tracks movement of the file(s) based on the drive paths and displays updates via API

### Development State: WIP
## Components
**FileWatcher:**

1. Paths: A Storage Driver that uses SQLite to store FILE URIs and will be used in FileTracker for obtaining and tracking the URIs.
2. FileTracker: A Service that checks the file paths mentioned in the paths component.
