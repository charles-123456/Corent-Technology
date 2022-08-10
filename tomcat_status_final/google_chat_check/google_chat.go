package google_chat_check

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
	"corent-go/corent/log"
)

func GoogleNotification(data map[string]string,ChatSpaceName string,ServiceAccPath string) error {
	// log.Info("Entered GoogleNotification\n")
	// path := "D:/Golang Tutorial/tomcat_status/google_chat_check/service_account.json"
	// dir,_:= os.Getwd()
	// path := dir+"\\service_account.json"
	path := ServiceAccPath
	ctx := context.Background()
	// client := getOauthClient(path)
	// conf := &jwt.Config{
	// 	Email:        "demo-729@sincere-venture-355410.iam.gserviceaccount.com",
	// 	PrivateKeyID: "796be9185264dbe3b992edc1b880472c816c0cbe",
	// 	PrivateKey:   []byte("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC6RGpZxt4IyrCY\nxShxDa66gqJWj0KPmx4HK27iaUT00bhCqMuODnGn8w/pPZUirNmkJg2ivWNFNmOe\nBdMr48NPDCWvGoiANkMGO9wRH4e+MctKvmsm4WAguixGydGFtyvsPDafLrbwc8w7\nHm8sNUNJZ+r8bdDElKj5DbAGPxS6ECNEtopQwNmxrit5Zu5eJpz2Kzhp0l9SLXNJ\ncrI4G17PetHOTIwPy16U56ehTv7zlftxSUHBI2wdCHD4TG2tyUCmq0veC7BAWnNj\nZk6WmrC5nqmztft8in6NDcpmhRSaj3+JsVtJjKtwtdr1/N4T1btJ/q6gGQxzIFAk\nhcoO+f5hAgMBAAECggEAFLwr+J1+Nw1VNlRwSZGD22hSaP9eofzbDWJE4DhmwDge\nYnFG46Pw/Ai/QzAjS25hdff0eaLKi7hfw4YTeWXL5p9MDJ+QBXHE6Us9lrxjUIVb\ng5EJ/ZSGLm57IlAyHkgDQDN9wx+gpsjFOemL56hwOpoyWvJglJmV68+dwjxAcuvK\n8WmZiXxQUM+40aTinbArHNqN6erNbsDqhUQqfoA9CAxx1WYtXJTqDDj25+9Dxcs7\ncH9DZulOLOzwW4PZuJ9WTYOWRZ8bkK8gw0vAYSaXA/LXlsQm0toeEmH2L3HG20gT\nuJhz8rB7pKJXb0wgGfYsnJkStF8kykbMz/S6M9gFNwKBgQD4wNAQyurYy6PlmAih\n6Fadpm2qSAgK5ZGr5zK4LRWB5fDQNKnCUCRwLXGTMUlWS+1oAFkRGH5rIOLZz5Qv\nv5NDNfnArBi8EDQ0p5PZZ1oRUAOgn5iM+uOJX/6fx3YQzSJ6xj5YAmRrIg6d3gix\nl3XJw1KKX+hEVCsrQ74pU2m4AwKBgQC/sZYMxoei/jdFf8IM/0SgVtueOMjKhKIM\n5BOqk1t/zt8WHaUB1MRDFuBYhBZxMWMW3Pk/Mo8axKFX6SWKE7hetxsRNusjxxCz\nmqBXyFEOfSxgn+fUDqD6PxyoAGrjRZgk72BBb2P2gjOCTdCmCPQk3vPj6jmeZs0s\nXF2mbqxcywKBgQCkussO09ICD5lSCgRtc/coH3awNVNeI/j75fdokDKV+zgmKTni\nPEBlKTL1TsZKJ63oGZLiB15wgy63Hwf7NtrGv4/NUCpxICnyVKdMaWzz2hEM5aOY\neO0FpFRyaxx8s9wJgg73KV5ms/8J/Ge1c5/FJVwb1rdxyGtuE0ZzB5ITEQKBgQCM\nxd1rhGAXUplEcI4Q/WVoWmDt0MWj88MNtHC803peYY1ysFJ9BcMbgbE/T8ErXxll\nsOFZh8eP4NabuJvYyqKa69z0x1/m5kldnDAkRvc/rKzqSIP3NscA/1gMCEJ2pKUW\nerQ1WZgPb45kIsEXLXwdl52Dwn6N1PDXov0jPCNYAQKBgQDDoi65um+J/4HJZnfW\nEuJV8BzQTpCAN9ZGYIFcAiG8YdG6X2VZGMppX1SIMmav6MpAT5xrHXebIpPE7t8b\nWL/gUYtsP6TFDVEI6fMVYqLJQYM5ZHQME3D4bSbSH7/ekMrWamEVJKsDl3xZzJ1D\nnLluMr5ybyU6RdH/EZD3Ue8W0Q==\n-----END PRIVATE KEY-----\n"),
	// 	TokenURL:     "https://oauth2.googleapis.com/token",
	// 	Scopes: []string{
	// 		"https://www.googleapis.com/auth/chat.bot",
	// 		"https://www.googleapis.com/auth/chat.messages.create",
	// 		"https://www.googleapis.com/auth/chat.spaces.create",
	// 		"https://www.googleapis.com/auth/chat.memberships",
	// 		"https://www.googleapis.com/auth/chat.memberships.app",
	// 	},
	// }
	// client := conf.Client(ctx)
	// a,_ := client.Get("https://mail.google.com/chat/u/0/#chat/space/AAAAwlgqHZg")
	// fmt.Println("client created", a)
	// cards := `{
	// 	        "cards": [
	//             {
	// 	               "header": {
	// 	                   "title": "title",
	// 	                    "subtitle": "subtitle"
	//                 },
	// 	                "sections": [
	// 	                    {
	// 	                       "widgets": ["hello"]
	// 	                    }
	// 	                ]
	// 	            }
	// 	        ]
	// 	    }`
	// postBody, _ := json.Marshal(cards)
	// responseBody := bytes.NewBuffer(postBody)
	// resp, err := client.Post("https://chat.googleapis.com/v1/space/AAAAwlgqHZg/messages", "application/json", responseBody)
	// if err != nil {
	// 	log.Fatalf("An Error Occured %v", err)
	// }
	// fmt.Println("response", resp)
	// defer resp.Body.Close()
	service, err := chat.NewService(ctx, option.WithCredentialsFile(path), option.WithScopes("https://www.googleapis.com/auth/chat.messages.create", "https://www.googleapis.com/auth/chat.spaces.create", "https://www.googleapis.com/auth/chat.memberships", "https://www.googleapis.com/auth/chat.memberships.app"))
	// service, err := chat.NewService(client)
	if err != nil {
		log.Error(err)
	}
	// log.Info("Normal Service Created \n")
	msgService := chat.NewSpacesMessagesService(service)
	msg := ChatCard(data)
	// msg := "hello"
	// log.Info("Now ChatCard Method called\n")
	Spaces := fmt.Sprintf("spaces/%v",ChatSpaceName)
	_, err = msgService.Create(Spaces, msg).Do()
	if err != nil {
		log.Error(err)
	}
	// log.Info("Message Service Created!!! %v\n", err)
	return nil
}

