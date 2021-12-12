# Movie-Gateway

# Structure

```
├── domain
│ └── movie
│     ├── client
│     └── handler
├── proto
│ └── movie
├── route
│ ├── gate
│ └── middleware
└── util

```

movie-service
https://github.com/luqman-v1/movie-service

 ## API



| Description | Method | Endpoint | Authorization | Parameters |
| :---: | :---: | :---: | :---: | :---: |
| List Movie | GET | /api/v1/movies | app-key | searchword (string) , pagination (string) |
| Detail Movie | GET | /api/v1/movies/:imdb_id | app-key |

