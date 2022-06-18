package postgres

import (
	pb "github.com/baxromumarov/my-services/post-service/genproto"
	"reflect"
	"testing"
)

func TestPostRepo_CreatePost(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Post
		want    pb.Post
		wantErr bool
	}{
		{
			name: "success case",
			input: pb.Post{
				Name: "test",
				Medias: []*pb.Media{
					{
						Type: "image",
						Link: "https://www.google.com/",
					},
					{
						Type: "video",
						Link: "https://www.youtube.com/",
					},
				},
			},
			want: pb.Post{
				Name: "test",
				Medias: []*pb.Media{
					{
						Type: "image",
						Link: "https://www.google.com/",
					},
					{
						Type: "video",
						Link: "https://www.youtube.com/",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CreatePost(&tt.input)
			if err != nil {
				t.Errorf("%s: expected error %v, got %v", tt.name, tt.wantErr, err)
			}
			got.Id = ""

			if !reflect.DeepEqual(&tt.want, got) {
				t.Fatalf("%s: expected %v, got %v", tt.name, tt.want, got)
			}
		})
	}
}

func TestPostRepo_GetByIdPost(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    pb.Post
		wantErr bool
	}{
		{
			name:  "success case",
			input: "242c4d1f-c171-4823-93b0-b38229195c5d",

			want: pb.Post{
				Name:      "test",
				CreatedAt: "2022-06-11 20:06:42.655835223+05:00",
				Medias: []*pb.Media{
					{
						Id:   "84c7b30a-60a9-4033-99da-35beed23dd28",
						Type: "image",
						Link: "https://www.google.com/",
					},
					{
						Id:   "32111db9-2e9f-4f69-9c07-76c584a4f8d8",
						Type: "video",
						Link: "https://www.youtube.com/",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByIdPost(tt.input)
			if err != nil {
				t.Errorf("%s: expected error %v, got %v", tt.name, tt.wantErr, err)
			}
			got.Id = ""
			got.UserId = ""

			if !reflect.DeepEqual(&tt.want, got) {
				t.Fatalf("%s: expected %v, got %v", tt.name, tt.want, got)
			}
		})
	}
}
