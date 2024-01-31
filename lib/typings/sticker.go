package typings

type MediaType string

const (
	IMAGE = "image"
	VIDEO = "video"
    TEKS = "teks"
)

type MetadataSticker struct {
	Author    string
	Pack      string
	KeepScale bool
	Removebg  any
	Circle    bool
}

type Sticker struct {
	File []byte
	Tipe MediaType
}
