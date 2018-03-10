package main

import (
	"flag"
	"log"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

func getVoice(lang *string) string {
	switch *lang {
		case "ca-fr":
			return "Chantal"
		case "fr":
			return "Celine"
		case "en":
			return "Joanna"
		case "de":
			fallthrough
		default:
			return "Marlene"
	}
}

func main() {

	textPtr := flag.String("text", "-", "text input (Default is - for stdin")
	filePtr := flag.String("input-file", "", "input file")
	outPtr := flag.String("out", "-", "audio outfile (Default is - for stdout")
	langPtr := flag.String("lang", "de", "language")

	flag.Parse()
	var text string
	var err error
	var t []byte
	if *textPtr == "-" {
		if *filePtr == "" {
			t, err = ioutil.ReadAll(os.Stdin)
		} else {
			t, err = ioutil.ReadFile(*filePtr)
		}
		if err != nil {
			log.Fatal(err)
		}
		text = string(t)
	} else {
		text = *textPtr
	}
	voice := getVoice(langPtr)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := polly.New(sess)

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text: &text,
		VoiceId: &voice,
	}

	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		log.Fatal(err)
	}
	var outFile io.Writer
	if *outPtr == "-" {
		outFile = os.Stdout
	} else {
		file, err := os.Create(*outPtr)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		outFile = file
	}
	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		log.Fatal(err)
	}
}
