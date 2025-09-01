// components/compare/AIRecommendation.tsx
import type { ComparisonItem } from "@/types/Compare/Comparison";
import { Check, X } from "lucide-react";
import { BsLightningChargeFill } from "react-icons/bs";
import { FaCircleCheck } from "react-icons/fa6";
import { IoIosStar } from "react-icons/io";
import { TiDelete } from "react-icons/ti";

type Analysis = {
  product: ComparisonItem["product"];
  score: number; // 0 - 100
  pros: string[];
  cons: string[];
};

export const AIRecommendation: React.FC<{ comparison: ComparisonItem[] }> = ({
  comparison,
}) => {
  if (!comparison || comparison.length === 0) return null;

  // Pick top product by productRating, tie-breaker: sellerScore
  const recommendedProduct = [...comparison].sort(
    (a, b) =>
      b.product.productRating - a.product.productRating ||
      b.product.sellerScore - a.product.sellerScore,
  )[0].product;

  // Build analysis cards (static scores & pros/cons for demo)
  const productAnalysis: Analysis[] = comparison.map((c) => ({
    product: c.product,
    score: c.product.aiMatchPercentage,
    pros: c.synthesis.pros,
    cons: c.synthesis.cons,
  }));

  return (
    <div className="bg-white rounded-2xl border border-gray-200 shadow-lg p-6 space-y-6">
      <div className="flex items-center gap-4">
        <div className="w-8 h-8 rounded-full bg-brand-saffron flex items-center justify-center">
          <IoIosStar className="w-5 h-5 text-[#090c11]" />
        </div>
        <h2 className="text-xl font-semibold text-[#090c11]">
          AI Recommendation
        </h2>
      </div>

      <div className="bg-gray-50 rounded-xl p-6">
        <div className="flex items-start gap-4 mb-6">
          <div className="w-12 h-12 rounded-full bg-[#ffd300] flex items-center justify-center">
            <BsLightningChargeFill className="w-6 h-6 text-[#090c11]" />
          </div>
          <div className="flex-1">
            <div className="flex items-center gap-2 mb-2">
              <span className="font-semibold text-[#090c11]">
                Our Top Pick:
              </span>
              <span className="font-semibold text-[#090c11]">
                {recommendedProduct.title}
              </span>
            </div>
            <p className="text-sm text-gray-600 mb-4">
              Based on price, ratings, and versatility, this product offers a
              strong balance of value and everyday utility.
            </p>
            <div className="flex items-center gap-4">
              <span className="text-sm text-[#ffd300] font-medium">
                ${recommendedProduct.price.usd}
              </span>
              <div className="flex items-center gap-1">
                <span className="text-sm font-medium text-[#090c11]">
                  {recommendedProduct.productRating}
                </span>
                <IoIosStar className="w-4 h-4 fill-[#090c11]" />
                <span className="text-sm text-gray-600">
                  ({recommendedProduct.sellerScore})
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {productAnalysis.map((analysis) => (
          <div key={analysis.product.id} className="space-y-4">
            <h3 className="font-semibold text-[#090c11]">
              {analysis.product.title}
            </h3>

            <div className="space-y-3">
              <div className="flex items-center gap-2">
                <FaCircleCheck className="w-4 h-4 text-[#090c11]" />
                <span className="text-sm font-medium text-[#090c11]">Pros</span>
              </div>
              {analysis.pros.map((pro, i) => (
                <div key={i} className="flex items-center gap-2 ml-6">
                  <Check className="w-3 h-3 text-[#090c11]" />
                  <span className="text-sm text-gray-600">{pro}</span>
                </div>
              ))}
            </div>

            <div className="space-y-3">
              <div className="flex items-center gap-2">
                <TiDelete className="w-4 h-4 text-[#090c11]" />
                <span className="text-sm font-medium text-[#090c11]">Cons</span>
              </div>
              {analysis.cons.map((con, i) => (
                <div key={i} className="flex items-center gap-2 ml-6">
                  <X className="w-3 h-3 text-red-500" />
                  <span className="text-sm text-gray-600">{con}</span>
                </div>
              ))}
            </div>

            <div className="space-y-2">
              <span className="text-xs font-medium text-[#090c11]">
                AI Score
              </span>
              <div className="w-full bg-gray-200 rounded-full h-2">
                <div
                  className="h-2 rounded-full bg-[#ffd300]"
                  style={{ width: `${analysis.score}%` }}
                />
              </div>
              <span className="text-xs text-gray-600">
                {analysis.score}% match
              </span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
