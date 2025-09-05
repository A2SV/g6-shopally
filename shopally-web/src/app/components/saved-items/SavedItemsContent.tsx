import { SavedItemUI } from "@/types/types";
import SavedItemCard from "./SavedItemCard";

interface SavedItemsContentProps {
  savedItems: SavedItemUI[];
  removeItem: (id: string) => void;
  updateItemPrice: (id: string, price: SavedItemUI["price"]) => void;
  onToggleAlert: (id: string) => void;
  onPlaceOrder: (
    productId: string,
    productTitle: string,
    price: SavedItemUI["price"]
  ) => void;
}

export default function SavedItemsContent({
  savedItems,
  removeItem,
  updateItemPrice,
  onToggleAlert,
  onPlaceOrder,
}: SavedItemsContentProps) {
  return (
    <div>
      {savedItems.length > 0 ? (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 sm:gap-5 lg:gap-6">
          {savedItems.map((item) => (
            <SavedItemCard
              key={item.id}
              {...item}
              onRemove={removeItem}
              onUpdatePrice={updateItemPrice}
              onToggleAlert={onToggleAlert}
              onPlaceOrder={onPlaceOrder}
            />
          ))}
        </div>
      ) : (
        <div className="text-center mt-8 sm:mt-12 lg:mt-16">
          <p className="text-gray-500 text-sm sm:text-base lg:text-lg">
            No saved items yet.
          </p>
        </div>
      )}
    </div>
  );
}
