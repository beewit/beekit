package weibo

import (
	"encoding/json"
	"fmt"
	"github.com/beewit/beekit/utils/uhttp"
)

type (
	// Weibo weibo
	Weibo struct {
		ID               int64       `json:"id"`                 // 用户UID
		IDStr            string      `json:"idstr"`              // 字符串型的用户UID
		ScreenName       string      `json:"screen_name"`        //用户昵称
		Name             string      `json:"name"`               // 友好显示名称
		Province         int         `json:"province"`           // 用户所在省级ID
		City             int         `json:"city"`               // 用户所在城市ID
		Location         string      `json:"location"`           // 用户所在地
		Description      string      `json:"description"`        // 用户个人描述
		URL              string      `json:"url"`                // 用户博客地址
		ProfileImageURL  string      `json:"profile_image_url"`  // 用户头像地址（中图），50×50像素
		ProfileURL       string      `json:"profile_url"`        // 用户的微博统一URL地址
		Domain           string      `json:"domain"`             // 用户的个性化域名
		Weihao           string      `json:"weihao"`             // 用户的微号
		Gender           string      `json:"gender"`             // 性别，m：男、f：女、n：未知
		FollowersCount   int         `json:"followers_count"`    // 粉丝数
		FriendsCount     int         `json:"friends_count"`      // 关注数
		StatusesCount    int         `json:"statuses_count"`     // 微博数
		FavouritesCount  int         `json:"favourites_count"`   // 收藏数
		CreatedAt        string      `json:"created_at"`         // 用户创建（注册）时间
		Following        bool        `json:"following"`          // 暂未支持
		AllowAllActMsg   bool        `json:"allow_all_act_msg"`  // 是否允许所有人给我发私信，true：是，false：否
		GeoEnabled       bool        `json:"geo_enabled"`        // 是否允许标识用户的地理位置，true：是，false：否
		Verified         bool        `json:"verified"`           // 是否是微博认证用户，即加V用户，true：是，false：否
		VerifiedType     int         `json:"verified_type"`      // 暂未支持
		Remark           string      `json:"remark"`             // 用户备注信息，只有在查询用户关系时才返回此字段
		Status           interface{} `json:"status"`             // 用户的最近一条微博信息字段 详细
		AllowAllComment  bool        `json:"allow_all_comment"`  // 是否允许所有人对我的微博进行评论，true：是，false：否
		AvatarLarge      string      `json:"avatar_large"`       // 用户头像地址（大图），180×180像素
		AvatarHD         string      `json:"avatar_hd"`          // 用户头像地址（高清），高清头像原图
		VerifiedReason   string      `json:"verified_reason"`    // 认证原因
		FollowMe         bool        `json:"follow_me"`          // 该用户是否关注当前登录用户，true：是，false：否
		OnlineStatus     int         `json:"online_status"`      // 用户的在线状态，0：不在线、1：在线
		BiFollowersCount int         `json:"bi_followers_count"` // 用户的互粉数
		Lang             string      `json:"lang"`               // 用户当前的语言版本，zh-cn：简体中文，zh-tw：繁体中文，en：英语
	}

	AccessToken struct {
		AccessToken string      `json:"access_token"`
		RemindIn    string      `json:"remind_in"`
		ExpiresIn   int         `json:"expires_in"`
		Uid         string      `json:"uid"`
	}
)

// NewWeibo new weibo
func NewWeibo() *Weibo {
	return &Weibo{}
}

// User user
func (w Weibo) User(accessToken, uID string) (Weibo, error) {
	result := Weibo{}
	url := fmt.Sprintf("https://api.weibo.com/2/users/show.json?access_token=%v&uid=%v",
		accessToken,
		uID,
	)
	body, err := uhttp.Cmd(uhttp.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

func (w Weibo) GetAccessToken(appKey, appSecret, redirectUri, code string) (AccessToken, error) {
	result := AccessToken{}
	url := fmt.Sprintf("https://api.weibo.com/oauth2/access_token?client_id=%v&client_secret=%v&grant_type=authorization_code" +
		"&redirect_uri=%v&code=%v",
		appKey,
		appSecret,
		redirectUri,
		code,
	)
	body, err := uhttp.Cmd(uhttp.Request{
		Method: "POST",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	//失败 {"error":"invalid_grant","error_code":21325,"request":"/oauth2/access_token","error_uri":"/oauth2/access_token","error_description":"invalid authorization code:85677e9ece16654ca1432e7269a3b86f"}
	println(string(body[:]))
	err = json.Unmarshal(body, &result)
	return result, err
}
