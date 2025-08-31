import { useState, useEffect, useCallback } from "react";
import {
  ProfileData,
  ProfileUpdateRequest,
  UserProfile,
  UserPreferences,
  ProfileStatistics,
} from "@/types/Profile/ProfileTypes";
import {
  fetchProfileData,
  updateProfile,
  uploadProfileImage,
  fetchUserStatistics,
  updateUserPreferences,
} from "@/utils/Profile/profileApi";
import { PROFILE_UPDATE_MESSAGES } from "@/utils/Profile/profileConstants";

interface UseProfileReturn {
  // Data
  profileData: ProfileData | null;
  user: UserProfile | null;
  preferences: UserPreferences | null;
  statistics: ProfileStatistics | null;
  
  // Loading states
  isLoading: boolean;
  isUpdating: boolean;
  isUploadingImage: boolean;
  
  // Error states
  error: string | null;
  
  // Actions
  updateProfileData: (updateData: ProfileUpdateRequest) => Promise<boolean>;
  uploadImage: (file: File) => Promise<boolean>;
  updatePreferences: (preferences: Partial<UserPreferences>) => Promise<boolean>;
  refreshProfile: () => Promise<void>;
  clearError: () => void;
}

export const useProfile = (): UseProfileReturn => {
  const [profileData, setProfileData] = useState<ProfileData | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [isUpdating, setIsUpdating] = useState(false);
  const [isUploadingImage, setIsUploadingImage] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Computed values
  const user = profileData?.user || null;
  const preferences = profileData?.preferences || null;
  const statistics = profileData?.statistics || null;

  // Fetch profile data
  const fetchData = useCallback(async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await fetchProfileData();
      setProfileData(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to fetch profile data");
    } finally {
      setIsLoading(false);
    }
  }, []);

  // Update profile data
  const updateProfileData = useCallback(async (updateData: ProfileUpdateRequest): Promise<boolean> => {
    try {
      setIsUpdating(true);
      setError(null);
      
      const response = await updateProfile(updateData);
      
      if (response.success) {
        // Update local state with new data
        setProfileData(prev => {
          if (!prev) return null;
          
          return {
            ...prev,
            user: response.user ? { ...prev.user, ...response.user } : prev.user,
            preferences: response.preferences ? { ...prev.preferences, ...response.preferences } : prev.preferences,
          };
        });
        
        return true;
      } else {
        setError(response.message || PROFILE_UPDATE_MESSAGES.ERROR);
        return false;
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : PROFILE_UPDATE_MESSAGES.ERROR);
      return false;
    } finally {
      setIsUpdating(false);
    }
  }, []);

  // Upload profile image
  const uploadImage = useCallback(async (file: File): Promise<boolean> => {
    try {
      setIsUploadingImage(true);
      setError(null);
      
      const imageUrl = await uploadProfileImage(file);
      
      // Update local state with new image URL
      setProfileData(prev => {
        if (!prev) return null;
        
        return {
          ...prev,
          user: {
            ...prev.user,
            profileImage: imageUrl,
          },
        };
      });
      
      return true;
    } catch (err) {
      setError(err instanceof Error ? err.message : PROFILE_UPDATE_MESSAGES.IMAGE_UPLOAD_ERROR);
      return false;
    } finally {
      setIsUploadingImage(false);
    }
  }, []);

  // Update preferences
  const updatePreferences = useCallback(async (newPreferences: Partial<UserPreferences>): Promise<boolean> => {
    try {
      setIsUpdating(true);
      setError(null);
      
      const updatedPreferences = await updateUserPreferences(newPreferences);
      
      // Update local state
      setProfileData(prev => {
        if (!prev) return null;
        
        return {
          ...prev,
          preferences: updatedPreferences,
        };
      });
      
      return true;
    } catch (err) {
      setError(err instanceof Error ? err.message : PROFILE_UPDATE_MESSAGES.PREFERENCES_ERROR);
      return false;
    } finally {
      setIsUpdating(false);
    }
  }, []);

  // Refresh profile data
  const refreshProfile = useCallback(async () => {
    await fetchData();
  }, [fetchData]);

  // Clear error
  const clearError = useCallback(() => {
    setError(null);
  }, []);

  // Initial data fetch
  useEffect(() => {
    fetchData();
  }, [fetchData]);

  return {
    // Data
    profileData,
    user,
    preferences,
    statistics,
    
    // Loading states
    isLoading,
    isUpdating,
    isUploadingImage,
    
    // Error states
    error,
    
    // Actions
    updateProfileData,
    uploadImage,
    updatePreferences,
    refreshProfile,
    clearError,
  };
};
