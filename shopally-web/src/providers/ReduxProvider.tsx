//src/providers/ReduxProvider.tsx
"use client";
import { store } from "@/lib/redux/store";
import { ReactNode } from "react";
import { Provider } from "react-redux";

function ReduxProvider({ children }: Readonly<{ children: ReactNode }>) {
  return <Provider store={store}>{children}</Provider>;
}
export default ReduxProvider;
