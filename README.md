### gator - a blog aggregator
#### prerequisites 
To be able to install and use the program, go [golang.com] and postgres sql [postgres.com] have to be installed. The respective websites have detailed information on the installation process.

#### installation
After downloading the repo go install . From the root of the repo on the local machine or go install {path/to/gator} will install the program and make it available across the terminal.

#### setup
To setup the program simply start a new database in postgres and paste the acces into a file at the root of the os named "~/.gatorconfig.json" which has to be created manually. The format should be like this:
```
```
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```
```
```

#### use
Calling blogAggregator from the CLI will activate the program. The following commands can be paired with blogAggregator:
- register {name}; registers a user to the database and logs that user in 
- login {name}; changes the current user of the database to name
- users; lists all users of the database and shows the current login
- addfeed {name}{url}; adds a feed to the database and creates a follow for the current user
- feeds; shows all feeds the current user follows
- follow {url}; connects current user to an existing feed by url
- following; lists all feeds the current user is following
- unfollow {url}; lets current user unfollow the given feed
- agg {duration}; scrapes the web for all existing feeds and stores the data in duration interval to the database
- browse {number}; shows the last posts for the feeds the current user follows.


