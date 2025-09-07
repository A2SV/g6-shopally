// src/app/components/home-page-component/page.tsx
"use client";
import { useSavedItems } from "@/hooks/useSavedItems";
import { Product } from "@/types/types";
import { Star } from "lucide-react";
import React, { useEffect, useState } from "react";

interface CardComponentProps {
  mode: "dark" | "light";
  product: Product;
}

const CardComponent: React.FC<CardComponentProps> = ({ mode, product }) => {
  const { saveItem } = useSavedItems();
  const [compareList, setCompareList] = useState<Product[]>([]);
  const [added, setAdded] = useState(false);

  const handleSaveItem = () => {
    const savedItem = {
      id: product.id,
      title: product.title,
      imageUrl: product.imageUrl,
      aiMatchPercentage: product.aiMatchPercentage,
      price: product.price,
      productRating: product.productRating ?? 0, // ✅ fallback to 0 if undefined
      sellerScore: product.sellerScore, // use sellerScore, not numberSold
      deliveryEstimate: product.deliveryEstimate,
      summaryBullets: product.summaryBullets,
      deeplinkUrl: product.deeplinkUrl,
    };

    saveItem(savedItem);
    alert("Item saved!");
  };

  function addToCompare() {
    const stored = localStorage.getItem("compareProduct");
    const list: Product[] = stored ? JSON.parse(stored) : [];

    // check if already exists
    if (list.some((p) => p.id === product.id)) {
      console.log("⚠️ Product already in compare list:", product.title);
      return;
    }

    // enforce max 4
    if (list.length >= 4) {
      console.log("⚠️ Cannot add more than 4 products to compare list");
      alert("You can only compare up to 4 products.");
      return;
    }

    const updated = [...list, product];
    localStorage.setItem("compareProduct", JSON.stringify(updated));
    setCompareList(updated);
    setAdded(true);

    console.log("✅ Added to compare list:", product.title, updated);

    // 🔔 Notify other components
    window.dispatchEvent(new Event("storage"));
  }

  function removeFromCompare() {
    const stored = localStorage.getItem("compareProduct");
    const list: Product[] = stored ? JSON.parse(stored) : [];

    const updated = list.filter((p) => p.id !== product.id);
    localStorage.setItem("compareProduct", JSON.stringify(updated));
    setCompareList(updated);
    setAdded(false);

    console.log("❌ Removed from compare list:", product.title, updated);

    // 🔔 Notify other components
    window.dispatchEvent(new Event("storage"));
  }

  useEffect(() => {
    const list = localStorage.getItem("compareProduct");
    if (list) {
      const stored: Product[] = JSON.parse(list);
      const exists = stored.some((p) => p.id === product.id);
      setAdded(exists);
      setCompareList(stored);
    }
  }, [product.id]);

  return (
    <div
      className={`w-full h-fit break-words border-[2px] rounded-[12px] ${
        mode === "dark"
          ? "border-[#262B32] bg-[#262B32]"
          : "border-[#E5E7EB] bg-[#FFFFFF]"
      }`}
    >
      <div className="h-fit grid place-items-center rounded-tl-[11px] rounded-tr-[11px] bg-amber-300 overflow-hidden">
        <img
          src={product.imageUrl}
          alt={product.title}
          className="w-full h-full object-cover"
        />
      </div>

      <div className="h-[60%] p-4">
        <h6
          className={`text-[14px] ${
            mode === "light" ? "text-[#262B32]" : "text-[#FFFFFF]"
          } font-semibold`}
        >
          {product.title}
        </h6>
        <div className="justify-between flex items-center mt-3 mb-4">
          <h6 className="text-[#FFD300] text-[18px]">${product.price.usd}</h6>
          <div className="flex items-center gap-1">
            <span
              className={`${
                mode === "dark" ? "text-white" : "text-[#262B32]"
              } text-[12px]`}
            >
              {product.productRating}
            </span>
            <Star className="w-3 h-3 text-yellow-400" />
          </div>
        </div>
        <button
          onClick={added ? removeFromCompare : addToCompare}
          className={` ${
            !added
              ? `bg-[#FFD300] ${
                  mode === "dark" ? "text-black" : "text-[#F3F4F6]"
                } w-full rounded-[5px] mb-2 h-[32px]`
              : `bg-[#757B81] ${
                  mode === "dark" ? "text-black" : "text-black"
                } w-full rounded-[5px] mb-2 h-[32px]`
          }`}
        >
          {!added ? "Add To Compare" : "Remove From Compare"}
        </button>
        <button className="border-[2px] w-full rounded-[5px] h-[36px] text-[#FFD300] border-[#FFD300]">
          <a href={product.deeplinkUrl} target="_blank" rel="noreferrer">
            Buy On AliExpress
          </a>
        </button>
        <button
          className="border-[2px] w-full mt-2 rounded-[5px] h-[36px] text-[#FFD300] border-[#FFD300]"
          onClick={handleSaveItem}
        >
          Save Item
        </button>
      </div>
    </div>
  );
};

export default CardComponent;
