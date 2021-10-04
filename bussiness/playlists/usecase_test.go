package playlists_test

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/playlists/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

var playlistsRepo mocks.Repository
var playlistDomain playlists.Domain
var playlistUseCase playlists.Usecase

func testSetup() {
	playlistDomain = playlists.Domain{
		Id:     1,
		Name:   "Happy",
		UserID: 1,
	}
	playlistUseCase = playlists.NewPlaylistUsecase(&playlistsRepo)
}

func TestInsert(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Valid Data", func(t *testing.T) {
		playlistsRepo.On("Insert", mock.Anything).Return(playlistDomain, nil).Once()
		playlistUseCase.Insert(playlistDomain)
	})
}

func TestGet(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Valid Data", func(t *testing.T) {
		playlistsRepo.On("Get", mock.Anything).Return([]playlists.Domain{}, nil).Once()
		playlistUseCase.Get(playlistDomain)
	})
}

func TestDelete(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Valid Data", func(t *testing.T) {
		playlistsRepo.On("Delete", mock.Anything).Return(playlistDomain, nil).Once()
		playlistUseCase.Delete(playlistDomain)
	})
}
