package token

type RefreshToken struct {
	ID    string `json:"-"`
	UID   string `json:"-"`
	Token string `json:"refreshToken"`
}

type IDToken struct {
	Token string `json:"idToken"`
}

type TokenPair struct {
	IDToken
	RefreshToken
}

type Token struct {
	Plaintext string
	Hash      string
}
