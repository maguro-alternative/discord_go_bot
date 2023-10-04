package commands

import (
	"fmt"
	"time"

	"github.com/maguro-alternative/discord_go_bot/bot_handler"

	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
)

func RecordCommand() *handlers.Command {
	/*
		start_recordコマンドの定義

		コマンド名: start_record
		説明: 録音を開始します
		オプション: なし
	*/
	return &handlers.Command{
		Name:        "test_start_record",
		Description: "録音を開始します",
		Options:     []*discordgo.ApplicationCommandOption{},
		Executor:    recordVoice,
	}
}

func createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version: 2,
			// Taken from Discord voice docs
			PayloadType:    0x78,
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}

func handleVoice(c chan *discordgo.Packet) {
	files := make(map[uint32]media.Writer)
	for p := range c {
		file, ok := files[p.SSRC]
		if !ok {
			var err error
			file, err = oggwriter.New(fmt.Sprintf("%d.ogg", p.SSRC), 48000, 2)
			if err != nil {
				fmt.Printf("failed to create file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
				return
			}
			files[p.SSRC] = file
		}
		// Construct pion RTP packet from DiscordGo's type.
		rtp := createPionRTPPacket(p)
		err := file.WriteRTP(rtp)
		if err != nil {
			fmt.Printf("failed to write to file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
		}
	}

	// Once we made it here, we're done listening for packets. Close all files
	for _, f := range files {
		f.Close()
	}
}

func recordVoice(s *discordgo.Session, i *discordgo.InteractionCreate) {
	/*
	録音の開始

	コマンドの実行結果を返す
	*/
	if i.Interaction.ApplicationCommandData().Name == "test_start_record" {
		vs, err := s.State.VoiceState(i.GuildID, i.Interaction.Member.User.ID)
		if err != nil {
			fmt.Println("failed to find voice state:", err)
			responsText(s, i, "ボイスチャンネルに接続していません")
			return
		}
		if vs == nil {
			responsText(s, i, "ボイスチャンネルに接続していません")
			return
		}
		responsText(s, i, "録音を開始します <#" + vs.ChannelID + ">")
		v, err := s.ChannelVoiceJoin(i.GuildID, vs.ChannelID, true, false)
		//fmt.Println(v)
		if err != nil {
			fmt.Println("failed to join voice channel:", err)
			responsText(s, i, "ボイスチャンネルに入ってください")
			return
		}
		go func() {
			time.Sleep(10 * time.Second)
			close(v.OpusRecv)
			v.Close()
		}()
		handleVoice(v.OpusRecv)
	}
}

func responsText(s *discordgo.Session, i *discordgo.InteractionCreate, contentText string) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentText,
		},
	})
	if err != nil {
		fmt.Printf("error responding to record command: %v\n", err)
		return err
	}
	return nil
}