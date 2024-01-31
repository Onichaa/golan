package random

import (
  "inc/lib" 
  "math/rand"
  "time"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "dare",
    As:       []string{"dare"},
    Tags:     "random",
    IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      data1 := []string{
		"Kirim pesan ke mantan kamu dan bilang 'aku masih suka sama kamu'",
		"Telfon crush/pacar sekarang dan ss ke pemain",
		"Pap ke salah satu anggota grup (bisa ke admin jga gpp)",
		`Bilang "KAMU CANTIK BANGET NGGAK BOHONG" ke cowo`,
		"Ss recent call whatsapp",
		"Drop emot 🤥 setiap ngetik di gc/pc selama 1 hari",
		"Kirim voice note bilang can i call u baby?",
		`Drop kutipan lagu/quote, terus tag member yang cocok buat kutipan itu`,
		"Pake foto aib teman mu di pp wa sampe 3 hari",
		"Ketik pake bahasa daerah 24 jam",
		`Ganti bio wa menjadi "OnlyforTakim" selama 5 jam`,
		"Chat ke kontak wa urutan sesuai %batre kamu, terus bilang ke dia 'i lucky to hv you'",
		"Prank chat mantan dan bilang 'i love u, pgn balikan'",
		"Record voice di grup ini bilang lu semua kayak babi",
		`Bilang "i hv crush on you, mau jadi pacarku gak?" ke lawan jenis yang terakhir bgt kamu chat (serah di wa/tele), tunggu dia bales, kalo udah ss drop ke sini`,
		"Sebutkan tipe pacar mu!",
		"Snap/post foto pacar/crush",
		"Teriak gajelas lalu kirim pake vn kesini",
		"Pap mukamu lalu kirim ke salah satu temanmu",
		"Kirim fotomu dengan caption, aku anak nya bobi",
		"Teriak pake kata kasar sambil vn trus kirim kesini",
		`Teriak "anjimm gabutt anjimmm " di depan rumah mu`,
		`Ganti nama jadi " aku anak tolol " selama 24 jam`,
		"Pura pura kerasukan, contoh : kerasukan maung, kerasukan belalang, kerasukan kulkas, dll",
	}
      rand.Seed(time.Now().UnixNano())
      index := rand.Intn(len(data1))
      hasil := data1[index]
      m.Reply(hasil)
    },
  })
}
