package mockingdemo

import (
	"golangtips/mockingdemo/mocks"
	"testing"
)

func TestFileSystem_PrintFileContent(t *testing.T) {
	type fields struct {
		conn FileConnector
	}

	connectorMock := mocks.FileConnector{}
	connectorMock.On("ReadFromFile", "/home/jc/go/src/Playground/mockingdemo/mockFileToread.txt").Return([]byte("salut"),nil)

	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test",
			fields{&connectorMock},
			args{fileName: "/home/jc/go/src/Playground/mockingdemo/mockFileToread.txt"},
			false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FileSystem{
				conn: tt.fields.conn,
			}
			if err := fs.PrintFileContent(tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("PrintFileContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
