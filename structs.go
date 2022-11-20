package main

import "time"

type IdentifyingRequest struct {
	Event string `json:"event"`
}

type MatrixWebhook struct {
	Body string `json:"body"`
	Key  string `json:"key"`
}

type MastodonSignUpEvent struct {
	Event     string    `json:"event"`
	CreatedAt time.Time `json:"created_at"`
	Object    struct {
		ID        string      `json:"id"`
		Username  string      `json:"username"`
		Domain    interface{} `json:"domain"`
		CreatedAt time.Time   `json:"created_at"`
		Email     string      `json:"email"`
		IP        string      `json:"ip"`
		Role      struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Color       string    `json:"color"`
			Position    int       `json:"position"`
			Permissions int       `json:"permissions"`
			Highlighted bool      `json:"highlighted"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		} `json:"role"`
		Confirmed     bool   `json:"confirmed"`
		Suspended     bool   `json:"suspended"`
		Silenced      bool   `json:"silenced"`
		Sensitized    bool   `json:"sensitized"`
		Disabled      bool   `json:"disabled"`
		Approved      bool   `json:"approved"`
		Locale        string `json:"locale"`
		InviteRequest string `json:"invite_request"`
		Ips           []struct {
			IP     string    `json:"ip"`
			UsedAt time.Time `json:"used_at"`
		} `json:"ips"`
		Account struct {
			ID             string        `json:"id"`
			Username       string        `json:"username"`
			Acct           string        `json:"acct"`
			DisplayName    string        `json:"display_name"`
			Locked         bool          `json:"locked"`
			Bot            bool          `json:"bot"`
			Discoverable   interface{}   `json:"discoverable"`
			Group          bool          `json:"group"`
			CreatedAt      time.Time     `json:"created_at"`
			Note           string        `json:"note"`
			URL            string        `json:"url"`
			Avatar         string        `json:"avatar"`
			AvatarStatic   string        `json:"avatar_static"`
			Header         string        `json:"header"`
			HeaderStatic   string        `json:"header_static"`
			FollowersCount int           `json:"followers_count"`
			FollowingCount int           `json:"following_count"`
			StatusesCount  int           `json:"statuses_count"`
			LastStatusAt   interface{}   `json:"last_status_at"`
			Noindex        bool          `json:"noindex"`
			Emojis         []interface{} `json:"emojis"`
			Fields         []interface{} `json:"fields"`
		} `json:"account"`
	} `json:"object"`
}

type MastodonReportEvent struct {
	Event     string    `json:"event"`
	CreatedAt time.Time `json:"created_at"`
	Object    struct {
		ID            string      `json:"id"`
		ActionTaken   bool        `json:"action_taken"`
		ActionTakenAt interface{} `json:"action_taken_at"`
		Category      string      `json:"category"`
		Comment       string      `json:"comment"`
		Forwarded     bool        `json:"forwarded"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     time.Time   `json:"updated_at"`
		Account       struct {
			ID        string      `json:"id"`
			Username  string      `json:"username"`
			Domain    interface{} `json:"domain"`
			CreatedAt time.Time   `json:"created_at"`
			Email     string      `json:"email"`
			IP        string      `json:"ip"`
			Role      struct {
				ID          int       `json:"id"`
				Name        string    `json:"name"`
				Color       string    `json:"color"`
				Position    int       `json:"position"`
				Permissions int       `json:"permissions"`
				Highlighted bool      `json:"highlighted"`
				CreatedAt   time.Time `json:"created_at"`
				UpdatedAt   time.Time `json:"updated_at"`
			} `json:"role"`
			Confirmed     bool        `json:"confirmed"`
			Suspended     bool        `json:"suspended"`
			Silenced      bool        `json:"silenced"`
			Sensitized    bool        `json:"sensitized"`
			Disabled      bool        `json:"disabled"`
			Approved      bool        `json:"approved"`
			Locale        string      `json:"locale"`
			InviteRequest interface{} `json:"invite_request"`
			Ips           []struct {
				IP     string    `json:"ip"`
				UsedAt time.Time `json:"used_at"`
			} `json:"ips"`
			Account struct {
				ID             string        `json:"id"`
				Username       string        `json:"username"`
				Acct           string        `json:"acct"`
				DisplayName    string        `json:"display_name"`
				Locked         bool          `json:"locked"`
				Bot            bool          `json:"bot"`
				Discoverable   bool          `json:"discoverable"`
				Group          bool          `json:"group"`
				CreatedAt      time.Time     `json:"created_at"`
				Note           string        `json:"note"`
				URL            string        `json:"url"`
				Avatar         string        `json:"avatar"`
				AvatarStatic   string        `json:"avatar_static"`
				Header         string        `json:"header"`
				HeaderStatic   string        `json:"header_static"`
				FollowersCount int           `json:"followers_count"`
				FollowingCount int           `json:"following_count"`
				StatusesCount  int           `json:"statuses_count"`
				LastStatusAt   string        `json:"last_status_at"`
				Noindex        bool          `json:"noindex"`
				Emojis         []interface{} `json:"emojis"`
				Fields         []interface{} `json:"fields"`
			} `json:"account"`
		} `json:"account"`
		TargetAccount struct {
			ID            string      `json:"id"`
			Username      string      `json:"username"`
			Domain        string      `json:"domain"`
			CreatedAt     time.Time   `json:"created_at"`
			Email         interface{} `json:"email"`
			IP            interface{} `json:"ip"`
			Role          interface{} `json:"role"`
			Confirmed     interface{} `json:"confirmed"`
			Suspended     bool        `json:"suspended"`
			Silenced      bool        `json:"silenced"`
			Sensitized    bool        `json:"sensitized"`
			Disabled      interface{} `json:"disabled"`
			Approved      interface{} `json:"approved"`
			Locale        interface{} `json:"locale"`
			InviteRequest interface{} `json:"invite_request"`
			Ips           interface{} `json:"ips"`
			Account       struct {
				ID             string        `json:"id"`
				Username       string        `json:"username"`
				Acct           string        `json:"acct"`
				DisplayName    string        `json:"display_name"`
				Locked         bool          `json:"locked"`
				Bot            bool          `json:"bot"`
				Discoverable   bool          `json:"discoverable"`
				Group          bool          `json:"group"`
				CreatedAt      time.Time     `json:"created_at"`
				Note           string        `json:"note"`
				URL            string        `json:"url"`
				Avatar         string        `json:"avatar"`
				AvatarStatic   string        `json:"avatar_static"`
				Header         string        `json:"header"`
				HeaderStatic   string        `json:"header_static"`
				FollowersCount int           `json:"followers_count"`
				FollowingCount int           `json:"following_count"`
				StatusesCount  int           `json:"statuses_count"`
				LastStatusAt   string        `json:"last_status_at"`
				Emojis         []interface{} `json:"emojis"`
				Fields         []struct {
					Name       string      `json:"name"`
					Value      string      `json:"value"`
					VerifiedAt interface{} `json:"verified_at"`
				} `json:"fields"`
			} `json:"account"`
		} `json:"target_account"`
		AssignedAccount      interface{} `json:"assigned_account"`
		ActionTakenByAccount interface{} `json:"action_taken_by_account"`
		Statuses             []struct {
			ID                 string      `json:"id"`
			CreatedAt          time.Time   `json:"created_at"`
			InReplyToID        interface{} `json:"in_reply_to_id"`
			InReplyToAccountID interface{} `json:"in_reply_to_account_id"`
			Sensitive          bool        `json:"sensitive"`
			SpoilerText        string      `json:"spoiler_text"`
			Visibility         string      `json:"visibility"`
			Language           string      `json:"language"`
			URI                string      `json:"uri"`
			URL                string      `json:"url"`
			RepliesCount       int         `json:"replies_count"`
			ReblogsCount       int         `json:"reblogs_count"`
			FavouritesCount    int         `json:"favourites_count"`
			EditedAt           time.Time   `json:"edited_at"`
			Content            string      `json:"content"`
			Reblog             interface{} `json:"reblog"`
			Account            struct {
				ID             string        `json:"id"`
				Username       string        `json:"username"`
				Acct           string        `json:"acct"`
				DisplayName    string        `json:"display_name"`
				Locked         bool          `json:"locked"`
				Bot            bool          `json:"bot"`
				Discoverable   bool          `json:"discoverable"`
				Group          bool          `json:"group"`
				CreatedAt      time.Time     `json:"created_at"`
				Note           string        `json:"note"`
				URL            string        `json:"url"`
				Avatar         string        `json:"avatar"`
				AvatarStatic   string        `json:"avatar_static"`
				Header         string        `json:"header"`
				HeaderStatic   string        `json:"header_static"`
				FollowersCount int           `json:"followers_count"`
				FollowingCount int           `json:"following_count"`
				StatusesCount  int           `json:"statuses_count"`
				LastStatusAt   string        `json:"last_status_at"`
				Emojis         []interface{} `json:"emojis"`
				Fields         []struct {
					Name       string      `json:"name"`
					Value      string      `json:"value"`
					VerifiedAt interface{} `json:"verified_at"`
				} `json:"fields"`
			} `json:"account"`
			MediaAttachments []interface{} `json:"media_attachments"`
			Mentions         []interface{} `json:"mentions"`
			Tags             []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"tags"`
			Emojis []interface{} `json:"emojis"`
			Card   struct {
				URL          string `json:"url"`
				Title        string `json:"title"`
				Description  string `json:"description"`
				Type         string `json:"type"`
				AuthorName   string `json:"author_name"`
				AuthorURL    string `json:"author_url"`
				ProviderName string `json:"provider_name"`
				ProviderURL  string `json:"provider_url"`
				HTML         string `json:"html"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Image        string `json:"image"`
				EmbedURL     string `json:"embed_url"`
				Blurhash     string `json:"blurhash"`
			} `json:"card"`
			Poll interface{} `json:"poll"`
		} `json:"statuses"`
		Rules []struct {
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"rules"`
	} `json:"object"`
}
