//src/lib/redux/api/userApiSlice.ts
import { ComparisonResponse, Product } from "@/types/Compare/Comparison";
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { AlertDeleteResponse } from "@/types/SavedItems/AlertDeleteResponse";
import { CreateAlert } from "@/types/SavedItems/SavedItems";
import { ProductResponse } from "@/types/types";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const userApi = createApi({
  reducerPath: "userApi",
  baseQuery: fetchBaseQuery({ baseUrl: "/api/" }),
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
      { query: string; priceMaxETB: number | null; minRating: number | null }
    >({
      query: (data) => ({
        url: "products/search",
        method: "POST",
        body: data,
      }),
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
