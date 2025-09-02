// src/utils/deviceId.ts
import { v4 as uuidv4 } from "uuid";

export function getDeviceId(): string {
  if (typeof window === "undefined") return "unknown";
  let deviceId = localStorage.getItem("deviceId");
  if (!deviceId) {
    deviceId = uuidv4();
    localStorage.setItem("deviceId", deviceId);
  }
  return deviceId;
}
