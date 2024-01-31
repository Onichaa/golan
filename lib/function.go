package lib


import (
"fmt"
  "os"
   "bytes"
   "mime/multipart"
  "io"
   "path/filepath"
  "io/ioutil"
  "encoding/json"
   "math/rand"
  "time"
  "net/http"
   "os/exec"
  "net/url"
  "strings"
"regexp"
 "github.com/PuerkitoBio/goquery"
)
/*
func Stiktele(query string) (map[string]interface{}, error) {
  resp, err := http.Get(fmt.Sprintf("https://getstickerpack.com/stickers?query=%s", query))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  doc, err := goquery.NewDocumentFromReader(resp.Body)
  if err !=nil {
    return nil, err
  }

  links := []string{}
  doc.Find("#stickerPacks > div > div:nth-child(3) > div > a").Each(func(i int, s *goquery.Selection) {
    link, _ := s.Attr("href")
    links = append(links, link)
  })

  randLink := links[rand.Intn(len(links))]
  resp2, err := http.Get(randLink)
  if err != nil {
    return nil, err
  }
  defer resp2.Body.Close()

  doc2, err := goquery.NewDocumentFromReader(resp2.Body)
  if err != nil {
    return nil, err
  }

  urls := []string{}
  doc2.Find("#stickerPack > div > div.row > div > img").Each(func(i int, s *goquery.Selection) {
    url, _ := s.Attr("src")
    urls = append(urls, strings.Split(url, "&d=")[0])
  })

  return map[string]interface{}{
    "creator": "Fajar Ihsana",
    "title":   doc2.Find("#intro > div > div > h1").Text(),
    "author":  doc2.Find("#intro > div > div > h5 > a").Text(),
    "author_link": doc2.Find("#intro > div > div > h5 > a").AttrOr("href", ""),
    "sticker": urls,
  }, nil
}
*/

type Result struct {
  Desc   string `csv:"desc"`
  Thumb  string `csv:"thumb"`
  HD     string `csv:"video_hd"`
  SD     string `csv:"video_sd"`
  URL    string `csv:"url"`
  Locale string `csv:"locale"`
}

func Fb(URL string) (*Result, error) {
  data := URL + "?locale=en"
  resp, err := http.PostForm("https://getmyfb.com/process", url.Values{"id": {data}})
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  doc, err := goquery.NewDocumentFromReader(resp.Body)
  if err != nil {
    return nil, err
  }

  thumb := doc.Find(".results-item img").AttrOr("src", "")
  desc := doc.Find(".results-item > .results-item-text").Text()
  hd := doc.Find(".results-download > .results-list > .results-list-item a").Eq(0).AttrOr("href", "")
  sd := doc.Find(".results-download > .results-list > .results-list-item a").Eq(1).AttrOr("href", "")

  result := &Result{
    Desc:   desc,
    Thumb:  thumb,
    HD:     hd,
    SD:     sd,
    URL:    URL,
    Locale: "en",
  }

  return result, nil
}

