// components/compare/ComparisonTable.tsx
import type { ComparisonItem } from "@/types/Compare/Comparison";
import { FC } from "react";
import { IoIosStar } from "react-icons/io";

interface ComparisonTableProps {
  comparison: ComparisonItem[];
}

export const ComparisonTable: FC<ComparisonTableProps> = ({ comparison }) => {
  if (!comparison || comparison.length === 0) return null;

  return (
    <div className="overflow-x-auto bg-white rounded-2xl border border-gray-200 shadow-lg">
      <div className="p-5">
        <h1 className="text-[#090c11] font-bold">Detailed Comparison</h1>
      </div>
      <table className="w-full text-left border-collapse min-w-[600px]">
        <thead>
          <tr className="bg-white">
            <th className="px-4 py-3 text-sm font-semibold text-[#262b32]">
              Feature
            </th>
            {comparison.map((c) => (
              <th
                key={c.product.id}
                className="px-4 py-3 text-sm font-semibold text-[#262b32] text-center"
              >
                {c.product.title}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="divide-y divide-gray-200">
          {/* Price */}
          <tr>
            <td className="px-4 py-3 text-sm font-semibold text-[#262b32]">
              Price (USD)
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-4 py-3 text-center text-gray-600"
              >
                ${c.product.price.usd}
              </td>
            ))}
          </tr>

          {/* Rating */}
          <tr>
            <td className="px-4 py-3 text-sm font-semibold text-[#262b32]">
              Rating
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-4 py-3 text-center text-gray-600"
              >
                <div className="flex items-center justify-center gap-1">
                  {c.product.productRating}
                  <IoIosStar className="w-4 h-4 fill-[#090c11]" />
                </div>
              </td>
            ))}
          </tr>

          {/* Delivery Estimate */}
          <tr>
            <td className="px-4 py-3 text-sm font-semibold text-[#262b32]">
              Delivery Estimate
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-4 py-3 text-center text-gray-600"
              >
                {c.product.deliveryEstimate}
              </td>
            ))}
          </tr>

          {/* Seller Score */}
          <tr>
            <td className="px-4 py-3 text-sm font-semibold text-[#262b32]">
              Seller Score
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-4 py-3 text-center text-gray-600"
              >
                {c.product.sellerScore}
              </td>
            ))}
          </tr>
        </tbody>
      </table>
    </div>
  );
};
