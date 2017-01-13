package service

import "github.com/ereminIvan/tablebooking/dto"

// IStorage interface for getting access to storage
type IStorage interface {
	Get(string, interface{}) error
	Write(string, interface{}) error
	Create(string, interface{}) error
	Update(string, interface{}) error
	Delete(string) error
}

// ISource interface for getting access to data source
type ISource interface {
	//Events
	GetEvent(id string) (dto.Event, error)
	CreateEvent(dto.Event) (error, string)
	UpdateEvent(dto.Event) error
	DeleteEvent(id string) error
	DeleteEvents() error
	GetEvents() (dto.Events, error)
	//Events Tables
	CreateEventTables(eventId string, tables dto.Tables) error
	//Guests
	CreateGuest(dto.Guest, dto.Event) error
	UpdateGuest(dto.Guest, dto.Event) error
	DeleteGuest(dto.Guest, dto.Event) error
	GetGuests(dto.Event) ([]dto.Guest, error)
	GetGuestByCode(string, dto.Event) (dto.Guest, error)
}

type storageClient struct {
	storage     IStorage
	randService IRandom
}

// NewStorage get new storage service
func NewStorage(s IStorage, r IRandom) ISource {
	return &storageClient{
		storage:     s,
		randService: r,
	}
}

// GetEvent Get event by title
func (s *storageClient) GetEvent(id string) (dto.Event, error) {
	e := dto.Event{}
	err := s.storage.Get("events/"+id, &e)
	return e, err
}

// CreateEvent Create single event with all attributes
func (s *storageClient) CreateEvent(e dto.Event) (error, string) {
	id := s.randService.Runes(15)
	return s.storage.Write("events/"+id, e), id
}

func (s *storageClient) CreateEventTables(eventId string, tables dto.Tables) error {
	return s.storage.Write("events/"+eventId+"/tables", tables)
}

// UpdateEvent Update given event
func (s *storageClient) UpdateEvent(e dto.Event) error {
	return s.storage.Write("events/"+e.Title, e)
}

// DeleteEvent Delete specified event
func (s *storageClient) DeleteEvent(id string) error {
	return s.storage.Delete("events/" + id)
}

// DeleteEvents Delete all events
func (s *storageClient) DeleteEvents() error {
	return s.storage.Delete("events")
}

// GetEvents Get all events
func (s *storageClient) GetEvents() (dto.Events, error) {
	es := dto.Events{}
	err := s.storage.Get("events", &es)
	return es, err
}

// GetGuestByCode Get guest by code and event
func (s *storageClient) GetGuestByCode(code string, e dto.Event) (dto.Guest, error) {
	g := dto.Guest{}
	err := s.storage.Get("events/"+e.Title+"/guests/"+code, &g)
	return g, err
}

// CreateGuest Create guest in event
func (s *storageClient) CreateGuest(g dto.Guest, e dto.Event) error {
	return s.storage.Write("events/"+e.Title+"/guests/"+g.Code, g)
}

// UpdateGuest update guest of given event
func (s *storageClient) UpdateGuest(g dto.Guest, e dto.Event) error {
	return s.storage.Update("events/"+e.Title+"/guests/"+g.Code, g)
}

// DeleteGuest delete guest of given event
func (s *storageClient) DeleteGuest(g dto.Guest, e dto.Event) error {
	s.storage.Delete("events/" + e.Title + "/guests/" + g.Code)
	return nil
}

// GetGuest get all guests of give event
func (s *storageClient) GetGuests(e dto.Event) ([]dto.Guest, error) {
	g := &[]dto.Guest{}
	err := s.storage.Get("events/"+e.Title+"/guests", g)
	return *g, err
}
