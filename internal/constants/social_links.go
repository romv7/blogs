package constants

type SocialLinkType uint

const (
	MASTODON SocialLinkType = iota
	FACEBOOK
	TWITTER
	REDDIT
	ARTSTATION
	YOUTUBE
	INSTAGRAM
	LINKEDIN
	DRIBBBLE
	DEVIANTART
	STEAM
	QUORA
)

func (s SocialLinkType) String() string {
	name := []string{
		"MASTODON", "FACEBOOK", "TWITTER", "REDDIT", "ARTSTATION",
		"YOUTUBE", "INSTAGRAM", "LINKEDIN", "DRIBBBLE", "DEVIANTART",
		"STEAM", "QUORA",
	}

	return name[s]
}
