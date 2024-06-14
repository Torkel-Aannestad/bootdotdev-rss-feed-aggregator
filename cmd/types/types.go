package types

import (
	"database/sql"
	"time"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func DatabaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt *time.Time
}

func DatabaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: sql.NullTime{Time: feed.LastFetchedAt, Valid: feed.LastFetchedAt.Valid},
	}
}
func DatabaseFeedsToFeeds(feeds []database.Feed) []Feed {
	feedsSlice := make([]Feed, len(feeds))
	for i, feed := range feeds {
		feedsSlice[i] = DatabaseFeedToFeed(feed)
	}
	return feedsSlice
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func DatabaseFeedFollowToFeedFollow(databaseFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        databaseFeedFollow.ID,
		CreatedAt: databaseFeedFollow.CreatedAt,
		UpdatedAt: databaseFeedFollow.UpdatedAt,
		UserID:    databaseFeedFollow.UserID,
		FeedID:    databaseFeedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	feedFollowsSlice := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		feedFollowsSlice[i] = DatabaseFeedFollowToFeedFollow(feedFollow)
	}
	return feedFollowsSlice
}
