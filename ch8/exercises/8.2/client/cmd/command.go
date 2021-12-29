package cmd

type Command struct {
	Filename string `json:",omitempty",`
	Data []byte `json:",omitempty"`
	FileSize int64 `json:",omitempty"`
	BytesRead int64 `json:",omitempty"`
	CommandName string
	UUID string
}
