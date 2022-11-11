package main

import (
	"fmt"
	"strconv"
	"time"
)

const PLAYER_1 = "X"
const PLAYER_2 = "O"
const MOVE_COMPLETE_LINE = "MOVE_COMPLETE_LINE"
const MOVE_TRUNCATE_RIVAL_LINE = "MOVE_TRUNCATE_RIVAL_LINE"
const MOVE_OTHER = "MOVE_OTHER"

func main() {
	tictactoe()
	worldWarIII()
}

func tictactoe() {
	board := getNewBoard()
	mockBoard(&board)
	printTictactoe(&board)
	for !isWinner(&board) {
		tally := getTally(&board)
		if playerMove(&board, &tally, PLAYER_1, PLAYER_2) {
			fmt.Println("PLAYER 1 WINS")
			break
		}
		if playerMove(&board, &tally, PLAYER_2, PLAYER_1) {
			fmt.Println("PLAYER 2 WINS")
			break
		}
		printTictactoe(&board)
	}
	printTictactoe(&board)
	fmt.Println("GAME OVER")
}

func getNewBoard() [3][3]string {
	var board [3][3]string
	board[0][0] = "7"
	board[0][1] = "8"
	board[0][2] = "9"
	board[1][0] = "4"
	board[1][1] = "5"
	board[1][2] = "6"
	board[2][0] = "1"
	board[2][1] = "2"
	board[2][2] = "3"

	return board
}

func mockBoard(board *[3][3]string) {
	board[0][0] = "X"
	board[0][1] = "O"
	board[0][2] = "O"
	board[1][0] = "O"
	board[1][1] = "X"
	board[1][2] = "O"
	board[2][0] = "X"
	board[2][1] = "X"
	// board[2][2] = "O"
}

func getTally(board *[3][3]string) map[string]map[string]int {
	tally := map[string]map[string]int{
		"h0": {PLAYER_1: 0, PLAYER_2: 0},
		"h1": {PLAYER_1: 0, PLAYER_2: 0},
		"h2": {PLAYER_1: 0, PLAYER_2: 0},
		"v0": {PLAYER_1: 0, PLAYER_2: 0},
		"v1": {PLAYER_1: 0, PLAYER_2: 0},
		"v2": {PLAYER_1: 0, PLAYER_2: 0},
		"d0": {PLAYER_1: 0, PLAYER_2: 0},
		"d1": {PLAYER_1: 0, PLAYER_2: 0},
	}

	for i := 0; i < 3; i++ {
		if board[0][i] == PLAYER_1 {
			tally["h0"][PLAYER_1]++
		} else if board[0][i] == PLAYER_2 {
			tally["h0"][PLAYER_2]++
		}
		if board[1][i] == PLAYER_1 {
			tally["h1"][PLAYER_1]++
		} else if board[1][i] == PLAYER_2 {
			tally["h1"][PLAYER_2]++
		}
		if board[2][i] == PLAYER_1 {
			tally["h2"][PLAYER_1]++
		} else if board[2][i] == PLAYER_2 {
			tally["h2"][PLAYER_2]++
		}
		if board[i][0] == PLAYER_1 {
			tally["v0"][PLAYER_1]++
		} else if board[i][0] == PLAYER_2 {
			tally["v0"][PLAYER_2]++
		}
		if board[i][1] == PLAYER_1 {
			tally["v1"][PLAYER_1]++
		} else if board[i][1] == PLAYER_2 {
			tally["v1"][PLAYER_2]++
		}
		if board[i][2] == PLAYER_1 {
			tally["v2"][PLAYER_1]++
		} else if board[i][2] == PLAYER_2 {
			tally["v2"][PLAYER_2]++
		}
		if board[i][i] == PLAYER_1 {
			tally["d0"][PLAYER_1]++
		} else if board[i][i] == PLAYER_2 {
			tally["d0"][PLAYER_2]++
		}
		if board[2-i][i] == PLAYER_1 {
			tally["d1"][PLAYER_1]++
		} else if board[2-i][i] == PLAYER_2 {
			tally["d1"][PLAYER_2]++
		}
	}

	return tally
}

