package fav

type FolderResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Count int `json:"count"`
		List  []struct {
			Id         int    `json:"id"`
			Fid        int    `json:"fid"`
			Mid        int    `json:"mid"`
			Attr       int    `json:"attr"`
			Title      string `json:"title"`
			FavState   int    `json:"fav_state"`
			MediaCount int    `json:"media_count"`
		} `json:"list"`
		Season interface{} `json:"season"`
	} `json:"data"`
}

type ContentResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Info struct {
			Id    int    `json:"id"`
			Fid   int    `json:"fid"`
			Mid   int    `json:"mid"`
			Attr  int    `json:"attr"`
			Title string `json:"title"`
			Cover string `json:"cover"`
			Upper struct {
				Mid       int    `json:"mid"`
				Name      string `json:"name"`
				Face      string `json:"face"`
				Followed  bool   `json:"followed"`
				VipType   int    `json:"vip_type"`
				VipStatue int    `json:"vip_statue"`
			} `json:"upper"`
			CoverType int `json:"cover_type"`
			CntInfo   struct {
				Collect int `json:"collect"`
				Play    int `json:"play"`
				ThumbUp int `json:"thumb_up"`
				Share   int `json:"share"`
			} `json:"cnt_info"`
			Type       int    `json:"type"`
			Intro      string `json:"intro"`
			Ctime      int    `json:"ctime"`
			Mtime      int    `json:"mtime"`
			State      int    `json:"state"`
			FavState   int    `json:"fav_state"`
			LikeState  int    `json:"like_state"`
			MediaCount int    `json:"media_count"`
		} `json:"info"`
		Medias  []Media `json:"medias"`
		HasMore bool    `json:"has_more"`
	} `json:"data"`
}

type Media struct {
	Id       int    `json:"id"`
	Type     int    `json:"type"`
	Title    string `json:"title"`
	Cover    string `json:"cover"`
	Intro    string `json:"intro"`
	Page     int    `json:"page"`
	Duration int    `json:"duration"`
	Upper    struct {
		Mid  int    `json:"mid"`
		Name string `json:"name"`
		Face string `json:"face"`
	} `json:"upper"`
	Attr    int `json:"attr"`
	CntInfo struct {
		Collect int `json:"collect"`
		Play    int `json:"play"`
		Danmaku int `json:"danmaku"`
	} `json:"cnt_info"`
	Link    string      `json:"link"`
	Ctime   int         `json:"ctime"`
	Pubtime int         `json:"pubtime"`
	FavTime int         `json:"fav_time"`
	BvId    string      `json:"bv_id"`
	Bvid    string      `json:"bvid"`
	Season  interface{} `json:"season"`
}
