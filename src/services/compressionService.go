package services

import (
	"errors"
	"fmt"

	"github.com/IngAngelMartinez/go-api-multitenant/src/models"
	"github.com/IngAngelMartinez/go-api-multitenant/src/repository"
)

type DocumentsService struct {
	DocumentsRepository *repository.DocumentsRepository
}

func (s *DocumentsService) GetAll() ([]models.Document, error) {

	documents, err := s.DocumentsRepository.GetAll()

	if err != nil {

		return nil, err
	}

	return documents, nil
}

func (s *DocumentsService) GetById(id string) (models.Document, error) {

	document, err := s.DocumentsRepository.GetById(id)

	if err != nil {

		return models.Document{}, err
	}

	return document, nil
}

func (s *DocumentsService) Create(document models.Document) (string, error) {

	id, err := s.DocumentsRepository.Create(document)

	if err != nil {

		return "", err
	}

	return id, nil
}

func (s *DocumentsService) Update(id string, document models.Document) error {

	modifiedCount, err := s.DocumentsRepository.Update(id, document)

	if err != nil {

		return err
	}

	if modifiedCount == 0 {

		return errors.New(fmt.Sprintf("Not found documents with id: %s", id))
	}

	return nil
}

func (s *DocumentsService) Delete(id string) error {

	deletedCount, err := s.DocumentsRepository.Delete(id)

	if err != nil {

		return err
	}

	if deletedCount == 0 {

		return errors.New(fmt.Sprintf("Not found documents with id: %s", id))
	}

	return nil
}
