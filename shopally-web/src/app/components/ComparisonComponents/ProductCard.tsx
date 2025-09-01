// //src/app/components/ComparePage/ProductCard.tsx
import { Product } from "@/types/Compare/Comparison";
import { FaHeart } from "react-icons/fa";
import { FaCheck } from "react-icons/fa6";
import { IoIosStar } from "react-icons/io";

export const ProductCard: React.FC<{ product: Product }> = ({ product }) => {
  return (
    <div className="bg-white rounded-2xl border border-gray-200 shadow p-6 space-y-6 w-full max-w-md mx-auto lg:max-w-none">
      <div className="aspect-square bg-gray-50 rounded-xl overflow-hidden">
        <img
          src={product.imageUrl}
          alt={product.title}
          className="w-full h-full object-cover"
        />
      </div>
      <div className="space-y-4">
        <h3 className="text-lg font-semibold text-[#090c11]">
          {product.title}
        </h3>
        <div className="flex items-center justify-between">
          <span className="text-2xl font-bold text-[#ffd300]">
            ${product.price.usd}
          </span>
          <div className="flex items-center gap-1 text-sm text-gray-600">
            <IoIosStar className="w-4 h-4 fill-[#090c11]" />
            <span>{product.productRating} (#)</span>
          </div>
        </div>
        <div className="space-y-3">
          <h4 className="text-sm font-medium text-[#090c11]">Key Features:</h4>
          {product.summaryBullets.map((feature, i) => (
            <div key={i} className="flex items-center gap-2">
              <FaCheck className="w-4 h-4 text-[#090c11]" />
              <span className="text-sm text-gray-600">{feature}</span>
            </div>
          ))}
        </div>
        <div className="flex gap-3">
          <button className="flex-1 bg-[#ffd300] text-[#090c11] font-medium py-3 px-6 rounded-xl hover:bg-opacity-90">
            Add to Cart
          </button>
          <button className="p-3 bg-gray-50 rounded-xl hover:bg-gray-100">
            <FaHeart className="w-5 h-5 text-[#090c11]" />
          </button>
        </div>
      </div>
    </div>
  );
};
