package author

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/romv7/blogs/internal/constants"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
)

const (
	StoragePlain           = storage.Plain
	authorInfoFileKey      = "author.toml"
	NMAX_AUTHOR_PARTITIONS = 32
)

var (
	ErrAttemptCreateHelperForUser = errors.New("attempting to create an AuthorHelper for an ordinary user")
	ErrAttemptSubscriptionToSelf  = errors.New("attempted to subscribe to self")
	ErrInvalidArgument            = errors.New("invalid argument")
	ErrNormalUserHasNoResourceId  = errors.New("normal user cannot have an author resource id")
)

type AuthorHelper struct {
	author     *pb.User
	authorInfo *AuthorInfo
	storage    storage.StorageDriver
}

type AuthorInfo struct {
	Bio         string                                `toml:"bio"`
	AltName     string                                `toml:"alt_name"`
	Stats       *AuthorStats                          `toml:"stats"`
	SocialLinks map[constants.SocialLinkType][]string `toml:"social_links"`
}

type AuthorStats struct {
	Subscriptions []string `toml:"subscriptions"`
}

type AuthorBlogMetadata struct {
	AuthorName   string `toml:"author"`
	BlogUuid     string `toml:"blog_uuid"`
	HeadlineText string `toml:"headline_text"`
	SummaryText  string `toml:"summary_text"`

	Attachments []string `toml:"attachments"`
	References  []string `toml:"references"`
}

func NewAuthorHelper(u *pb.User, s storage.StorageDriverType) (out *AuthorHelper) {
	out = &AuthorHelper{}

	if u.Type != pb.User_T_AUTHOR {
		log.Panic(ErrAttemptCreateHelperForUser)
	}

	out.author = u

	switch s {
	case StoragePlain:
		out.storage = storage.NewPlainStorage(u.StoragePath)
	default:
		log.Panic(storage.ErrorInvalidStorageDriver)
	}

	p := &bytes.Buffer{}

	// Create a default author.toml for the new author possible author.
	if !out.storage.Contains(authorInfoFileKey) {

		out.authorInfo = &AuthorInfo{
			Stats:       &AuthorStats{},
			SocialLinks: map[constants.SocialLinkType][]string{},
		}

		if err := toml.NewEncoder(p).Encode(out.authorInfo); err != nil {
			log.Panic(err)
		}

		if err := out.storage.Put(authorInfoFileKey, p.Bytes()); err != nil {
			log.Panic(err)
		}

	} else {
		out.authorInfo = &AuthorInfo{}

		var fileb []byte

		if b, err := out.storage.Get(authorInfoFileKey); err != nil {
			log.Panic(err)
		} else {
			fileb = b
		}

		p.Write(fileb)

		if _, err := toml.NewDecoder(p).Decode(out.authorInfo); err != nil {
			log.Panic(err)
		}
	}

	return
}

func (ah *AuthorHelper) GetAuthorMetadata() *AuthorInfo {
	return ah.authorInfo
}

func (ah *AuthorHelper) DeleteAuthorMetadata() (err error) {
	return ah.storage.Remove(authorInfoFileKey)
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

func (ah *AuthorHelper) AddSocialLink(s constants.SocialLinkType, url string) {
	ah.authorInfo.SocialLinks[s] = append(ah.authorInfo.SocialLinks[s], url)
}

func (ah *AuthorHelper) SaveAuthorMetadata() {
	b := new(bytes.Buffer)

	if err := toml.NewEncoder(b).Encode(ah.authorInfo); err != nil {
		log.Panic(err)
	}

	if err := ah.storage.Put(authorInfoFileKey, b.Bytes()); err != nil {
		log.Panic(err)
	}
}

func (ah *AuthorHelper) SaveAuthorPost(p *pb.Post) error {
	if p.Uuid == "" || p.HeadlineText == "" || p.State == nil || p.User == nil {
		log.Panic(ErrInvalidArgument)
	}

	metadata := &AuthorBlogMetadata{
		ah.author.Name,
		p.Uuid,
		p.HeadlineText,
		p.SummaryText,
		p.Attachments,
		p.Refs,
	}

	b := new(bytes.Buffer)

	if err := toml.NewEncoder(b).Encode(metadata); err != nil {
		log.Panic(err)
	}

	out := new(bytes.Buffer)
	P := []byte(fmt.Sprintf("%s\n-->\n%s",
		base64.RawURLEncoding.EncodeToString(b.Bytes()),
		base64.RawURLEncoding.EncodeToString([]byte(p.Content)),
	))

	zw := gzip.NewWriter(out)
	zw.Name = p.Uuid
	zw.Comment = p.HeadlineText
	zw.ModTime = p.State.CreatedAt.AsTime().UTC()

	if _, err := zw.Write(P); err != nil {
		log.Panic(err)
	}

	zw.Close()
	zw.Flush()

	if err := ah.storage.Put(ah.GetBlogPostFileKey(p.Uuid), out.Bytes()); err != nil {
		return err
	}

	return nil
}

func (ah *AuthorHelper) GetAuthorPostMetadata(p *pb.Post) (m *AuthorBlogMetadata, content string, err error) {
	var fileb bytes.Buffer

	if p, err := ah.storage.Get(ah.GetBlogPostFileKey(p.Uuid)); err != nil {
		return nil, "", err
	} else {
		if _, err := fileb.Write(p); err != nil {
			log.Panic(err)
		}
	}

	zr, err := gzip.NewReader(&fileb)
	if err != nil {
		log.Panic(err)
	}

	defer zr.Close()

	var uncom bytes.Buffer

	if _, err := io.Copy(&uncom, zr); err != nil {
		log.Panic(err)
	}

	m = &AuthorBlogMetadata{}
	data := bytes.Split(uncom.Bytes(), []byte("\n-->\n"))

	var b64data, contentb, dest []byte

	b64data = data[0]
	contentb = data[1]

	dest = make([]byte, base64.RawStdEncoding.DecodedLen(len(b64data)))

	if _, err := base64.RawStdEncoding.Decode(dest, b64data); err != nil {
		log.Panic(err)
	}

	if _, err = toml.NewDecoder(bytes.NewBuffer(dest)).Decode(m); err != nil {
		return nil, "", err
	}

	dest = make([]byte, base64.RawStdEncoding.DecodedLen(len(contentb)))

	if _, err := base64.RawStdEncoding.Decode(dest, contentb); err != nil {
		log.Panic(err)
	}

	return m, string(dest), nil
}

func (ah *AuthorHelper) DeletePostMetadata(p *pb.Post) error {
	return ah.storage.Remove(ah.GetBlogPostFileKey(p.Uuid))
}

func (ah *AuthorHelper) GetBlogPostFileKey(uuid string) string {
	return fmt.Sprintf("blogs/%s.gz",
		base64.RawStdEncoding.Strict().EncodeToString([]byte(uuid)),
	)
}

func AuthorRootResourceId(u *pb.User) string {
	if u.Type != pb.User_T_AUTHOR {
		log.Panic(ErrNormalUserHasNoResourceId)
	}

	uuidSum := uint64(0)

	for _, p := range []byte(strings.ReplaceAll(u.Uuid, "-", "")) {
		uuidSum += uint64(p)
	}

	return fmt.Sprintf("%d/%s", (uuidSum+u.Id)%NMAX_AUTHOR_PARTITIONS, u.Uuid)
}
