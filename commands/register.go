package commands

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/spf13/cobra"
	"gogit-sync/git"
	"gogit-sync/utils"
	"strings"
)

func RegisterMirror(cmd *cobra.Command, args []string) {
	fmt.Println("Let's register a new repository...")
	fmt.Printf("Name: ")

	repo := &git.Repository{}

	_, err := fmt.Scanln(&repo.Name)
	utils.CheckError(err)

	var isRemote bool

	for isRemote, err = getRemote(); err != nil; isRemote, err = getRemote() {
		fmt.Println("Invalid Remote Mode.")
	}

	if isRemote {
		getRemoteInfo(repo)
	} else {

	}


}

func getRemote() (bool, error) {
	fmt.Printf("Remote (Y/N): ")
	var remote string

	_, err := fmt.Scanln(&remote)

	utils.CheckError(err)

	switch strings.ToUpper(remote) {
	case "Y":
		return true, nil
	case "N":
		return false, nil
	default:
		return false, errors.New("invalid remote")
	}
}

func getRemoteInfo(repository *git.Repository) {
	fmt.Printf("Remote URL: ")

	var remote string

	_, err := fmt.Scanln(&remote)

	utils.CheckError(err)

	repository.Location = git.Location{
		Type: git.REMOTE,
		Value: remote,
	}

	var auth git.Auth

	for auth, err := getAuthMode(); err != nil; auth, err = getAuthMode() {
		fmt.Printf("Invalid, try again.")
	}

	repository.Auth =
}

func getAuthMode() (git.Auth, error) {
	fmt.Printf("Auth mode: [SSH, USER, NONE]")

	var authmode git.AuthMode

	_,err := fmt.Scanln(&authmode)

	utils.CheckError(err)

	if git.NONE != authmode && git.SSH != authmode && git.USER != authmode {
		return git.Auth{}, errors.New("invalid")
	}

	if authmode == git.NONE {
		return git.Auth{
			Type: git.NONE,
		}, nil
	}

	if authmode == git.SSH {

	}

	return authmode, nil
}