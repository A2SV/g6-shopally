// src/lib/firebase.ts
import { getApps, initializeApp } from "firebase/app";
import { getMessaging, Messaging } from "firebase/messaging";

const firebaseConfig = {
  apiKey: "AIzaSyAlycppLAEkwEUAB_OtYvzfazz7P-VQ94s",
  authDomain: "shopally-944a4.firebaseapp.com",
  projectId: "shopally-944a4",
  storageBucket: "shopally-944a4.firebasestorage.app",
  messagingSenderId: "999033337399",
  appId: "1:999033337399:web:1a51e362f1a4f67ac131b2",
  measurementId: "G-E4ZGHQT7R0",
};

// ✅ Initialize Firebase app safely
const app = !getApps().length ? initializeApp(firebaseConfig) : getApps()[0];

// ✅ Messaging should only be initialized on the client
let messaging: Messaging | null = null;
if (typeof window !== "undefined" && "serviceWorker" in navigator) {
  try {
    messaging = getMessaging(app);
  } catch (err) {
    console.warn(
      "⚠️ Firebase Messaging not supported in this environment:",
      err
    );
  }
}

export { app, messaging };
