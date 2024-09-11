package management

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type TaxonomyResponse struct {
	Taxonomy Taxonomy `json:"taxonomy"`
}

type TaxonomyRequest struct {
	Taxonomy TaxonomyInput `json:"taxonomy"`
}

// Taxonomy represents the taxonomy in contentstack.
type Taxonomy struct {
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	Name             	string              `json:"name,omitempty"`
	UID               string              `json:"uid,omitempty"`
	Description       string              `json:"description"`
}

// TaxonomyInput is used to create or update a taxonomy
type TaxonomyInput struct {
	Name       *string         `json:"name,omitempty"`
	UID         *string         `json:"uid,omitempty"`
	Description *string         `json:"description,omitempty"`
}

func (si *StackInstance) TaxonomyCreate(ctx context.Context, input TaxonomyInput) (*Taxonomy, error) {
	data, err := serializeInput(TaxonomyRequest{Taxonomy: input})
	if err != nil {
		return nil, err
	}

	resp, err := si.client.post(
		ctx,
		"/v3/taxonomies/",
		url.Values{},
		si.headers(),
		data,
	)
	if err != nil {
		return nil, err
	}

	result := &TaxonomyResponse{}
	if err = si.client.processResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result.Taxonomy, nil
}

func (si *StackInstance) TaxonomyUpdate(ctx context.Context, uid string, input TaxonomyInput) (*Taxonomy, error) {
	data, err := serializeInput(TaxonomyRequest{Taxonomy: input})
	if err != nil {
		return nil, err
	}

	resp, err := si.client.put(
		ctx,
		fmt.Sprintf("/v3/taxonomies/%s", uid),
		url.Values{},
		si.headers(),
		data,
	)
	if err != nil {
		return nil, err
	}

	result := &TaxonomyResponse{}
	if err = si.client.processResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result.Taxonomy, nil
}

func (si *StackInstance) TaxonomyDelete(ctx context.Context, uid string) error {
	resp, err := si.client.delete(
		ctx,
		fmt.Sprintf("/v3/taxonomies/%s", uid),
		url.Values{},
		si.headers(),
		nil,
	)

	if err != nil {
		return err
	}

	result := &TaxonomyResponse{}
	if err = si.client.processResponse(resp, &result); err != nil {
		return err
	}

	return nil
}

func (si *StackInstance) TaxonomyFetch(ctx context.Context, uid string) (*Taxonomy, error) {
	resp, err := si.client.get(
		ctx,
		fmt.Sprintf("/v3/taxonomies/%s", uid),
		url.Values{},
		si.headers(),
	)
	if err != nil {
		return nil, err
	}

	result := &TaxonomyResponse{}
	if err = si.client.processResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result.Taxonomy, nil
}

func (si *StackInstance) TaxonomyFetchAll(ctx context.Context) ([]Taxonomy, error) {
	resp, err := si.client.get(
		ctx,
		"/v3/taxonomies",
		url.Values{},
		si.headers(),
	)
	if err != nil {
		return nil, err
	}

	result := struct {
		Taxonomys []Taxonomy `json:"taxonomies"`
	}{}

	if err = si.client.processResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Taxonomys, nil
}
