// src/types/Compare/Comparisons.ts
export interface Product {
  id: string;
  title: string;
  imageUrl: string;
  aiMatchPercentage: number;
  price: {
    etb: number;
    usd: number;
    fxTimestamp: string;
  };
  productRating?: number;
  sellerScore: number;
  deliveryEstimate: string;
  summaryBullets: string[];
  deeplinkUrl: string;

  // Fields you don’t send to backend
  description?: string;
  productSmallImageUrls?: string[];
  numberSold?: number;
  taxRate?: number;
  discount?: number;
} // ✅ reuse main Product

export interface ComparisonResponse {
  data: {
    products: ComparisonItem[];
    overallComparison: OverallComparison;
  };
  error?: string | null;
}

export interface ComparisonItem {
  product: Product;
  synthesis: Synthesis;
}

export interface Price {
  etb: number;
  usd: number;
  fxTimestamp: string; // ISO date string
}

export interface Synthesis {
  pros: string[];
  cons: string[];
  isBestValue: boolean;
  features: Record<string, string>;
}

export interface OverallComparison {
  bestValueProduct: string;
  bestValueLink: string;
  bestValuePrice: Price;
  keyHighlights: string[];
  summary: string;
}

export interface ComparePayload {
  products: Array<
    Pick<
      Product,
      | "id"
      | "title"
      | "imageUrl"
      | "aiMatchPercentage"
      | "price"
      | "productRating"
      | "sellerScore"
      | "deliveryEstimate"
      | "summaryBullets"
      | "deeplinkUrl"
    >
  >;
}
