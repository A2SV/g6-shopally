//src/types/savedItems/AlertCreateResponse.ts
export interface AlertCreateResponse {
  data: {
    alertId?: string;
    status: string;
  } | null;
  error: string | null;
}
