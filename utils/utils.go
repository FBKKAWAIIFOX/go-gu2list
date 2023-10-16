package utils

import (
	"context"
	"fmt"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/shopspring/decimal"
	"gu2list/commands"
	"gu2list/utils/VCoin"
	"log"
	"math"
	"time"
)

func DiscordClient(Token string) /**state.State*/ {
	s := state.New("Bot " + Token)
	r := cmdroute.NewRouter()
	r.Use(cmdroute.Deferrable(s, cmdroute.DeferOpts{
		Timeout: 60 * time.Second,
	}))

	r.AddFunc("vprice", func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		Price, Update := VCoin.Price()
		if Price != "" && Update != "" {
			purchasePrice, _ := decimal.NewFromString(Price)
			CalculatedPrice, _ := purchasePrice.Float64()
			return &api.InteractionResponseData{
				Embeds: &[]discord.Embed{
					{
						Title:     "廢土村民錠價格 <YourEmojiHere>",
						Timestamp: discord.NowTimestamp(),
						Color:     5763719,
						Fields: []discord.EmbedField{
							{
								Name:   "價格",
								Value:  fmt.Sprintf("<YourEmojiHere> %v  (手續費後: %v)", Price, math.Round(CalculatedPrice*1.015)),
								Inline: false,
							},
							{
								Name:   "更新時間",
								Value:  Update,
								Inline: false,
							},
						},
					},
				},
			}
		}
		return nil
	})
	s.AddInteractionHandler(r)
	s.AddIntents(gateway.IntentGuilds | gateway.IntentGuildMembers | gateway.IntentDirectMessages | gateway.IntentGuildMessages)

	s.AddHandler(func(c *gateway.ReadyEvent) {
		fmt.Println(c.User.Tag())
	})

	if err := overwriteCommands(s); err != nil {
		fmt.Println(err)
	}

	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("Failed to Open:", err)
	}

	select {}
}

func overwriteCommands(s *state.State) error {
	return cmdroute.OverwriteCommands(s, commands.Commands)
}
