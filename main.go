package main

import (
	"fmt"
	"time"
)

const PLAYER_1 = "X"
const PLAYER_2 = "O"

func main() {
	tictactoe()
	worldWarIII()
}

func tictactoe() {
	var board [3][3]string
	mockBoard(&board)
	printTictactoe(board)
	for !isWinner(board) {
		move(&board, PLAYER_1, PLAYER_2)
		move(&board, PLAYER_2, PLAYER_1)
		printTictactoe(board)
	}
}

func mockBoard(board *[3][3]string) {
	// board[0][0] = "X"
	board[0][1] = "O"
	board[0][2] = "O"
	board[1][0] = "O"
	board[1][1] = "X"
	board[1][2] = "O"
	board[2][0] = "X"
	board[2][1] = "X"
	board[2][2] = "O"
}

func move(board *[3][3]string, player string, rival string) {
	if completeLine(board, player, rival) {
		fmt.Println("PLAYER", player, "WINS")
		return
	}
	/* if truncateRivalLine(board, player) {
		return
	}
	if expandLine(board, player) {
		return
	}
	if startLine(board, player) {
		return
	} */
}

func completeLine(board *[3][3]string, player string, rival string) bool {
	h1Player := 0
	h1Rival := 0
	/*h2 := 0
	h3 := 0
	v1 := 0
	v2 := 0
	v3 := 0
	d1 := 0
	d2 := 0 */
	for i := 0; i < 3; i++ {
		if board[0][i] == player {
			h1Player++
		} else if board[0][i] == rival {
			h1Rival++
		}
	}
	if h1Player == 2 && h1Rival == 0 {
		for i := 0; i < 3; i++ {
			if board[0][i] != player && board[0][i] != rival {
				board[0][i] = player
				return true
			}
		}
	}
	return false
}

func isWinner(board [3][3]string) bool {
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

func printTictactoe(board [3][3]string) { // TODO by ref
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
