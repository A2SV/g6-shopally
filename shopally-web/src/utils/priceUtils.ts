// Utility functions for price formatting, prioritizing ETB for Ethiopian users

// Current USD to ETB conversion rate (you can update this or fetch from an API)
const USD_TO_ETB_RATE = 57.5; // Approximate rate, should be updated regularly

export interface Price {
  etb: number;
  usd: number;
  fxTimestamp: string;
}

/**
 * Formats price to always show ETB for Ethiopian users
 * Falls back to USD conversion if ETB is not available or is 0
 */
export function formatPriceForEthiopia(price: Price): string {
  // If ETB is available and greater than 0, use it
  if (price.etb && price.etb > 0) {
    return `${Number(price.etb).toFixed(2)} ETB`;
  }
  
  // If ETB is not available or is 0, convert USD to ETB
  if (price.usd && price.usd > 0) {
    const convertedETB = price.usd * USD_TO_ETB_RATE;
    return `${convertedETB.toFixed(2)} ETB`;
  }
  
  // Fallback if both are unavailable
  return "Price not available";
}

/**
 * Gets the ETB value, converting from USD if necessary
 */
export function getETBValue(price: Price): number {
  // If ETB is available and greater than 0, use it
  if (price.etb && price.etb > 0) {
    return price.etb;
  }
  
  // If ETB is not available or is 0, convert USD to ETB
  if (price.usd && price.usd > 0) {
    return price.usd * USD_TO_ETB_RATE;
  }
  
  // Fallback
  return 0;
}

/**
 * Formats price with currency symbol for display
 */
export function formatPriceWithCurrency(price: Price, currency: 'ETB' | 'USD' = 'ETB'): string {
  if (currency === 'ETB') {
    return formatPriceForEthiopia(price);
  }
  
  // For USD display
  if (price.usd && price.usd > 0) {
    return `$${Number(price.usd).toFixed(2)}`;
  }
  
  return "Price not available";
}

/**
 * Updates the USD to ETB conversion rate
 */
export function updateConversionRate(newRate: number): void {
  // This could be called when fetching real-time exchange rates
  // For now, we'll use a static rate
  console.log(`Conversion rate updated to: ${newRate}`);
}
