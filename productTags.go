package woocommerce

import "net/http"

type ProductTagService service

type ProductTag struct {
	Id          int             `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Slug        string          `json:"slug,omitempty"`
	Count       int             `json:"count,omitempty"`
	Links       ProductTagLinks `json:"_links,omitempty"`
}

type ProductTagLinks struct {
	Self       []Link `json:"self,omitempty"`
	Collection []Link `json:"collection,omitempty"`
}

type Link struct {
	Url string `json:"href,omitempty"`
}

type BatchProductTagsUpdate struct {
	Create *[]ProductTag `json:"create,omitempty"`
	Update *[]ProductTag `json:"update,omitempty"`
	Delete *[]int        `json:"delete,omitempty"`
}

type BatchProductTagsUpdateResponse struct {
	Create *[]ProductTag `json:"create,omitempty"`
	Update *[]ProductTag `json:"update,omitempty"`
	Delete *[]ProductTag `json:"delete,omitempty"`
}

// Create a product tag. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-product-tag
func (service *ProductTagService) Create(productTag *ProductTag) (*ProductTag, *http.Response, error) {
	req, err := service.client.NewRequest("POST", "/products/tags", nil, productTag)
	if err != nil {
		return nil, nil, err
	}

	createdProductTag := new(ProductTag)
	response, err := service.client.Do(req, createdProductTag)

	if err != nil {
		return nil, response, err
	}

	return createdProductTag, response, nil
}

// Get a product tag. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-product-tag
func (service *ProductTagService) Get(productTagID string) (*ProductTag, *http.Response, error) {
	req, err := service.client.NewRequest("GET", "/product/tags/"+productTagID, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	productTag := new(ProductTag)
	response, err := service.client.Do(req, productTag)

	if err != nil {
		return nil, response, err
	}

	return productTag, response, nil
}

// List all product tags. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-product-tags
func (service *ProductTagService) List(opts *ListProductParams) (*[]ProductTag, *http.Response, error) {
	req, err := service.client.NewRequest("GET", "/products/tags", opts, nil)
	if err != nil {
		return nil, nil, err
	}

	productTags := new([]ProductTag)
	response, err := service.client.Do(req, productTags)

	if err != nil {
		return nil, response, err
	}

	return productTags, response, nil
}

// Update a product tag. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-product-tag
func (service *ProductTagService) Update(productTagID string, product *ProductTag) (*ProductTag, *http.Response, error) {
	_url := "/products/tags/" + productTagID
	req, _ := service.client.NewRequest("PUT", _url, nil, product)

	updatedProductTag := new(ProductTag)
	response, err := service.client.Do(req, updatedProductTag)

	if err != nil {
		return nil, response, err
	}

	return updatedProductTag, response, nil
}

// Delete a product tag. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-product-tag
func (service *ProductTagService) Delete(productTagID string) (*ProductTag, *http.Response, error) {
	_url := "/products/tags/" + productTagID + "?force=true" // Force must be set
	req, _ := service.client.NewRequest("DELETE", _url, nil, nil)

	productTag := new(ProductTag)
	response, err := service.client.Do(req, productTag)

	if err != nil {
		return nil, response, err
	}

	return productTag, response, nil
}

// Batch update product tags. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-product-tags
func (service *ProductTagService) Batch(opts *BatchProductTagsUpdate) (*BatchProductTagsUpdateResponse, *http.Response, error) {
	_url := "/products/tags/batch"
	req, _ := service.client.NewRequest("POST", _url, nil, opts)

	productTags := new(BatchProductTagsUpdateResponse)
	response, err := service.client.Do(req, productTags)

	if err != nil {
		return nil, response, err
	}

	return productTags, response, nil
}
