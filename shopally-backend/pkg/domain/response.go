package domain

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type AliProduct struct {
	AppSalePrice        string `json:"app_sale_price"`
	OriginalPrice       string `json:"original_price"`
	ProductDetailURL    string `json:"product_detail_url"`
	Discount            string `json:"discount"`
	ProductMainImageURL string `json:"product_main_image_url"`
	TaxRate             string `json:"tax_rate"`
	ProductID           int64  `json:"product_id"`
	ShipToDays          string `json:"ship_to_days"`
	EvaluateRate        string `json:"evaluate_rate"`
	SalePrice           string `json:"sale_price"`
	ProductTitle        string `json:"product_title"`

	TargetSalePrice            string `json:"target_sale_price"`
	TargetAppSalePrice         string `json:"target_app_sale_price"`
	TargetSalePriceCurrency    string `json:"target_sale_price_currency"`
	TargetAppSalePriceCurrency string `json:"target_app_sale_price_currency"`

	ProductSmallImageURLs struct {
		String []string `json:"string"`
	} `json:"product_small_image_urls"`
	SecondLevelCategoryName     string `json:"second_level_category_name"`
	FirstLevelCategoryName      string `json:"first_level_category_name"`
	OriginalPriceCurrency       string `json:"original_price_currency"`
	TargetOriginalPriceCurrency string `json:"target_original_price_currency"`
	TargetOriginalPrice         string `json:"target_original_price"`
	LastestVolume               int    `json:"lastest_volume"`
	SalePriceCurrency           string `json:"sale_price_currency"`
}

type SgResp struct {
	AliexpressResp struct {
		RespResult struct {
			Result struct {
				CurrentRecordCount int `json:"current_record_count"`
				TotalRecordCount   int `json:"total_record_count"`
				CurrentPageNo      int `json:"current_page_no"`
				Products           struct {
					Product []AliProduct `json:"product"`
				} `json:"products"`
			} `json:"result"`
		} `json:"resp_result"`
	} `json:"aliexpress_affiliate_product_query_response"`
}

type GeminiRequest struct {
	Contents        []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}
