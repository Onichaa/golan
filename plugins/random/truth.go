package random

import (
  "inc/lib" 
  "math/rand"
  "time"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "truth",
    As:       []string{"truth"},
    Tags:     "random",
    IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {
      
      data1 := []string{
        "Pernah suka sama siapa aja? berapa lama?",
        "Kalau boleh atau kalau mau, di GC/luar GC siapa yang akan kamu jadikan sahabat?(boleh beda/sama jenis)",
        "Apa ketakutan terbesar kamu?",
        "Pernah suka sama orang dan merasa orang itu suka sama kamu juga?",
        "Siapa nama mantan pacar teman mu yang pernah kamu sukai diam-diam?",
        "Pernah gak nyuri uang nyokap atau bokap? Alasannya?",
        "Hal yang bikin seneng pas lu lagi sedih apa?",
        "Pernah cinta bertepuk sebelah tangan? Kalo pernah sama siapa? Rasanya gimana, bro?",
        "Pernah jadi selingkuhan orang?",
        "Hal yang paling ditakutin?",
        "Siapa orang yang paling berpengaruh kepada kehidupanmu?",
        "Hal membanggakan apa yang kamu dapatkan di tahun ini?",
        "Siapa orang yang bisa membuatmu marah?",
        "Siapa orang yang pernah buatmu marah?",
        "(bgi yg muslim) Pernah ga solat seharian?",
        "Siapa yang paling mendekati tipe pasangan idealmu di sini?",
        "Suka mabar(main bareng) sama siapa?",
        "Pernah nolak orang? Alasannya kenapa?",
        "Sebutkan kejadian yang bikin kamu sakit hati yang masih di ingat?",
        "Pencapaian yang udah didapet apa aja ditahun ini?",
        "Kebiasaan terburuk lo pas di sekolah apa?",
      }
      rand.Seed(time.Now().UnixNano())
      index := rand.Intn(len(data1))
      hasil := data1[index]
      m.Reply(hasil)
    },
  })
}
