"use client";

import { useEffect, useState } from "react";

interface ToggleSwitchProps {
  checked: boolean; // controlled
  onChange: () => void;
}

export default function ToggleSwitch({ checked, onChange }: ToggleSwitchProps) {
  const [enabled, setEnabled] = useState(checked);

  // Keep local state in sync with controlled prop
  useEffect(() => {
    setEnabled(checked);
  }, [checked]);

  return (
    <button
      onClick={() => {
        setEnabled(!enabled); // immediate animation
        onChange(); // update parent state
      }}
      className={`w-10 h-5 flex items-center rounded-full p-1 transition-colors duration-300 ${
        enabled ? "bg-[#FFD300]" : "bg-gray-300"
      }`}
    >
      <div
        className={`bg-white w-4 h-4 rounded-full shadow-md transform transition-transform duration-300 ${
          enabled ? "translate-x-5" : "translate-x-0"
        }`}
      />
    </button>
  );
}
