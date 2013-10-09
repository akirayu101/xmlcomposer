package crawl

import (
    "net/http"
    "regexp"
    "strings"
    "encoding/base64"
    "log"
    "io/ioutil"
    "strconv"
    "fmt"
)
//google api
func Crawl_google(page string) string{
    client := &http.Client{}
    req, err := http.NewRequest("GET",page,nil)
    if err != nil{
        log.Fatalln(err)
    }
    req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh;Intel Mac OS X 10_8_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.95 Safari/537.36")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    return string(body)
}

func DecodeGoogleImage(page string){
    v, _ := regexp.Compile(`jpeg;base64,[A-Za-z0-9+/\\]+`)
    result := v.FindAllStringSubmatch(page,15)
    for i, v := range result{
        str := strings.Replace(v[0],`\x3d`,`=`,-1)
        data, err := base64.StdEncoding.DecodeString(str[12:])
        if err != nil {
            log.Fatal("error:", err)
        }

        err = ioutil.WriteFile(strconv.Itoa(i)+".jpeg", data, 0644)
    }
}

//facebook api
type AccessToken struct {
	Token string
	Expiry int64
}

func readHttpBody(response *http.Response) string {
	fmt.Println("Reading body")
	body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
	return string(body)
}

//Converts a code to an Auth_Token
func GetAccessToken(client_id string, code string, secret string, callbackUri string) AccessToken {
	fmt.Println("GetAccessToken")
	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
	response, err := http.Get("https://graph.facebook.com/oauth/access_token?grant_type=client_credentials&client_id=494717897285592&client_secret=30228f96eb5f6692d88bb41c184c6b59")
	if err == nil {
		auth := readHttpBody(response)
	    var token AccessToken
	    tokenArr := strings.Split(auth, "&")
	    token.Token = strings.Split(tokenArr[0], "=")[1]
	    expireInt,err := strconv.Atoi(strings.Split(tokenArr[1],"=")[1])
	    if (err == nil) {
		    token.Expiry = int64(expireInt)
		}
	    return token
	}
	var token AccessToken
	return token
}

