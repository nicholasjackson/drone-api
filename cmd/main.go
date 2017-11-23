package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/nats-io/nats"
	messages "github.com/nicholasjackson/drone-messages"
)

var intro = `                     __   ______   __                __                                         
                    /  | /      \ /  |              /  |                                        
  _______   ______  $$ |/$$$$$$  |$$/           ____$$ |  ______    ______   _______    ______  
 /       | /      \ $$ |$$ |_ $$/ /  | ______  /    $$ | /      \  /      \ /       \  /      \ 
/$$$$$$$/ /$$$$$$  |$$ |$$   |    $$ |/      |/$$$$$$$ |/$$$$$$  |/$$$$$$  |$$$$$$$  |/$$$$$$  |
$$      \ $$    $$ |$$ |$$$$/     $$ |$$$$$$/ $$ |  $$ |$$ |  $$/ $$ |  $$ |$$ |  $$ |$$    $$ |
 $$$$$$  |$$$$$$$$/ $$ |$$ |      $$ |        $$ \__$$ |$$ |      $$ \__$$ |$$ |  $$ |$$$$$$$$/ 
/     $$/ $$       |$$ |$$ |      $$ |        $$    $$ |$$ |      $$    $$/ $$ |  $$ |$$       |
$$$$$$$/   $$$$$$$/ $$/ $$/       $$/          $$$$$$$/ $$/        $$$$$$/  $$/   $$/  $$$$$$$/


`

var nc *nats.Conn

var yellow = color.New(color.FgYellow)
var green = color.New(color.FgGreen)
var red = color.New(color.FgRed)

func main() {

	yellow.Println(intro)

	var err error
	nc, err = nats.Connect("nats://192.168.1.113:4222")
	if err != nil {
		log.Fatal("Unable to connect to nats server")
	}

	fmt.Println("Selfi-Drone command line\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()

		if strings.Index(txt, messages.CommandConnect) == 0 {
			handleCommand(txt, 1)
			continue
		}

		if strings.Index(txt, messages.CommandUp) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandDown) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandBackward) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandForward) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandLeft) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandRight) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandClockwise) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandCounterClockwise) == 0 {
			handleCommand(txt, 2)
			continue
		}

		if strings.Index(txt, messages.CommandTakeOff) == 0 {
			handleCommand(txt, 1)
			continue
		}

		if strings.Index(txt, messages.CommandLand) == 0 {
			handleCommand(txt, 1)
			continue
		}

		red.Println("## unknown command!\n")
	}
}

func handleCommand(txt string, params int) {
	val := strings.Split(txt, " ")

	optional := ""
	optionalValue := ""
	if params == 2 {
		optional = " [int]"
		optionalValue = val[1]
	}

	if len(val) < params {
		red.Printf("## command should be \"%s\"!\n\n", val[0], optional)
		return
	}

	green.Printf("## sending message: %s %s...\n\n", val[0], optionalValue)
	sendMessage(val[0], optionalValue)
}

func sendMessage(command string, value string) {
	m := messages.Flight{
		Command: command,
	}

	if value != "" {
		val, _ := strconv.ParseInt(value, 10, 64)
		m.Value = int(val)
	}

	nc.Publish(messages.MessageFlight, m.EncodeMessage())
}
