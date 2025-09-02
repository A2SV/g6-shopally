//src/types/Compare/Comparisons.ts
export interface ComparisonResponse {
  data: {
    comparison: ComparisonItem[];
  };
  error?: string | null;
}

export interface ComparisonItem {
  product: Product;
  synthesis: Synthesis;
}

export interface Product {
  id: string;
  title: string;
  imageUrl: string;
  aiMatchPercentage: number;
  price: {
    etb: number;
    usd: number;
    fxTimestamp: string; // ISO date string
  };
  productRating: number;
  sellerScore: number;
  deliveryEstimate: string;
  summaryBullets: string[];
  deeplinkUrl: string;
}

export interface Synthesis {
  pros: string[];
  cons: string[];
  isBestValue: boolean;
  features?: Record<string, string>; // flexible key-value for features
}
