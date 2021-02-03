package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Game holds the state of the game, current position of pieces, total moves, and last played move
type Game struct {
	board         [][]int
	numberOfMoves int
	lastMove      Move
}

//Move represents coordinates for one move played - x for horizontal, y for vertical
type Move struct {
	x int
	y int
}

const boardHeight = 6
const boardWidth = 7
const playerCount = 2
const winConditionCount = 4
const maximumMoves = boardHeight * boardWidth

var (
	errNonExistingColumn = errors.New("ileggal move, column does not existis on board")
	errColumnFilled      = errors.New("ileggal move, column is filled already")
)

func main() {
	//save possible directions for checking a connect4 - up, down, left, right, and diagonals
	up := Move{x: 1, y: 0}
	down := Move{x: -1, y: 0}
	right := Move{x: 0, y: 1}
	left := Move{x: 0, y: -1}
	ur := Move{x: 1, y: 1}
	dl := Move{x: -1, y: -1}
	dr := Move{x: 1, y: -1}
	ul := Move{x: -1, y: 1}

	//organize the directions in relevant format, up+down, left+right, and correct diagonals
	var directions = [][]Move{
		{up, down},
		{right, left},
		{ur, dl},
		{dr, ul},
	}
	g := &Game{}
	g.numberOfMoves = 0
	g.board = make([][]int, boardWidth)
	for i := 0; i < boardWidth; i++ {
		g.board[i] = make([]int, boardHeight)
		for j := 0; j < boardHeight; j++ {
			g.board[i][j] = 0
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	minimumMoves := playerCount*winConditionCount - 1
	for scanner.Scan() {
		moves := parseInputToMoves(scanner.Text())
		for _, move := range moves {
			if err := g.playNextMove(move); err != nil {
				fmt.Println("Can't make a move ", move, err)
				continue
			}
			g.printBoard()
			if g.numberOfMoves > minimumMoves {
				win := g.checkForWin(directions)
				if win {
					fmt.Println("WINNER: Player ", g.getPieceAtPosition(g.lastMove))
					g.printBoard()
					return
				}

				if g.numberOfMoves == maximumMoves {
					fmt.Println("DRAW")
					g.printBoard()
					return
				}
			}

		}

	}
}

func (g *Game) checkForWin(directions [][]Move) bool {
	player := g.getPieceAtPosition(g.lastMove)
	for i := 0; i < 4; i++ {
		var inARow = 1
		for j := 0; j < 2; j++ {
			checkingCell := Move{g.lastMove.x, g.lastMove.y}
			for player == g.getPieceAtPosition(checkingCell) {
				checkingCell.x += directions[i][j].x
				checkingCell.y += directions[i][j].y
				if checkingCell.x >= boardWidth || checkingCell.x < 0 || checkingCell.y >= boardHeight || checkingCell.y < 0 {
					break
				}
				if g.getPieceAtPosition(checkingCell) == player {
					inARow++
				}
				if inARow >= winConditionCount {
					return true
				}
			}
		}
	}
	return false
}

func (g *Game) getPieceAtPosition(move Move) int {
	return g.board[move.x][move.y]
}

func parseInputToMoves(input string) []int {
	var retVal []int
	moves := strings.Split(input, " ")
	for _, move := range moves {
		number, err := strconv.Atoi(move)
		if err != nil {
			fmt.Print("Not a possible move: ", move)
		}
		retVal = append(retVal, number)
	}
	return retVal
}

func (g *Game) playNextMove(column int) error {
	if column >= boardWidth {
		return errNonExistingColumn
	}

	verticalPos := 0
	for ; verticalPos < boardHeight; verticalPos++ {
		if g.board[column][verticalPos] == 0 {
			break
		}
	}
	if verticalPos >= boardHeight {
		return errColumnFilled
	}

	currentPlayer := g.numberOfMoves%playerCount + 1
	g.board[column][verticalPos] = currentPlayer
	g.numberOfMoves++
	g.lastMove = Move{x: column, y: verticalPos}

	return nil
}

func (g *Game) printBoard() {
	for i := boardHeight - 1; i >= 0; i-- {
		for j := 0; j < boardWidth; j++ {
			fmt.Print(g.board[j][i])
		}
		fmt.Print("\n")
	}
}
