package woocommerce

import (
	"net/http"
	"strconv"
)

type ProductVariationService service

// ProductVariation object. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#product-properties
type ProductVariation struct {
	Id                int                  `json:"id,omitempty"`
	DateCreated       string               `json:"date_created,omitempty"`
	DateCreatedGmt    string               `json:"date_created_gmt,omitempty"`
	DateModified      string               `json:"date_modified,omitempty"`
	DateModifiedGmt   string               `json:"date_modified_gmt,omitempty"`
	Description       string               `json:"description,omitempty"`
	Permalink         string               `json:"permalink,omitempty"`
	Sku               string               `json:"sku,omitempty"`
	Price             string               `json:"price,omitempty"`
	RegularPrice      string               `json:"regular_price,omitempty"`
	SalePrice         string               `json:"sale_price,omitempty"`
	DateOnSaleFrom    string               `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGmt string               `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      string               `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGmt   string               `json:"date_on_sale_to_gmt,omitempty"`
	OnSale            bool                 `json:"on_sale,omitempty"`
	Status            string               `json:"status,omitempty"`
	Purchasable       bool                 `json:"purchasable,omitempty"`
	Virtual           bool                 `json:"virtual,omitempty"`
	Downloadable      bool                 `json:"downloadable,omitempty"`
	DownloadLimit     int                  `json:"download_limit,omitempty"`
	DownloadExpiry    int                  `json:"download_expiry,omitempty"`
	TaxStatus         string               `json:"tax_status,omitempty"`
	TaxClass          string               `json:"tax_class,omitempty"`
	ManageStock       bool                 `json:"manage_stock,omitempty"`
	StockQuantity     int                  `json:"stock_quantity,omitempty"`
	StockStatus       string               `json:"stock_status,omitempty"`
	Backorders        string               `json:"backorders,omitempty"`
	BackordersAllowed bool                 `json:"backorders_allowed,omitempty"`
	Backordered       bool                 `json:"backordered,omitempty"`
	Weight            string               `json:"weight,omitempty"`
	ShippingClass     string               `json:"shipping_class,omitempty"`
	ShippingClassId   string               `json:"shipping_class_id,omitempty"`
	MenuOrder         int                  `json:"menu_order,omitempty"`
	Image             *Image               `json:"image,omitempty"`
	Dimensions        *ProductDimensions   `json:"dimensions,omitempty"`
	Downloads         *[]ProductDownload   `json:"downloads,omitempty"`
	Attributes        *[]ProductAttributes `json:"attributes,omitempty"`
	MetaData          *[]MetaData          `json:"meta_data,omitempty"`
}

type ListProductVariationParams struct {
	Context       string `url:"context,omitempty"`
	Page          int    `url:"page,omitempty"`
	PerPage       int    `url:"per_page,omitempty"`
	Search        string `url:"search,omitempty"`
	After         string `url:"after,omitempty"`
	Before        string `url:"before,omitempty"`
	Exclude       *[]int `url:"exclude,omitempty"`
	Include       *[]int `url:"include,omitempty"`
	Offset        int    `url:"offset,omitempty"`
	Order         string `url:"order,omitempty"`
	OrderBy       string `url:"orderby,omitempty"`
	Parent        *[]int `url:"parent,omitempty"`
	ParentExclude *[]int `url:"parent_exclude,omitempty"`
	Slug          string `url:"slug,omitempty"`
	Status        string `url:"status,omitempty"`
	IncludeStatus string `url:"include_status,omitempty"`
	ExcludeStatus string `url:"exclude_status,omitempty"`
	Sku           string `url:"sku,omitempty"`
	TaxClass      string `url:"tax_class,omitempty"`
	OnSale        bool   `url:"on_sale,omitempty"`
	MinPrice      string `url:"min_price,omitempty"`
	MaxPrice      string `url:"max_price,omitempty"`
	StockStatus   string `url:"stock_status,omitempty"`
	Virtual       bool   `url:"virtual,omitempty"`
	Downloadable  bool   `url:"downloadable,omitempty"`
}

type DeleteProductVariationParams struct {
	Force string `json:"force,omitempty"`
}

type BatchProductVariationUpdate struct {
	Create *[]ProductVariation `json:"create,omitempty"`
	Update *[]ProductVariation `json:"update,omitempty"`
	Delete *[]int              `json:"delete,omitempty"`
}

type BatchProductVariationUpdateResponse struct {
	Create *[]ProductVariation `json:"create,omitempty"`
	Update *[]ProductVariation `json:"update,omitempty"`
	Delete *[]ProductVariation `json:"delete,omitempty"`
}

// Create a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-product
func (service *ProductVariationService) Create(productId int, variation *ProductVariation) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + strconv.Itoa(productId) + "/variations"
	req, _ := service.client.NewRequest("POST", _url, nil, variation)

	createdVariation := new(ProductVariation)
	response, err := service.client.Do(req, createdVariation)

	if err != nil {
		return nil, response, err
	}

	return createdVariation, response, nil
}

// Get a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-product
func (service *ProductVariationService) Get(productID int, variationID string) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + strconv.Itoa(productID) + "/variations/" + variationID
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	variation := new(ProductVariation)
	response, err := service.client.Do(req, variation)

	if err != nil {
		return nil, response, err
	}

	return variation, response, nil
}

// List products. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-products
func (service *ProductVariationService) List(productId int, opts *ListProductVariationParams) ([]ProductVariation, *http.Response, error) {
	_url := "/products/" + strconv.Itoa(productId) + "/variations"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	variations := new([]ProductVariation)
	response, err := service.client.Do(req, variations)

	if err != nil {
		return nil, response, err
	}

	return *variations, response, nil
}

// Update a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-product
func (service *ProductVariationService) Update(productID int, variation *ProductVariation) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + strconv.Itoa(productID) + "/variations/" + strconv.Itoa(variation.Id)
	req, _ := service.client.NewRequest("PUT", _url, nil, variation)

	updatedVariant := new(ProductVariation)
	response, err := service.client.Do(req, updatedVariant)

	if err != nil {
		return nil, response, err
	}

	return updatedVariant, response, nil
}

// Delete a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-product-variation
func (service *ProductVariationService) Delete(productId int, variantId int, opts *DeleteProductParams) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + strconv.Itoa(productId) + "/variations/" + strconv.Itoa(variantId)
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	variation := new(ProductVariation)
	response, err := service.client.Do(req, variation)

	if err != nil {
		return nil, response, err
	}

	return variation, response, nil
}

// Batch update products. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-products
func (service *ProductVariationService) Batch(productId int, opts *BatchProductVariationUpdate) (*BatchProductVariationUpdateResponse, *http.Response, error) {
	_url := "/products/" + strconv.Itoa(productId) + "/variations/batch"
	req, _ := service.client.NewRequest("POST", _url, nil, opts)

	variants := new(BatchProductVariationUpdateResponse)
	response, err := service.client.Do(req, variants)

	if err != nil {
		return nil, response, err
	}

	return variants, response, nil
}
