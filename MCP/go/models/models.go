package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CreateEcommerceOrderResponse represents the CreateEcommerceOrderResponse schema from the OpenAPI specification
type CreateEcommerceOrderResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// EcommerceAddress represents the EcommerceAddress schema from the OpenAPI specification
type EcommerceAddress struct {
	Country string `json:"country,omitempty"` // Country of the billing address.
	Line1 string `json:"line1,omitempty"` // Address line 1 of the billing address.
	Line2 string `json:"line2,omitempty"` // Address line 2 of the billing address.
	Postal_code string `json:"postal_code,omitempty"` // Postal/ZIP code of the billing address.
	State string `json:"state,omitempty"` // State/province of the billing address.
	City string `json:"city,omitempty"` // City of the billing address.
	Company_name string `json:"company_name,omitempty"` // Company name of the customer
}

// CreateEcommerceCustomerResponse represents the CreateEcommerceCustomerResponse schema from the OpenAPI specification
type CreateEcommerceCustomerResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// GetEcommerceCustomerResponse represents the GetEcommerceCustomerResponse schema from the OpenAPI specification
type GetEcommerceCustomerResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data EcommerceCustomer `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"` // Email address
	Id string `json:"id,omitempty"` // Unique identifier for the email address
	TypeField string `json:"type,omitempty"` // Email type
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// EcommerceOrder represents the EcommerceOrder schema from the OpenAPI specification
type EcommerceOrder struct {
	Tracking []TrackingItem `json:"tracking,omitempty"`
	Billing_address EcommerceAddress `json:"billing_address,omitempty"` // An object representing a shipping or billing address.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Customer LinkedEcommerceCustomer `json:"customer,omitempty"` // The customer this entity is linked to.
	Discounts []EcommerceDiscount `json:"discounts,omitempty"`
	Order_number string `json:"order_number,omitempty"` // Order number, if any.
	Sub_total string `json:"sub_total,omitempty"` // Sub-total amount, normally before tax.
	Total_tax string `json:"total_tax,omitempty"` // Total tax, if any.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Status string `json:"status,omitempty"` // Current status of the order.
	Payment_status string `json:"payment_status,omitempty"` // Current payment status of the order.
	Shipping_address EcommerceAddress `json:"shipping_address,omitempty"` // An object representing a shipping or billing address.
	Total_amount string `json:"total_amount,omitempty"` // Total amount due.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Total_discount string `json:"total_discount,omitempty"` // Total discount, if any.
	Fulfillment_status string `json:"fulfillment_status,omitempty"` // Current fulfillment status of the order.
	Id string `json:"id"` // A unique identifier for an object.
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for this order.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Shipping_cost string `json:"shipping_cost,omitempty"` // Shipping cost, if any.
	Line_items []EcommerceOrderLineItem `json:"line_items,omitempty"`
	Note string `json:"note,omitempty"` // Note for the order.
}

// EcommerceOrderLineItem represents the EcommerceOrderLineItem schema from the OpenAPI specification
type EcommerceOrderLineItem struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Options []interface{} `json:"options,omitempty"`
	Sku string `json:"sku,omitempty"` // The SKU of the product or variant associated with the line item.
	Total_amount string `json:"total_amount"` // The total amount for the product(s) or variant associated with the line item, including taxes and discounts.
	Description string `json:"description,omitempty"` // The description of the product or variant associated with the line item.
	Tax_rate string `json:"tax_rate,omitempty"` // The tax rate applied to the product or variant associated with the line item.
	Variant_id string `json:"variant_id,omitempty"` // A unique identifier for the variant of the product associated with the line item, if applicable.
	Name string `json:"name"` // The name of the product or variant associated with the line item.
	Quantity string `json:"quantity"` // The quantity of the product or variant associated with the line item.
	Tax_amount string `json:"tax_amount,omitempty"` // The total tax amount applied to the product or variant associated with the line item.
	Product_id string `json:"product_id,omitempty"` // A unique identifier for the product associated with the line item.
	Unit_price string `json:"unit_price,omitempty"` // The unit price of the product or variant associated with the line item.
	Discounts []EcommerceDiscount `json:"discounts,omitempty"`
}

// Website represents the Website schema from the OpenAPI specification
type Website struct {
	Id string `json:"id,omitempty"` // Unique identifier for the website
	TypeField string `json:"type,omitempty"` // The type of website
	Url string `json:"url"` // The website URL
}