func playerMove(board *[3][3]string, tally *map[string]map[string]int, player string, rival string) bool {
	move := chooseMoveBasedOnTally(tally, player, rival)
	for m, row := range move {
		switch m {
		case MOVE_COMPLETE_LINE:
			completeLine(board, row, player)
			return true
		case MOVE_TRUNCATE_RIVAL_LINE:
			truncateRivalLine(board, row, player)
			return false
		case MOVE_OTHER:
			randomMove(board, player)
			return false
		default:
			randomMove(board, player)
			return false
		}
	}

	return false
}

func chooseMoveBasedOnTally(tally *map[string]map[string]int, player string, rival string) map[string]string {
	rows := [8]string{"h0", "h1", "h2", "v0", "v1", "v2", "d0", "d1"}
	move := make(map[string]string)

	for _, row := range rows {
		if (*tally)[row][player] == 2 && (*tally)[row][rival] == 0 {
			move[MOVE_COMPLETE_LINE] = row
			return move
		}
		if (*tally)[row][player] == 0 && (*tally)[row][rival] == 2 {
			move[MOVE_TRUNCATE_RIVAL_LINE] = row
		}
	}
	if len(move) == 0 {
		move[MOVE_OTHER] = ""
	}

	return move
}

func completeLine(board *[3][3]string, row string, player string) {
	direction := row[0]
	no, _ := strconv.Atoi(string(row[1]))
	for i := 0; i < 3; i++ {
		switch direction {
		case 'h':
			if board[no][i] != PLAYER_1 && board[no][i] != PLAYER_2 {
				board[no][i] = player
			}
		case 'v':
			if board[i][no] != PLAYER_1 && board[i][no] != PLAYER_2 {
				board[i][no] = player
			}
		case 'd':
			var j int
			if no == 0 {
				j = 0
			} else {
				j = 3 - i
			}
			if board[j][j] != PLAYER_1 && board[j][j] != PLAYER_2 {
				board[j][j] = player
			}
		}
	}
}

func truncateRivalLine(board *[3][3]string, player string, rival string) {

}

func randomMove(board *[3][3]string, player string) {

}

func isWinner(board *[3][3]string) bool {
	if board[0][0] == board[0][1] && board[0][1] == board[0][2] {
		return true
	}
	if board[1][0] == board[1][1] && board[1][1] == board[0][2] {
		return true
	}
	if board[2][0] == board[2][1] && board[2][1] == board[0][2] {
		return true
	}
	if board[0][0] == board[1][0] && board[1][0] == board[2][0] {
		return true
	}
	if board[0][1] == board[1][1] && board[1][1] == board[2][1] {
		return true
	}
	if board[0][2] == board[1][2] && board[1][2] == board[2][2] {
		return true
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}

	return false
}

func printTictactoe(board *[3][3]string) {
	fmt.Println(board[0][0], "|", board[0][1], "|", board[0][2])
	fmt.Println(board[1][0], "|", board[1][1], "|", board[1][2])
	fmt.Println(board[2][0], "|", board[2][1], "|", board[2][2])
	fmt.Println()
}

func worldWarIII() {
	fmt.Println("INITIALISING WORLD WAR III")
	fmt.Println("60 SECONDS TO MISSILE LAUNCH")
	time.Sleep(time.Second * 10)
	fmt.Println("50 SECONDS TO MISSILE LAUNCH")
	time.Sleep(time.Second * 10)
	fmt.Println("40 SECONDS TO MISSILE LAUNCH")
	time.Sleep(time.Second * 10)
	fmt.Println("30 SECONDS TO MISSILE LAUNCH")
	time.Sleep(time.Second * 10)
	fmt.Println("20 SECONDS TO MISSILE LAUNCH")
	time.Sleep(time.Second * 10)
	fmt.Println("10 SECONDS TO MISSILE LAUNCH")
	time.Sleep(time.Second * 10)
	fmt.Println("LAUNCH IS A GO")
	time.Sleep(time.Second * 1)
	fmt.Println("THE WORLD IS DOOMED")
	time.Sleep(time.Second * 1)
	fmt.Println("GAME OVER")
}
