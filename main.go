package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"pk/aesx"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "decrypt",
		Aliases: []string{"d", "D"},
		Usage:   "Is the decrypt mode? or default encrypt mode",
	},
	&cli.StringFlag{
		Name:     "key",
		Aliases:  []string{"k"},
		Usage:    "The key for AES process",
		Required: true,
	},
	&cli.StringFlag{
		Name:     "value",
		Aliases:  []string{"v"},
		Usage:    "The text value that you want to encrypt or decrypt",
		Required: true,
	},
	&cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "The file path for the output text",
		Value:   "./.env",
	},
}

//go:generate ./pk --key a475a27Bb1028f140bc2a7c843318afD --value "This is a secret!"
func main() {
	app := &cli.App{
		Flags: flags,
		Name:  "pk",
		Usage: "the pk is the .env of secret mode for  boot",
		Action: func(cCtx *cli.Context) error {
			isDecryptMode := cCtx.Bool("decrypt")
			key := cCtx.String("key")
			value := cCtx.String("value")
			output := cCtx.String("output")
			// log.Printf("[PARAMS] isDecryptMode=%t key=%s value=%s output=%s\n", isDecryptMode, key, value, output)
			var (
				text string
				err  error
			)
			if isDecryptMode {
				text, err = aesx.DecryptAES([]byte(key), value)
				CheckError(err)
			} else {
				text, err = aesx.EncryptAES([]byte(key), value)
				CheckError(err)
			}
			file, err := os.Create(output)
			CheckError(err)
			defer func(file *os.File) {
				_ = file.Close()
			}(file)
			if _, err = file.Write([]byte(text)); err != nil {
				log.Fatal("[ERROR] Can't Write to file :", output, err)
			}
			fmt.Println()
			log.Printf("Success! Value=%s OutputFile = %s\n\n", text, output)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
