package convert

import (
	"inc/lib"
	"inc/lib/api"
	"inc/lib/typings"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(s|sticker|stiker)",
		As:       []string{"sticker"},
		Tags:     "convert",
		IsPrefix: true,
		IsMedia:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			data, _ := client.WA.Download(m.Media)

			s := api.StickerApi(&typings.Sticker{
				File: data,
				Tipe: func() typings.MediaType {
					if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
						return typings.IMAGE
					} else {
						return typings.VIDEO
					}
				}(),
			}, &typings.MetadataSticker{
				Author:    m.PushName,
				Pack:      "https://s.id/ryuubot",
				KeepScale: true,
				Removebg:  "false",
				Circle: func() bool {
					if m.Querry == "-c" {
						return true
					} else {
						return false
					}
				}(),
			})

			client.SendSticker(m.From, s.Build(), m.ID)

		},
	})
}