func Instagram(URL string) ([]map[string]string, error) {
  client := &http.Client{}
  resp, err := client.Get("https://indown.io/")
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  bodyBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  body := string(bodyBytes)
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
  if err != nil {
    return nil, err
  }

  referer := doc.Find("input[name=referer]").AttrOr("value", "")
  locale := doc.Find("input[name=locale]").AttrOr("value", "")
  token := doc.Find("input[name=_token]").AttrOr("value", "")

	fmt.Println(locale, referer, token)
	
  params := url.Values{}
  params.Set("link", URL)
  params.Set("referer", referer)
  params.Set("locale", locale)
  params.Set("_token", token)

  req, err := http.NewRequest("POST", "https://indown.io/download", strings.NewReader(params.Encode()))
  if err != nil {
    return nil, err
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", fmt.Sprint(len(params.Encode())))
  req.Header.Add("Cookie", strings.Join(resp.Header["Set-Cookie"], "; "))

  resp, err = client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  bodyBytes, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  body = string(bodyBytes)
  doc, err = goquery.NewDocumentFromReader(strings.NewReader(body))
  if err != nil {
    return nil, err
  }

  result := []map[string]string{}
  doc.Find("#result video").Each(func(i int, s *goquery.Selection) {
    thumbnail, _ := s.Attr("poster")
    videoURL, _ := s.Find("source").Attr("src")

    result = append(result, map[string]string{
      "type":      "video",
      "thumbnail": thumbnail,
      "url":       videoURL,
    })
  })

  doc.Find("#result img").Each(func(i int, s *goquery.Selection) {
    imageURL, _ := s.Attr("src")

    result = append(result, map[string]string{
      "type": "image",
      "url":  imageURL,
    })
  })

  return result, nil
}


func Pixiv(query string) (string, error) {
    urlParsed, err := url.Parse("https://www.pixiv.net/ajax/search/artworks/"+ query)
  if err != nil {
    return "", err
  }

    req, err := http.NewRequest("GET", urlParsed.String() , nil)
    println(urlParsed.String())
    req.Header.Set("Content-Type", "application/json")

  if err != nil {
    return "", err
  }

    client := &http.Client{}
    resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
    defer resp.Body.Close()

    bodyText, err := io.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
    var result map[string]interface{}
    err = json.Unmarshal(bodyText, &result)
  if err != nil {
    return "", err
  }
    body, ok := result["body"].(map[string]interface{})
    if !ok {
        return "", err
    }

    novel, ok := body["illustManga"].(map[string]interface{})

    data, ok := novel["data"].([]interface{})
    if !ok {
        return "", err
    }
    if len(data) == 0 {
        return "No data found", err
    }

    dataElement, ok := data[rand.Intn(len(data))].(map[string]interface{})
    if !ok {
        return "", err
    }
    title, ok := dataElement["title"].(string)
    if !ok {
        return "", err
    }
    tagsInterface, ok := dataElement["tags"].([]interface{})
    if !ok {
        // handle the error
    }

    tagsString := make([]string, len(tagsInterface))
    for i, v := range tagsInterface {
        tagsString[i], ok = v.(string)
        if !ok {
            // handle the error
        }
    }

    res := "Pixiv > "+ query + "\nTitle: "+ title + "\nAlternatives Title: " + dataElement["alt"].(string) + "\nTags: " + strings.Join(tagsString, ", ") + "\n" + dataElement["url"].(string)


  return res, err
}

func Capcutdl(url string) ([]byte, error) {
  resp, err := http.Get(url)
if err != nil {
  return nil, err
}
  defer resp.Body.Close()


  re := regexp.MustCompile(`\d+`)
  token := re.FindString(strings.Split(resp.Request.URL.String(), "?")[0])

  if token == "" {
    return nil, nil
  }

  downloadURL := fmt.Sprintf("https://ssscap.net/api/download/%s", token)
  req, err := http.NewRequest(http.MethodGet, downloadURL, nil)
if err != nil {
  return nil, err
}

  req.Header.Set("Cookie", "sign=2cbe441f7f5f4bdb8e99907172f65a42; device-time=1685437999515")

  client := &http.Client{}
  downResp, err := client.Do(req)
if err != nil {
  return nil, err
}
  defer downResp.Body.Close()

  data, err := ioutil.ReadAll(downResp.Body)
if err != nil {
  return nil, err
}

 return data, nil
}


var c = time.Now()
func GetUptime() time.Time{
  return c
}

func ShortUrl(query string) (string, error) {
  client := &http.Client{}
  data := url.Values{}
  data.Set("url", query)

  req, err := http.NewRequest("POST", "https://tinyurl.com/api-create.php", strings.NewReader(data.Encode()))
  if err != nil {
    return "", err
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}


func ToAudio(buffer []byte, ext string) ([]byte, error) {
  tmpFile := fmt.Sprintf("%d.%s", time.Now().UnixNano(), ext)
  outFile := fmt.Sprintf("%d.%s.mp3", time.Now().UnixNano(), ext)

  err := ioutil.WriteFile(tmpFile, buffer, 0644)
  if err != nil {
    return nil, err
  }
  defer os.Remove(tmpFile)

  cmd := exec.Command("ffmpeg",
    "-y",
    "-i", tmpFile,
    "-vn",
    "-ac", "2",
    "-b:a", "128k",
    "-ar", "44100",
    "-f", "mp3",
    outFile,
  )
  var stderr bytes.Buffer
  cmd.Stderr = &stderr
  err = cmd.Run()
  if err != nil {
    return nil, fmt.Errorf("error running ffmpeg: %s\n%s", err, stderr.String())
  }
  defer os.Remove(outFile)

  audioData, err := ioutil.ReadFile(outFile)
  if err != nil {
    return nil, err
  }

  return audioData, nil
}

func GetRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  rand.Seed(time.Now().UnixNano())

  b := make([]byte, length)
  for i := range b {
    b[i] = charset[rand.Intn(len(charset))]
  }
  return string(b)
}

func Upload(mediaPath string) (string, error) {
    if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
        return "", fmt.Errorf("File not found")
    }

    media, err := os.Open(mediaPath)
    if err != nil {
        return "", err
    }
    defer media.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    part, err := writer.CreateFormFile("files[]", filepath.Base(mediaPath))
    if err != nil {
        return "", err
    }

    _, err = io.Copy(part, media)
    if err != nil {
        return "", err
    }

    err = writer.Close()
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", "https://pomf.lain.la/upload.php", body)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    // Parse the response body
    type ResponseData struct {
        Files []struct {
            URL string `json:"url"`
        } `json:"files"`
    }
    var data ResponseData
    err = json.Unmarshal(responseData, &data)
    if err != nil {
        return "", err
    }

    if len(data.Files) == 0 {
        return "", fmt.Errorf("Failed to retrieve file URL")
    }

    return data.Files[0].URL, nil
}

