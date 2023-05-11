package author

import (
	"bytes"
	"errors"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/romv7/blogs/internal/pb"
)

var (
	ErrAttemptCreateHelperForUser = errors.New("attempting to create an AuthorHelper for an ordinary user")
	ErrAttemptSubscriptionToSelf  = errors.New("attempted to subscribe to self")
	ErrInvalidArgument            = errors.New("invalid argument")
)

type AuthorHelper struct {
	author     *pb.User
	authorInfo *AuthorInfo
}

type AuthorInfo struct {
	Bio         string                      `toml:"bio"`
	AltName     string                      `toml:"alt_name"`
	Stats       *authorStats                `toml:"stats"`
	SocialLinks map[SocialLinkType][]string `toml:"social_links"`
}

type authorStats struct {
	Subscriptions []string `toml:"subscriptions"`
}

type AuthorBlogMetadata struct {
	AuthorName   string `toml:"author"`
	BlogUuid     string `toml:"blog_uuid"`
	HeadlineText string `toml:"headline_text"`
	SummaryText  string `toml:"summary_text"`

	Images      []string `toml:"images"`
	Attachments []string `toml:"attachments"`
	References  []string `toml:"references"`
}

func NewAuthorHelper(u *pb.User) *AuthorHelper {
	if u.Type != pb.User_T_AUTHOR {
		log.Panic(ErrAttemptCreateHelperForUser)
	}

	inf := &AuthorInfo{}

	return &AuthorHelper{u, inf}
}

func (ah *AuthorHelper) GetAuthorMetadata() *AuthorInfo {
	return ah.authorInfo
}

func (ah *AuthorHelper) SetBio(text string) {
	ah.authorInfo.Bio = text
}

func (ah *AuthorHelper) SetAltName(alt string) {
	ah.authorInfo.AltName = alt
}

func (ah *AuthorHelper) SubscribeTo(u *pb.User) {
	if u.Uuid == ah.author.Uuid {
		log.Panic(ErrAttemptSubscriptionToSelf)
	}
}

func (ah *AuthorHelper) AddSocialLink(s SocialLinkType, url string) {
	ah.authorInfo.SocialLinks[s] = append(ah.authorInfo.SocialLinks[s], url)
}

func (ah *AuthorHelper) SaveAuthorMetadata() {
	b := new(bytes.Buffer)

	if err := toml.NewEncoder(b).Encode(ah.authorInfo); err != nil {
		log.Panic(err)
	}
}

func (ah *AuthorHelper) SaveAuthorPost(p *pb.Post) {
	if p.Uuid == "" || p.HeadlineText == "" || p.State == nil || p.User == nil {
		log.Panic(ErrInvalidArgument)
	}

	metadata := &AuthorBlogMetadata{
		ah.author.Name,
		p.Uuid,
		p.HeadlineText,
		p.SummaryText,
		p.Images,
		p.Attachments,
		p.Refs,
	}

	b := new(bytes.Buffer)

	if err := toml.NewEncoder(b).Encode(metadata); err != nil {
		log.Panic(err)
	}

}
