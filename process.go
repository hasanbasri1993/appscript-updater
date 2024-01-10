package main

import (
	"app-script-go/auth"
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/script/v1"
	"io/ioutil"
	"log"
	"time"
)

func process(scriptId string, count int) {
	b, err := ioutil.ReadFile("secret/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/script.projects")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := auth.GetClient(context.Background(), config)

	srv, err := script.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve script Client %v", err)
	}

	fmt.Println("Begin: ", count+1, scriptId)
	time.Sleep(1 * time.Second) // prevent race condition
	// Read the new script content from a file
	newScriptContent, err := ioutil.ReadFile("src/main.js")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	manifestContent, err := ioutil.ReadFile("src/appsscript.json")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	ctx := context.Background()

	// Create an update request
	request := &script.Content{
		ScriptId: scriptId,
		Files: []*script.File{
			{
				Name:   "main",
				Type:   "SERVER_JS",
				Source: string(newScriptContent),
			},

			{
				Name:   "appsscript",
				Type:   "JSON",
				Source: string(manifestContent),
			},
			// Include other files here as needed
		},
	}
	// Execute the update request
	_, err = srv.Projects.UpdateContent(scriptId, request).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to update script: %v", err)
	}
	//fmt.Printf("Project created: %s\n", createOp.ScriptId)
}
