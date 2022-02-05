package langs

type Language interface {
	Flag() string
	ChooseLanguageText() string
	WelcomeMenu() string
}
