// src/utils/Compare/Mappers.ts
import { ComparisonItem } from "@/types/Compare/Comparison";

export function mapComparisonToUIProducts(items: ComparisonItem[]) {
  return items.map((item) => ({
    id: item.product.id,
    name: item.product.title, // old components use "name"
    image: item.product.imageUrl,
    price: item.product.price.usd, // keep USD
    rating: item.product.productRating,
    reviews: Math.round(item.product.sellerScore * 100), // fake "reviews"
    features: item.product.summaryBullets,
    specs: {
      warranty: item.product.deliveryEstimate, // re-use delivery as warranty
    },
    // keep synthesis for AIRecommendation
    pros: item.synthesis.pros,
    cons: item.synthesis.cons,
    aiScore: item.product.aiMatchPercentage,
  }));
}
