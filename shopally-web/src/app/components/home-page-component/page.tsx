import { useSavedItems } from "@/hooks/useSavedItems";
import { Product } from "@/types/types";
import { Star } from "lucide-react";
import React, { useEffect, useState } from "react";

interface CardComponentProps {
  mode: "dark" | "light";
  product: Product;
  onClick?: () => void; // ✅ added
}

const CardComponent: React.FC<CardComponentProps> = ({ mode, product, onClick }) => {
  const { saveItem } = useSavedItems();
  const [compareList, setCompareList] = useState<Product[]>([]);
  const [added, setAdded] = useState(false);

  const handleSaveItem = () => {
    saveItem(product);
    alert("Item saved!");
  };

  function addToCompare() {
    const stored = localStorage.getItem("compareProduct");
    const list: Product[] = stored ? JSON.parse(stored) : [];

    if (list.some((p) => p.id === product.id)) return;
    if (list.length >= 4) {
      alert("You can only compare up to 4 products.");
      return;
    }

    const updated = [...list, product];
    localStorage.setItem("compareProduct", JSON.stringify(updated));
    setCompareList(updated);
    setAdded(true);
    window.dispatchEvent(new Event("storage"));
  }

  function removeFromCompare() {
    const stored = localStorage.getItem("compareProduct");
    const list: Product[] = stored ? JSON.parse(stored) : [];

    const updated = list.filter((p) => p.id !== product.id);
    localStorage.setItem("compareProduct", JSON.stringify(updated));
    setCompareList(updated);
    setAdded(false);
    window.dispatchEvent(new Event("storage"));
  }

  useEffect(() => {
    const list = localStorage.getItem("compareProduct");
    if (list) {
      const stored: Product[] = JSON.parse(list);
      setAdded(stored.some((p) => p.id === product.id));
      setCompareList(stored);
    }
  }, [product.id]);

  return (
    <div
      onClick={onClick}
      className={`w-full h-fit border-[2px] rounded-[12px] cursor-pointer transition hover:shadow-lg ${
        mode === "dark" ? "border-[#262B32] bg-[#262B32]" : "border-[#E5E7EB] bg-[#FFFFFF]"
      }`}
    >
      {/* Image */}
      <div className="h-fit grid place-items-center rounded-tl-[11px] rounded-tr-[11px] bg-amber-300 overflow-hidden">
        <img src={product.imageUrl} alt={product.title} className="w-full h-full object-cover" />
      </div>

      {/* Content */}
      <div className="h-[60%] p-4">
        <h6 className={`text-[14px] ${mode === "light" ? "text-[#262B32]" : "text-white"} font-semibold`}>
          {product.title}
        </h6>

        <div className="flex justify-between items-center mt-3 mb-4">
          <h6 className="text-[#FFD300] text-[18px]">${product.price.usd}</h6>
          <div className="flex items-center gap-1">
            <span className={`${mode === "dark" ? "text-white" : "text-[#262B32]"} text-[12px]`}>
              {product.productRating}
            </span>
            <Star className="w-3 h-3 text-yellow-400" />
          </div>
        </div>

        {/* Compare */}
        <button
          onClick={(e) => {
            e.stopPropagation();
            added ? removeFromCompare() : addToCompare();
          }}
          className={`${
            added ? "bg-[#757B81] text-black" : `bg-[#FFD300] ${mode === "dark" ? "text-black" : "text-[#F3F4F6]"}`
          } w-full rounded-[5px] mb-2 h-[32px]`}
        >
          {!added ? "Add To Compare" : "Remove From Compare"}
        </button>

        {/* Buy */}
        <button
          onClick={(e) => e.stopPropagation()}
          className="border-[2px] w-full rounded-[5px] h-[36px] text-[#FFD300] border-[#FFD300]"
        >
          <a href={product.deeplinkUrl} target="_blank" rel="noreferrer">
            Buy On AliExpress
          </a>
        </button>

        {/* Save */}
        <button
          onClick={(e) => {
            e.stopPropagation();
            handleSaveItem();
          }}
          className="border-[2px] w-full mt-2 rounded-[5px] h-[36px] text-[#FFD300] border-[#FFD300]"
        >
          Save Item
        </button>
      </div>
    </div>
  );
};

export default CardComponent;
