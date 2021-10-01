package appconfig

/*
{
    "items": [
        {
          "name": "{key-name}"
        },
        ...
    ],
    "@nextLink": "{relative uri}"
}
*/
type GetKeysResponse struct {
}

func (a *Client) GetAllKeys() []string {
	resp := a.FetchKeys()
	keys := make([]string, 0)
	for _, key := range resp.Items {
		keys = append(keys, key.Name)
	}
	return keys
}

/*
{
  "etag": "4f6dd610dd5e4deebc7fbaef685fb903",
  "key": "{key}",
  "label": "{label}",
  "content_type": null,
  "value": "example value",
  "last_modified": "2017-12-05T02:41:26+00:00",
  "locked": "false",
  "tags": {
    "t1": "value1",
    "t2": "value2"
  }
}
*/

type KeyValue struct {
	Etag         string   `json:"etag"`
	Key          string   `json:"key"`
	Label        string   `json:"label"`
	ContentType  string   `json:"content_type"`
	Value        string   `json:"value"`
	LastModified string   `json:"last_modified"`
	Locked       bool     `json:"locked"`
	Tags         []string `json:"tags"`
}

func (a *Client) GetKeyValue(key string) (KeyValue, error) {

	return KeyValue{}, nil

}

func (a *Client) GetKeyValueWithLabel(key string, label string) (KeyValue, error) {

	return KeyValue{}, nil
}
