# 30-04-2026

* I changed everything in the code to postgres (check issue with PR's on)
* Migrated the database data with sqlite3 dump file, which i cleansed for DDL statements so that only insert statements where left. Then i piped the data directly into the database in the docker container.

## Crontab & backups

* Im not sure that the our crontab or backup.sh is shown in the docs but here they are.
* We have been using RClone for backing up the sqllite database and moving it to a shared google drive
* The migration to postgres has made this redundant, which reminded me that this might not be in the documentation

``` Crontab
0 0 * * * /home/appuser/backups/backup_db.sh >> /home/appuser/backups/backup.log 2>&1 
```

``` .sh
#!/bin/bash

DB_NAME="/home/appuser/_data/whoknows.db"
BACKUP_DIR="/home/appuser/backups"
REMOTE_GDRIVE="gdrive:"

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="$BACKUP_DIR/whoknows_$TIMESTAMP.db"
mkdir -p "$BACKUP_DIR"

# Lav SQLite backup
sqlite3 "$DB_NAME" ".backup '$BACKUP_FILE'"

# Upload til Google Drive
rclone copy "$BACKUP_FILE" "$REMOTE_GDRIVE"

# Slet gamle backups ældre end 7 dage
find "$BACKUP_DIR" -type f -name "*.db" -mtime +7 -delete

```
