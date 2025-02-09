# Fam assignment
click [here](http://34.57.235.77/) to view site

### Cron
The cronjob runs the script `main.py` every 3rd minute. The script calls youtube api for 50 videos with `Query` specified in the environment variable, after a publish time (it uses the latest time from db and has a default value of current -1 hour) and inserts them in the database.

> [!NOTE]
> API key is not used if it reaches the maximum quota before the day ends and is reset after 24 hours

### CLI
One can manage keys using the CLI tool, use `--help` to know more.

#### Installation
```
cd cli-tool
./configure --prefix=<dir>
make install
```

There are two commands `list` and `add`, which will list all api keys along with their usage and add keys respectively.

#### How are keys used ?
The keys are stored as a list of object with their number of usage, the script uses the api key with least usage and updates it's value. Once the value reaches 100, the key is no longer used.
The keys are then reset after 24 hours of first use

> Example: ytbutler add key-1 key-2 ....

### API
the endpoint is as follows

`/api/videos?page=${currentPage}&sort=${currSort}&title=${filter.title}&start=${filter.start}&end=${filter.end}&limit=${limit}`

```curl
curl "http://34.57.235.77/api/videos?page=1&limit=1"
```
```json
{
    "currentPage": 1,
    "totalPages": 375,
    "sorted_in": "published_at_desc",
    "filter": {
        "title": "",
        "start": "1970-01-01T00:00:00Z",
        "end": "2025-02-09T14:42:37.65Z"
    },
    "videos": [
    {
        "id": "BIMX86gQmhs",
        "title": "Cats",
        "description": "Sample description",
        "published_at": "2025-02-09T14:38:46Z",
        "thumbnail": "https://i.ytimg.com/vi/BIMX86gQmhs/mqdefault.jpg",
        "url": "https://www.youtube.com/watch?v=BIMX86gQmhs"
    },
  ]
}
```

the available sorts are:
- `published_at_desc`
- `title_asc`

filters are (all are optional):
- `title`: keywords in the title (string) (default is all)
- `start`: the start date to filter (unix timestamp) (default is 01-01-1970)
- `end`  : the end date to filter (unix timestamp) (default is the time of request)

### Starting the server
clone the repository, install the cli tool by following installation and add keys. Then run command `docker compose up -d`, the website will be accessible on port `80`

### Techstack
- `golang`    : for backend and cli
- `next`      : for frontend
- `python`    : to fetch latest videos and update database using cron
- `postgresql`: for database
- `docker`    : to containerize the application

### Scaling the application
Since the main focus of the application is storing metadata of youtube videos, it will get very large after some time. One can do the following to make it scalable:
- Use a cursor based pagination
- Shard the database
- Use orchestrator to manage deployments and scale it
- Use cacheing to cache the data so that the database is not overloaded
