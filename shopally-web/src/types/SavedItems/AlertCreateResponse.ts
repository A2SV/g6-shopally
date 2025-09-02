export interface AlertCreateResponse {
  data: {
    alertId?: string;
    status: string;
  } | null;
  error: string | null;
}
