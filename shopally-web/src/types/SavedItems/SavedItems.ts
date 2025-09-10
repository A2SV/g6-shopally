export interface CreateAlert {
  productId: string;
  deviceId: string;
  productTitle: string;
  currentPriceETB: number;
}

export interface UpdateProductResponse {
  data: {
    updated_price_usd: number;
    updated_price_etb: number;
  } | null;
  error: {
    error: string | null;
  };
}
