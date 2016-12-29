package service

import (
	"github.com/ereminIvan/tablebooking/model"
)

type IStorage interface {
	Get(string, interface{}) error
	Write(string, interface{}) error
	Create(string, interface{}) error
	Update(string, interface{}) error
	Delete(string) error
}

type ISource interface {
	//Events
	GetEvent(string) (model.Event, error)
	CreateEvent(model.Event) error
	UpdateEvent(model.Event) error
	DeleteEvent(model.Event) error
	DeleteEvents() error
	GetEvents() (model.Events, error)
	//Guests
	GetGuest(string) (model.Guest, error)
	CreateGuest(guest model.Guest, code string) error
	UpdateGuest(model.Guest) error
	DeleteGuest(model.Guest) error
	GetGuests() ([]model.Guest, error)
	GetGuestByCode(code string) (model.Guest, error)
}

type storageClient struct {
	storage IStorage
}

func NewStorage(s IStorage) ISource {
	return &storageClient{
		storage: s,
	}
}

func (s *storageClient) GetEvent(title string) (model.Event, error) {
	e := model.Event{}
	err := s.storage.Get("events/"+title, &e)
	return e, err
}

// CreateEvent Create single event with all attributes
func (s *storageClient) CreateEvent(e model.Event) error {
	return s.storage.Write("events/"+e.Title, e)
}

func (s *storageClient) UpdateEvent(e model.Event) error {
	return s.storage.Write("events/"+e.Title, e)
}

func (s *storageClient) DeleteEvents() error {
	return s.storage.Delete("events")
}

func (s *storageClient) DeleteEvent(e model.Event) error {
	return s.storage.Delete("events/" + e.Title)
}

func (s *storageClient) GetEvents() (model.Events, error) {
	es := model.Events{}
	err := s.storage.Get("events", &es)
	return es, err
}

func (s *storageClient) GetGuest(string) (model.Guest, error) {
	return model.Guest{}, nil
}

func (s *storageClient) GetGuestByCode(code string) (model.Guest, error) {
	e := model.Guest{}
	err := s.storage.Get("guests/"+code, &e)
	return e, err
}

func (s *storageClient) CreateGuest(g model.Guest, code string) error {
	return s.storage.Write("guests/"+code, g)
}

func (s *storageClient) UpdateGuest(g model.Guest) error {
	return nil
}

func (s *storageClient) DeleteGuest(g model.Guest) error {
	return nil
}

func (s *storageClient) GetGuests() ([]model.Guest, error) {
	return []model.Guest{}, nil
}
