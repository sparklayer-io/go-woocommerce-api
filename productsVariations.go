package woocommerce

import (
	"net/http"
)

type ProductVariantsService service

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
	Downloads         *[]ProductDownloads  `json:"downloads,omitempty"`
	Attributes        *[]ProductAttributes `json:"attributes,omitempty"`
	MetaData          *[]MetaData          `json:"meta_data,omitempty"`
}

type DeleteProductVariantParams struct {
	Force    string `json:"force,omitempty"`
	Reassign int    `json:"reassign,omitempty"`
}

type BatchProductVariantUpdate struct {
	Create *[]Product `json:"create,omitempty"`
	Update *[]Product `json:"update,omitempty"`
	Delete *[]int     `json:"delete,omitempty"`
}

type BatchProductVariantUpdateResponse struct {
	Create *[]Product `json:"create,omitempty"`
	Update *[]Product `json:"update,omitempty"`
	Delete *[]Product `json:"delete,omitempty"`
}

type ProductDownload struct {
	DownloadId         string `json:"download_id,omitempty"`
	DownloadUrl        string `json:"download_url,omitempty"`
	ProductId          int    `json:"product_id,omitempty"`
	ProductName        string `json:"product_name,omitempty"`
	DownloadName       string `json:"download_name,omitempty"`
	OrderId            int    `json:"order_id,omitempty"`
	OrderKey           string `json:"order_key,omitempty"`
	DownloadsRemaining string `json:"downloads_remaining,omitempty"`
	AccessExpires      string `json:"access_expires,omitempty"`
	AccessExpiresGmt   string `json:"access_expires_gmt,omitempty"`
	File               *File  `json:"file,omitempty"`
	Links              *Links `json:"_links,omitempty"`
}

// Create a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-product
func (service *ProductVariantsService) Create(product *Product) (*Product, *http.Response, error) {
	_url := "/products"
	req, _ := service.client.NewRequest("POST", _url, nil, product)

	createdProduct := new(Product)
	response, err := service.client.Do(req, createdProduct)

	if err != nil {
		return nil, response, err
	}

	return createdProduct, response, nil
}

// Get a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-product
func (service *ProductVariantsService) Get(productID string, variationID string) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + productID + "/variations/" + variationID
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	variation := new(ProductVariation)
	response, err := service.client.Do(req, variation)

	if err != nil {
		return nil, response, err
	}

	return variation, response, nil
}

// Get a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-product
func (service *ProductVariantsService) GetAll(productID string) ([]ProductVariation, *http.Response, error) {
	_url := "/products/" + productID + "/variations"
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	var variations []ProductVariation
	//variations := new(Product)
	response, err := service.client.Do(req, variations)

	if err != nil {
		return nil, response, err
	}

	return variations, response, nil
}

// List products. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-products
func (service *ProductVariantsService) List(opts *ListProductParams) (*[]ProductVariation, *http.Response, error) {
	_url := "/products"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	products := new([]ProductVariation)
	response, err := service.client.Do(req, products)

	if err != nil {
		return nil, response, err
	}

	return products, response, nil
}

// Update a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-product
func (service *ProductVariantsService) Update(productID string, product *Product) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + productID
	req, _ := service.client.NewRequest("PUT", _url, nil, product)

	updatedVariant := new(ProductVariation)
	response, err := service.client.Do(req, updatedVariant)

	if err != nil {
		return nil, response, err
	}

	return updatedVariant, response, nil
}

// Delete a product. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-product
func (service *ProductVariantsService) Delete(productID string, opts *DeleteProductParams) (*ProductVariation, *http.Response, error) {
	_url := "/products/" + productID
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	product := new(ProductVariation)
	response, err := service.client.Do(req, product)

	if err != nil {
		return nil, response, err
	}

	return product, response, nil
}

// Batch update products. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-products
func (service *ProductVariantsService) Batch(opts *BatchProductUpdate) (*BatchProductVariantUpdateResponse, *http.Response, error) {
	_url := "/products/batch"
	req, _ := service.client.NewRequest("POST", _url, nil, opts)

	products := new(BatchProductVariantUpdateResponse)
	response, err := service.client.Do(req, products)

	if err != nil {
		return nil, response, err
	}

	return products, response, nil
}
