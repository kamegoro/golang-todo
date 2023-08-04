package usecase

import (
	"errors"
	"reflect"
	"testing"

	"example.com/mod/domain"
	mockDomain "example.com/mod/mocks/domain"
)

func TestNewTodoUsecase(t *testing.T) {
	type args struct {
		tr domain.TodoRepository
	}
	mockTodoRepo := mockDomain.NewTodoRepository(t)
	tests := []struct {
		name string
		args args
		want domain.TodoUsecase
	}{
		{
			name: "NewTodoUsecase_Success",
			args: args{tr: mockTodoRepo},
			want: &todoUsecase{todoRepo: mockTodoRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoUsecase(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoUsecase_AllGet(t *testing.T) {
	type fields struct {
		todoRepo domain.TodoRepository
	}
	type testStruct struct {
		name    string
		fields  fields
		want    []domain.Todo
		wantErr bool
	}
	tests := []testStruct{}
	{
		AllGet_Error := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			mockTodoRepo.On("AllGet").Return(nil, errors.New("dummy error"))

			return testStruct{
				name:    "AllGet_Error",
				fields:  fields{todoRepo: mockTodoRepo},
				want:    nil,
				wantErr: true,
			}
		}

		AllGet_Success := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			want := []domain.Todo{
				{ID: 23, Text: "dummy todo", Completed: false},
				{ID: 24, Text: "dummy todo 2", Completed: true},
			}
			mockTodoRepo.On("AllGet").Return(want, nil)

			return testStruct{
				name:    "AllGet_Error",
				fields:  fields{todoRepo: mockTodoRepo},
				want:    want,
				wantErr: false,
			}
		}

		tests = append(tests, AllGet_Error())
		tests = append(tests, AllGet_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			got, err := tr.AllGet()
			if (err != nil) != tt.wantErr {
				t.Errorf("todoUsecase.AllGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoUsecase.AllGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoUsecase_StatusUpdate(t *testing.T) {
	type fields struct {
		todoRepo domain.TodoRepository
	}
	type args struct {
		id int
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		StatusUpdate_Error := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			id := 123
			mockTodoRepo.On("StatusUpdate", id).Return(errors.New("dummy error"))

			return testStruct{
				name:    "StatusUpdate_Error",
				fields:  fields{todoRepo: mockTodoRepo},
				args:    args{id: id},
				wantErr: true,
			}
		}

		StatusUpdate_Success := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			id := 123
			mockTodoRepo.On("StatusUpdate", id).Return(nil)

			return testStruct{
				name:    "StatusUpdate_Success",
				fields:  fields{todoRepo: mockTodoRepo},
				args:    args{id: id},
				wantErr: false,
			}
		}

		tests = append(tests, StatusUpdate_Error())
		tests = append(tests, StatusUpdate_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			if err := tr.StatusUpdate(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("todoUsecase.StatusUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoUsecase_Store(t *testing.T) {
	type fields struct {
		todoRepo domain.TodoRepository
	}
	type args struct {
		todo domain.Todo
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		Store_Error := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			todo := domain.Todo{ID: 1}
			mockTodoRepo.On("Store", todo).Return(errors.New("dummy error"))

			return testStruct{
				name:    "Store_Error",
				fields:  fields{todoRepo: mockTodoRepo},
				args:    args{todo: todo},
				wantErr: true,
			}
		}
		Store_Success := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			todo := domain.Todo{ID: 2}
			mockTodoRepo.On("Store", todo).Return(nil)

			return testStruct{
				name:    "Store_Success",
				fields:  fields{todoRepo: mockTodoRepo},
				args:    args{todo: todo},
				wantErr: false,
			}
		}

		tests = append(tests, Store_Error())
		tests = append(tests, Store_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			if err := tr.Store(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("todoUsecase.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoUsecase_Delete(t *testing.T) {
	type fields struct {
		todoRepo domain.TodoRepository
	}
	type args struct {
		id int
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		Delete_Error := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			id := 1
			mockTodoRepo.On("Delete", id).Return(errors.New("dummy error"))

			return testStruct{
				name:    "Delete_Error",
				fields:  fields{todoRepo: mockTodoRepo},
				args:    args{id: id},
				wantErr: true,
			}
		}
		Delete_Success := func() testStruct {
			mockTodoRepo := mockDomain.NewTodoRepository(t)
			id := 1
			mockTodoRepo.On("Delete", id).Return(nil)

			return testStruct{
				name:    "Delete_Success",
				fields:  fields{todoRepo: mockTodoRepo},
				args:    args{id: id},
				wantErr: false,
			}
		}

		tests = append(tests, Delete_Error())
		tests = append(tests, Delete_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			if err := tr.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("todoUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
