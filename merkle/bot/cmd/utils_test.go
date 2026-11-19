package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/slack-go/slack"
)

func TestUploadFileToSlack(t *testing.T) {
	for _, apiError := range []string{"", "channel_not_found"} {
		name := apiError
		if name == "" {
			name = "success"
		}
		t.Run(name, func(t *testing.T) {
			const content = "address,amount\n0x123,42\n"
			var completed atomic.Bool
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/files.getUploadURLExternal":
					if r.FormValue("filename") != "airdrop.csv" || r.FormValue("length") != fmt.Sprint(len(content)) {
						t.Error("upload request did not preserve the filename and size")
					}
					fmt.Fprintf(w, `{"ok":true,"upload_url":%q,"file_id":"F123"}`, "http://"+r.Host+"/upload")
				case "/upload":
					file, header, err := r.FormFile("file")
					if err != nil {
						t.Errorf("read uploaded file: %v", err)
						http.Error(w, "invalid upload", http.StatusBadRequest)
						return
					}
					defer file.Close()
					data, err := io.ReadAll(file)
					if err != nil || string(data) != content || header.Filename != "airdrop.csv" {
						t.Errorf("unexpected upload: filename=%q content=%q error=%v", header.Filename, data, err)
					}
					fmt.Fprint(w, `{"ok":true}`)
				case "/files.completeUploadExternal":
					completed.Store(true)
					for key, want := range map[string]string{
						"channel_id":      "C123",
						"thread_ts":       "123.456",
						"initial_comment": "Rewards ready",
					} {
						if got := r.FormValue(key); got != want {
							t.Errorf("%s = %q, want %q", key, got, want)
						}
					}
					var files []slack.FileSummary
					if err := json.Unmarshal([]byte(r.FormValue("files")), &files); err != nil {
						t.Errorf("decode file metadata: %v", err)
					} else if len(files) != 1 || files[0].ID != "F123" || files[0].Title != "Airdrop" {
						t.Errorf("unexpected file metadata: %+v", files)
					}
					if apiError != "" {
						fmt.Fprintf(w, `{"ok":false,"error":%q}`, apiError)
						return
					}
					fmt.Fprint(w, `{"ok":true,"files":[{"id":"F123","title":"Airdrop"}]}`)
				default:
					t.Errorf("unexpected Slack request: %s", r.URL.Path)
					http.NotFound(w, r)
				}
			}))
			defer server.Close()
			ctx := CommandContext{
				Client: slack.New("test-token", slack.OptionAPIURL(server.URL+"/")),
			}
			err := uploadFileToSlack(ctx, "C123", "123.456", "airdrop.csv", "Airdrop", "Rewards ready", bytes.NewBufferString(content))
			if apiError == "" && err != nil {
				t.Fatalf("upload failed: %v", err)
			}
			if apiError != "" {
				var slackError slack.SlackErrorResponse
				if !errors.As(err, &slackError) || slackError.Err != apiError {
					t.Fatalf("error = %v, want %s", err, apiError)
				}
			}
			if !completed.Load() {
				t.Error("upload was not completed")
			}
		})
	}
}
