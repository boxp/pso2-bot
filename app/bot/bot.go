package bot

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/boxp/pso2-bot/app/models"
	"github.com/boxp/pso2-bot/app/services"
	"github.com/revel/revel"
)

type Bot struct {
	Me            anaconda.User
	Api           *anaconda.TwitterApi
	TwitterStream *anaconda.Stream
}

func NewBot() (*Bot, error) {

	consumer_key := revel.Config.StringDefault("twitter.consumer_key", "")
	consumer_secret := revel.Config.StringDefault("twitter.consumer_secret", "")

	access_token := revel.Config.StringDefault("twitter.access_token", "")
	access_token_secret := revel.Config.StringDefault("twitter.access_token_secret", "")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)

	api := anaconda.NewTwitterApi(access_token, access_token_secret)

	twitterStream := api.UserStream(nil)

	me, err := api.GetSelf(nil)
	if err != nil {
		log.Fatalf("Failed to getself %v\n", err)
	}

	b := &Bot{me, api, twitterStream}

	return b, err
}

func (b Bot) Reply(tweet anaconda.Tweet, context string) {
	s := "@" + tweet.User.ScreenName + " " + context

	_, err := b.Api.PostTweet(s, nil)
	if err != nil {
		log.Printf("Failed to Reply %v\n", err)
	}
}

func (b Bot) RegisterArks(tweet anaconda.Tweet, shipStr string) {
	ship, err := strconv.Atoi(shipStr)
	if err != nil {
		log.Fatalf("Failed to parse shipStr %v\n", err)
	}
	name := tweet.User.Name
	screenName := tweet.User.ScreenName

	arks := models.Arks{Ship: ship, Name: name, ScreenName: screenName}

	services.CreateArks(arks)

	context := "Ship" + shipStr + "で登録しました。"
	b.Reply(tweet, context)
}

func (b Bot) SearchArks(tweet anaconda.Tweet, shipStr string) {
	context := ""
	ship, err := strconv.Atoi(shipStr)
	if err != nil {
		log.Fatalf("Failed to parse shipStr %v\n", err)
	}

	services.DeleteExpiredArks()

	arkses := services.SearchArksByShip(ship)

	if len(arkses) > 0 {
		context = "Ship" + shipStr + "の人権獲得者です！\n"

		for _, arks := range arkses {
			context = context + arks.Name + "(" + arks.ScreenName + ")\n"
		}
	} else {
		context = "現在Ship" + shipStr + "に人権獲得者は居ません"
	}

	b.Reply(tweet, context)

}

func (b Bot) OnReply(tweet anaconda.Tweet) {

	toRegisterRegexp := regexp.MustCompile(`Ship(\d+)で拉致[らさ]れた！`)
	toSearchRegexp := regexp.MustCompile(`Ship(\d+)$`)

	t := strings.Split(tweet.Text, " ")[1]

	switch {
	case toRegisterRegexp.MatchString(t):
		// 新規登録
		shipStr := toRegisterRegexp.ReplaceAllString(t, `$1`)
		b.RegisterArks(tweet, shipStr)
	case toSearchRegexp.MatchString(t):
		shipStr := toSearchRegexp.ReplaceAllString(t, `$1`)
		b.SearchArks(tweet, shipStr)
	}
}

func (b Bot) Start() {

	// Replyの正規表現
	r := regexp.MustCompile("@" + b.Me.ScreenName + " ")

	// TwitterStream listner
	for {
		x := <-b.TwitterStream.C
		switch tweet := x.(type) {
		case anaconda.Tweet:
			if r.MatchString(tweet.Text) {
				b.OnReply(tweet)
			}
		default:
			log.Printf("Received unhandled event: %s\n", tweet)
		}
	}
}

func InitBot() {
	b, err := NewBot()
	if err != nil {
		log.Fatalf("Failed to construct Bot %v\n", err)
	}
	go b.Start()
}
