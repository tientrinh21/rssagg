package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tientrinh21/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedtoFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed get feeds: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}

func (apiCfg *apiConfig) handlerGetPostsByFeed(w http.ResponseWriter, r *http.Request) {
	feedIDString := chi.URLParam(r, "feedID")
	feedID, err := uuid.Parse(feedIDString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed id: %v", err))
		return
	}

	posts, err := apiCfg.DB.GetPostsByFeed(r.Context(), database.GetPostsByFeedParams{
		FeedID: feedID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to get posts: %v", err))
	}
	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