// LinkedEcommerceCustomer represents the LinkedEcommerceCustomer schema from the OpenAPI specification
type LinkedEcommerceCustomer struct {
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Company_name string `json:"company_name,omitempty"` // Company name of the customer
	Emails []Email `json:"emails,omitempty"`
	First_name string `json:"first_name,omitempty"` // First name of the customer
	Id string `json:"id"` // The ID of the customer this entity is linked to.
	Last_name string `json:"last_name,omitempty"` // Last name of the customer
	Name string `json:"name,omitempty"` // Full name of the customer
}

// PassThroughQuery represents the PassThroughQuery schema from the OpenAPI specification
type PassThroughQuery struct {
	Example_downstream_property string `json:"example_downstream_property,omitempty"` // All passthrough query parameters are passed along to the connector as is (?pass_through[search]=leads becomes ?search=leads)
}

// LinkedEcommerceOrder represents the LinkedEcommerceOrder schema from the OpenAPI specification
type LinkedEcommerceOrder struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Status string `json:"status,omitempty"` // Current status of the order.
	Total string `json:"total,omitempty"` // The total amount of the order.
}

// DeleteProductResponse represents the DeleteProductResponse schema from the OpenAPI specification
type DeleteProductResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// UpdateEcommerceOrderResponse represents the UpdateEcommerceOrderResponse schema from the OpenAPI specification
type UpdateEcommerceOrderResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// UpdateProductResponse represents the UpdateProductResponse schema from the OpenAPI specification
type UpdateProductResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	TypeField string `json:"type,omitempty"` // The type of phone number
	Area_code string `json:"area_code,omitempty"` // The area code of the phone number, e.g. 323
	Country_code string `json:"country_code,omitempty"` // The country code of the phone number, e.g. +1
	Extension string `json:"extension,omitempty"` // The extension of the phone number
	Id string `json:"id,omitempty"` // Unique identifier of the phone number
	Number string `json:"number"` // The phone number
}

