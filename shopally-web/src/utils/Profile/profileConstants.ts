import { CategoryOption, CurrencyOption, RegionOption } from "@/types/Profile/ProfileTypes";

export const PRODUCT_CATEGORIES: CategoryOption[] = [
  { value: "Electronics", label: "Electronics", translatedLabel: "Electronics" },
  { value: "Fashion", label: "Fashion", translatedLabel: "Fashion" },
  { value: "Home & Garden", label: "Home & Garden", translatedLabel: "Home & Garden" },
  { value: "Sports & Outdoors", label: "Sports & Outdoors", translatedLabel: "Sports & Outdoors" },
  { value: "Books", label: "Books", translatedLabel: "Books" },
  { value: "Beauty & Health", label: "Beauty & Health", translatedLabel: "Beauty & Health" },
  { value: "Automotive", label: "Automotive", translatedLabel: "Automotive" },
  { value: "Toys & Games", label: "Toys & Games", translatedLabel: "Toys & Games" },
];

export const CURRENCY_OPTIONS: CurrencyOption[] = [
  { value: "USD", label: "USD", symbol: "$" },
  { value: "EUR", label: "EUR", symbol: "€" },
  { value: "GBP", label: "GBP", symbol: "£" },
  { value: "JPY", label: "JPY", symbol: "¥" },
  { value: "CAD", label: "CAD", symbol: "C$" },
  { value: "AUD", label: "AUD", symbol: "A$" },
  { value: "CHF", label: "CHF", symbol: "CHF" },
  { value: "CNY", label: "CNY", symbol: "¥" },
  { value: "INR", label: "INR", symbol: "₹" },
  { value: "BRL", label: "BRL", symbol: "R$" },
];

export const REGION_OPTIONS: RegionOption[] = [
  { value: "US", label: "US", flag: "🇺🇸" },
  { value: "CA", label: "CA", flag: "🇨🇦" },
  { value: "GB", label: "GB", flag: "🇬🇧" },
  { value: "DE", label: "DE", flag: "🇩🇪" },
  { value: "FR", label: "FR", flag: "🇫🇷" },
  { value: "IT", label: "IT", flag: "🇮🇹" },
  { value: "ES", label: "ES", flag: "🇪🇸" },
  { value: "JP", label: "JP", flag: "🇯🇵" },
  { value: "AU", label: "AU", flag: "🇦🇺" },
  { value: "BR", label: "BR", flag: "🇧🇷" },
  { value: "IN", label: "IN", flag: "🇮🇳" },
  { value: "CN", label: "CN", flag: "🇨🇳" },
  { value: "KR", label: "KR", flag: "🇰🇷" },
  { value: "MX", label: "MX", flag: "🇲🇽" },
  { value: "NL", label: "NL", flag: "🇳🇱" },
  { value: "SE", label: "SE", flag: "🇸🇪" },
  { value: "NO", label: "NO", flag: "🇳🇴" },
  { value: "DK", label: "DK", flag: "🇩🇰" },
  { value: "FI", label: "FI", flag: "🇫🇮" },
  { value: "CH", label: "CH", flag: "🇨🇭" },
];

export const LANGUAGE_OPTIONS = [
  { value: "English", label: "English" },
  { value: "Amharic", label: "አማርኛ" },
  { value: "Spanish", label: "Español" },
  { value: "French", label: "Français" },
  { value: "German", label: "Deutsch" },
  { value: "Italian", label: "Italiano" },
  { value: "Portuguese", label: "Português" },
  { value: "Russian", label: "Русский" },
  { value: "Chinese", label: "中文" },
  { value: "Japanese", label: "日本語" },
  { value: "Korean", label: "한국어" },
];

export const PROFILE_IMAGE_MAX_SIZE = 5 * 1024 * 1024; // 5MB
export const PROFILE_IMAGE_ACCEPTED_TYPES = [
  "image/jpeg",
  "image/jpg", 
  "image/png",
  "image/webp"
];

export const PROFILE_UPDATE_MESSAGES = {
  SUCCESS: "Profile updated successfully!",
  NO_CHANGES: "No changes detected. Nothing to update.",
  ERROR: "Failed to update profile. Please try again.",
  IMAGE_UPLOAD_SUCCESS: "Profile image uploaded successfully!",
  IMAGE_UPLOAD_ERROR: "Failed to upload image. Please try again.",
  PREFERENCES_SAVED: "Preferences saved successfully!",
  PREFERENCES_ERROR: "Failed to save preferences. Please try again.",
} as const;

export const VALIDATION_MESSAGES = {
  REQUIRED_FIELD: "This field is required",
  INVALID_EMAIL: "Please enter a valid email address",
  PASSWORD_MIN_LENGTH: "Password must be at least 8 characters long",
  PASSWORD_MISMATCH: "Passwords do not match",
  INVALID_IMAGE_TYPE: "Please select a valid image file (JPEG, PNG, or WebP)",
  IMAGE_TOO_LARGE: "Image file size must be less than 5MB",
} as const;
