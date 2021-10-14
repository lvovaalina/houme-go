package forge

type Project struct {
	BucketName string `json:"bucketName"`
	FileName   string `json:"fileName"`
	Urn        string `json:"urn"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

type Object struct {
	BucketKey string `json:"bucketKey"`
	ObjectKey string `json:"objectKey"`
}

type ObjectsResponse struct {
	Objects []Object `json:"items"`
}