// EcommerceOrdersFilter represents the EcommerceOrdersFilter schema from the OpenAPI specification
type EcommerceOrdersFilter struct {
	Email string `json:"email,omitempty"` // Customer email address to filter on
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// Department represents the Department schema from the OpenAPI specification
type Department struct {
	Code string `json:"code,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"` // Department name
	Parent_id string `json:"parent_id,omitempty"` // Parent ID
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// GetProductsResponse represents the GetProductsResponse schema from the OpenAPI specification
type GetProductsResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []EcommerceProduct `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetProductResponse represents the GetProductResponse schema from the OpenAPI specification
type GetProductResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data EcommerceProduct `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// UpdateEcommerceCustomerResponse represents the UpdateEcommerceCustomerResponse schema from the OpenAPI specification
type UpdateEcommerceCustomerResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// EcommerceStore represents the EcommerceStore schema from the OpenAPI specification
type EcommerceStore struct {
	Admin_url string `json:"admin_url,omitempty"` // The store's admin login URL
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Id string `json:"id"` // A unique identifier for an object.
	Name string `json:"name,omitempty"` // The store's name
	Store_url string `json:"store_url,omitempty"` // The store's website URL
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// EcommerceCustomersFilter represents the EcommerceCustomersFilter schema from the OpenAPI specification
type EcommerceCustomersFilter struct {
	Email string `json:"email,omitempty"` // Customer email address to filter on
	Phone_number string `json:"phone_number,omitempty"` // Customer phone number to filter on
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Name string `json:"name,omitempty"` // The name of the address.
	Phone_number string `json:"phone_number,omitempty"` // Phone number of the address
	Longitude string `json:"longitude,omitempty"` // Longitude of the address
	Street_number string `json:"street_number,omitempty"` // Street number
	Email string `json:"email,omitempty"` // Email address of the address
	StringField string `json:"string,omitempty"` // The address string. Some APIs don't provide structured address data.
	Fax string `json:"fax,omitempty"` // Fax number of the address
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Line2 string `json:"line2,omitempty"` // Line 2 of the address
	State string `json:"state,omitempty"` // Name of state
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	Notes string `json:"notes,omitempty"` // Additional notes
	County string `json:"county,omitempty"` // Address field that holds a sublocality, such as a county
	Contact_name string `json:"contact_name,omitempty"` // Name of the contact person at the address
	City string `json:"city,omitempty"` // Name of city.
	Id string `json:"id,omitempty"` // Unique identifier for the address.
	Line3 string `json:"line3,omitempty"` // Line 3 of the address
	Latitude string `json:"latitude,omitempty"` // Latitude of the address
	Line1 string `json:"line1,omitempty"` // Line 1 of the address e.g. number, street, suite, apt #, etc.
	Line4 string `json:"line4,omitempty"` // Line 4 of the address
	Website string `json:"website,omitempty"` // Website of the address
	Salutation string `json:"salutation,omitempty"` // Salutation of the contact person at the address
	Postal_code string `json:"postal_code,omitempty"` // Zip code or equivalent.
	TypeField string `json:"type,omitempty"` // The type of address.
}

// EcommerceCustomer represents the EcommerceCustomer schema from the OpenAPI specification
type EcommerceCustomer struct {
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"` // An array of phone numbers for the customer.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Id string `json:"id"` // A unique identifier for an object.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Last_name string `json:"last_name,omitempty"` // Last name of the customer
	Name string `json:"name,omitempty"` // Full name of the customer
	Status string `json:"status,omitempty"` // The current status of the customer
	Addresses []map[string]interface{} `json:"addresses,omitempty"` // An array of addresses for the customer.
	Company_name string `json:"company_name,omitempty"` // Company name of the customer
	Emails []Email `json:"emails,omitempty"` // An array of email addresses for the customer.
	Orders []LinkedEcommerceOrder `json:"orders,omitempty"`
	First_name string `json:"first_name,omitempty"` // First name of the customer
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// GetEcommerceCustomersResponse represents the GetEcommerceCustomersResponse schema from the OpenAPI specification
type GetEcommerceCustomersResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []EcommerceCustomer `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// DeleteEcommerceCustomerResponse represents the DeleteEcommerceCustomerResponse schema from the OpenAPI specification
type DeleteEcommerceCustomerResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// GetEcommerceOrderResponse represents the GetEcommerceOrderResponse schema from the OpenAPI specification
type GetEcommerceOrderResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data EcommerceOrder `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetEcommerceOrdersResponse represents the GetEcommerceOrdersResponse schema from the OpenAPI specification
type GetEcommerceOrdersResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []EcommerceOrder `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// DeleteEcommerceOrderResponse represents the DeleteEcommerceOrderResponse schema from the OpenAPI specification
type DeleteEcommerceOrderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// EcommerceProduct represents the EcommerceProduct schema from the OpenAPI specification
type EcommerceProduct struct {
	Status string `json:"status,omitempty"` // The current status of the product (active or archived).
	Weight_unit string `json:"weight_unit,omitempty"` // The unit of measurement for the weight of the product.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Description string `json:"description,omitempty"` // A detailed description of the product.
	Categories []map[string]interface{} `json:"categories,omitempty"` // An array of categories for the product, used for organization and searching.
	Id string `json:"id"` // A unique identifier for an object.
	Images []map[string]interface{} `json:"images,omitempty"` // An array of image URLs for the product.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Name string `json:"name,omitempty"` // The name of the product as it should be displayed to customers.
	Price string `json:"price,omitempty"` // The price of the product.
	Weight string `json:"weight,omitempty"` // The weight of the product.
	Options []map[string]interface{} `json:"options,omitempty"` // An array of options for the product.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Variants []map[string]interface{} `json:"variants,omitempty"`
	Sku string `json:"sku,omitempty"` // The stock keeping unit of the product.
	Inventory_quantity string `json:"inventory_quantity,omitempty"` // The quantity of the product in stock.
	Tags []string `json:"tags,omitempty"` // An array of tags for the product, used for organization and searching.
}

// GetStoreResponse represents the GetStoreResponse schema from the OpenAPI specification
type GetStoreResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data EcommerceStore `json:"data"`
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CreateProductResponse represents the CreateProductResponse schema from the OpenAPI specification
type CreateProductResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
}

// GetStoresResponse represents the GetStoresResponse schema from the OpenAPI specification
type GetStoresResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []EcommerceStore `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// EcommerceDiscount represents the EcommerceDiscount schema from the OpenAPI specification
type EcommerceDiscount struct {
	Amount string `json:"amount,omitempty"` // The fixed amount of the discount.
	Code string `json:"code,omitempty"` // The code used to apply the discount.
	Percentage string `json:"percentage,omitempty"` // The percentage of the discount.
}

// TrackingItem represents the TrackingItem schema from the OpenAPI specification
type TrackingItem struct {
	Number string `json:"number"` // The tracking number associated with the shipment, which can be used to track the progress of the delivery.
	Provider string `json:"provider"` // The name or code of the carrier or shipping company that is handling the shipment.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Url string `json:"url,omitempty"` // The URL of the carrier's tracking page, which can be used to view detailed information about the shipment's progress.
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Id string `json:"id"` // Unique identifier for the custom field.
	Name string `json:"name,omitempty"` // Name of the custom field.
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
}
