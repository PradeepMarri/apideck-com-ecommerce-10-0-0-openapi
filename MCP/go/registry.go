package main

import (
	"github.com/ecommerce-api/mcp-server/config"
	"github.com/ecommerce-api/mcp-server/models"
	tools_stores "github.com/ecommerce-api/mcp-server/tools/stores"
	tools_customers "github.com/ecommerce-api/mcp-server/tools/customers"
	tools_orders "github.com/ecommerce-api/mcp-server/tools/orders"
	tools_products "github.com/ecommerce-api/mcp-server/tools/products"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_stores.CreateStoresoneTool(cfg),
		tools_customers.CreateCustomersallTool(cfg),
		tools_customers.CreateCustomersoneTool(cfg),
		tools_orders.CreateOrdersallTool(cfg),
		tools_orders.CreateOrdersoneTool(cfg),
		tools_products.CreateProductsallTool(cfg),
		tools_products.CreateProductsoneTool(cfg),
	}
}
