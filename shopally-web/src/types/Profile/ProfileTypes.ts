export interface UserProfile {
  id: string;
  fullName: string;
  email: string;
  profileImage?: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserPreferences {
  preferredCategories: string[];
  defaultCurrency: string;
  shippingRegion: string;
  languagePreference: string;
  darkModeEnabled: boolean;
}

export interface ProfileStatistics {
  productsCompared: number;
  ordersPlaced: number;
  totalSaved: number;
  aiMatchRate: number;
  wishlistItems: number;
  reviews: number;
}

export interface ProfileUpdateRequest {
  fullName?: string;
  email?: string;
  profileImage?: string;
  preferences?: Partial<UserPreferences>;
}

export interface ProfileUpdateResponse {
  success: boolean;
  message: string;
  user?: UserProfile;
  preferences?: UserPreferences;
}

export interface ProfileData {
  user: UserProfile;
  preferences: UserPreferences;
  statistics: ProfileStatistics;
}

export interface ProfileFormData {
  fullName: string;
  email: string;
  password?: string;
  selectedCategories: string[];
  currency: string;
  region: string;
  language: string;
}

export interface CategoryOption {
  value: string;
  label: string;
  translatedLabel: string;
}

export interface CurrencyOption {
  value: string;
  label: string;
  symbol: string;
}

export interface RegionOption {
  value: string;
  label: string;
  flag?: string;
}
