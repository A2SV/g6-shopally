//src/lib/redux/api/userApiSlice.ts
import { ComparisonResponse } from "@/types/Compare/Comparison";
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { AlertDeleteResponse } from "@/types/SavedItems/AlertDeleteResponse";
import { CreateAlert } from "@/types/SavedItems/SavedItems";
import { ProductResponse, Product } from "@/types/types";
import { getOrCreateDeviceId } from "@/utils/deviceId";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getLanguage } from "../languageBridge";

export const userApi = createApi({
  reducerPath: "userApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "/api/v1",
    prepareHeaders: (headers) => {
      const langCode = getLanguage();
      if (langCode) {
        headers.set("Accept-Language", langCode);
      }

      const deviceId = getOrCreateDeviceId();
      headers.set("x-device-id", deviceId);

      return headers;
    },
  }),
  endpoints: (builder) => ({
    createAlert: builder.mutation<AlertCreateResponse, CreateAlert>({
      query: (data) => ({
        url: "alerts",
        method: "POST",
        body: data,
      }),
    }),
    deleteAlert: builder.mutation<AlertDeleteResponse, { id: string }>({
      query: (data) => ({
        url: `alerts/${data.id}`,
        method: "DELETE",
      }),
    }),
    searchProducts: builder.mutation<
      ProductResponse,
      { query: string; priceMaxETB?: number | null; minRating?: number | null; language?: string }
    >({
      query: ({ query, priceMaxETB, minRating, language }) => {
        const params = new URLSearchParams({ q: query });
        if (priceMaxETB) params.append("priceMaxETB", String(priceMaxETB));
        if (minRating) params.append("minRating", String(minRating));
        if (language) params.append("language", language);

        return {
          url: `/search?${params.toString()}`, // ✅ use query params
          method: "GET",
        };
      },
    }),

    compareProducts: builder.mutation<
      ComparisonResponse,
      { products: Product[] }
    >({
      query: (data) => ({
        url: "compare",
        method: "POST",
        body: data,
      }),
    }),
  }),
});

export const {
  useCreateAlertMutation,
  useDeleteAlertMutation,
  useSearchProductsMutation,
  useCompareProductsMutation,
} = userApi;
