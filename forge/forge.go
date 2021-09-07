package forge

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	baseUrl      = "https://developer.api.autodesk.com/"
	clentId      = "HjKvq1UCmrAOpHBp97m0lfgYEGEB7E2V"
	clientSecret = "1wuJ3vwrxbqnuxjP"
	grantType    = "client_credentials"
	scope        = "data:read data:write data:create bucket:create bucket:read bucket:delete"
	token        = ""
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func Authentificate() string {
	data := url.Values{}
	data.Set("client_id", clentId)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", grantType)
	data.Set("scope", scope)

	resp, err := http.PostForm(baseUrl+"authentication/v1/authenticate", data)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result AuthResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	//log.Println("TOKEN: " + result.AccessToken)

	return result.AccessToken
}

func GetAuthToken(force bool) string {
	if token == "" || force {
		token = Authentificate()
	}

	return token
}

func GetBuckets() {
	token := GetAuthToken(false)

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"oss/v2/buckets", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	//if (response.status == "")

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)
}

func DeleteBucket(bucketKey string) {
	token := GetAuthToken(false)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", baseUrl+"oss/v2/buckets/"+bucketKey, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)
}

func CreateBucket(bucketKey string) {
	var jsonData = []byte(`{
		"bucketKey":"` + bucketKey + `",
		"policyKey": "transient"
	}`)

	token := GetAuthToken(false)

	client := &http.Client{}

	req, err := http.NewRequest("POST", baseUrl+"oss/v2/buckets", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(resp.Status)
	log.Println(bodyString)
}

func GetBucketObjects(bucketKey string) {
	token := GetAuthToken(false)

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"oss/v2/buckets/"+bucketKey+"/details", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)
}

func UploadFileToBucket(bucketKey string, filePath string, fileName string) {
	token := GetAuthToken(false)

	binaryData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
		log.Println("err")
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"PUT", baseUrl+"oss/v2/buckets/"+bucketKey+"/objects/"+fileName,
		bytes.NewBuffer(binaryData))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "text/plain; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)
	log.Println(resp.Status)
}

func TranslateFile(bucketKey string, fileName string) {
	encodedUrn := getFileUrn(bucketKey, fileName)
	log.Println("ENCODED URN: ", encodedUrn)
	var jsonData = []byte(`{
		"input": {
			"urn": "` + encodedUrn + `"
		},
		"output": {
			"formats": [
				{
					"type": "svf",
					"views": ["2d", "3d"]
				}
			]
		}
	}`)

	token := GetAuthToken(false)

	client := &http.Client{}

	req, err := http.NewRequest("POST", baseUrl+"modelderivative/v2/designdata/job", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(resp.Status)
	log.Println(bodyString)
}

func GetTranslationStatus(bucketKey string, fileName string) {
	urn := getFileUrn(bucketKey, fileName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"modelderivative/v2/designdata/"+urn+"/manifest", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)
}

func getFileUrn(bucketKey string, fileName string) string {
	objectKey := "urn:adsk.objects:os.object:" + bucketKey + "/" + fileName
	log.Println("OBJECT KEY: ", objectKey)
	encodedUrn := base64.StdEncoding.EncodeToString([]byte(objectKey))
	return encodedUrn
}
