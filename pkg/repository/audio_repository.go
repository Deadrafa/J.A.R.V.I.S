package repository

type AudioDownloader interface {
	Download(fileID string) (string, error)
}

type SpeechRecognizer interface {
	Recognize(audioPath string) (string, error)
}
