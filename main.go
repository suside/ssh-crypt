package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/suside/ssh-crypt/lib"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	idRsaDefault = os.ExpandEnv(
		strings.Join([]string{"$HOME", ".ssh", "id_rsa"}, string(os.PathSeparator)))
	authKeysDefault = os.ExpandEnv(
		strings.Join([]string{"$HOME", ".ssh", "authorized_keys"}, string(os.PathSeparator)))

	version = "master"
	app     = kingpin.New(
		"ssh-crypt",
		fmt.Sprintf("Encrypt file with your ssh keys. (%s)", version),
	)
	idRsaPath = app.Flag(
		"identity-file",
		"File from which the identity (private key) is read.",
	).Short('i').Default(idRsaDefault).ExistingFile()

	edit       = app.Command("edit", "Edit vault content")
	vaultFileE = edit.Arg(
		"vault-file",
		"File where encrypted content is stored.",
	).Required().String()
	authKeysPath = edit.Flag(
		"authorized-keys",
		"File from which the public keys are read.",
	).Short('a').Default(authKeysDefault).ExistingFile()
	fromStdin = edit.Flag(
		"stdin",
		"Read from standard input instead from $EDITOR",
	).Short('s').Bool()

	view       = app.Command("view", "Print vault content")
	vaultFileV = view.Arg(
		"vault-file",
		"File where encrypted content is stored.",
	).Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case edit.FullCommand():
		vault := lib.Vault{}
		vault.ReadAuthorizedKeys(*authKeysPath)
		if *fromStdin {
			vault.ReadStdIn()
		} else {
			err := vault.DecryptVaultWithKey(*vaultFileE, *idRsaPath)
			if err != nil {
				log.Fatal(err)
			}
			vault.EditVaultFile()
		}
		vault.StoreSecuredVault(*vaultFileE)
	case view.FullCommand():
		vault := lib.Vault{}
		err := vault.DecryptVaultWithKey(*vaultFileV, *idRsaPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(vault.Plaintext))
	}
}
