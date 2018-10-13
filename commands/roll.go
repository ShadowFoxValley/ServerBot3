package commands

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func (data Commands) roll() {
	var elements = strings.Split(data.message.Content, " ")
	var first, second, result, elemLen = 0, 100, 0, len(elements)

	log.Print(elements)

	if elemLen == 3 {
		second, _ = strconv.Atoi(elements[2])
		log.Print(elements)
	} else if elemLen == 4 {
		first, _ = strconv.Atoi(elements[2])
		second, _ = strconv.Atoi(elements[3])
	}

	if first > second {
		first, second = second, first
	}

	if first == second {
		result = first
	} else {
		result = rand.Intn(second-first) + first
	}

	data.mainSession.ChannelMessageSend(data.channelId, "Result: "+strconv.Itoa(result))
}