func Blackbox(query string) (string, error) {
	url := "https://www.blackbox.ai/api/chat"
	userId := "7a492784-ba58-4b52-aa3b-14a2a9cdd0a9"
	userInput := query
	cookies := "sessionId=2ed66013-1238-4b3d-8569-2e385720f97c; g_state={\"i_l\":0}; __Host-next-auth.csrf-token=49832ab2932a2dfa1698e676bd02574f468b97068b28226b1a39dafc6840d415|19eced3245bd787c855a757a04ca9fab1c536497dde3aa3e1de287e9e2804f31; __Secure-next-auth.callback-url=https://www.blackbox.ai; __Secure-next-auth.session-token=eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..sVjspDVopruvcAuT.cCyydm_HRZsJ00-Rdft1YrsUKZOL7YQ_RpFNTGCt6l8jUiBcfAkc-KEhl51zrtAY3b1gDUVpS3crCLWsJZU1U3Vcz8v6-7rAsE077XOfINEiu8AyGDPsv_0dQdTov9C58J3--iSj8ZKfoxvTnIvlkYcMOCqAAC7RXK6Jixqp2W1MBLtaS7fe2HbgnZeI_y0HW06IzAPS7_kCoLBKvgclQ1d24xowFweRpU73v4rGuizrkmSzRAgsh06Q-jcMBWFAinNRyj04UjNy-aXzqfLAGqIpARSntDYQAzPmMnrqmMYfFO6R9jdVcmXrQ3qV4CXip0bwzSSjpPpOP9xHFo53CPGNADOog0vuESmo-7aOJO5YKrfcEw.6lMTen--sdyMsaIqcZP2MQ"

	payload := []byte(fmt.Sprintf(`{
		"messages": [{
			"id": "EQbmkyx",
			"content": "%s",
			"role": "user"
		}],
		"previewToken": null,
		"userId": "%s"
	}`, userInput, userId))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookies)
	req.Header.Set("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 12; M2010J19SG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.52 Mobile Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func UploadV2(mediaPath string) (string, error) {
    if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
        return "", fmt.Errorf("File not found")
    }

    media, err := os.Open(mediaPath)
    if err != nil {
        return "", err
    }
    defer media.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    part, err := writer.CreateFormFile("files[]", filepath.Base(mediaPath))
    if err != nil {
        return "", err
    }

    _, err = io.Copy(part, media)
    if err != nil {
        return "", err
    }

    err = writer.Close()
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", "https://pomf.lain.la/upload.php", body)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    // Parse the response body
    type ResponseData struct {
        Files []struct {
            URL string `json:"url"`
        } `json:"files"`
    }
    var data ResponseData
    err = json.Unmarshal(responseData, &data)
    if err != nil {
        return "", err
    }

    if len(data.Files) == 0 {
        return "", fmt.Errorf("Failed to retrieve file URL")
    }

    return data.Files[0].URL, nil
}
