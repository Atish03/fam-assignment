# Fam assignment
### Cron
The cronjob runs the script `main.py` every 2nd minute. The script calls youtube api for 50 videos with `Query` after a publish time (it uses the latest time from db and has a default value of current -1 hour) specified in the environment variable and inserts them in the database.

> [!NOTE]
> API key is delete if it reaches the maximum limit before the day ends

### CLI
One can manage keys using the CLI tool, use `--help` to know more.

#### Installation
```
cd cli-tool
./configure --prefix=<dir>
make install
```

There are two commands `list` and `add`, which will list all api keys along with their usage and one can add keys.

> Example: ytbutler add key-1 key-2 ....

### API
the endpoint is as follows

`/api/videos?page=${currentPage}&sort=${currSort}&title=${filter.title}&start=${filter.start}&end=${filter.end}`

the available sorts are:
- `published_at_desc`
- `title_asc`

filters are:
- `title`: keywords in the title (default is all)
- `start`: the start date to filter (default is 01-01-1970 by default)
- `end`: the end date to filter (default is the time time of request)

### Starting the server
clone the repository and run command `docker compose up -d` and the website is accessible on port `8080`

