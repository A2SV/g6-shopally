//src/providers/AuthProvider.tsx
"use client";
import type { Session } from "next-auth";
import { SessionProvider } from "next-auth/react";
import { ReactNode } from "react";
interface AuthProviderProps {
  children: ReactNode;
  session?: Session | null;
}

const AuthProvider = ({ children, session }: AuthProviderProps) => {
  return <SessionProvider session={session}>{children}</SessionProvider>;
};

export default AuthProvider;
