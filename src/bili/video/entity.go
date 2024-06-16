package video

type Info struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizId   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionId int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Premiere           interface{} `json:"premiere"`
		TeenageMode        int         `json:"teenage_mode"`
		IsChargeableSeason bool        `json:"is_chargeable_season"`
		IsStory            bool        `json:"is_story"`
		NoCache            bool        `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		Staff []struct {
			Mid   int    `json:"mid"`
			Title string `json:"title"`
			Name  string `json:"name"`
			Face  string `json:"face"`
			Vip   struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelUriHans       string `json:"img_label_uri_hans"`
					ImgLabelUriHant       string `json:"img_label_uri_hant"`
					ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
			} `json:"vip"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official"`
			Follower   int `json:"follower"`
			LabelStyle int `json:"label_style"`
		} `json:"staff"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
			Honor []struct {
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				Desc               string `json:"desc"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
		LikeIcon string `json:"like_icon"`
	} `json:"data"`
}

type PlayInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Dash              struct {
			Duration       int     `json:"duration"`
			MinBufferTime  float64 `json:"minBufferTime"`
			MinBufferTime1 float64 `json:"min_buffer_time"`
			Video          []struct {
				Id            int      `json:"id"`
				BaseUrl       string   `json:"baseUrl"`
				BaseUrl1      string   `json:"base_url"`
				BackupUrl     []string `json:"backupUrl"`
				BackupUrl1    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType1     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate1    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap1 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase1 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				Id            int      `json:"id"`
				BaseUrl       string   `json:"baseUrl"`
				BaseUrl1      string   `json:"base_url"`
				BackupUrl     []string `json:"backupUrl"`
				BackupUrl1    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType1     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate1    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap1 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase1 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"audio"`
			Dolby struct {
				Type  int `json:"type"`
				Audio []struct {
					Id            int      `json:"id"`
					BaseUrl       string   `json:"baseUrl"`
					BaseUrl1      string   `json:"base_url"`
					BackupUrl     []string `json:"backupUrl"`
					BackupUrl1    []string `json:"backup_url"`
					Bandwidth     int      `json:"bandwidth"`
					MimeType      string   `json:"mimeType"`
					MimeType1     string   `json:"mime_type"`
					Codecs        string   `json:"codecs"`
					Width         int      `json:"width"`
					Height        int      `json:"height"`
					FrameRate     string   `json:"frameRate"`
					FrameRate1    string   `json:"frame_rate"`
					Sar           string   `json:"sar"`
					StartWithSap  int      `json:"startWithSap"`
					StartWithSap1 int      `json:"start_with_sap"`
					SegmentBase   struct {
						Initialization string `json:"Initialization"`
						IndexRange     string `json:"indexRange"`
					} `json:"SegmentBase"`
					SegmentBase1 struct {
						Initialization string `json:"initialization"`
						IndexRange     string `json:"index_range"`
					} `json:"segment_base"`
					Codecid int `json:"codecid"`
				} `json:"audio"`
			} `json:"dolby"`
			Flac struct {
				Display bool `json:"display"`
				Audio   struct {
					Id            int      `json:"id"`
					BaseUrl       string   `json:"baseUrl"`
					BaseUrl1      string   `json:"base_url"`
					BackupUrl     []string `json:"backupUrl"`
					BackupUrl1    []string `json:"backup_url"`
					Bandwidth     int      `json:"bandwidth"`
					MimeType      string   `json:"mimeType"`
					MimeType1     string   `json:"mime_type"`
					Codecs        string   `json:"codecs"`
					Width         int      `json:"width"`
					Height        int      `json:"height"`
					FrameRate     string   `json:"frameRate"`
					FrameRate1    string   `json:"frame_rate"`
					Sar           string   `json:"sar"`
					StartWithSap  int      `json:"startWithSap"`
					StartWithSap1 int      `json:"start_with_sap"`
					SegmentBase   struct {
						Initialization string `json:"Initialization"`
						IndexRange     string `json:"indexRange"`
					} `json:"SegmentBase"`
					SegmentBase1 struct {
						Initialization string `json:"initialization"`
						IndexRange     string `json:"index_range"`
					} `json:"segment_base"`
					Codecid int `json:"codecid"`
				} `json:"audio"`
			} `json:"flac"`
		} `json:"dash"`
		Durl []struct {
			Order     int      `json:"order"`
			Length    int      `json:"length"`
			Size      int      `json:"size"`
			Ahead     string   `json:"ahead"`
			Vhead     string   `json:"vhead"`
			URL       string   `json:"url"`
			BackupURL []string `json:"backup_url"`
		} `json:"durl"`
		SupportFormats []struct {
			Quality        int      `json:"quality"`
			Format         string   `json:"format"`
			NewDescription string   `json:"new_description"`
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Codecs         []string `json:"codecs"`
		} `json:"support_formats"`
		HighFormat   interface{} `json:"high_format"`
		LastPlayTime int         `json:"last_play_time"`
		LastPlayCid  int         `json:"last_play_cid"`
	} `json:"data"`
}