func ChatCard(data map[string]string) *chat.Message {

	var widgets []*chat.WidgetMarkup
	for key, value := range data {
		widgets = append(widgets, &chat.WidgetMarkup{KeyValue: &chat.KeyValue{TopLabel: key, Content: value}})
	}
	fmt.Printf("All chat service widget created %v\n", widgets)
	cards := `{
       "cards": [
           {
               "sections": [
                   {
                       "widgets": []
                   }
               ]
           }
       ]
   }`
	// outputString := fmt.Sprintf(cards)
	// fmt.Printf("Output string created %v\n", cards)
	var message chat.Message
	json.Unmarshal([]byte(cards), &message)

	message.Cards[0].Sections[0].Widgets = widgets
	// fmt.Printf("Chat msg values widgets %v %v", widgets, &message)
	return &message
}

func StartingPoint(data map[string]string,ChatSpaceName string,ServiceAccPath string) {
	// log.Info("Starting GoogleNotification method Called!!!\n")
	GoogleNotification(map[string]string{"Charlie Says": data["data"]},ChatSpaceName,ServiceAccPath)
	// go func(){
	// 	log.Fatal(http.ListenAndServeTLS("linuxmigration.corenttechnology.com:777","E:\\SurpaasSetup/\\keys\\corenttechnology.pkcs12","E:\\SurpaasSetup\\keys\\paasswd.txt",nil))
	// }()
	
}