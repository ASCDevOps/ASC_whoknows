# 20-03-2026

* We have have discussed and came to the conclusion, that every group member reviewing PR's is ineffective and takes a massive amount of.
* We now have a CI file that can check for linting errors.
* "Own your code" is having production ready code when pushing into master, that is everyones responsibility.

# 20-03-2026
* Compose is now up and running with our continous delivery pipeline. Deployment will be incorporated soon.
* I have created 2 seperate setups for docker compose. Dev utilizes dockerfile, while prod uses ghcr.
* We are currently loading .env file by SCP'ing it onto the server, compose file is also scp'ed.
* Instead of SCP'ing, it should be automated by continous deployment whill will come.

# 23-03-2026
* I have created a script on the vm, that backups the database with ".backup" moves it into backup folder and moves it to gdrive with rclone.
* I have setup crontab to run each day at midnight (00:00).
* The gdrive is a completely new google account that each of us have access to (discord)
