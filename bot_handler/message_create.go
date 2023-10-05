// messageCreate.go
package botHandler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // メッセージが作成されたときに実行する処理
	//u := m.Author

    fmt.Println(m.Content)

    if(m.Author.Bot == false){
        s.ChannelMessageSend(m.ChannelID, m.Content)
    }
}