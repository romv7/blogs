package constants

// Provides a set of enum values for what are the available social link types
// that can be added into the users info. This is not an exhaustive list,
// if there is an unknown social link type that a user wants to add, what they can do
// is issue a ticket so that we could add the social link type they want.
type SocialLinkType string

const (
	MASTODON   SocialLinkType = "S_MASTODON"
	FACEBOOK                  = "S_FACEBOOK"
	TWITTER                   = "S_TWITTER"
	REDDIT                    = "S_REDDIT"
	ARTSTATION                = "S_ARTSTATION"
	YOUTUBE                   = "S_YOUTUBE"
	INSTAGRAM                 = "S_INSTAGRAM"
	LINKEDIN                  = "S_LINKEDIN"
	DRIBBBLE                  = "S_DRIBBBLE"
	DEVIANTART                = "S_DEVIANTART"
	STEAM                     = "S_STEAM"
	QUORA                     = "S_QUORA"
	TWITCH                    = "S_TWITCH"
	DISCORD                   = "S_DISCORD"
)
