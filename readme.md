# TicTacToe API

## Take home task

#### SUN Golang Developer Take Home Assessment

##### Description

Your task is to create a Golang web server that allows players to play a single game of TicTacToe using a pre-existing TicTacToe Go module that already contains all necessary game logic. The server should expose endpoints to allow players to make moves, retrieve the current game state, and reset the game. The webserver should support the following REST endpoints:

`GET /game`

This endpoint should return JSON including:
- The current player’s turn.
- If there is a winner then the winner of the game. 
- A representation of the game board.

`POST /game/move`

This endpoint requires the sending of JSON that contains the following:
- The player making the move.
- The row to place the player’s symbol.
- The column to place the player’s symbol.

On successful move this endpoint should return the updated state of the game in the same JSON format as what is returned by GET /game.

On failed move this endpoint should return JSON containing a descriptive error message of what went wrong.

`DELETE /game`

This endpoint should delete the existing game and reset the game i.e. a new game with an empty board would be returned on the next GET /game call.

##### Requirements

- Use The Golang programming language.

- Utilize the existing TicTacToe Go module found here: https://github.com/chrisfregly/tictactoe
NOTE: you should NOT be writing any TicTacToe game logic at any point in this task but instead utilize the provided Go module above.

- You may use external libraries when writing the web server but it is recommended that you use Golang standard libraries and/or well known + maintained libraries whenever possible. 

- Include instructions on how to run, play, and test your server in a README.md file.

##### Submission

Please create and send us a GitHub public repo containing your source code and README.md file.

If you have any questions or are uncertain about specific requirements please do NOT email asking for clarification but instead use your best judgment and include any assumptions you made in your README.md file.

##### Evaluation

Your code will be evaluated on the following factors:

- Code runs and works to specifications.

- Code is easily readable and well structured.

- REST API standards are followed.

- Documentation is clear with descriptions on how to run and use.

- There is appropriate test coverage.

## Instructions

### Start the Server API

### API Routes

`GET /game`

201 Creates a new game and returns it's token and initial state
```
{ 
    "token": <Game ID Token>,
    "next-player": <Player with the current MOVE>,
    "winner": <if present, shows the winner of the game>,
    "board": <current representation of the game board>
}
```

`GET /game/<id>`

**Responses:**

200 Current state of the game
```
{ 
    "token": <Game ID Token>,
    "next-player": <Player with the current MOVE>,
    "winner": <if present, shows the winner of the game>,
    "board": <current representation of the game board>
}
```

404 Game not found

`POST /game/<id>/move`

**Request:**
```
{
    "player": <player to make the move. Accepts 'X' or 'O'>,
    "row": <row to place the move. Must be within board bounds>,
    "column": <column to place the move. Must be within board bounds>
}
```

**Responses:**

200 Current state of the game
```
{ 
    "token": <Game ID Token>,
    "next-player": <Player with the current MOVE>,
    "winner": <if present, shows the winner of the game>,
    "board": <current representation of the game board>,
}
```

404 Game not found

401 Bad request
- Game is already finished
- Out of bounds for row/column
- Invalid player
- Location is not empty
- Invalid player's turn


`DELETE /game/<id>`

**Responses:**

200 Game deleted

404 Game not found


### Examples
Example of a game with a winner:

```
curl curl --request GET \
    --url http://localhost:3000/game

Response:
{
	"token": "269f9469-7c43-4680-8497-bdc9a94823c0",
	"next-player": "X",
	"winner": "",
	"board": [
		["-","-","-"],
		["-","-","-"],
		["-","-","-"]
    ]
}
```

Check game state
```
curl curl --request GET \
    --url http://localhost:3000/game/269f9469-7c43-4680-8497-bdc9a94823c0

Response:
{
	"token": "269f9469-7c43-4680-8497-bdc9a94823c0",
	"next-player": "X",
	"winner": "",
	"board": [
		["-","-","-"],
		["-","-","-"],
		["-","-","-"]
    ]
}
```

Make a move X
```
curl --request POST \
  --url http://localhost:3000/game/269f9469-7c43-4680-8497-bdc9a94823c0/move \
  --header 'Content-Type: application/json' \
  --data '{
	"player": "X",
	"row": 1,
	"column": 1
}'

Response:
{
	"token": "269f9469-7c43-4680-8497-bdc9a94823c0",
	"next-player": "O",
	"winner": "",
	"board": [
		["-","-","-"],
		["-","X","-"],
		["-","-","-"]
    ]
}
```
Make a move O
```
curl --request POST \
  --url http://localhost:3000/game/269f9469-7c43-4680-8497-bdc9a94823c0/move \
  --header 'Content-Type: application/json' \
  --data '{
	"player": "O",
	"row": 0,
	"column": 1
}'

Response:
{
	"token": "269f9469-7c43-4680-8497-bdc9a94823c0",
	"next-player": "X",
	"winner": "",
	"board": [
		["-","O","-"],
		["-","X","-"],
		["-","-","-"]
    ]
}
```

## Tests

The tests cover API Calls and the TicTacToe Wrapper.

```make tests```

or coverage:

```make coverage```

  
## Considerations

- No persistent storage was used. The server is stateful. It stores each game in memory. Restarting the server clear all games.
- Included a token in the api to access different games
- There is a wrapper for the TicTacToe lib. This way the system can replace the lib if required.
- No Authentication
- No throtling control
- Simple Concurrency control with Mutex Lock/Unlock and RLock/RUnlock

## Design

- No persistent storage
- Games State in memory
- Wrapper for TicTacToe lib
- Each request handler defines it's own structs to parse the json request
- There is a middleware as an example to set the request ID in the response for tracing
- Unit tests for the wrapper
- API Tests
- Using Golang CHI router. 
  - Lightweight
  - fast
  - Keeps the HTTP Handler request signature
- Code is verbose. It is a demo. Usually there are helpers to abstract repetitive code. Example: HTTP Responses/Logging
- Both access log and application logs are pushed to STDOUT. In prod, they should be split into different outputs.
- TicTacToe is a 3x3 board game. No code to change the board size.
  
## Improvements

 - Rewrite middleware to use the same log across the whole application.
 - Improve loggin/tracing.
 - Improve how to include the RequestID. For example, use Zerolog and include the RequestID in the middleware.
 - Create a cli to execute maintenance tasks.
 - 