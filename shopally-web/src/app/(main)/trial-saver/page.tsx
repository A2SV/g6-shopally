// shopally-web/src/app/(main)/trial-saver/page.tsx
import TrialSaverClient from "@/app/components/TrialCards/TrialSaverClient";
import { fetchProducts } from "@/utils/Home/productsUtils";

export default async function ProductsPage() {
  const products = await fetchProducts();

  return (
    <div>
      <div>
        <TrialSaverClient products={products} />
      </div>
    </div>
  );
}
