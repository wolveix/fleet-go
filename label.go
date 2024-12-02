package fleet

import (
	"net/http"
	"strconv"
	"time"
)

type Label struct {
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Query               string    `json:"query"`
	LabelType           string    `json:"label_type"`
	LabelMembershipType string    `json:"label_membership_type"`
	HostCount           int       `json:"host_count"`
	DisplayText         string    `json:"display_text"`
	Count               int       `json:"count"`
	HostIDs             []int     `json:"host_ids"`
}

// AddLabelsToHost adds the given labels to the corresponding host.
func (s *Service) AddLabelsToHost(hostID int, labels ...string) error {
	req := struct {
		Labels []string `json:"labels"`
	}{Labels: labels}

	if _, err := s.makeRequest(http.MethodPost, "hosts/"+strconv.Itoa(hostID)+"/labels", req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteLabelsFromHost deletes the given labels from the corresponding host.
func (s *Service) DeleteLabelsFromHost(hostID int, labels ...string) error {
	req := struct {
		Labels []string `json:"labels"`
	}{Labels: labels}

	if _, err := s.makeRequest(http.MethodDelete, "hosts/"+strconv.Itoa(hostID)+"/labels", req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteLabel deletes the label with the corresponding ID.
func (s *Service) DeleteLabel(id int) error {
	if _, err := s.makeRequest(http.MethodDelete, "labels/id/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// DeleteLabelByName deletes the label with the corresponding name.
func (s *Service) DeleteLabelByName(name string) error {
	if _, err := s.makeRequest(http.MethodDelete, "labels/"+name, nil, nil); err != nil {
		return err
	}

	return nil
}

// FindLabels returns all labels from the Fleet API.
func (s *Service) FindLabels() ([]*Label, error) {
	resp := struct {
		Labels []*Label `json:"labels"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "labels", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Labels, nil
}
