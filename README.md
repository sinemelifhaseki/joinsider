# joinsider

* Used sqlite3 as database, storage.db was used as the schema of the project consisting of all tables. 
* Due to conflicts while connecting front-end to back-end, I tried to see the results of my handlers and model methods in Postman requests as JSON responses.
* Used sql queries explicitly to obtain data for test since I could not use front-end for CRUD operations.
* Tried to implement TeamDetails.vue as well, but for now it only shows the win, loss, draw, etc. stats for a team.

* To test code I added following leagues and teams: 
```sql
INSERT INTO leagues(name, current_week, total_week, created_at, updated_at, actions) VALUES("Bundesliga", 1, 6, "2020-11-20", "2020-11-20", "Created");
INSERT INTO leagues(name, current_week, total_week, created_at, updated_at, actions) VALUES("La Liga", 1, 6, "2020-11-20", "2020-11-20", "Created");
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(1, "Bayern Munchen","Bundesliga", "2020-11-20", "Created")
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(1, "Schalke","Bundesliga", "2020-11-20", "Created")
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(1, "RB Leipzig","Bundesliga", "2020-11-20", "Created")
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(1, "Borussia Dortmund","Bundesliga", "2020-11-20", "Created")

INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(2, "Atletico Madrid","La Liga", "2020-11-20", "Created")
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(2, "Real Madrid","La Liga", "2020-11-20", "Created")
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(2, "Valencia","La Liga", "2020-11-20", "Created")
INSERT INTO teams(league_id, name, league_nme, updated, actions) VALUES(2, "Barcelona","La Liga", "2020-11-20", "Created")

```
For creating, updating, deleting queries I did not write SQL queries explicitly, they can be found within related methods to be executed upon calls.
