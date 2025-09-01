import { SavedItemUI } from "@/types/types";
import SavedItemCard from "./SavedItemCard";

interface SavedItemsContentProps {
  savedItems: SavedItemUI[];
  removeItem: (id: string) => void;
  updateItemPrice: (id: string, price: SavedItemUI["price"]) => void;
  onToggleAlert: (id: string) => void;
}

export default function SavedItemsContent({
  savedItems,
  removeItem,
  updateItemPrice,
  onToggleAlert,
}: SavedItemsContentProps) {
  return (
    <div>
      {savedItems.length > 0 ? (
        savedItems.map((item) => (
          <SavedItemCard
            key={item.id}
            {...item}
            onRemove={removeItem}
            onUpdatePrice={updateItemPrice}
            onToggleAlert={onToggleAlert}
          />
        ))
      ) : (
        <p className="text-gray-500 text-center mt-10">No saved items yet.</p>
      )}
    </div>
  );
}
