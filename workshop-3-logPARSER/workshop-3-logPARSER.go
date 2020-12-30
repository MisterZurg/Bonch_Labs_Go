package main

import (
	"bufio"
	"fmt"
	"github.com/Knetic/govaluate"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

//User_To_Server(UID, ORDER, ELEMENT)  Examp: User_To_Server(30, 1, "6")
//Server_To_User(UID, RESULT)            Examp: Server_To_User(27, "= 8")
var strNum = 1

func main() {
	userHashmap := make(map[int64]map[int64]string)
	result := make(map[int64]string)
	file, err := os.Open("workshop-3-logPARSER/IKB_2_10.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	myScan := bufio.NewScanner(file)
	for myScan.Scan() {
		line := myScan.Text()
		line = strings.Replace(strings.Replace(line, strconv.Itoa(strNum), "", 1), ": ", "", 1)
		strNum++
		if strings.Contains(line, "User_To_Server") {
			// Replacing garbage
			userInfo := strings.Replace(strings.Replace(line, "User_To_Server(", "", 1), ")", "", 1)
			userInfo = strings.ReplaceAll(userInfo, " ", "")
			userInfo = strings.ReplaceAll(userInfo, "\"", "")
			userInfoArr := strings.Split(userInfo, ",")
			userID, _ := strconv.ParseInt(userInfoArr[0], 10, 10)
			userOpOrd, _ := strconv.ParseInt(userInfoArr[1], 10, 10)

			// Cheking nested key
			data, ok := userHashmap[userID]
			if !ok {
				data = make(map[int64]string)
				data[userOpOrd] = userInfoArr[2]
				userHashmap[userID] = data
			} else {
				userHashmap[userID][userOpOrd] = userInfoArr[2]
			}
		} else {
			// Replacing garbage
			userInfo := strings.Replace(strings.Replace(line, "Server_To_User(", "", 1), ")", "", 1)
			userInfo = strings.ReplaceAll(userInfo, " ", "")
			userInfo = strings.ReplaceAll(userInfo, "\"", "")
			userInfoArr := strings.Split(userInfo, ",")
			userID, _ := strconv.ParseInt(userInfoArr[0], 10, 10)
			result[userID] = userInfoArr[1]
		}
	}
	// Сортировка
	keys := make([]int, 0, len(userHashmap))
	for k := range userHashmap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		resultStringSlice := []string(nil)
		fmt.Println(k, userHashmap[int64(k)])
		for kk := 1; kk < 9; kk++ {
			resultStringSlice = append(resultStringSlice, userHashmap[int64(k)][int64(kk)])
		}
		resultString := strings.Join(resultStringSlice, "")
		// Avoiding -- -+
		resultString = strings.ReplaceAll(resultString, "--", "+")
		resultString = strings.ReplaceAll(resultString, "+-", "-")
		// Avoiding *-
		replaceCounter := strings.Count(resultString, "*-")
		for z := 0; z < replaceCounter; z++ {
			myRune := []rune(resultString)
			tmpRune := []rune("*(-")

			for i := 1; i < utf8.RuneCountInString(resultString); i++ {
				if string(myRune[i]) == "-" && string(myRune[i-1]) == "*" {
					//-9*-5*-5--8
					tmp := myRune[i+1]
					myRune[i+1] = ')' //-9*-)*-5--8
					i++
					tmpRune = append(tmpRune, tmp)
					resultString = strings.Replace(string(myRune), "*-", string(tmpRune), 1)
					break
				}

			}
		}
		//fmt.Println(resultString)
		expression, err := govaluate.NewEvaluableExpression(resultString)
		if err != nil {
			panic(err)
		}
		res, err := expression.Evaluate(nil)
		fmt.Println("U>", resultString)
		fmt.Println("Gotten result>", res)
		fmt.Println("Serv to UID  >", k, result[int64(k)])
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	}
}
