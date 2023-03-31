package telegram

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

// Helper: interactive sign-in using terminal
type interactiveTerminalAuth struct {
	phone string
}

// Assert: implements auth.UserAuthenticator
var _ auth.UserAuthenticator = interactiveTerminalAuth{}

// Input: phone number
func (a interactiveTerminalAuth) Phone(_ context.Context) (string, error) {
	return promptStdin("Phone: ")
}

func (a interactiveTerminalAuth) Password(_ context.Context) (string, error) {
	// TODO: Use "golang.org/x/crypto/ssh/terminal" to read password
	// bytePwd, err := terminal.ReadPassword(0)
	return promptStdin("Telegram cloud password: ")
}

func (a interactiveTerminalAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	return promptStdin("2FA code: ")
}
func (c interactiveTerminalAuth) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, fmt.Errorf("not implemented")
}

func (c interactiveTerminalAuth) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}

func promptStdin(prompt string) (string, error) {
	fmt.Print(prompt)
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(value), nil
	}
}
