package postgres

import (
	"reflect"
	"testing"

	pb "github.com/baxromumarov/my-services/api-gateway/genproto"
)

func TestUserRepo_InsertAd(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Address
		want    pb.Address
		wantErr bool
	}{
		{
			name: "success case",
			input: pb.Address{
				City:       "Mosco",
				Country:    "Russi",
				District:   "Peterho",
				PostalCode: "1234",
			},
			want: pb.Address{
				City:       "Mosco",
				Country:    "Russi",
				District:   "Peterho",
				PostalCode: "1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.InsertAd(&tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: expected error %v, got %v", tt.name, tt.wantErr, err)
				// return
			}
			got.Id = ""
			if !reflect.DeepEqual(&tt.want, got) {
				t.Fatalf("%s: expected %v, got %v", tt.name, tt.want, got)
			}
		})
	}

}

func TestUserRepo_Insert(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.User
		want    pb.User
		wantErr bool
	}{
		{
			name: "success case",
			input: pb.User{
				FirstName:    "Baxrom",
				LastName:     "Umarov",
				Email:        "baxromumarov10@gmail.com",
				Bio:          "I am a developer",
				PhoneNumbers: []string{"+79998887766", "+79998887767", "+79998887768"},
			},
			want: pb.User{
				FirstName:    "Baxrom",
				LastName:     "Umarov",
				Email:        "baxromumarov10@gmail.com",
				Bio:          "I am a developer",
				PhoneNumbers: []string{"+79998887766", "+79998887767", "+79998887768"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Insert(&tt.input)
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
func TestUserRepo_GetById(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.ById
		want    pb.User
		wantErr bool
	}{
		{
			name: "success case",
			input: pb.ById{
				Id: "ff79dd5e-ea74-455f-8897-d271b9e1baa5",
			},
			want: pb.User{
				FirstName: "deserunt mollit in",
				LastName:  "consectetur est amet commodo dolor",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetById(&tt.input)
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
